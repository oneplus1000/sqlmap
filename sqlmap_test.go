package sqlmap

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

type SetYear struct {
	ID          int    `sqlmap:"Id"`
	CompanyName string `sqlmap:"CompanyName"`
	K           string
}

func TestSqlMap(t *testing.T) {
	db, err := sql.Open("mysql", "root:123456@/setsDb")
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()
	rows, err := db.Query("select Id,CompanyName from SetYear")
	if err != nil {
		t.Error(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var sy SetYear
		Scan(rows, &sy)
		fmt.Printf("%+v\n", sy)
	}
}
