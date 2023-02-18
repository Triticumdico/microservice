package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello Word")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Oops", http.StatusBadRequest)
			return
		}
		fmt.Printf("Hello %s", d)
		fmt.Fprintf(w, "Hello %s", d)
	})
	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye Word")
	})

	http.ListenAndServe(":9000", nil)
}
