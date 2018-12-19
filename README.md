##  RESTful Pet information collector ##

# Installation
Create a database and use https://github.com/golang-migrate/migrate to create the tables with the following command:
$ migrate -database  postgres://username:postgres@localhost:5432/database_name  -verbose -path migrations up
