# The GraphQL Proxy

## Run The Project

Simply run this command

```
make run
```

## Run The Project Using Docker

### Build The Project

```
docker build -t go-grpc-graphql:latest .
```

### Run The Docker Image

To run the docker image, you need to use volume

```
docker run --env PRODUCT_HOST=host.docker.internal -p 8080:8080 go-grpc-graphql
```
