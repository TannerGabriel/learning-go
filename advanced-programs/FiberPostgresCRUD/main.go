package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber"
	"github.com/tannergabriel/learning-go/advanced-programs/FiberPostgresCRUD/database"
	"github.com/tannergabriel/learning-go/advanced-programs/FiberPostgresCRUD/item"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)

	app.Get("/api/v1/item", item.GetItems)
	app.Get("/api/v1/item/:id", item.GetItem)
	app.Post("/api/v1/item", item.NewItem)
	app.Delete("/api/v1/item/:id", item.DeleteItem)
}

func initDatabase() {
	var connectionString string

	// Get environment variable for connection string
	host := os.Getenv("HOST")
	if "" == host {
		connectionString = "host=localhost port=5432 user=postgres dbname=pq-demo password=example sslmode=disable"
	} else {
		connectionString = "host=" + host + " port=5432 user=postgres dbname=pq-demo password=example sslmode=disable"
	}

	var err error
	database.DBConn, err = gorm.Open("postgres", connectionString)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
}

func createTable() {
	db := database.DBConn

	// Delete table if you changed/updated the schema
	// deleteQuery := `DROP TABLE IF EXISTS items`
	// db.Exec(deleteQuery)

	query := `
	CREATE TABLE IF NOT EXISTS items
	(
	 id serial NOT NULL,
	 Title character varying NOT NULL,
	 Owner character varying,
	 Rating integer,
	 created_at date,
	 updated_at date,
	 deleted_at date,
	 CONSTRAINT pk_books PRIMARY KEY (id )
	)
	WITH (
	 OIDS=FALSE
	);
	ALTER TABLE items
	 OWNER TO postgres;`

	db.Exec(query)
}

func main() {
	app := Setup()
	app.Listen(3000)

	defer database.DBConn.Close()
}

// Setup fiber app and database
func Setup() *fiber.App {
	app := fiber.New()
	initDatabase()
	createTable()

	setupRoutes(app)

	return app
}
