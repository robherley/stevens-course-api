package graphapi

import (
	"fmt"

	"github.com/robherley/stevens-course-api/stevens"
	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
	"github.com/robherley/stevens-course-api/mongo"
)

var requireType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Requirement",
	Description: "The requirement(s) for a specific course section. TODO: describe arg/val/op",
	Fields: graphql.Fields{
		"control": &graphql.Field{
			Type: graphql.String,
		},
		"argument": &graphql.Field{
			Type: graphql.String,
		},
		"value1": &graphql.Field{
			Type: graphql.String,
		},
		"operator": &graphql.Field{
			Type: graphql.String,
		},
		"value2": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var meetingType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Meeting",
	Description: "The meeting time for a specific class section.",
	Fields: graphql.Fields{
		"day": &graphql.Field{
			Type: graphql.String,
		},
		"startTime": &graphql.Field{
			Type: graphql.String,
		},
		"endTime": &graphql.Field{
			Type: graphql.String,
		},
		"site": &graphql.Field{
			Type: graphql.String,
		},
		"building": &graphql.Field{
			Type: graphql.String,
		},
		"room": &graphql.Field{
			Type: graphql.String,
		},
		"activity": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var courseType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Course",
	Description: "A Course at Stevens",
	Fields: graphql.Fields{
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"section": &graphql.Field{
			Type: graphql.String,
		},
		"callNumber": &graphql.Field{
			Type: graphql.String,
		},
		"minCredit": &graphql.Field{
			Type: graphql.String,
		},
		"maxCredit": &graphql.Field{
			Type: graphql.String,
		},
		"currentEnrollment": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"startDate": &graphql.Field{
			Type: graphql.String,
		},
		"endDate": &graphql.Field{
			Type: graphql.String,
		},
		"instructor": &graphql.Field{
			Type: graphql.String,
		},
		"meetings": &graphql.Field{
			Type: &graphql.List{
				OfType: meetingType,
			},
		},
		"requirements": &graphql.Field{
			Type: &graphql.List{
				OfType: requireType,
			},
		},
	},
})

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"getCourse": &graphql.Field{
			Type:        courseType,
			Description: "Get a course at Stevens by Call Number and Semester",
			Args: graphql.FieldConfigArgument{
				"callnumber": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "The callnumber of a course.",
				},
				"semester": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "The semester of course.",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				cn := p.Args["callnumber"].(string)
				sem := p.Args["semester"].(string)

				s, c, err := mongo.SafeGetSemester(sem)
				if err != nil {
					return nil, err
				}
				defer s.Close()

				var result stevens.Course
				err = c.FindId(cn).One(&result)
				if err != nil {
					return nil, err
				}

				return result, nil
			},
		},
		"searchCourse": &graphql.Field{
			Type: &graphql.List{
				OfType: courseType,
			},
			Description: "Search for courses by section.",
			Args: graphql.FieldConfigArgument{
				"section": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "The section to look for.",
				},
				"semester": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "The semester of course.",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				sec := p.Args["section"].(string)
				sem := p.Args["semester"].(string)

				s, c, err := mongo.SafeGetSemester(sem)
				if err != nil {
					return nil, err
				}
				defer s.Close()

				var results []stevens.Course
				err = c.Pipe([]bson.M{{"$match": bson.M{"section": bson.M{"$regex": fmt.Sprintf(".*%s.*", sec)}}}}).All(&results)
				if err != nil {
					return nil, err
				}

				return results, nil
			},
		},
	},
})
