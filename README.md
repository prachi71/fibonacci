# Fibonacci series using SQL 
Sample project using Postgress in Golang to generate a Fibonacci series, find nth value

##### Requirements
Expose a Fibonacci sequence generator through a web API that memoizes intermediate values. The web API should expose operations to (a) fetch the Fibonacci number given an ordinal (e.g. Fib(11) == 89, Fib(12) == 144), (b) fetch the number of memoized results less than a given value (e.g. there are 12 intermediate results less than 120), and (c) clear the data store.

The web API must be written in Go, and Postgres must be used as the data store for the memoized results. Please include tests for your solution, and a README.md describing how to build and run it.

Bonus points:
- Use dockertest.
- Include a Makefile
- Include some data on performance

### Rest Api

_GIN_ : For rest services </br>
Swagger : Api docs :

### Makefile : Key targets are in bold
make help</br></br>
build : build a local binary , get the dependencies and then run a go build </br></br>
build : build a local binary in bin directory </br></br>
run : run the app locally without docker via run.sh</br></br>
run : run locally without docker called from run.sh</br></br>
pg-start : start pg db in docker</br></br>
pg-stop : stop docker container for pg</br></br>
ps : docker ps list the running images in docker</br></br>
**container** : run docker compose that will build the docker image for the app , start pg and start the app</br></br>
**stop-container** : stop docker containers started by docker-compose</br></br>
login-pg : login / bash shell to the pg container</br></br>
login-app : login / bash shell to the app container</br></br>
**test** : Runs unit tests.</br></br>
cover : Generates a test coverage report</br></br>
clean : clean build/run artifacts including remove coverage report and the binary.</br></br>
deps : Update/refresh dependencies from go.mod</br></br>

### Optional docker cmds , wrappers are provided in the make file ( container/stop-container)

docker-compose up --build </br></br>
docker-compose -f docker-compose.yml stop

### Swagger details

1) make container </br>
2) browse to http://localhost:8080/swagger/index.html

### Transient store

Postgres table is used as scratch pad only. Transaction begin/rollback strategy is used.

### Misc

Rest response includes time taken to process the request.


