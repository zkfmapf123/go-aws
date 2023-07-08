package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "it's working")
}

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	// err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatalln("ListenAndServeTLS error : ", err)
	}
}
