# Virtual Experience API

## Run Project

```bash

docker-compose up -d # start mongodb instance
go build -o api
./api
```

## Shut down mongoDB Container

```bash
docker-compose down
```

## Testing 

### Cypress tests 

There are some cypress tests that make API requests to the server. However, after spending sometime writing these tests I realized cypress does a bad job at handling multiple requests in one test and this was not a good solution for testing the API.

To run the cypress tests:
```bash
docker-compose up -d
go build -o api
./api

# In a new terminal window
cd integration_tests/; npm i && $(npm bin)/cypress run; cd -
docker-compose down
```

## Unit Testing
Added a set of unit tests to help me figure out some tricky boolean logic

Run unit tests with:
```bash
go test
```

## Testing API Endpoint with Static Pages

You can test the endpoints with the static html files I have created. I think this is better solution to test if the API is working as expected. However, not ideal because it is manual. If I had more time, I would write cypress tests against these pages, so I can test the stack end to end.

Visit this url after you have the webserver running to test the API: http://0.0.0.0:8080