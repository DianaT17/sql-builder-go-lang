package sqlbuilder

import (
    "testing"
)

type User struct {
    ID    int    `sql:"INT PRIMARY KEY"`
    Name  string `sql:"VARCHAR(100)"`
    Email string `sql:"VARCHAR(100)"`
}

func TestCreateTable(t *testing.T) {
    qb := QueryBuilder{}
    query := qb.CreateTable("users", User{})
    expected := "CREATE TABLE users (ID INT PRIMARY KEY, Name VARCHAR(100), Email VARCHAR(100));"
    if query != expected {
        t.Errorf("expected %s, got %s", expected, query)
    }
}

func TestSelect(t *testing.T) {
    qb := QueryBuilder{}
    query := qb.Select("users", User{})
    expected := "SELECT ID, Name, Email FROM users;"
    if query != expected {
        t.Errorf("expected %s, got %s", expected, query)
    }
}

func TestInsert(t *testing.T) {
    qb := QueryBuilder{}
    query := qb.Insert("users", User{})
    expected := "INSERT INTO users (ID, Name, Email) VALUES (?, ?, ?);"
    if query != expected {
        t.Errorf("expected %s, got %s", expected, query)
    }
}

func TestUpdate(t *testing.T) {
    qb := QueryBuilder{}
    query := qb.Update("users", User{}, "ID = ?")
    expected := "UPDATE users SET ID = ?, Name = ?, Email = ? WHERE ID = ?;"
    if query != expected {
        t.Errorf("expected %s, got %s", expected, query)
    }
}
