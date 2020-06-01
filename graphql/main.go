package main

import (
	"io/ioutil"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type query struct{}

func (_ *query) Hello() string {
	return "Hello, world!"
}

func main() {
	s, err := getSchema("./api/schema.graphql")
	if err != nil {
		panic(err)
	}

	schema := graphql.MustParseSchema(s, &query{})
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
