package mongo

import (
	"log"
	"os"

	mgo "gopkg.in/mgo.v2"
)

//GetSemester returns the session and a reference to the post collection. Credit: https://github.com/nilstgmd/graphql-starter-kit/blob/master/mongo/mongo.go
func GetSemester(s string) (*mgo.Session, *mgo.Collection) {
	session, err := mgo.Dial(os.Getenv("MONGO_URI"))
	if err != nil {
		log.Fatal(err)
	}

	collection := session.DB("semester").C(s)
	return session, collection
}
