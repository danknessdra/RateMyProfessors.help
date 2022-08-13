package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// return ranking of courses based off school, and course name

// first return courses that contain course name

type Course struct {
	Department string  `json:"department"`
	Course     string  `json:"course"`
	Prof       string  `json:"prof"`
	Rating     float64 `json:"rating"`
	Size       int     `json:"size"`
}

// custom print function
func (course Course) Format(f fmt.State, c rune) {
	f.Write([]byte(course.Department + "\n"))
	f.Write([]byte(course.Course + "\n"))
	f.Write([]byte(course.Prof + "\n"))

	// kinda/kinda dont fully understand format float, but hopefully it means no exp, and truncate to only 1 place
	f.Write([]byte(strconv.FormatFloat(course.Rating, 'f', 1, 32) + "\n"))
	f.Write([]byte(strconv.Itoa(course.Size) + "\n"))
}

// file =../data/courses.json
/*
error handling

os.open school json

if err != nil
return N/A to webserver
*/
func GetJson(file string) []Course {

	jsonFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
	}

	// research what ioutil.ReadALL does
	byteValue, _ := io.ReadAll(jsonFile)

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

type ByRanking []Course

func (a ByRanking) Len() int      { return len(a) }
func (a ByRanking) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// make it rank if ranks are equal, compare number of ratins
func (a ByRanking) Less(i, j int) bool { return a[i].Rating < a[j].Rating }

func GetRanked(courses []Course, input string) []Course {
	sort.Sort(ByRanking(courses))

	var ranked []Course
	for i := 0; i < len(courses); i++ {
		if strings.Compare(courses[i].Course, input) == 0 {
			ranked = append(ranked, courses[i])
			//fmt.Println(courses[i])
		}
	}

	sort.Sort(ByRanking(ranked))
	return ranked
}
