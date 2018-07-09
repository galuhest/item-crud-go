// Package ini dibuat untuk membungkus pemanggilan terhadap database.
package crud

import (
    "fmt"
    "os"
    "strconv"
  "encoding/json"
    "database/sql"
    "github.com/joho/godotenv"
    _ "github.com/go-sql-driver/mysql"

)

// Struct ini adalah object kembalian dari
// seluruh function dalam package ini. 
type Response struct {
    Status string `json:"status"` 
    Payload map[string]string `json:"payload,omitempty"`
}

// init akan membaca file .env jika ada
func init() {
    godotenv.Load()
}

// Function ini akan membuat koneksi terhadap database
// berdasarkan parameter dalam .env
func ConnectDb() (*sql.DB, error) {
    db_config := fmt.Sprintf("%s:%s@/%s",os.Getenv("DB_USER"),os.Getenv("DB_PASSWORD"),os.Getenv("DATABASE"))
    db, err := sql.Open("mysql", db_config)
    return db, err
}

// GetItem akan mengembalikan nama dari user
// berdasarkan id yang diberikan. Argumen pertama
// adalah database yang digunakan, dan parameter kedua
// adalah id dari user yang dicari.
func GetItem(db *sql.DB, id int) (string, error){
    stmtOut, err := db.Prepare("SELECT name FROM item WHERE id = ?")
    if err != nil { 
        return "", err // proper error handling instead of panic in your app
    }
    defer stmtOut.Close()
    
    var name string

    err = stmtOut.QueryRow(id).Scan(&name)
    if err != nil {
        return "", err
    }
    payload := make(map[string]string)
    payload["name"] = name
    response := Response{"OK",payload}
    js, err := json.Marshal(response)
    if err != nil {
        return "", err
    }
    return string(js), err
}

// CreateItem akan memasukan user baru kedalam database.
// Function ini menerima object database sebagai parameter pertama,
// dan nama user baru dari parameter kedua.
func CreateItem(db *sql.DB, name string) (string, error) {

    stmtIns, err := db.Prepare("INSERT INTO item (name) VALUES(?)") // ? = placeholderl
    if err != nil {
        return "", err 
    }
    defer stmtIns.Close() // Close the statement when we leave main() / the program terminates
    
    _, err = stmtIns.Exec(name) // Insert tuples (i, i^2)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    stmtOut, err := db.Prepare("SELECT LAST_INSERT_ID()")
    if err != nil {
        return "", err
    }
    defer stmtOut.Close()
    
    var id int

    err = stmtOut.QueryRow().Scan(&id)
    if err != nil {
        return "", err
    }

    payload := make(map[string]string)
    payload["id"] = strconv.Itoa(id)
    response := Response{"OK",payload}
    js, err := json.Marshal(response)
    if err != nil {
        return "", err
    }
    return string(js), err
}

// UpdateItem akan mengganti nama user dengan nama baru.
// Function ini menerima 3 (tiga) parameter. Parameter
// pertama adalah object database, parameter kedua adalah id user
// yang ingin diganti namanya, dan parameter ketiga adalah nama baru.
func UpdateItem(db *sql.DB, id int, name string) (string, error)    {
    stmtIns, err := db.Prepare("update item set name = ? where id = ?") // ? = placeholderl
    if err != nil {
        return "", err
    }
    defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

    _, err = stmtIns.Exec(name,id) // Insert tuples (i, i^2)
    if err != nil {
        return "", err
    }
    response := Response{"OK",make(map[string]string)}
    js, err := json.Marshal(response)
    if err != nil {
        return "", err
    }
    return string(js), err  
}

// DeleteItem akan mengapus data user dari database berdasarkan id.
// Function ini menerima object database sebagai parameter pertama,
// dan id user yang ingin dihapus sebagai parameter kedua.
func DeleteItem(db *sql.DB, id int) (string, error) {

    stmtIns, err := db.Prepare("delete from item where id = ?") // ? = placeholderl
    if err != nil {
        return "", err
    }
    defer stmtIns.Close() // Close the statement when we leave main() / the program terminates
    
    _, err = stmtIns.Exec(id) // Insert tuples (i, i^2)
    if err != nil {
        return "", err
    }

    response := Response{"OK",make(map[string]string)}
    js, err := json.Marshal(response)
    if err != nil {
    return "", err
    }
    return string(js), err  
}