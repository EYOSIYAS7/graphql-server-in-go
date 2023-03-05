package main

import (
	"fmt"
	"log"
	"net/http"

	graph "github.com/EYOSIYAS7/gptGraphql/Graph"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Movie is a struct representing a movie

// ConnectDB is a function to connect to Postgres database

func main() {
    // Connect to the database
  
  
    rootQuery := graph.RootQuery
    Mutation := graph.Mutation 
    

    schema, err := graphql.NewSchema(graphql.SchemaConfig{
        Query: rootQuery,
        Mutation: Mutation,
    })

    if err != nil {
        log.Fatal("Error creating schema: ", err)
    }

    // Define the GraphQL handler
    graphqlHandler := handler.New(&handler.Config{
        Schema:   &schema,
        Pretty:   true,
        GraphiQL: true,
    })

    // Serve the GraphQL API
    http.Handle("/graphql", graphqlHandler)
    fmt.Println("Listening on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
