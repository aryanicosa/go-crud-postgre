# go-crud-postgre

In this project we will create CRUD application with GO and Postgresql

1. go mod init github.com/aryanicosa/go-crud-postgre
"go modules is dependency manager, it is similar to the package.json in nodejs"

2. Installing dependency (if never installed before)
- `go get -u github.com/gorilla/mux` -> for implementing a request router and dispatcher for matching incoming requests to their respective handle
- `go get -u github.com/lib/pq` -> A pure Go postgre driver for Go database/sql package
- `go get github.com/joho/godotenv` -> For saving environment variables to keep sensitive data

3. Create database
*switch to root account
- createdb go_crud_1
- psql -d go_crud_1
- \conninfo

- CREATE TABLE users (
    userid SERIAL PRIMARY KEY,
    name TEXT,
    age INT,
    location TEXT
);

check your database and table, if it is okay let just move to the code

4. Create models to store database schema
- cd models -> `touch model.go`

5. Create middleware
- cd .. -> cd middleware -> `touch handlers.go`

6. Create router
- cd .. -> cd router -> `touch router.go`

7. Create main.go
- cd .. -> `touch main.go`

7. Create the `.env` file
- ` touch .env` # do not forget to include this file to git ignore

write `POSTGRE_DB="postgres://username:password@hostname/databasename"` in the .env file

8. Run the project
- `go run main.go`

OR 

Build the project
- `go build` then run `./go-crud`