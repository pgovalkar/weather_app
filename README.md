# Weather API App

This app will import the seattle weather report which is in csv format and the create http server to query the data

## Pre-requisites 
### For local setup
. golang 1.21.0 version should be installed
. After git clone please update the full path off seattleweather.csv file in `util/import.go` file

### For running as container
. docker is installed is on your local machine

## Next steps, clone the git repo  
## Testing Locally
1. Run `go run main.go` from the root folder of the repo
2. Open browser and connect to `localhost:8080`
3. To filter query can use below examples
```
 - Limit results to a number `http://localhost:8080/query?limit=5`
 - By date: `http://localhost:8080/query?date=2012-06-04`
 - By weather type: `http://localhost:8080/query?weather=rain`
 - multi-query filtering, eg. `http://localhost:8080/query?weather=rain&limit=5`
```

## Running as Container
1. Make sure csv file path is updated to `/weather/seattleweather.csv`
2. Create  docker image using Dockerfile under root directory <br />
   `docker build -t weather .` <br />
2. Create container eg: <br />
   `docker run -d -p 2200:8080 --name weather weather` <br />
3. For testing follow steps 2 & 3 from Testing Locally <br />

## Testing the APi using script
1. Create  docker image using Dockerfile under testing directory <br />
   `docker build -t weather-testing .` <br />
2. Create bridge network so bothe containers can talk to each other <br />
   `docker network create -d bridge mynetwork` <br />
3. Create containers using bridge network eg: <br />
   `docker run -d -p 2200:8080 --name weather --network mynetwork weather` <br />
   `docker run -d --name weather-test --network mynetwork weather-testing` <br />
4. The weather-testing container will connect to weather container and execute the script which is undet testing. <br />
5. To download the result, run <br />
`for i in 1 2 3 4 5 ; do docker cp weather-test:/testing/test$i.out . ; done`
6. Can also test using browser on `localhost:2200`

