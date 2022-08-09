package main

import (
	"net/http"
	"fmt"
)

type node struct {
	Course		string `json:"course"`
	Professor	string `json:"prof"`
	Rating		float32 `json:"rating"`
}

var nodes = []node{
for i in range(0,6):
	courses = ["Math", "Physics", "English", "Chemistry", "History", "PE" ]
	professors = ["jeff", "saul", "white", "jesse", "raymond", "bhat"]
	rating = [3.2, 3.4, 4.6, 4.2, 5, 0.0]
	
	print("\{Course : \"{0}\", Professor : \"{1}\", Rating : \"{2}\"".format(courses[i], professors[i], rating[i]))
}

func main() {
}
/*
for i in range(0,3):
	courses = ["Math", "Physics", "English", "Chemistry", "History", "PE" ]
	professors = ["jeff", "saul", "white"]
	rating = [5.2, 3.4, 5.6]
	
	print("{0}, {1},{2}".format(courses[i], professors[i], rating[i]))

for i in range(0,6):
	courses = ["Math", "Physics", "English", "Chemistry", "History", "PE" ]
	professors = ["jeff", "saul", "white", "jesse", "raymond", "bhat"]
	rating = [3.2, 3.4, 4.6, 4.2, 5, 0.0]
	
	print("Course : \"{0}\", Professor : \"{1}\", Rating : \"{2}\"".format(courses[i], professors[i], rating[i]))
*/
