package graph

import (
	model "github.com/EYOSIYAS7/gptGraphql/Model"
	"github.com/EYOSIYAS7/gptGraphql/dbconnection"
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

// Define the GraphQL schema
var db *gorm.DB
func init()  {
	db = dbconnection.ConnectDB()
}
 
	var movieType = graphql.NewObject(graphql.ObjectConfig{
        Name: "Movie",
        Fields: graphql.Fields{
            "id": &graphql.Field{
                Type: graphql.Int,
            },
            "title": &graphql.Field{
                Type: graphql.String,
            },
            "rating": &graphql.Field{
                Type: graphql.Float,
            },
        },
    })

   var  RootQuery = graphql.NewObject(graphql.ObjectConfig{
        Name: "Query",
        Fields: graphql.Fields{
            "movies": &graphql.Field{
                Type: graphql.NewList(movieType),
                Resolve: func(params graphql.ResolveParams) (interface{}, error) {
                    var movies [] model.Movie
                    db.Find(&movies)
                    return movies, nil
                },
            },
            "movie": &graphql.Field{
                Type: movieType,
                Description: "find a movie",

                Args:graphql.FieldConfigArgument{
                    "Id": &graphql.ArgumentConfig{
                        Type: graphql.NewNonNull(graphql.Int),
                    },

                },
                Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                    // get the id from the args 
                    movie := model.Movie{
                        Id: p.Args["Id"].(int),
                    }
                    db.First(&movie)

                    return movie, nil
                },
            },
        },
    })

   var Mutation = graphql.NewObject((graphql.ObjectConfig{

        Name: "Mutation",
        Fields: graphql.Fields{
            "addMovies": &graphql.Field {

                Type: movieType,
                Description: "Add a new movie",
                Args: graphql.FieldConfigArgument{
                    "Title": &graphql.ArgumentConfig{
                        Type: graphql.NewNonNull(graphql.String),

                    },
                    "Rating": &graphql.ArgumentConfig{
                        Type: graphql.NewNonNull(graphql.Float),

                    },
                },
                Resolve: func(p graphql.ResolveParams) (interface{}, error) {

                   
                    movies := model.Movie {
                        Title: p.Args["Title"].(string),
                        Rating: p.Args["Rating"].(float64),
                    }

                     db.Create(&movies)
 
                    return movies , nil
                },
                   
            },
            "deleteMovie": &graphql.Field{
                Type: movieType,
                Description: " Delete a movie from the database",
                Args: graphql.FieldConfigArgument{
                    "Id" :&graphql.ArgumentConfig{
                        Type: graphql.NewNonNull(graphql.Int),
                    },
                },
                Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                    movie := model.Movie{
                        Id:p.Args["Id"].(int),
                    }

                    db.Delete(&movie)

                    return movie, nil
                },
            },
            "updateMovie": &graphql.Field{
                Type: movieType,
                Description: "Update a movie from the database",
                Args: graphql.FieldConfigArgument{
                    "ID" : &graphql.ArgumentConfig{
                        Type: graphql.NewNonNull(graphql.Int),
                    },
                    "Title" : &graphql.ArgumentConfig{
                        Type: graphql.NewNonNull(graphql.String),
                    },
                    "Rating" : &graphql.ArgumentConfig{
                        Type: graphql.NewNonNull(graphql.Float),
                    },
                   
                }, 
                Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                    movie := model.Movie{
                        Id:p.Args["ID"].(int),
                        Title: p.Args["Title"].(string),
                        Rating: p.Args["Rating"].(float64),
                    }
                    title := movie.Title
                    rating := movie.Rating
                    db.Model(&movie.ID).Update(title,rating)
                  
                    db.Save(movie)
                    return movie, nil
                },
            },
        },


    }))