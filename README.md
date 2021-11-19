# greenlight

|Method | Url Pattern |  Handler |  Action | Permission |
|---|---|---|---|---|
| GET | /v1/healthcheck | healthcheckHandler | Show application information  | |
| GET | /v1/movies | listMoviesHandler | Show the details of all movies  | movies:read |
| POST | /v1/movies | createMovieHandler | Create a new movie  | movies:write|
| GET | /v1/movies/:id | showMovieHandler | Show the details of a specific movie |movies:read |
| PATCH | /v1/movies/:id | updateMovieHandler | Update the details of a specific movie |  movies:write |
| DELETE | /v1/movies/:id | deleteMovieHandler | Delete a specific movie   | movies:write|
| POST | /v1/users | registerUserHandler | Register a new user  | |
| PUT | /v1/users/activated | activateUserHandler | Activate a specific user |  |
| POST | /v1/tokens/authentication | createAuthenticationTokenHandler | Generate a new authentication token  | |
| GET | /debug/vars | expvar.Handler() | Display application metrics  | |

## Database Postgres

CREATE DATABASE greenlight;  
CREATE ROLE greenlight WITH LOGIN PASSWORD 'pa55word';  
CREATE EXTENSION IF NOT EXISTS citext;  
DSN : postgres://greenlight:pa55word@localhost/greenlight

## Maintenance Command


```    
go mod tidy
go mod verify  
go fmt ./...  
go vet ./...
go test -race -vet=off ./...

go install honnef.co/go/tools/cmd/staticcheck@latest
which staticcheck
   
```
