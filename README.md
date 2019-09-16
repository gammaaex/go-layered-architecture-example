# go-layered-architecture-example

## Usage
1. `cd docker`
1. `cp .env.example .env`
1. `docker-compose build`
1. `docker-compose up -d`
1. Access to `localhost:8888`
1. `docker-compose down`

### You can create user by HTTP POST Request
```
curl -X POST -H "Content-Type: application/json" -d '{"name":"user"}' localhost:8888/users
```

### Container Service List
|service|port|  
|:---:|:---:|  
|API Server(Apache)|8888|
|phpMyAdmin|8080|
|MySQL|13306|

## Self Build
1. Edit code
1. `cd docker`
1. `docker-compose build`
1. `docker-compose up -d`
1. `docker-compose exec go bash`
1. `make fmt`
1. `make build`
1. `exit`
1. `docker-compose down`
