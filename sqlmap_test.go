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
	Date        string `sqlmap:"Date"`
	Year        int    `sqlmap:"Year"`
}

func TestSqlMap(t *testing.T) {
	db, err := sql.Open("mysql", "root:123456@/setsDb")
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select Id,CompanyName,Date,Year from SetYear where Id = 35 ")
	if err != nil {
		t.Error(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var sy SetYear
		Scan(rows, &sy)
		//str := FlatStructToStr(&sy)
		fmt.Printf("%#v\n", sy)
	}
}

/*
func TestNormalSql(t *testing.T) {
	db, err := sql.Open("mysql", "root:123456@/setsDb")
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()
	rows, err := db.Query("select Id,CompanyName,Date,Year from SetYear where Id = 35 ")
	if err != nil {
		t.Error(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		datas := make([]interface{}, 4)
		err = rows.Scan(datas...)
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Printf("->%d\n", datas[1].(int))
	}
}*/
