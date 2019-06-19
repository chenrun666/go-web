package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "chenrun"
	password = "123456"
	dbname   = "chitchat"
)

func connectDB() *sql.DB {
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, host, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connect Success !!!")

	return db
}

func query(db *sql.DB) {
	var id, name, email, password string

	rows, err := db.Query("select id, name, email, password from users where id = $1;", 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name, &email, &password)

		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(id, name, email, password)

}

func insertUser(db *sql.DB) {
	stmt, err := db.Prepare("insert into users(uuid, name, email, password) values($1, $2, $3, $4);")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = stmt.Exec(2345000999, "penghuixian", "17610780919@163.com", "lloveyou")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("insert into users success !!!")

}

func updateUser(db *sql.DB) {
	stmt, err := db.Prepare("update users set name=$1 where id=$2")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = stmt.Exec("shuaidong", 9)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("update user success !!!")
}

func delectUser(db *sql.DB) {
	stmt, err := db.Prepare("delete from users where id=$1")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = stmt.Exec(2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("delete form user success !!!")
}

func main() {

	db := connectDB()
	defer db.Close()

	// query(db)
	// insertUser(db)
	// updateUser(db)
	delectUser(db)
}
