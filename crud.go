package main

import (
	"fmt"
	"os"
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"

)

func main() {
	err := godotenv.Load()
  if err != nil {
    panic(err.Error())
  }

	fmt.Printf("Hello, world!\n")
	db_config := fmt.Sprintf("%s:%s@/%s",os.Getenv("DB_USER"),os.Getenv("DB_PASSWORD"),os.Getenv("DATABASE"))
	db, err := sql.Open("mysql", db_config)
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Prepare statement for inserting data
	stmtIns, err := db.Prepare("INSERT INTO item (name) VALUES(?)") // ? = placeholderl
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT name FROM item WHERE id = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	// Insert square numbers for 0-24 in the database
	for i := 0; i < 25; i++ {
		_, err = stmtIns.Exec(i) // Insert tuples (i, i^2)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}

	var squareNum string // we "scan" the result in here

	// Query the square-number of 13
	err = stmtOut.QueryRow(13).Scan(&squareNum) // WHERE number = 13
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The square number of 13 is: %s\n", squareNum)

	// Query another number.. 1 maybe?
	err = stmtOut.QueryRow(2).Scan(&squareNum) // WHERE number = 1
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The square number of 1 is: %s\n", squareNum)
}