#  RESTful Pet information collector ##

## Requirements
Postgres database
Go
.env file with the keys set

## Installation
Create a database and use https://github.com/golang-migrate/migrate to create the tables with the following command:
`migrate -database  postgres://username:postgres@localhost:5432/database_name  -verbose -path migrations up`

## Test
`go test -v`

## Build application
`go build`

## Run
`source .env; ./pet-rego`

The application will now be listening on localhost:8000.
Example interaction can be found in examples.sh

## Notes
A pet can be optionally created with a food type, if a pet is created and the food preference is not specified, then a food type will be defaulted in, for some species. If a food type is specified, or the pet species is not in the list, then no food type will be recorded.
All previously exisitng records in the database will be updated, such that if the species is one of dog, cat, snake, or chicken, then their food type will be updated.

