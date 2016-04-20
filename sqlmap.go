package sqlmap

import (
	"database/sql"
	"reflect"
)

//Scan scan rows to struct
func Scan(rows *sql.Rows, obj interface{}) error {

	cols, err := rows.Columns()
	if err != nil {
		return err
	}
	elem := reflect.ValueOf(obj).Elem()
	values := make([]interface{}, len(cols))
	num := elem.NumField()
	for c, col := range cols {
		i := 0
		for i < num {
			if elem.Type().Field(i).Tag.Get("sqlmap") == col {
				v := elem.Field(i).Addr().Interface()
				values[c] = v
				break
			}
			i++
		}
	}
	err = rows.Scan(values...)
	if err != nil {
		return err
	}
	return nil
}
