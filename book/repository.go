package book

import (
	"context"
	"golang-graphql-go/infrastructure"
)

func GetBookByName(name string) (result interface{}) {
	var book Book
	db, err := infrastructure.InitDB()
	if err == nil {
		db.Where("Name = ?", name).Find(&book)
	}
	//data.Decode(&book)
	infrastructure.DisconectionDB(db)
	return book
}

func GetBookList(ctx context.Context, limit int) (result interface{}) {
	//var book Book
	var books []Book
	db, err := infrastructure.InitDB()
	if err == nil {
		db.Find(&books)
	}

	infrastructure.DisconectionDB(db)

	return books
}
