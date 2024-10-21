package main

import (
	"net/http"

	"errors"

	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// DEFINING STRUCT HERE FOR BOOKS
type book struct {
	ID     string `json:"id"`
	TITLE  string `json:"title"`
	AUTHOR string `json:"author"`
	QTY    int    `json:"qty"`
}

var books = []book{
	{ID: "1", TITLE: "Harry Potter and the Sorcerors Apprentice", AUTHOR: "J.K Rowling", QTY: 2},
	{ID: "2", TITLE: "Harry Potter Order of the Phoenix", AUTHOR: "J.K Rowling", QTY: 3},
	{ID: "3", TITLE: "Harry Potter and the Deathly Hallows", AUTHOR: "J.K Rowling", QTY: 5},
}

func getBooks(c *gin.Context) {
	// COVER ALL FUNCTION THAT JUST RETURNS ALL BOOKS IN JSON FORMAT.
	c.IndentedJSON(http.StatusOK, books)
}

func get_book_by_id(id string) (*book, error) {

	// GETS ALL BOOKS
	for i, b := range books {

		// IF BOOK ID IS EQUAL TO ARG ID, RETURN.
		if b.ID == id {
			return &books[i], nil
		}
	}

	// IF ABOVE DOES NOT FIND BOOK, RETURNS BOOK NOT FOUND.
	return nil, errors.New("book not found")
}

func bookByID(c *gin.Context) {
	// GETS ID FROM ARGUMENTS E.G. "?id=2"
	id := c.Param("id")

	// PASSES THAT ID TO EARLIER MADE FUNCTION.
	book, err := get_book_by_id(id)

	// IF WE GET ERROR, STATE BOOK NOT FOUND.
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	// IF NOT ERROR, STATUS IS OKAY AND IT SHOWS THE BOOK IN JSON FORMAT.
	c.IndentedJSON(http.StatusOK, book)
}

func createBook(c *gin.Context) {

	// CREATING NEWBOOK USING BOOK
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// APPENDING NEWBOOK TO BOOKS DEFINED ABOVE.
	books = append(books, newBook)

	// SHOWING STATUS AND NEW BOOK CREATED IN JSON FORMAT.
	c.IndentedJSON(http.StatusCreated, newBook)
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	// IF NOT OKAY, INFORM USER MISSING ID.
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing ID query"})
		return
	}

	book, err := get_book_by_id(id)

	// IF ERROR IS NOT NONE, STATE BOOK IS NOT FOUND.
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	// CHECK BOOK QTY, IF 0 CAN'T BE CHECKED OUT.
	if book.QTY <= 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not available"})
		return
	}
	// PASSED ALL CHECKS ABOVE, REDUCE BOOK QTY BY 1 AND MARK AS CHECKED OUT.
	book.QTY -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	// IF NOT OKAY, INFORM USER MISSING ID.
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing ID query"})
		return
	}

	book, err := get_book_by_id(id)

	// IF ERROR IS NOT NONE, STATE BOOK NOT FOUND.
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	// PASSED ALL ABOVE TESTS, RETURN BOOK ADD 1 QTY AND INFORM USER BOOK IS CHECKED BACK IN.
	book.QTY += 1
	c.IndentedJSON(http.StatusOK, gin.H{"message": "The book: " + book.TITLE + " has been returned."})
	c.IndentedJSON(http.StatusOK, book)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the license key from the environment variable
	licenseKey := os.Getenv("BOOK_API_LICENSE_KEY")
	if licenseKey == "" {
		log.Fatal("License key is not set")
	}

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("book-api"),            // Replace with your application name
		newrelic.ConfigLicense(licenseKey),            // Replace with your license key
		newrelic.ConfigDistributedTracerEnabled(true), // Optional: enables distributed tracing
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("New Relic startup success")
	// DEFINE ROUTER USING GIN
	router := gin.Default()

	router.Use(nrgin.Middleware(app))
	// GET ROUTING
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookByID)

	// POST ROUTING
	router.POST("/books", createBook)

	// PATCH ROUTING
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/checkin", returnBook)

	// RUN ROUTER
	router.Run(":8080")
}
