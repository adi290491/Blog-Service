# Blog Service

This is a blog service api implemented with Go and gRPC and utilized MongoDB as the database for managing the blogs. The purpose of this project is to demonstrate working with gRPC and Golang.

## Setup Requirement
* Go 1.18+
* Docker

## Usage
* Start MongoDB from docker
  ```bash
    docker compose up
  ```
* Start the server
  ```bash
    cd server/
    ./server
  ```
* Start the client
  ```bash
    cd client/
    ./client
  ```
## Future Scope
* Development with a microservice architecture
* Front-end
* Concurrency
