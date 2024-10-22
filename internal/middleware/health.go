package middleware

import (
    "encoding/json"
    "github.com/d3code/zlog"
    "net/http"
)

type Health struct {
    Status     string `json:"status"`
    Connection string `json:"connection"`
}

func ServerHealth() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        w.Header().Set("Content-Type", "application/json")
        resp, err := json.Marshal(Health{Status: "ok"})
        if err != nil {
            zlog.Log.Errorf("Failed to marshal response: %v", err)
            return
        }

        _, err = w.Write(resp)
        if err != nil {
            zlog.Log.Errorf("Failed to write response: %v", err)
            return
        }
    }
}
