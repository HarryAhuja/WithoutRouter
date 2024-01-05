package main

import (
	"encoding/json"
	"net/http"
)

// http package has things like in-flight,canonical form etc used in uber
func main() {

	http.HandleFunc("/get", func(resp http.ResponseWriter, req *http.Request) {
		maps := map[string]bool{}
		maps["harry"] = true
		bytes, _ := json.Marshal(maps)
		resp.WriteHeader(200)
		resp.Write(bytes)
	})
	http.ListenAndServe(":8000", nil)
}
