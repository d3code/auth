package util

import "strings"

func GetPlaceholders(count int, positions ...int) string {
    if count <= 0 {
        return "()"
    }

    // Convert positions to a map for O(1) lookups
    positionMap := make(map[int]bool)
    for _, pos := range positions {
        positionMap[pos] = true
    }

    placeholders := make([]string, count)
    for i := 0; i < count; i++ {
        if positionMap[i] {
            placeholders[i] = "UUID_TO_BIN(?)"
        } else {
            placeholders[i] = "?"
        }
    }

    return "(" + strings.Join(placeholders, ",") + ")"
}
