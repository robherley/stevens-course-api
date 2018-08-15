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
	if os.Getenv("MONGO_URI") == "" {
		os.Setenv("MONGO_URI", "localhost:27017")
	}
	sem := stevens.FetchSemester("2018F")
	fmt.Printf("inserting semester '%s' into db...\n", sem.Semester)
	sem.InsertToDB()
	fmt.Println("finished insertion")

	h := handler.New(&handler.Config{
		Schema:   &graphapi.Schema,
		Pretty:   true,
		GraphiQL: false,
	})

	http.Handle("/", h)
	http.Handle("/play", http.StripPrefix("/play", http.FileServer(http.Dir("./static"))))
	http.ListenAndServe(":8080", nil)
}
