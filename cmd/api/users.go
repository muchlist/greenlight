package main

import (
	"errors"
	"github.com/muchlist/greenlight/internal/data"
	"github.com/muchlist/greenlight/internal/validator"
	"net/http"
)

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := &data.User{
		Name:      input.Name,
		Email:     input.Email,
		Activated: false,
	}

	// Use the Password.Set() method to generate and store the hashed and plaintext
	// passwords.
	err = user.Password.Set(input.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	v := validator.New()

	if data.ValidateUser(v, user); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Users.Insert(user)
	if err != nil {
		switch {
		// If we get a ErrDuplicateEmail error, use the v.AddError() method to manually
		// add a message to the validator instance, and then call our
		// failedValidationResponse() helper.
		case errors.Is(err, data.ErrDuplicateEmail):
			v.AddError("email", "a user with this email address already exists")
			app.failedValidationResponse(w, r, v.Errors)
		default:
			app.serverErrorResponse(w, r, err)
		}

		return
	}

	// Launch a goroutine which runs an anonymous function that sends the welcome email.
	//go func() {
	//	// Run a deferred function which uses recover() to catch any panic, and log an
	//	// error message instead of terminating the application.
	//	defer func() {
	//		if err := recover(); err != nil {
	//			app.logger.PrintError(fmt.Errorf("%s", err), nil)
	//		}
	//	}()
	//	err = app.mailer.Send(user.Email, "user_welcome.tmpl", user)
	//	if err != nil {
	//		// Importantly, if there is an error sending the email then we use the
	//		// app.logger.PrintError() helper to manage it, instead of the
	//		// app.serverErrorResponse() helper like before.
	//		app.logger.PrintError(err, nil)
	//	}
	//}()

	// Use the background helper to execute an anonymous function that sends the welcome
	// email.
	app.background(func() {
		err = app.mailer.Send(user.Email, "user_welcome.tmpl", user)
		if err != nil {
			app.logger.PrintError(err, nil)
		}
	})

	// client a 202 Accepted status code.
	// This status code indicates that the request has been accepted for processing, but
	// the processing has not been completed.
	err = app.writeJSON(w, http.StatusAccepted, envelope{"user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
