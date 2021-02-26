#Vanila make file

#export PG credentials for local runs outside docker
export POSTGRES_USER=PG_USER
export POSTGRES_PASSWORD=PG_PASS
export POSTGRES_DB=PG_DATABASE
export POSTGRES_HOST=localhost
export POSTGRES_PORT=5432

# Go related commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test ./...
GOGET=$(GOCMD) get -u -v

# Name of actual binary to create
BINARY = fibonacci

.PHONY : help
help : Makefile
	@sed -n 's/^##//p' $<

## build : build a local binary , get the dependencies and then run a go build
.PHONY: build
build: deps localbuild

## build : build a local binary in bin directory
.PHONY: localbuild
localbuild:
	$(GOCMD) build -o bin/${BINARY} .

## run : run the app locally without docker via run.sh
.PHONY: run
run:
	./run.sh

## run : run locally without docker called from run.sh
.PHONY: runLocal
runLocal: build
	bin/${BINARY}

## pg-start : start pg db in docker
.PHONY: pg-start
pg-start:
	docker run -d --rm -P -p 5432:5432 -e POSTGRES_USER=PG_USER -e POSTGRES_PASSWORD=PG_PASS -e POSTGRES_DB=PG_DATABASE --name pg postgres:latest

## pg-stop : stop docker container for pg
.PHONY: pg-stop
pg-stop:
	docker stop pg

## ps : docker ps list the running images in docker
.PHONY: ps
ps:
	docker ps

## container : run docker compose that will build the docker image for the app , start pg and start the app
.PHONY: container
container:
	docker-compose up --build

## stop-container : stop docker containers started by docker-compose
.PHONY: stop-container
stop-container:
	docker-compose -f docker-compose.yml stop

## login-pg : login / bash shell to the pg container
.PHONY: login-pg
login-pg:
	docker-compose -f docker-compose.yml exec database /bin/bash

## login-app : login / bash shell to the app container
.PHONY: login-app
login-app:
	docker-compose -f docker-compose.yml exec server /bin/sh

## test : Runs unit tests.
.PHONY: test
test:
	$(GOTEST) -v

## cover : Generates a test coverage report
.PHONY: cover
cover:
	${GOCMD} test -coverprofile=coverage.out ./... && ${GOCMD} tool cover -html=coverage.out

## clean : clean build/run artifacts including remove coverage report and the binary.
.SILENT: clean
.PHONY: clean
clean:
	$(GOCLEAN)
	@rm -f ${BINARY}-$(OS)-${GOARCH}
	@rm -f coverage.out
	@rm -rf bin

## deps : Update/refresh dependencies from go.mod
.PHONY: deps
deps:
	${GOCMD} mod download

