## Overview
This is a simple project about cqrs. It's including: how to connect to mongodb, work with gRPC and HTTP/REST. Code structure follows clean code (not fully, I guess :D).

## Setup project
1. clone `.env.example` and change its name to `.env`
2. create mongodb database with `event` collection
3. This project uses `go dep` to manage dependencies. So after install `go dep` run `dep ensure`.
4. RUN `go run cmd/main.go` or `go build cmd/main.go && ./main`
5. Use whatever you like to test RESTful API (I personally like using insomnia). For event api I create 2 endpoints:
    - `curl -X GET 'http://localhost:8080/v1/event'` (for event list)
    - `curl -X POST 'http://localhost:8080/v1/event' -d '{"body": "test 3ewjrhewjfk"}'` (for create)

## Explain
#### pb
this folder contains our gRPC proto files. I use v1 folder for versioning api. Each `*.proto` file represents APIs relevant to an entity inside `storage` folder. After create an `proto` file run `./protoc-gen.sh {version} {entity}` (ex: `./protoc-gen.sh v1 event`) to generate others files.
#### storage
this is where our entities live. Each file declares entity struct and methods to interact with it.
#### service
If you familiar with repository-service pattern, this is where we write business logic.
#### handler
handler is api endpoints
#### cmd
main package, it calls run server gRPC and HTTP/REST