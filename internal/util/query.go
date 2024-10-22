package util

import (
    "strings"
)

func QueryLike(query string) string {
    query = strings.TrimSpace(query)
    if query == "" {
        return "%"
    }
    return "%" + query + "%"
}
