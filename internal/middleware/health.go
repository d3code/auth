package middleware

import (
    "encoding/json"
    "github.com/d3code/auth/internal/db"
    "github.com/d3code/auth/internal/model"
    "github.com/d3code/zlog"
    "net/http"
)

type Health struct {
    Status       string            `json:"status"`
    Connection   string            `json:"connection"`
    DatabaseSize []model.TableSize `json:"database_size"`
}

func ServerHealth() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Cache-Control", "no-cache")

        tables, err := db.SelectTableSizes()
        if err != nil {
            zlog.Log.Errorf("Failed to get table sizes: %v", err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        resp, err := json.Marshal(Health{
            Status:       "ok",
            Connection:   r.Host,
            DatabaseSize: tables,
        })
        if err != nil {
            zlog.Log.Errorf("Failed to marshal response: %v", err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        _, err = w.Write(resp)
        if err != nil {
            zlog.Log.Errorf("Failed to write response: %v", err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
    }
}
