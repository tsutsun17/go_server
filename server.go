package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Tatsuemon/go_server/models"
	"github.com/codegangsta/negroni"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		user := models.NewUser("Naoyoshi Aikawa", 29)
		userStr, err := json.Marshal(user)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(userStr))
	})
	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":8080")
}
