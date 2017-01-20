package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Response struct {
	Status   string `json:"status"`
	Resource string `json:"resource"`
	Message  string `json:"message"`
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func log(message string) {
	fmt.Println("[", time.Now().Format("2006-01-02 15:04:05"), "]", message)
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(200000)
	w.Header().Add("Content Type", "application/json")
	var response Response
	if r.Method == "POST" {
		resource := resourceParse(r.Form.Get("resource"))
		if resource.Valid {
			response = Response{
				Status:   "ok",
				Resource: resource.Resource,
			}
		} else {
			response = Response{
				Status:  "failed",
				Message: "invalid resource",
			}
		}
	} else {
		response = Response{
			Status:   "ok",
			Resource: Uuid(),
		}
	}
	json.NewEncoder(w).Encode(response)
}
