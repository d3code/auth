package model

import "time"

type User struct {
    Id       string    `json:"id"`
    Username string    `json:"username"`
    Password string    `json:"password"`
    Scope    string    `json:"scope"`
    Active   bool      `json:"active"`
    Created  time.Time `json:"created"`
}
