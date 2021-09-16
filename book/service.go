package book

import (
	"fmt"
	"golang-graphql-go/infrastructure"
	"net/http"
)

/* Rest API */
func RestApiGetBook(w http.ResponseWriter, r *http.Request) {
	//var book Book
	var books []Book
	db, err := infrastructure.InitDB()
	if err == nil {
		db.Find(&books)
	}

	infrastructure.DisconectionDB(db)
	fmt.Println(books)
	//return
}
