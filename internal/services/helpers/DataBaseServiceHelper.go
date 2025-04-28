package helpers

import (
	"database/sql"
	"fmt"
	"log"
)

func displayRows(rows *sql.Rows) {
	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))

	for rows.Next() {
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			log.Fatal(err)
		}

		for i, col := range columns {
			val := values[i]
			var v interface{}
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			fmt.Printf("%s: %v\n", col, v)
		}
	}
}
