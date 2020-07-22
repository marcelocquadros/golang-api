package book

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/marcelocquadros/database"
)

type Book struct {
	*gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating string `json:"rating"`
}

func GetBooks(c *fiber.Ctx) {
	var books []Book
	database.DBConn.Find(&books)
	c.JSON(books)
}

func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	var book Book
	if err := database.DBConn.First(&book, id).Error; err != nil {
		c.Status(404).Send("Product not found with given id")
		return
	}

	c.JSON(book)
}

func NewBook(c *fiber.Ctx) {
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(400).Send(err)
		return
	}

	database.DBConn.Create(&book)

	c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	var book Book
	if err := database.DBConn.First(&book, id).Error; err != nil {
		c.Status(404).Send("Product not found with given id")
		return
	}
	database.DBConn.Delete(&book)
}
