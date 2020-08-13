package item

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tannergabriel/learning-go/advanced-programs/FiberPostgresCRUD/database"
)

type Item struct {
	gorm.Model
	Title  string `json:"Title"`
	Owner  string `json:"Owner"`
	Rating int    `json:"Rating"`
}

func GetItems(c *fiber.Ctx) {
	db := database.DBConn
	var items []Item
	db.Find(&items)
	c.JSON(items)
}

func GetItem(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var item Item
	db.Find(&item, id)
	c.JSON(item)
}

func NewItem(c *fiber.Ctx) {
	db := database.DBConn

	item := new(Item)
	if err := c.BodyParser(item); err != nil {
		fmt.Println(err)
		c.Status(503).Send(err)
		return
	}

	db.Create(&item)
	c.JSON(item)
}

func DeleteItem(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var item Item
	db.First(&item, id)
	if item.Title == "" {
		c.Status(500).Send("No item found with given ID")
		return
	}
	db.Delete(&item)
	c.Send("Item successfully deleted")
}
