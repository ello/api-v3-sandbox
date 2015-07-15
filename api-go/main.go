package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type user struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

func main() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/users", usersHandler).Methods("GET")
	log.Println("Listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func usersHandler(res http.ResponseWriter, req *http.Request) {
	log.Println("Responding to /users request")
	log.Println(req.UserAgent())

	users := []user{
		user{Username: "rtyer", Name: "Ryan Tyer"},
		user{Username: "jayzes", Name: "Jay Zeschin"},
	}

	outgoingJSON, error := json.Marshal(users)

	if error != nil {
		log.Println(error.Error())
		http.Error(res, error.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Fprint(res, string(outgoingJSON))
}
