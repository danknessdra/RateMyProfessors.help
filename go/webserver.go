package main
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

type node struct {
	Course		string `json:"course"`
	Professor	string `json:"prof"`
	Rating		float32 `json:"rating"`
}

var nodes = []node{
	{Course : "Math", Professor : "jeff", Rating : "3.2"}
	{Course : "Physics", Professor : "saul", Rating : "3.4"}
	{Course : "English", Professor : "white", Rating : "4.6"}
	{Course : "Chemistry", Professor : "jesse", Rating : "4.2"}
	{Course : "History", Professor : "raymond", Rating : "5"}
	{Course : "PE", Professor : "bhat", Rating : "0.0"}
}

/*
for i in range(0,6):
	courses = ["Math", "Physics", "English", "Chemistry", "History", "PE" ]
	professors = ["jeff", "saul", "white", "jesse", "raymond", "bhat"]
	rating = [3.2, 3.4, 4.6, 4.2, 5, 0.0]
	
	print("\t{{Course : \"{0}\", Professor : \"{1}\", Rating : \"{2}\"}}".format(courses[i], professors[i], rating[i]))
*/
