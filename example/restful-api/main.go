package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type book struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	books map[int]*book
	seq   = 1
)

func main() {
	books = make(map[int]*book)
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())

	// Routes
	e.POST("/book", createBook)
	e.GET("/books", listBook)
	e.GET("/book/:id", getBook)
	e.PUT("/book/:id", updateBook)
	e.DELETE("/book/:id", deleteBook)

	// Start server
	e.Start(":8080")
}

//----------
// Handlers
//----------
func createBook(c echo.Context) error {
	b := &book{}
	c.Bind(b)
	b.ID = seq
	books[b.ID] = b
	seq++
	return c.JSON(http.StatusCreated, b)
}

func listBook(c echo.Context) error {
	return c.JSON(http.StatusOK, books)
}

func getBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, books[id])
}

func updateBook(c echo.Context) error {
	b := new(book)
	c.Bind(b)
	id, _ := strconv.Atoi(c.Param("id"))
	books[id].Name = b.Name
	return c.JSON(http.StatusOK, books[id])
}

func deleteBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(books, id)
	return c.NoContent(http.StatusNoContent)
}
