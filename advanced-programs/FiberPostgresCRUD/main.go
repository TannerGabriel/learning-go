package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber"
	"github.com/tannergabriel/advanced-programs/FiberPostgresCRUD/database"
	"github.com/tannergabriel/advanced-programs/FiberPostgresCRUD/item"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
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
	host := goDotEnvVariable("HOST")
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

func goDotEnvVariable(key string) string {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return os.Getenv(key)
}

func createTable() {
	db := database.DBConn

	query := `
	CREATE TABLE IF NOT EXISTS items
	(
	 id serial NOT NULL,
	 Title character varying NOT NULL,
	 Author character varying,
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
	app := fiber.New()
	initDatabase()
	createTable()

	setupRoutes(app)
	app.Listen(3000)

	defer database.DBConn.Close()
}
