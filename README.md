# cast0r

This is a test repository.

## Local Development
 Create a local mongo container on host network. You can use the following command
 docker run -d -p 27017:27017 --name example-mongo mongo:latest

Run using go run ./main.go --env dev [possible values are dev/test]

## Using Docke-compose