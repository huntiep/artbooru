package db

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect(path string) error {
    db, err := sql.Open("sqlite3", path)
    if err != nil {
        return err
    }
    DB = db
    return nil
}
