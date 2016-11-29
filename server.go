// Author: Gary McHugh
// Date: November 28th, 2016
// Adapted from: http://blog.narenarya.in/build-rest-api-go-mysql.html

package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//use your username and password for the database here
	//db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/Library")

	//my username was: root
	//my password was: hello
	db, err := sql.Open("mysql", "root:hello@tcp(127.0.0.1:3306)/Library")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	//error handling
	if err != nil {
		fmt.Print(err.Error())
	}
	//declare a struct with the database variables
	type Books struct {
		ISBN      string
		Title     string
		Author    string
		Category  string
		Copies    int
		LoanLimit int
	}
	router := gin.Default()

	// GET a books details
	router.GET("/book/:isbn", func(c *gin.Context) {
		var (
			books  Books
			result gin.H
		)
		id := c.Param("isbn")
		//get a single row from the database using a sql statement
		row := db.QueryRow("select ISBN, title, author, category, copies, loan_limit from books where ISBN = ?;", id)
		err = row.Scan(&books.ISBN, &books.Title, &books.Author, &books.Category, &books.Copies, &books.LoanLimit)
		if err != nil {
			// If no results send null
			result = gin.H{
				"result": nil,
				"count":  0,
			}
		} else {
			//return the books
			result = gin.H{
				"result": books,
				"count":  1,
			}
		}
		c.JSON(http.StatusOK, result)
	})

	// GET all books
	router.GET("/books", func(c *gin.Context) {
		var (
			books Books
			//declare an array of books to hold all the values
			bookss []Books
		)
		rows, err := db.Query("select ISBN, title, author, category, copies, loan_limit from books;")
		if err != nil {
			fmt.Print(err.Error())
		}
		//loop over the database for each row
		for rows.Next() {
			err = rows.Scan(&books.ISBN, &books.Title, &books.Author, &books.Category, &books.Copies, &books.LoanLimit)
			bookss = append(bookss, books)
			if err != nil {
				fmt.Print(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": bookss,
			"count":  len(bookss),
		})
	})

	// POST new book details
	//this is used for adding a book to the database
	//this was not implemented but does function
	router.POST("/book", func(c *gin.Context) {
		var buffer bytes.Buffer
		ISBN := c.PostForm("ISBN")
		title := c.PostForm("title")
		author := c.PostForm("author")
		category := c.PostForm("category")
		copies := c.PostForm("copies")
		loan_limit := c.PostForm("loan_limit")
		stmt, err := db.Prepare("insert into books (ISBN, title, author, category, copies, loan_limit) values(?,?,?,?,?,?);")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(ISBN, title, author, category, copies, loan_limit)

		if err != nil {
			fmt.Print(err.Error())
		}

		// Fastest way to append strings
		buffer.WriteString(ISBN)
		buffer.WriteString(" ")
		buffer.WriteString(title)
		buffer.WriteString(" ")
		buffer.WriteString(author)
		buffer.WriteString(" ")
		buffer.WriteString(category)
		buffer.WriteString(" ")
		buffer.WriteString(copies)
		buffer.WriteString(" ")
		buffer.WriteString(loan_limit)
		defer stmt.Close()
		name := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf(" %s successfully created", name),
		})
	})

	// PUT - update a books details
	//this is for updating a book in the database
	//this was not implemented but is functional
	router.PUT("/book", func(c *gin.Context) {
		var buffer bytes.Buffer
		ISBN := c.Query("ISBN")
		title := c.PostForm("title")
		author := c.PostForm("author")
		category := c.PostForm("category")
		copies := c.PostForm("copies")
		loan_limit := c.PostForm("loan_limit")
		stmt, err := db.Prepare("update books set ISBN= ?, title= ?, author= ?, category= ?, copies= ?, loan_limit= ? where ISBN= ?;")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(ISBN, title, author, category, copies, loan_limit)
		if err != nil {
			fmt.Print(err.Error())
		}

		// Fastest way to append strings
		buffer.WriteString(ISBN)
		buffer.WriteString(" ")
		buffer.WriteString(title)
		buffer.WriteString(" ")
		buffer.WriteString(author)
		buffer.WriteString(" ")
		buffer.WriteString(category)
		buffer.WriteString(" ")
		buffer.WriteString(copies)
		buffer.WriteString(" ")
		buffer.WriteString(loan_limit)
		defer stmt.Close()
		name := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully updated to %s", name),
		})
	})

	// Delete resources
	//this is to delete a specific book from the database
	//this was not implemented but is fully functional
	router.DELETE("/book", func(c *gin.Context) {
		ISBN := c.Query("ISBN")
		stmt, err := db.Prepare("delete from books where ISBN= ?;")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(ISBN)
		if err != nil {
			fmt.Print(err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully deleted book: %s", ISBN),
		})
	})
	//run on port 5000
	router.Run(":5000")
}
