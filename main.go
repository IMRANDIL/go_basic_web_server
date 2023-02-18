package main

import (
	"fmt"
	"log"
	"net/http"
)



func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found",http.StatusNotFound)
		return
	}


}




func main () {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Printf("server started on port :8000\n")

	err :=http.ListenAndServe(":8000",nil)

		if err != nil {
			log.Fatal(err)
		}
}