package util

import (
    "database/sql"
    "github.com/d3code/zlog"
)

func CloseRows(rows *sql.Rows) {
    if rows == nil {
        return
    }

    err := rows.Close()
    if err != nil {
        zlog.Log.Error(err)
    }
}

func ClosePrepare(prepare *sql.Stmt) {
    prepareError := prepare.Close()
    if prepareError != nil {
        zlog.Log.Error(prepareError)
    }
}
