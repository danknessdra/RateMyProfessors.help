package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
error handling

os.open school json

if err != nil
return N/A to webserver
*/

// return ranking of courses based off school, and course name

// first return courses that contain course name

type Course struct {
	Course     string  `json:"course"`
	Department string  `json:"department"`
	Prof       string  `json:"prof"`
	Rating     float64 `json:"rating"`
	Size       int     `json:"size"`
}

/*
type Courses struct {
	Courses []Course `json:"courses"`
}
*/

type ByRanking []Course

func (a ByRanking) Len() int           { return len(a) }
func (a ByRanking) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRanking) Less(i, j int) bool { return a[i].Rating < a[j].Rating }

// custom print function

func (course Course) Format(f fmt.State, c rune) {
	f.Write([]byte(course.Course + "\n"))
	f.Write([]byte(course.Department + "\n"))
	f.Write([]byte(course.Prof + "\n"))
	f.Write([]byte(strconv.FormatFloat(course.Rating, 'f', 1, 32) + "\n"))
	f.Write([]byte(strconv.Itoa(course.Size)))
}

// file =../data/courses.json
func GetJson(file string) []Course {

	jsonFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
	}

	// research what ioutil.ReadALL does
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// courses var that is the struct Courses( which is a list of course structs )
	var courses []Course

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'course' which we defined above

	err_marsh := json.Unmarshal(byteValue, &courses)

	if err_marsh != nil {
		fmt.Println("error:", err)
	}

	return courses
}

func GetCourse(courses []Course) {
	sort.Sort(ByRanking(courses))

	fmt.Println("Department: " + courses[0].Department)

	for i := 0; i < len(courses); i++ {
		var unranked []Course
		if strings.Compare(courses[i].Course, "MATH 1A") == 0 {
			unranked = append(unranked, courses[i])
			fmt.Println(courses[i])
		}
		// kinda/kinda dont fully understand format float, but hopefully it means no exp, and truncate to only 1 place
	}
}
