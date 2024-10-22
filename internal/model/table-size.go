package model

type TableSize struct {
    Database string   `json:"database"`
    Size     *float64 `json:"size"`
}
