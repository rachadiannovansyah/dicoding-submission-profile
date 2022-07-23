package main

import (
	"fmt"
	"log"
	"net/http"
)

// func ReadStaticAsset(w http.ResponseWriter, r *http.Request) {
// 	data, err := ioutil.ReadFile("static/")
// 	if err != nil {
// 		http.Error(w, "Couldn't read file", http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
// 	w.Write(data)
// }

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
