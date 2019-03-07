package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Data struct {
	rid    int
	reward string
}

var slice []Data

//golang操作mysql实现curd基础操作
func main() {
	db, error := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/pu_localhost?charset=utf8")
	if error != nil {
		log.Fatal(error)
	}
	//select
	rows, error := db.Query("SELECT * FROM rewards")
	for rows.Next() {
		var rid int
		var reward string
		error = rows.Scan(&rid, &reward)
		a := Data{rid, reward}
		slice = append(slice, a)
	}

	fmt.Println(slice)
	for _, v := range slice {
		fmt.Println(v)
	}
	//update
	update, error := db.Prepare("UPDATE rewards SET reward =? WHERE rid = ?")
	res, _ := update.Exec("tiger", 5)
	id, error := res.RowsAffected()
	fmt.Println(id)
	//insert
	insert, error := db.Prepare("INSERT rewards SET reward = ?")
	insetRes, _ := insert.Exec("Docker")
	fmt.Println(insetRes)
	insertId, _ := insetRes.LastInsertId()
	fmt.Println(insertId)
	//delete
	delete, error := db.Prepare("DELETE FROM rewards WHERE rid = ?")
	deleteRes, _ := delete.Exec(12)
	fmt.Println(deleteRes)
}
