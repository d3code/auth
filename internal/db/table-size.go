package db

import (
    _ "embed"
    "github.com/d3code/auth/internal/config"
    "github.com/d3code/auth/internal/model"
    "github.com/d3code/auth/internal/util"
    "github.com/d3code/zlog"
)

//go:embed sql/table/table-size.sql
var SqlTableSizes string

func SelectTableSizes() ([]model.TableSize, error) {
    d := config.DatabaseConnection()

    rows, err := d.Query(SqlTableSizes)
    defer util.CloseRows(rows)
    if err != nil {
        zlog.Log.Error(err)
        return nil, err
    }

    var tableSizes []model.TableSize
    for rows.Next() {
        var tableSize model.TableSize
        scanError := rows.Scan(&tableSize.Database, &tableSize.Size)
        if scanError != nil {
            zlog.Log.Error(scanError)
            return nil, scanError
        }
        tableSizes = append(tableSizes, tableSize)
    }

    return tableSizes, nil
}
