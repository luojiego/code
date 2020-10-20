package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"reflect"
)

func main() {
	dns := "root:123456@tcp(192.168.196.16:3307)/game_db"

	db, err := sqlx.Connect("mysql", dns)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	count := 0
	err = db.Get(&count, "SELECT COUNT(*) FROM account_data")
	if err != nil {
		panic(err)
	}

	var ids []int
	test(ids)
	err = db.Select(&ids, "SELECT user_id FROM account_data LIMIT 10")
	if err != nil {
		panic("ids: " + err.Error())
	}
	fmt.Println(ids)
}

func test(v interface{}) {
	fmt.Println(reflect.TypeOf(v))
}
