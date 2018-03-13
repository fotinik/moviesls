# Movie Serverless Spike

Spike out:
* AWS serverless API
* GoLang
* Dynamo DB

![alt text](https://github.com/fotinik/moviesls/blob/master/moviesls.png)

### Pre-requistes 
* sls installed
* aws credentials

### Build

```
$ make 
```

### Deploy

```
$ sls deploy 
```

### Add movie

```
export BASE_DOMAIN=<your url eg. "https://23235.execute-api.us-east-1.amazonaws.com/dev">
curl -H "Content-Type: application/json" -X POST ${BASE_DOMAIN}/movies -d '{"movieId": "4", "name": "Mad max"}'
```

### Get movies

```
curl -H "Content-Type: application/json" -X GET ${BASE_DOMAIN}/movies
```

### Logs

```
$ sls logs -f movies
```

### Dependencies
```
  $ go get -u github.com/aws/aws-sdk-go/aws
  $ go get -u github.com/aws/aws-sdk-go/aws/session
  $ go get -u github.com/guregu/dynamo
```

