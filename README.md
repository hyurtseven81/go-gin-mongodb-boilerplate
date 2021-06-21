# Golang gin mongodb boilerplate

This project aims to creating simple REST api with Golang's gin framework on top of mongodb.

Basically it's built on top of the following ideas:

- MongoDB's official driver (There is no need to use ORM)
- Repository pattern to abstract MongoDB related parts from the other layers
- Database connection pooling
- Parallelism with goroutines (see resources/task/service/List)
- Auth middleware for JWT decoding and handling
- To be ready for protobufs (see resources/task/pb)
- Custom error handling
- Unit tests + integration tests
- [air](https://github.com/cosmtrek/air) support for live reload
- Docker support
- Docker compose support
- Healthcheck support
- Follows [12factor.net](https://12factor.net) standards


## Quick setup

Either running docker compose directly: 

```bash
docker compose up
```

or by building from scratch

```bash
go get
air
```

### Tests

To run the tests you need to have a running mongodb database. To do that:

```bash
docker compose up -f docker-compose.mongodb.yaml up -d
````

After that you can run 
```bash
go test ./...
```