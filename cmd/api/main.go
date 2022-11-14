package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handleNewsletter(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "GET METHOD\n")
	case "POST":
		w.WriteHeader(http.StatusOK)
		/* Making json formated response*/
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Status OK"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}
}

func main() {
	http.HandleFunc("/newsletter", handleNewsletter)
	http.ListenAndServe(":5000", nil)
}
