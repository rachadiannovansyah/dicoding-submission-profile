package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Now Listening on 8080")
	http.HandleFunc("/", serveFiles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveFiles(w http.ResponseWriter, r *http.Request) {

	p := "." + r.URL.Path
	if p == "./" {
		p = "./static/index.html"
	}
	fmt.Println(p)
	http.ServeFile(w, r, p)
}
