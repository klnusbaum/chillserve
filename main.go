package main

import (
	"net/http"
	"github.com/klnusbaum/chillserve/handlers"
)

func main() {

	ch := handlers.NewRandomChillHandler("chill", "super chill", "chilly freeze")

	http.Handle("/chill", ch)
	http.ListenAndServe(":8080", nil)
}
