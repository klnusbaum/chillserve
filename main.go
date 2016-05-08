package main

import (
	"net/http"
)

func main() {

	ch := ChillHandler{
		ChillPhrases: []string{"chill", "super chill", "chilly freeze"},
	}

	http.Handle("/chill", ch)
	http.ListenAndServe(":8080", nil)
}
