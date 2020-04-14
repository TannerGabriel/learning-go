package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/godb?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	dropTable(db)

	createTable(db)

	username := "johndoe"
	password := "secret"

	insertUser(db, username, password)

	queryUser(db, 1)

	queryAllUser(db)

	deleteUser(db, 1)
}

func dropTable(db *sql.DB) {
	if _, err := db.Exec("DROP TABLE IF EXISTS `users`"); err != nil {
		log.Fatal(err)
	}
}

func createTable(db *sql.DB) {
	query := `
		CREATE TABLE users (
			id INT AUTO_INCREMENT,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			created_at DATETIME,
			PRIMARY KEY (id)
		);`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

func insertUser(db *sql.DB, username string, password string) {
	createdAt := time.Now()

	result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	fmt.Println(id)
}

func queryUser(db *sql.DB, userID int) {
	var (
		id        int
		username  string
		password  string
		createdAt time.Time
	)

	query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
	if err := db.QueryRow(query, userID).Scan(&id, &username, &password, &createdAt); err != nil {
		log.Fatal(err)
	}

	fmt.Println(id, username, password, createdAt)
}

func queryAllUser(db *sql.DB) {
	type user struct {
		id        int
		username  string
		password  string
		createdAt time.Time
	}

	rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var u user

		err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", users)
}

func deleteUser(db *sql.DB, id int) {
	_, err := db.Exec(`DELETE FROM users WHERE id = ?`, id)
	if err != nil {
		log.Fatal(err)
	}
}
