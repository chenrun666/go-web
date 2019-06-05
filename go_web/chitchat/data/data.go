package data

import (
	"database/sql"
	"time"
)

// Db 全局变量的数据连接的句柄
var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "dbname=chitchat sslmode=disable")
	if err != nil {
		panic(err)
	}
	return
}

// Thread 数据库对应的数据结构
type Thread struct {
	ID        int
	UUID      string
	Topic     string
	UserID    int
	CreatedAT time.Time
}
