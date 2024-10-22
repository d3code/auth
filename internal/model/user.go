package model

import "time"

type User struct {
    Id         string    `json:"id"`
    Username   string    `json:"username"`
    Email      string    `json:"email"`
    Password   string    `json:"password"`
    NameFamily string    `json:"name_family"`
    NameGiven  string    `json:"name_given"`
    Scope      string    `json:"scope"`
    Active     bool      `json:"active"`
    Created    time.Time `json:"created"`
}
