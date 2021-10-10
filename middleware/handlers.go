package middleware

import (
	"database/sql"
	"encoding/json" //encode and decode json to struct
	"fmt"
	"log"
	"net/http" //access request and response object of the api
	"os"       //use to read environment variable
    "strconv"

	"github.com/aryanicosa/go-crud-postgre/models"
    
    "github.com/gorilla/mux" // used to get the params from the route

	"github.com/joho/godotenv" // package used to read the .env file
	_ "github.com/lib/pq"      // postgres golang driver
)

//respnse format
type response struct {
	ID 		int64 	`json:"id,omitempty"`
	Message string 	`json:"messages,omitempty"`
}

//create connection to database
func createConnection() *sql.DB {
	//load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	//open connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_DB"))

	if err != nil {
		panic(err)
	}

	//check connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
    // return the connection
    return db
}

// CreateUser create a user in the postgres db
func CreateUser(w http.ResponseWriter, r *http.Request) {
    // set the header to content type x-www-form-urlencoded
    // Allow all origin to handle cors issue
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    // create an empty user of type models.User
    var user models.User

    // decode the json request to user
    err := json.NewDecoder(r.Body).Decode(&user)

    if err != nil {
        log.Fatalf("Unable to decode the request body.  %v", err)
    }

    // call insert user function and pass the user
    insertID := insertUser(user)

    // format a response object
    res := response{
        ID:      insertID,
        Message: "User created successfully",
    }

    // send the response
    json.NewEncoder(w).Encode(res)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    //get the useid from params, key is "id"
    params := mux.Vars(r)

    //convert id form string to int
    id, err := strconv.Atoi(params["id"])

    if err != nil {
        log.Fatalf("unable to convert the string into int, value : %v", err)
    }

    //call the getUser function
    user, err := getUser(int64(id))

    if err != nil {
        log.Fatalf("unable to get user, value : %v", err)
    }

    //send response
    json.NewEncoder(w).Encode(user)
}

//------------------------- handler functions ----------------
// insert one user in the DB
func insertUser(user models.User) int64 {

    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    // create the insert sql query
    // returning userid will return the id of the inserted user
    sqlStatement := `INSERT INTO users (name, location, age) VALUES ($1, $2, $3) RETURNING userid`

    // the inserted id will store in this id
    var id int64

    // execute the sql statement
    // Scan function will save the insert id in the id
    err := db.QueryRow(sqlStatement, user.Name, user.Location, user.Age).Scan(&id)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    fmt.Printf("Inserted a single record %v", id)

    // return the inserted id
    return id
}

//function to get a user by id
func getUser(id int64) (models.User, error) {
    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    //create a user of model.User type
    var user models.User

    //query the user
    sqlStatement := `SELECT * FROM users where userid=$1`

    //execute the quey
    row := db.QueryRow(sqlStatement, id)

    //unmarshall the row object to user struct
    err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Location)

    switch err {
    case sql.ErrNoRows:
        fmt.Println("No row were returned!")
        return user, nil
    case nil:
        return user, nil
    default:
        log.Fatalf("Unable to scan the row, %v", err)
    }

    //return empty on error
    return user, err
}