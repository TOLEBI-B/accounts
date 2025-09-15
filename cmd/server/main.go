package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "accounts/internal/db"
    "accounts/internal/handlers"
)

func main() {
    dbURL := os.Getenv("DATABASE_URL")
    if dbURL == "" {
        log.Fatal("DATABASE_URL is not set")
    }

    store, err := db.NewStore(dbURL)
    if err != nil {
        log.Fatalf("failed to connect db: %v", err)
    }

    http.HandleFunc("/accounts", handlers.GetAccounts(store))

    fmt.Println("Server started at :9000")
    log.Fatal(http.ListenAndServe(":9000", nil))
}
