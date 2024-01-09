package main

import (
	"net/http"

	"github.com/andrewvo148/orders-manage/internal/rest"
)

func main() {
	// Start the HTTP server on port 8080
	http.ListenAndServe(":8080", rest.Run())

}
