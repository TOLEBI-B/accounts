package handlers

import (
    "encoding/json"
    "net/http"
    "accounts/internal/db"
)

func GetAccounts(store *db.Store) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        rows, err := store.DB.Query("SELECT id, name, balance FROM accounts")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer rows.Close()

        type Account struct {
            ID      int     `json:"id"`
            Name    string  `json:"name"`
            Balance float64 `json:"balance"`
        }

        var accounts []Account
        for rows.Next() {
            var acc Account
            if err := rows.Scan(&acc.ID, &acc.Name, &acc.Balance); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            accounts = append(accounts, acc)
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(accounts)
    }
}
