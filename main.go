package main

import (
	"net/http"
	"os"

	"github.com/graphql-go/handler"
	"github.com/robherley/stevens-course-api/graphql"
)

func main() {
	os.Setenv("MONGO_URI", "localhost:27017")
	// sem := stevens.FetchSemester("2018F")
	// fmt.Printf("inserting semester '%s' into db...\n", sem.Semester)
	// sem.InsertToDB()
	// fmt.Println("finished insertion")

	h := handler.New(&handler.Config{
		Schema:   &graphapi.Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/", h)
	http.ListenAndServe(":8080", nil)
}
