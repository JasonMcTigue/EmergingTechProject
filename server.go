// server.go
//
// REST APIs with Go and MySql.
//
// Usage:
//
//   # run go server in the background
//   $ go run server.go

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Library struct {
	ISBN      string
	Title     string
	Author    string
	Category  string
	Copies    int
	LoanLimit int
}

//Handle all requests
func Handler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "text/html")
	webpage, err := ioutil.ReadFile("index.html")
	if err != nil {
		http.Error(response, fmt.Sprintf("home.html file error %v", err), 500)
	}
	fmt.Fprint(response, string(webpage))
}

// Respond to URLs of the form /generic/...
func APIHandler(response http.ResponseWriter, request *http.Request) {

	//Connect to database
	db, e := sql.Open("mysql", "root:hello@tcp(localhost:3306)/Library")
	if e != nil {
		fmt.Print(e)
	}

	//set mime type to JSON
	response.Header().Set("Content-type", "application/json")

	err := request.ParseForm()
	if err != nil {
		http.Error(response, fmt.Sprintf("error parsing url %v", err), 500)
	}

	//can't define dynamic slice in golang
	var result = make([]string, 1000)

	switch request.Method {
	case "GET":
		st, err := db.Prepare("select * from Books")
		if err != nil {
			fmt.Print(err)
		}
		rows, err := st.Query()
		if err != nil {
			fmt.Print(err)
		}
		i := 0
		for rows.Next() {
			var isbn string
			var title string
			var author string
			var category string
			var copies int
			var loanlimit int
			err = rows.Scan(&isbn, &title, &author, &category, &copies, &loanlimit)
			library := &Library{ISBN: isbn, Title: title, Author: author, Category: category, Copies: copies, LoanLimit: loanlimit}
			b, err := json.Marshal(library)
			if err != nil {
				fmt.Println(err)
				return
			}
			result[i] = fmt.Sprintf("%s", string(b))
			i++
		}
		result = result[:i]

	default:
	}

	json, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send the text diagnostics to the client.
	fmt.Fprintf(response, "%v", string(json))
	//fmt.Fprintf(response, " request.URL.Path   '%v'\n", request.Method)
	db.Close()
}

func main() {
	port := 1234
	var err string
	portstring := strconv.Itoa(port)

	mux := http.NewServeMux()
	mux.Handle("/api/", http.HandlerFunc(APIHandler))
	mux.Handle("/", http.HandlerFunc(Handler))

	// Start listing on a given port with these routes on this server.
	log.Print("Listening on port " + portstring + " ... ")
	errs := http.ListenAndServe(":"+portstring, mux)
	if errs != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}
