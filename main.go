package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	m := http.NewServeMux()

	var n int64
	n = 0

	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n++
		w.Write([]byte("request number: " + fmt.Sprintf("%d", n)))
		
		if n == 10 {
			log.Printf("request number: %d", n)
		}

		if n == 1000000000 {
			n = 0
		}
	})
	
	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", m)
}