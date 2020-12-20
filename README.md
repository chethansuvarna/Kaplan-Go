### prerequisites
1. Go must be installed in your system refer this link to install https://golang.org/doc/install
2. Docker must be installed in your system refer this link to install https://docs.docker.com/engine/install/

### Setup
1. Download the repo
2. Run `go build` to install dependencies
3. Run `go test ./...` to run unit tests
5. Run `go run main.go` to start the server

### To Build Docker Image
1. Run `docker build -t <Image Name> .` 
2. Run `docker run -it -p 8085:8080  <Image Name>:latest`

### To Generate Mock 
`mockgen -destination=population/mock/service.go -package=mocks Kaplan-Go/population Service`

### Pros
1. As I have used GoRoutines the API calls to third party API are concurrent
2. To manage huge number of API calls we can use Load Balancers
3. If the third party API is not available it throws error but we can use caching system like redis to cache data.

### Questions Answered
1.What if the requirement for the new endpoint was to also allow the consumer to compare populations for any given date. How would you modify your implementation?
`I will create New Endpoint for that and new service layer to compare population for the countries I can reuse the method written to get population for the countries`
2. What if we have a database of users and we wanted to make our API smarter by defaulting comparisons to always include the population of the current user's country. How could we accomplish this?
`I will store all the countries population in cache and return that by fetching user's country population`
3. What if we wanted to keep a tally of the most frequently requested countries and have this be available to consumers. How could we accomplish this?
`I would give ranking for the most requested countries by incrementing values each time a request is made`
