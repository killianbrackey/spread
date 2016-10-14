package main

import (
	"github.com/killianbrackey/bookstore/models"
	"fmt"
	"net/http"
)

func main() {
	models.InitDB("mysql://sezzle:Testing123!@#@127.0.0/bookstore")

	http.HandleFunc("/books", booksIndex)
	http.ListenAndServe(":3000", nil)

}

func booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	bks, err := models.AllBooks()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, $s,.2f]n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}
