// Author: Gary McHugh
// Date: September 20th, 2016
// Adapted from: https://github.com/emerging-technologies/example-webapp

package main

import "gopkg.in/macaron.v1"

func main() {
	//use macaron framework
	m := macaron.Classic()
	m.Use(macaron.Renderer())

	// when root loads, load the following html pages
	m.Get("/", func(tx *macaron.Context) {
		tx.HTML(200, "head")
		tx.HTML(200, "nav")
		tx.HTML(200, "index")
		tx.HTML(200, "footer")
	})

	//only thing that changes is the content
	m.Get("/bookaroom", func(tx *macaron.Context) {
		tx.HTML(200, "head")
		tx.HTML(200, "nav")
		tx.HTML(200, "bookaroom")
		tx.HTML(200, "footer")
	})

	m.Get("/openingHours", func(tx *macaron.Context) {
		tx.HTML(200, "head")
		tx.HTML(200, "nav")
		tx.HTML(200, "openingHours")
		tx.HTML(200, "footer")
	})

	m.Get("/aboutUs", func(tx *macaron.Context) {
		tx.HTML(200, "head")
		tx.HTML(200, "nav")
		tx.HTML(200, "aboutUs")
		tx.HTML(200, "footer")
	})

	m.Get("/contactUs", func(tx *macaron.Context) {
		tx.HTML(200, "head")
		tx.HTML(200, "nav")
		tx.HTML(200, "contactUs")
		tx.HTML(200, "footer")
	})

	//used or the individual books
	m.Get("/booksA", func(tx *macaron.Context) {
		tx.HTML(200, "head")
		tx.HTML(200, "nav")
		tx.HTML(200, "booksA")
		tx.HTML(200, "footer")
	})

	m.Get("/booksC", func(tx *macaron.Context) {
		tx.HTML(200, "head")
		tx.HTML(200, "nav")
		tx.HTML(200, "booksC")
		tx.HTML(200, "footer")
	})

	m.Get("/BooksE", func(tx *macaron.Context) {
		tx.HTML(200, "head")
		tx.HTML(200, "nav")
		tx.HTML(200, "BooksE")
		tx.HTML(200, "footer")
	})

	m.Get("/BooksF", func(tx *macaron.Context) {
		tx.HTML(200, "head")
		tx.HTML(200, "nav")
		tx.HTML(200, "BooksF")
		tx.HTML(200, "footer")
	})

	m.Get("/BooksM", func(tx *macaron.Context) {
		tx.HTML(200, "head")
		tx.HTML(200, "nav")
		tx.HTML(200, "BooksM")
		tx.HTML(200, "footer")
	})

	m.Get("/BooksPhysics", func(tx *macaron.Context) {
		tx.HTML(200, "head")
		tx.HTML(200, "nav")
		tx.HTML(200, "BooksPhysics")
		tx.HTML(200, "footer")
	})

	m.Get("/BooksPsycol", func(tx *macaron.Context) {
		tx.HTML(200, "head")
		tx.HTML(200, "nav")
		tx.HTML(200, "BooksPsycol")
		tx.HTML(200, "footer")
	})

	//a form to add a book to the database
	m.Get("/addBooks", func(tx *macaron.Context) {
		tx.HTML(200, "head")
		tx.HTML(200, "nav")
		tx.HTML(200, "addBooks")
		tx.HTML(200, "footer")
	})

	//load the sign in page
	m.Get("/signIn", func(tx *macaron.Context) {
		tx.HTML(200, "signIn")
	})

	m.Run(4001)
}
