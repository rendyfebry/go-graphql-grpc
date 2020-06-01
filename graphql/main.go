package main

import (
	"io/ioutil"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"google.golang.org/grpc"

	"github.com/rendyfebry/go-graphql-grpc/graphql/resolver"
)

func main() {
	s, err := getSchema("./api/schema.graphql")
	if err != nil {
		panic(err)
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	schema := graphql.MustParseSchema(s, resolver.NewResolver(conn))
	http.Handle("/query", &relay.Handler{
		Schema: schema,
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getSchema(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
