# Introduction  

webginrest adalah skeleton web golang using framwork [gin](https://github.com/gin-gonic/gin).  

## Getting Started  

## Prerequisites  

1. Install latest [golang](https://golang.org/doc/install)  

## Installation process  

1. Clone Repository  

```sh
$ git clone git@github.com:rulzi/webginrest.git
```  

2. Get Depedency repository  

```sh  
$ go get ./...
```  

3. Open config.yml in project root dan setting your environment  

```sh  
$ cp app/config/config.yml.example app/config/config.yml  
```  

## Usage  

1. use golang run

```sh  
$ go run cmd/main.go 
```  

## Makefile  

- build application golang to make sure application runing well  

```sh  
$ make build  
```  

- build application and generate binary image  

```sh  
$ make build-generate  
```  

## Key Directory

* `cmd`: Main Golang
* `server`: server file (edit if necessary)
* `server/config`: Config File
* `server/provider`: Provide function to communicate with gin
* `server/provider`: Provide function service
* `server/server`: Run Server
* `database`: File for setting Migration Database
* `app/helpers`: Helper File
* `app/middleware`: Middleware File
* `app/routers`: Router app
* `app/utils`: Utilities File
* `app/pkg/controller`: All Controllrt Layer (This layer will act as the presenter. Decide how the data will presented. Could be as REST API, or HTML File. This layer also will accept the input from user. Sanitize the input and sent it to Usecase layer.)
* `app/pkg/form_validation`: Struct File Validation
* `app/pkg/models`: All Models layer (Same as Entities, will used in all layer. This layer, will store any Object’s Struct and its method)
* `app/pkg/transformers`: All Transformer layer
* `app/pkg/repository`: Repository layer (Repository will store any Database handler. Querying, or Creating/ Inserting into any database will stored here. This layer will act for CRUD to database only. No business process happen here. Only plain function to Database.)
* `app/pkg/usecase`: All Usecase Layer (This layer will act as the business process handler. Any process will handled here. This layer will decide, which repository layer will use. And have responsibility to provide data to serve into delivery. Process the data doing calculation or anything will done here.)