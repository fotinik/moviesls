package main

import (
    "fmt"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    MovieService "moviesls/services"
)


func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received request: ", request.HTTPMethod)

	if request.HTTPMethod == "POST" {
		return createMovie(request)
	}

	return allMovies()
}

func allMovies() (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received retrieve all movies request")
	err, result := MovieService.RetrieveAllMovies()
	return events.APIGatewayProxyResponse{Body: result, StatusCode: 200}, err
}


func createMovie(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received create movie request: ", request)
	body := request.Body
	err := MovieService.StoreMovie(body)
	return events.APIGatewayProxyResponse{Body: body, StatusCode: 200}, err
}

func main() {
	lambda.Start(Handler)
}
