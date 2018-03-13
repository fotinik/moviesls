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
	body, err := json.Marshal(results)
	result := string(body)
	return err, result
}

func StoreMovie(body string) error {
	fmt.Println("Received body: ", body)
	db := dynamo.New(session.New(), &aws.Config{})
	table := db.Table("Movies")
	var m Movie
	err := json.Unmarshal([]byte(body), &m)
	fmt.Println("Json err: ", err)
	fmt.Println("Json: ", m)
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
