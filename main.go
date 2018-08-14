package main

import (
	"fmt"
	"os"

	"github.com/robherley/stevens-course-api/stevens"
)

func main() {
	sem := stevens.FetchSemester("2018F")
	os.Setenv("MONGO_URI", "localhost:27017")
	fmt.Printf("inserting semester '%s' into db...\n", sem.Semester)
	sem.InsertToDB()
	fmt.Println("finished insertion")
}
