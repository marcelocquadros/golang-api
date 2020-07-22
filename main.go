package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/marcelocquadros/book"
	"github.com/marcelocquadros/database"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello from golang")
}

func setupRoutes(app *fiber.App) {
	app.Get("/books", book.GetBooks)
	app.Get("/books/:id", book.GetBook)
	app.Post("/books", book.NewBook)
	app.Delete("/books/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed to connect to database ")
	}
	fmt.Println("Database connection successfull")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database migrated")
}

func main() {
	app := fiber.New()
	setupRoutes(app)

	initDatabase()
	defer database.DBConn.Close()
	app.Listen(8080)
}
