package main
<<<<<<< HEAD
import (
    "fmt"
    "log"
    "net/http"
)
func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
=======

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
<<<<<<< HEAD
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
=======
>>>>>>> 83695b6fa284c9f008a6a9f875b144175f5c9d07
>>>>>>> 690a314d3e5a121cb5b0506109a5ce5abbd27af8
