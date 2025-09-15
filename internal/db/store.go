package db

import (
    "database/sql"
    _ "github.com/lib/pq"
)

type Store struct {
    DB *sql.DB
}

func NewStore(connStr string) (*Store, error) {
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }
    return &Store{DB: db}, nil
}
