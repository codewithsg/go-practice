package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, res *http.Request) {
		log.Println("Hello World!")
		d, err := io.ReadAll(res.Body)
		if err != nil {
			http.Error(rw, "Oops error has occured", http.StatusBadRequest)
			return
		}
		// log.Printf("Data %s", d)
		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println(("Bye World!"))
	})

	http.ListenAndServe(":7373", nil)
}
