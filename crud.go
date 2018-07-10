// Package ini dibuat untuk membungkus pemanggilan terhadap database.
package crud

import (
    "fmt"
    "os"
    "strconv"
    "database/sql"
    "github.com/joho/godotenv"
    _ "github.com/go-sql-driver/mysql"
)

// Struct untuk membatasi DB
type MyDb struct {
    db *sql.DB
}

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
func ConnectDb() (*MyDb, error) {
    db_config := fmt.Sprintf("%s:%s@/%s",os.Getenv("DB_USER"),os.Getenv("DB_PASSWORD"),os.Getenv("DATABASE"))
    db, err := sql.Open("mysql", db_config)
    mydb := &MyDb{db : db}
    return mydb, err
}

func (m *MyDb) CloseDb() (error)    {
    return m.db.Close()
}

// GetItem akan mengembalikan nama dari user
// berdasarkan id yang diberikan. Argumen pertama
// adalah database yang digunakan, dan parameter kedua
// adalah id dari user yang dicari.
func (m *MyDb) GetItem(id int) (*Response, error){
    response := &Response{}
    stmtOut, err := m.db.Prepare("SELECT name FROM item WHERE id = ?")
    if err != nil { 
        return response, err
    }
    defer stmtOut.Close()
    
    var name string

    err = stmtOut.QueryRow(id).Scan(&name)
    if err != nil {
        return response, err
    }
    payload := make(map[string]string)
    payload["name"] = name
    response = &Response{"OK",payload}
    return response, err
}

// CreateItem akan memasukan user baru kedalam database.
// Function ini menerima object database sebagai parameter pertama,
// dan nama user baru dari parameter kedua.
func (m *MyDb) CreateItem(name string) (*Response, error) {
    response := &Response{}
    stmtIns, err := m.db.Prepare("INSERT INTO item (name) VALUES(?)")
    if err != nil {
        return response, err 
    }
    defer stmtIns.Close()
    
    _, err = stmtIns.Exec(name)
    if err != nil {
        return response, err
    }
    stmtOut, err := m.db.Prepare("SELECT LAST_INSERT_ID()")
    if err != nil {
        return response, err
    }
    defer stmtOut.Close()
    
    var id int

    err = stmtOut.QueryRow().Scan(&id)
    if err != nil {
        return response, err
    }

    payload := make(map[string]string)
    payload["id"] = strconv.Itoa(id)
    response = &Response{"OK",payload}

    return response, err
}

// UpdateItem akan mengganti nama user dengan nama baru.
// Function ini menerima 3 (tiga) parameter. Parameter
// pertama adalah object database, parameter kedua adalah id user
// yang ingin diganti namanya, dan parameter ketiga adalah nama baru.
func (m *MyDb) UpdateItem(id int, name string) (*Response, error)    {
    response := &Response{}
    stmtIns, err := m.db.Prepare("update item set name = ? where id = ?")
    if err != nil {
        return response, err
    }
    defer stmtIns.Close()

    _, err = stmtIns.Exec(name,id)
    if err != nil {
        return response, err
    }
    response = &Response{"OK",make(map[string]string)}
    return response, err  
}

// DeleteItem akan mengapus data user dari database berdasarkan id.
// Function ini menerima object database sebagai parameter pertama,
// dan id user yang ingin dihapus sebagai parameter kedua.
func (m *MyDb) DeleteItem(id int) (*Response, error) {
    response := &Response{}
    stmtIns, err := m.db.Prepare("delete from item where id = ?")
    if err != nil {
        return response, err
    }
    defer stmtIns.Close()
    
    _, err = stmtIns.Exec(id)
    if err != nil {
        return response, err
    }

    response = &Response{"OK",make(map[string]string)}
    return response, err  
}