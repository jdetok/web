package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

func ConvertTypes(colTypes []*sql.ColumnType) ([]any, error){
	count := len(colTypes)
	scanArgs := make([]any, count)

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
	return scanArgs, nil
}

func AssignName(arg any, data map[string]any, colType *sql.ColumnType) (any, error) {
	if z, ok := arg.(*sql.NullBool); ok {
		if z.Valid {
			data[colType.Name()] = z.Bool
		}
	} else if z, ok := arg.(*sql.NullString); ok {
		if z.Valid {
			data[colType.Name()] = z.String
		}
	} else if z, ok := arg.(*sql.NullInt64); ok {
		if z.Valid {
			data[colType.Name()] = z.Int64
		}
	} else if z, ok := arg.(*sql.NullFloat64); ok {
		if z.Valid {
			data[colType.Name()] = z.Float64
		}
	} else if data[colType.Name()] == nil {
		data[colType.Name()] = nil
	}
	return data[colType.Name()], nil
}

// 	TODO - FIGURE OUT HOW TO RETAIN THE ORDER OF FIELDS WHEN CONVERTED TO JSON
func RowsToJSON(rows *sql.Rows, indent bool) ([]byte, error) {
	colTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}
	
	finalRows := []any{};

	for rows.Next() {
		scanArgs, err := ConvertTypes(colTypes)
		if err != nil {
			fmt.Println("Error in convert types function")
		}

		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}

		masterData := map[string]any{}
		for i, v := range colTypes {
			masterData[v.Name()], err = AssignName(scanArgs[i], masterData, v)
			if err != nil {
				return nil, err
			}
		}
		finalRows = append(finalRows, masterData)
	}
// indented json if indent == true
	if indent {
		js, err :=  json.MarshalIndent(finalRows, "", "  ")	
		if err != nil {
			return nil, err
		}
		return js, nil
	}
	
// unindented json if indent == false
	js, err :=  json.Marshal(finalRows)	
	if err != nil {
		return nil, err
	}
	return js, nil
}
