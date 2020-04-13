package main

import (
	"URL-Shortener/handler"
	"net/http"
)

func main() {






	http.ListenAndServe(":3000", http.HandlerFunc(handler.CustomUrls))
}

