package routing

import "net/http"

// HomePageHandler function is used in main.go
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(`{\"message\" : \"Starting application\"}`))
}
