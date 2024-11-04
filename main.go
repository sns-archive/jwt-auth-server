package main

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	id       string
	name     string
	email    string
	password string
}

func main() {
	helloStruct := User{
		name: "うんち💩",
	}
	// NOTE: +vでvalueだけでなく、keyも表示できる
	fmt.Printf("%+v\n", helloStruct)

	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Addr:                 fmt.Sprintf("%s:%d", "127.0.0.1", 3307), // 127.0.01:3306
		DBName:               "sns_archive_jwt",
		ParseTime:            true,
		Net:                  "tcp",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		slog.Error(err.Error())
		return
	}
	// 接続確認するため
	if err := db.Ping(); err != nil {
		slog.Error(err.Error())
		return
	}
	xdb := sqlx.NewDb(db, "mysql")
	defer db.Close()

	sql := `INSERT INTO users (id, username, email, password) VALUES (?, ?, ?, ?)`

	id := "123e4567-e89b-12d3-a456-426614174000" // 例としてUUIDを使用
	name := "うんち💩"
	email := "example@example.com"
	password := "securepassword"

	result, err := xdb.Exec(sql, id, name, email, password)
	fmt.Printf("%+v\n", result)
	if err != nil {
		slog.Error(err.Error())
		return
	}
}

// func insertUsers() {
// }
