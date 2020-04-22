package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "example"
	dbname   = "pq-demo"
)

type Book struct {
	ID              int
	Name            string
	Author          string
	PublicationDate time.Time
	Pages           int
}

// Create the books table in the database
func createTable() {
	query := `
	CREATE TABLE books
	(
	 id serial NOT NULL,
	 name character varying NOT NULL,
	 author character varying,
	 pages integer,
	 publication_date date,
	 CONSTRAINT pk_books PRIMARY KEY (id )
	)
	WITH (
	 OIDS=FALSE
	);
	ALTER TABLE books
	 OWNER TO postgres;`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

// Drop the book table from the database
func dropTable() {
	query := `DROP TABLE books;`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

// Get a specific book using the book id
func getBook(bookID int) (Book, error) {
	res := Book{}

	var id int
	var name string
	var author string
	var pages int
	var publicationDate pq.NullTime

	err := db.QueryRow(`SELECT id, name, author, pages, publication_date FROM books where id = $1`, bookID).Scan(&id, &name, &author, &pages, &publicationDate)
	if err == nil {
		res = Book{ID: id, Name: name, Author: author, Pages: pages, PublicationDate: publicationDate.Time}
	}

	return res, err
}

// Get all books from the database
func getAllBooks() ([]Book, error) {
	books := []Book{}

	rows, err := db.Query(`SELECT id, name, author, pages, publication_date FROM books order by id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var author string
		var pages int
		var publicationDate pq.NullTime

		err = rows.Scan(&id, &name, &author, &pages, &publicationDate)
		if err != nil {
			return books, err
		}

		currentBook := Book{ID: id, Name: name, Author: author, Pages: pages}
		if publicationDate.Valid {
			currentBook.PublicationDate = publicationDate.Time
		}

		books = append(books, currentBook)
	}

	return books, err
}

// Insert a book into the database
func insertBook(name, author string, pages int, publicationDate time.Time) (int, error) {
	var bookID int
	err := db.QueryRow(`INSERT INTO books(name, author, pages, publication_date) VALUES($1, $2, $3, $4) RETURNING id`, name, author, pages, publicationDate).Scan(&bookID)

	if err != nil {
		return 0, err
	}

	fmt.Printf("Last inserted ID: %v\n", bookID)
	return bookID, err
}

// Update a book in the database
func updateBook(id int, name, author string, pages int, publicationDate time.Time) (int, error) {
	res, err := db.Exec(`UPDATE books set name=$1, author=$2, pages=$3, publication_date=$4 where id=$5 RETURNING id`, name, author, pages, publicationDate, id)
	if err != nil {
		return 0, err
	}

	rowsUpdated, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsUpdated), err
}

// Remove a book from the database
func removeBook(bookID int) (int, error) {
	res, err := db.Exec(`delete from books where id = $1`, bookID)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsDeleted), nil
}

var db *sql.DB

func main() {
	// Database connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error

	// Connect to the database
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// Check the database connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	createTable()

	insertBook("Hello World", "Author", 100, time.Now())

	fmt.Println(getAllBooks())

	updateBook(1, "Hello World 2!", "Author 2", 205, time.Now())

	fmt.Println(getBook(1))

	removeBook(1)

	dropTable()
}
