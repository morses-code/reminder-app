package server

import (
	"net/http"

	"github.com/morses-code/reminder-app/api"
)

func Server() {
	http.HandleFunc("/hello", api.Hello)
	http.ListenAndServe(":8080", nil)
}
