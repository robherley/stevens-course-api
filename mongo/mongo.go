package mongo

import (
	"fmt"
	"os"
	"sort"

	mgo "gopkg.in/mgo.v2"
)

//GetSemester returns the session and a reference to the post collection.
func GetSemester(s string) (*mgo.Session, *mgo.Collection, error) {
	session, err := mgo.Dial(os.Getenv("MONGO_URI"))
	if err != nil {
		panic(err)
	}

	collection := session.DB("semester").C(s)
	return session, collection, nil
}

//SafeGetSemester throws an error when trying to get a bad semester not in db
func SafeGetSemester(s string) (*mgo.Session, *mgo.Collection, error) {
	session, err := mgo.Dial(os.Getenv("MONGO_URI"))
	if err != nil {
		panic(err)
	}

	cnames, err := session.DB("semester").CollectionNames()
	if err != nil {
		panic(err)
	}

	sort.Strings(cnames)
	i := sort.SearchStrings(cnames, s)
	if i < len(cnames) && cnames[i] == s {
		collection := session.DB("semester").C(s)
		return session, collection, nil
	}
	err = fmt.Errorf("The collection '%s' does not exist", s)
	return session, nil, err
}
