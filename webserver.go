package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "html/index.html")
}

func main() {

	router := httprouter.New()
	router.GET("/", Index)
	router.ServeFiles("/js/*filepath", http.Dir("js"))
	router.ServeFiles("/css/*filepath", http.Dir("css"))

	router.POST("/get_rank", getRank)
	// router.GET("/get_courses", getCourses) API

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))

	// UNCOMMENT THIS AND COMMENT ABOVE TO GET TLS WORKING
	// log.Fatal(http.ListenAndServeTLS("127.0.0.1:8080", "/etc/sslratemyprofessor.crt", "/etc/ssl/private/ratemyprofessor.key", router))

}

func getRank(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var data InputCall
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	ranked := GetRanked(GetJson("./data/courses.json"), data.Course)
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(ranked)
}

func getCourses(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ranked := GetJson("./data/courses.json")
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(writer).Encode(ranked)
}

type InputCall struct {
	School string
	Course string
}
