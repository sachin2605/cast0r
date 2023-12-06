# cast0r

This is a test repository.

### Description

Simple REST API perfom CRUD operation of Fruits. 


### Required
1. Gin
2. MongoDB
2. Gin-Swagger


## Local Development
 Create a local mongo container on host network. You can use the following command
 ```
 docker run -d -p 27017:27017 --name example-mongo mongo:latest
 ```
Run using go run ./main.go  

Update MongoURI in .env file to change mongoURI path.

## Using Docker-compose
```
docker compose up -d
```

## Swagger
 visit Swagger UI on localhost at http://localhost:8080/swagger/index.html 

 Generate using swagget . Full docs at [Gin-Swagger GitHub](https://github.com/swaggo/gin-swagger)