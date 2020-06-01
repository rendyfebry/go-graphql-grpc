package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"google.golang.org/grpc"

	"github.com/rendyfebry/go-graphql-grpc/graphql/internal/resolver"
)

func main() {
	// Get env variables
	productHost := os.Getenv("PRODUCT_HOST")
	if productHost == "" {
		productHost = "localhost"
	}

	productPort := os.Getenv("PRODUCT_PORT")
	if productPort == "" {
		productPort = "50051"
	}

	// Read graphql schema
	s, err := getSchema("./api/schema.graphql")
	if err != nil {
		panic(err)
	}

	fmt.Println("Product :", productHost)
	fmt.Println("Product :", productPort)

	// Set up a connection to the server.
	grpcAddress := fmt.Sprintf("%s:%s", productHost, productPort)
	conn, err := grpc.Dial(grpcAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	schema := graphql.MustParseSchema(s, resolver.NewResolver(conn))
	http.Handle("/query", &relay.Handler{
		Schema: schema,
	})

	log.Println("GraphQL service started!")
	log.Println("URL: localhost:8080")
	log.Println("gRPC Service addr:", grpcAddress)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getSchema(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
