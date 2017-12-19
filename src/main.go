package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	// HTTP
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)


	// HTTPS
	// err := http.ListenAndServeTLS(“:8080, fullchain.pem, privkey.pem, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
