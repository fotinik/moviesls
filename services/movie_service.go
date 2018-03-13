package services

import (
	"encoding/json"
	"github.com/guregu/dynamo"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"fmt"
)

func RetrieveAllMovies() (error, string) {
	table := getMoviesTable()
	var results []Movie
	err := table.Scan().All(&results)
	jsonResult, err := json.Marshal(results)
	return err, string(jsonResult)
}

func StoreMovie(request string) error {
	table := getMoviesTable()
	var m Movie
	err := json.Unmarshal([]byte(request), &m)
	fmt.Println("Error parsing movie: ", err)
	fmt.Println("Storing new movie: ", m)
	err = table.Put(m).Run()
	return err
}

func getMoviesTable() dynamo.Table {
	db := dynamo.New(session.New(), &aws.Config{})
	table := db.Table("Movies")
	return table
}

type Movie struct {
	MovieId string `dynamo:"movieId"`
	Name string `dynamo:"name"`
}
