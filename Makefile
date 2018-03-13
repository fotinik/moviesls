build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/movies movies/main.go
#	env GOOS=linux go build -ldflags="-s -w" -o bin/services services/movie_service.go