package crud

import "testing"
import "gopkg.in/DATA-DOG/go-sqlmock.v1"

func TestGetUser(t *testing.T)  {
    rows := sqlmock.NewRows([]string{"name"}).
    AddRow("item N")
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()
    mock.ExpectPrepare("SELECT (.*) FROM (.*)").ExpectQuery().WithArgs(1).WillReturnRows(rows)
    mydb := &MyDb{db : db}
    res, err := mydb.GetItem(1)
    if err != nil {
        t.Errorf("something goes wrong 500")
    }
    if res != "{\"status\":\"OK\",\"payload\":{\"name\":\"item N\"}}" {
        t.Errorf("wrong json")
    }
    // we make sure that all expectations were met
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expectations: %s", err)
    }
}

func TestCreateUser(t *testing.T)   {
    var err error
    rows := sqlmock.NewRows([]string{"id"}).
    AddRow("1")

    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()
    mock.ExpectPrepare("INSERT (.*)").ExpectExec().WillReturnResult(sqlmock.NewResult(1, 1))
    mock.ExpectPrepare("SELECT (.*)").ExpectQuery().WillReturnRows(rows)
    // fmt.Println(CreateItem(db,"item N"))
    mydb := &MyDb{db : db}
    res, err := mydb.CreateItem("item N")
    if err != nil {
        t.Errorf("something goes wrong 500")
    }
    if res != "{\"status\":\"OK\",\"payload\":{\"id\":\"1\"}}" {
        t.Errorf("wrong json")
    }
    // we make sure that all expectations were met
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expectations: %s", err)
    }
}

func TestUpdateUser(t *testing.T)   {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()
    mock.ExpectPrepare("update (.*)").ExpectExec().WithArgs("item N new!", 1).WillReturnResult(sqlmock.NewResult(1, 1))
    // fmt.Println(CreateItem(db,"item N"))
    mydb := &MyDb{db : db}
    res, err := mydb.UpdateItem(1, "item N new!")
    if err != nil  {
        t.Errorf("something goes wrong 500")
    }
    if res != "{\"status\":\"OK\"}" {
        t.Errorf("wrong json")
    }
    // we make sure that all expectations were met
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expectations: %s", err)
    }
}

func TestDeleteUser(t *testing.T)   {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()
    mock.ExpectPrepare("delete (.*)").ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
    // fmt.Println(CreateItem(db,"item N"))
    mydb := &MyDb{db : db}
    res, err := mydb.DeleteItem(1)
    if err != nil {
        t.Errorf("something goes wrong")
    }
    if res != "{\"status\":\"OK\"}" {
        t.Errorf("wrong json")
    }
    // we make sure that all expectations were met
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expectations: %s", err)
    }
} 