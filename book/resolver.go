package book

import (
	"context"

	"github.com/graphql-go/graphql"
)

var bookType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Book",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"book": &graphql.Field{
				Type:        bookType,
				Description: "Get book by name",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var result interface{}
					name, ok := p.Args["name"].(string)
					if ok {
						// Find product
						result = GetBookByName(name)
					}
					return result, nil
				},
			},
			"list": &graphql.Field{
				Type:        graphql.NewList(bookType),
				Description: "Get book list",
				Args: graphql.FieldConfigArgument{
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					var result interface{}
					limit, _ := params.Args["limit"].(int)
					result = GetBookList(context.Background(), limit)
					return result, nil
				},
			},
			"findAll": &graphql.Field{
				Type:        graphql.NewList(bookType),
				Description: "get all book",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					//var result interface
					resurt := GetBookList(context.Background(), -1)
					return resurt, nil

				},
			},
		},
	})

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
		//Mutation: mutationType,
	},
)
