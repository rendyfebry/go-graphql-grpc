# GraphQL & gRPC Implementation on Golang

## Run The Project

Read each service documentation to run the service manually

- [Product Service](./product/README.md)
- [GraphQL Service](./graphql/README.md)

## Run The Project Using Docker Compose

Run all containers using `docker-compose`

```
docker-compose up --build
```

## Usage

### Get All Product

GraphQL

```
query {
    products{
        id
        name
        price
    }
}
```

CURL

```
curl -L -X POST 'http://localhost:8080/query' \
-H 'Content-Type: application/json' \
--data-raw '{"query":"query {\n    products{\n        id\n        name\n        price\n    }\n}","variables":{}}'
```
