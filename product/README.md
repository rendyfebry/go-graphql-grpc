# The gRPC Microservice (Product)

## Run The Project

Simply run this command

```
make run
```

## Run The Project Using Docker

### Build The Project

```
docker build -t go-grpc-product:latest .
```

### Run The Docker Image

To run the docker image, you need to use volume

```
docker run -p 50051:50051 go-grpc-product
```
