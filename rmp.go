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

	err := http.ListenAndServeTLS(":8081", "/etc/ssl/ratemyprofessors.crt", "/etc/ssl/private/ratemyprofessors.key", router)

	if err != nil {
		fmt.Println(err)
		fmt.Println("WARNING: using http instead of https")
		log.Fatal(http.ListenAndServe(":8081", router))
	}
}

func getRank(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var data InputCall
	err := json.NewDecoder(request.Body).Decode(&data)
	ranked := GetRanked(GetJson("./data/DeAnzaCourses.json"), data.Course)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if data.School == "UC Berkeley" {
		ranked = GetRanked(GetJson("./data/BerkeleyCourses.json"), data.Course)
	}
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(ranked)
}

func getCourses(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ranked := GetJson("./data/DeAnzaCourses.json")
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(writer).Encode(ranked)
}

type InputCall struct {
	School string
	Course string
}
