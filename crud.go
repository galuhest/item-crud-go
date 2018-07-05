package main

import (
	"fmt"
	"os"
  "encoding/json"
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"

)

type Response struct {
	Status string `json:"status"` 
	Payload map[string]string `json:"payload,omitempty"`
}

func get(id int) string {
	err := godotenv.Load()
  if err != nil {
    panic(err.Error())
  }

	db_config := fmt.Sprintf("%s:%s@/%s",os.Getenv("DB_USER"),os.Getenv("DB_PASSWORD"),os.Getenv("DATABASE"))
	db, err := sql.Open("mysql", db_config)
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	
	stmtOut, err := db.Prepare("SELECT name FROM item WHERE id = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()
	
	var name string

	err = stmtOut.QueryRow(id).Scan(&name)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	payload := make(map[string]string)
	payload["name"] = name
	response := Response{"OK",payload}
	js, err := json.Marshal(response)
	if err != nil {
    panic(err.Error())
  }
	return string(js)
}

func create(name string)	string {
	err := godotenv.Load()
  if err != nil {
    panic(err.Error())
  }

	db_config := fmt.Sprintf("%s:%s@/%s",os.Getenv("DB_USER"),os.Getenv("DB_PASSWORD"),os.Getenv("DATABASE"))
	db, err := sql.Open("mysql", db_config)
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	stmtIns, err := db.Prepare("INSERT INTO item (name) VALUES(?)") // ? = placeholderl
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates
	
	_, err = stmtIns.Exec(name) // Insert tuples (i, i^2)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	stmtOut, err := db.Prepare("SELECT LAST_INSERT_ID()")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()
	
	var id int

	err = stmtOut.QueryRow().Scan(&id)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	payload := make(map[string]string)
	payload["id"] = string(id)
	response := Response{"OK",payload}
	js, err := json.Marshal(response)
	if err != nil {
    panic(err.Error())
  }
	return string(js)
}

func update(id int, name string) string	{
	err := godotenv.Load()
  if err != nil {
    panic(err.Error())
  }

	db_config := fmt.Sprintf("%s:%s@/%s",os.Getenv("DB_USER"),os.Getenv("DB_PASSWORD"),os.Getenv("DATABASE"))
	db, err := sql.Open("mysql", db_config)
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}

	defer db.Close()

	stmtIns, err := db.Prepare("update item set name = \"?\" where id = ?") // ? = placeholderl
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	_, err = stmtIns.Exec(id,name) // Insert tuples (i, i^2)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	response := Response{"OK",make(map[string]string)}
	js, err := json.Marshal(response)
	if err != nil {
    panic(err.Error())
  }
	return string(js)	
}

func delete(id int)	string	{
	err := godotenv.Load()
  if err != nil {
    panic(err.Error())
  }

	db_config := fmt.Sprintf("%s:%s@/%s",os.Getenv("DB_USER"),os.Getenv("DB_PASSWORD"),os.Getenv("DATABASE"))
	db, err := sql.Open("mysql", db_config)
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	stmtIns, err := db.Prepare("delete from item where id = ?") // ? = placeholderl
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates
	
	_, err = stmtIns.Exec(id) // Insert tuples (i, i^2)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	response := Response{"OK",make(map[string]string)}
	js, err := json.Marshal(response)
	if err != nil {
    panic(err.Error())
  }
	return string(js)	
}