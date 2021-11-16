# greenlight

## Database Postgres
CREATE DATABASE greenlight;  
CREATE ROLE greenlight WITH LOGIN PASSWORD 'pa55word';  
CREATE EXTENSION IF NOT EXISTS citext;  
DSN : postgres://greenlight:pa55word@localhost/greenlight  

## Migration  
migrate create -seq -ext=.sql -dir=./migrations create_movies_table  
migrate create -seq -ext=sql -dir=migrations create_movies_table << windows

migrate -path=./migrations -database=$GREENLIGHT_DB_DSN up  
migrate -path=./migrations -database=$GREENLIGHT_DB_DSN down 1

migrate -path=./migrations -database=$EXAMPLE_DSN version  
migrate -path=./migrations -database=$EXAMPLE_DSN force 1