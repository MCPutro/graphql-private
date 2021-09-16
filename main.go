package main

import (
	"golang-graphql-go/book"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/graphql-go/handler"
)

func main() {
	//initialisasi table
	// db, _ := infrastructure.InitDB()
	// var tablebuku book.Book
	// db.AutoMigrate(&tablebuku)
	// infrastructure.DisconectionDB(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	route := chi.NewRouter()

	/** GraphQL **/
	graphQL := handler.New(&handler.Config{
		Schema:   &book.Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	//route.Use(middleware.Logger)
	route.Handle("/query", graphQL)
	route.Get("/get/book/{bookname}", book.RestApiGetBook)

	log.Println("Server ready at 8080")
	log.Fatal(http.ListenAndServe(":"+port, route))
}
