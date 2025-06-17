package db

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
)

type OrderedRow []OrderedField 

type OrderedField struct {
	Key string
	Value any
}


// writes bytes directly to ensure the order of the json objects is the same as the select order
// was originally using a map but it reordered the columns
func (row OrderedRow) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('{')

	// m := make(map[string]any, len(row))
	for i, field := range row {
		keyJSON, err := json.Marshal(field.Key)
		if err != nil {
			return nil, errors.New("json marshal error: key")
		}
		valJSON, err := json.Marshal(field.Value)
		if err != nil {
			return nil, err
		}
// separate objects with if not first i
		if i > 0 {
			buf.WriteByte(',')
		}
// write the json lines to 
		fmt.Fprintf(&buf, "%s:%s", keyJSON, valJSON)
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

// called in RowsToJSON
func convertTypes(colTypes []*sql.ColumnType) ([]any, error){
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

// called in RowsToJSON
func getDBType(arg any, data map[string]any, col *sql.ColumnType) (any, error) {
	if z, ok := arg.(*sql.NullBool); ok {
		if z.Valid {
			data[col.Name()] = z.Bool
		}
	} else if z, ok := arg.(*sql.NullString); ok {
		if z.Valid {
			data[col.Name()] = z.String
		}
	} else if z, ok := arg.(*sql.NullInt64); ok {
		if z.Valid {
			data[col.Name()] = z.Int64
		}
	} else if z, ok := arg.(*sql.NullFloat64); ok {
		if z.Valid {
			data[col.Name()] = z.Float64
		}
	} else if data[col.Name()] == nil {
		data[col.Name()] = nil
	}
	return data[col.Name()], nil
}

func RowsToJSON(rows *sql.Rows, indent bool) ([]byte, error) {
	colTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}
	
	finalRows := []any{};

	for rows.Next() {
		scanArgs, err := convertTypes(colTypes)
		if err != nil {
			fmt.Println("Error in convert types function")
		}

		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}

		// masterData := map[string]any{}
		var rowOrdered OrderedRow
		for i, v := range colTypes {
			val, err := getDBType(scanArgs[i], map[string]any{}, v)
			if err != nil {
				return nil, err
			}

			rowOrdered = append(rowOrdered,  OrderedField{
				Key: v.Name(),
				Value: val,
			})
		}

		finalRows = append(finalRows, rowOrdered)
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