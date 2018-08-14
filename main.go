package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/graphql-go/handler"
	"github.com/robherley/stevens-course-api/graphql"
	"github.com/robherley/stevens-course-api/stevens"
)

func main() {
	sem := stevens.FetchSemester("2018F")
	os.Setenv("MONGO_URI", "localhost:27017")
	fmt.Printf("inserting semester '%s' into db...\n", sem.Semester)
	sem.InsertToDB()
	fmt.Println("finished insertion")

	h := handler.New(&handler.Config{
		Schema:   &graphapi.Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}
