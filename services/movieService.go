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
	var results []movie
	err := table.Scan().All(&results)

	if err != nil {
		fmt.Println("Error retrieving movies: ", err)
	}

	jsonResult, err := json.Marshal(results)
	return err, string(jsonResult)
}

func StoreMovie(request string) error {
	table := getMoviesTable()
	var m movie
	err := json.Unmarshal([]byte(request), &m)

	if err != nil {
		fmt.Println("Error parsing movie: ", err)
	}

	fmt.Println("Storing new movie: ", m)
	err = table.Put(m).Run()
	return err
}

func getMoviesTable() dynamo.Table {
	s, err := session.NewSession(&aws.Config{})
	if err != nil {
		fmt.Println("Error dynamo session: ", err)
	}

	db := dynamo.New(s, &aws.Config{})
	table := db.Table("Movies")
	return table
}

type movie struct {
	MovieId string `dynamo:"movieId"`
	Name string `dynamo:"name"`
}
