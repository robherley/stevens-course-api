package stevens

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"

	"github.com/robherley/stevens-course-api/mongo"
	"github.com/robherley/stevens-course-api/util"
)

// Semester struct for semesters
type Semester struct {
	Semester   string   `xml:"Semester,attr" json:"semester"`
	CourseList []Course `xml:"Course" json:"courses"`
}

// Course struct for courses
type Course struct {
	Section           string        `xml:"Section,attr" json:"section"`
	Title             string        `xml:"Title,attr" json:"title"`
	CallNumber        string        `xml:"CallNumber,attr" json:"callNumber" bson:"_id"`
	MinCredit         string        `xml:"MinCredit,attr" json:"minCredit"`
	MaxCredit         string        `xml:"MaxCredit,attr" json:"maxCredit"`
	CurrentEnrollment string        `xml:"CurrentEnrollment,attr" json:"currentEnrollment"`
	Status            string        `xml:"Status,attr" json:"status"`
	StartDate         string        `xml:"StartDate,attr" json:"startTime"`
	EndDate           string        `xml:"EndDate,attr" json:"endTime"`
	Instructor        string        `xml:"Instructor1,attr" json:"instructor"`
	Meetings          []Meeting     `xml:"Meeting" json:"meetings"`
	Requirements      []Requirement `xml:"Requirement" json:"requirements"`
}

// Meeting struct for meetings
type Meeting struct {
	Day       string `xml:"Day,attr" json:"day"`
	StartTime string `xml:"StartTime,attr" json:"startTime"`
	EndTime   string `xml:"EndTime,attr" json:"endTime"`
	Site      string `xml:"Site,attr" json:"site"`
	Building  string `xml:"Building,attr" json:"building"`
	Room      string `xml:"Room,attr" json:"room"`
	Activity  string `xml:"Activity,attr" json:"activity"`
}

// Requirement struct for requirements
type Requirement struct {
	Control  string `xml:"Control,attr" json:"control"`
	Argument string `xml:"Argument,attr" json:"argument"`
	Value1   string `xml:"Value1,attr" json:"value1"`
	Operator string `xml:"Operator,attr" json:"operator"`
	Value2   string `xml:"Value2,attr" json:"value2"`
}

// FetchSemester grabs a semester object for a given year
func FetchSemester(s string) Semester {
	body, err := util.ByteRequest(fmt.Sprintf("https://web.stevens.edu/scheduler/core/core.php?cmd=getxml&term=%s", s))
	if err != nil {
		panic(err)
	}
	var sm Semester
	err = xml.Unmarshal(body, &sm)
	if err != nil {
		panic(err)
	}
	return sm
}

// ToJSON Returns the entire semester as a json object
func (sm Semester) ToJSON() []byte {
	json, err := json.Marshal(&sm)
	if err != nil {
		panic(err)
	}
	return json
}

// InsertToDB adds the semester to mongo
func (sm Semester) InsertToDB() {
	s, c := mongo.GetSemester(sm.Semester)
	defer s.Close()

	for _, cr := range sm.CourseList {
		_, err := c.UpsertId(cr.CallNumber, cr)
		if err != nil {
			log.Fatal(err)
		}
	}
}
