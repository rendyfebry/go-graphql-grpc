package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	h "google.golang.org/grpc/health"
	hpb "google.golang.org/grpc/health/grpc_health_v1"

	ppb "github.com/rendyfebry/go-graphql-grpc/product/api"
	p "github.com/rendyfebry/go-graphql-grpc/product/internal/grpc"
)

var (
	env    *string
	prefix *string
	host   *string
	port   *int
)

func main() {
	// Init flag
	env = flag.String("env", "development", "Environment")
	host = flag.String("host", "0.0.0.0", "Host for http listener")
	port = flag.Int("port", 50051, "Port for http listener")
	prefix = flag.String("prefix", "/v1", "Prefix for http route")
	flag.Parse()

	// Init gRPC server
	grpcServer := grpc.NewServer()
	ppb.RegisterProductSvcServer(grpcServer, p.NewServer())
	hpb.RegisterHealthServer(grpcServer, h.NewServer())

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Panic(err.Error())
	}

	log.Println("Service started!")
	log.Println(fmt.Sprintf("URL: %s:%d", *host, *port))

	if err := grpcServer.Serve(lis); err != nil {
		log.Panic(err.Error())
	}
}
