package db

import (
	"database/sql"
	"encoding/json"
)

func RowsToJSON(rows *sql.Rows) ([]byte, error) {
	colTypes, err := rows.ColumnTypes()

	if err != nil {
		return nil, err
	}

	count := len(colTypes)
	
	finalRows := []interface{}{};

	for rows.Next() {
		scanArgs := make([]interface{}, count)

		for i, v := range colTypes {
			switch v.DatabaseTypeName() {
			case "VARCHAR", "TEXT", "UUID", "TIMESTAMP":
				scanArgs[i] = new(sql.NullString)
			case "BOOL":
				scanArgs[i] = new(sql.NullBool)
			case "INT4", "INT8", "INT":
				scanArgs[i] = new(sql.NullInt64)
			case "FLOAT", "FLOAT8", "FLOAT4":
				scanArgs[i] = new(sql.NullFloat64)
			default:
				scanArgs[i] = new(sql.NullString)
			}
		}

		err := rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}

		masterData := map[string]any{}

		for i, v := range colTypes {
			if z, ok := scanArgs[i].(*sql.NullBool); ok {
				if z.Valid {
					masterData[v.Name()] = z.Bool
				} else {
					masterData[v.Name()] = nil
				}	
				continue
			}

			if z, ok := scanArgs[i].(*sql.NullString); ok {
				if z.Valid {
					masterData[v.Name()] = z.String
				} else {
					masterData[v.Name()] = nil
				}	
				continue
			}

			if z, ok := scanArgs[i].(*sql.NullInt64); ok {
				if z.Valid {
					masterData[v.Name()] = z.Int64
				} else {
					masterData[v.Name()] = nil
				}	
				continue
			}

			if z, ok := scanArgs[i].(*sql.NullFloat64); ok {
				if z.Valid {
					masterData[v.Name()] = z.Float64
				} else {
					masterData[v.Name()] = nil
				}	
				continue
			}

			masterData[v.Name()] = scanArgs[i]
		}

		finalRows = append(finalRows, masterData)
	}

	js, err :=  json.Marshal(finalRows)
	if err != nil {
		return nil, err
	}

	return js, nil
}