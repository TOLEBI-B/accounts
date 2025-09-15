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
        log.Fatalf("failed to connect db pgbd: %v", err)
    }

    // существующий роут
    http.HandleFunc("/accounts", handlers.GetAccounts(store))

    // health-check
    http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("ok"))
    })

    // корневой маршрут
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("accounts service is running"))
    })

    fmt.Println("Server started at :9000")
    log.Fatal(http.ListenAndServe(":9000", nil))
}
