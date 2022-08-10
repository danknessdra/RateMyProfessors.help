package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "html/index.html")
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	//router.GET("/hello/:name", Hello)
	router.ServeFiles("/js/*filepath", http.Dir("js"))
	router.ServeFiles("/css/*filepath", http.Dir("css"))

	log.Fatal(http.ListenAndServe(":8080", router))
}

/*
type node struct {
	Course    string  `json:"course"`
	Professor string  `json:"prof"`
	Rating    float32 `json:"rating"`
}

var nodes = []node{
	{Course : "Math", Professor : "jeff", Rating : "3.2"},
	{Course : "Physics", Professor : "saul", Rating : "3.4"},
	{Course : "English", Professor : "white", Rating : "4.6"},
	{Course : "Chemistry", Professor : "jesse", Rating : "4.2"},
	{Course : "History", Professor : "raymond", Rating : "5"},
	{Course : "PE", Professor : "bhat", Rating : "0.0"},
}

*/

/* generates sample nodes
for i in range(0,6):
	courses = ["Math", "Physics", "English", "Chemistry", "History", "PE" ]
	professors = ["jeff", "saul", "white", "jesse", "raymond", "bhat"]
	rating = [3.2, 3.4, 4.6, 4.2, 5, 0.0]

	print("\t{{Course : \"{0}\", Professor : \"{1}\", Rating : \"{2}\"}},".format(courses[i], professors[i], rating[i]))
*/
