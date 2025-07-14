package nethttp

import (
	"fmt"
	"net/http"
)

// fungsi handler function untuk routing `/`
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome di bellshade: home hadler\n")
}

// fungsi handler function untuk routing `/about`
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "about: bellshade golang \n")
}

// middleware sederhana yang digunakan untuk logging
// setiap request
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("request diterima: %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

func main() {
	http.HandleFunc("/", loggingMiddleware(homeHandler))
	http.HandleFunc("/about", loggingMiddleware(aboutHandler))
	fmt.Println("server berjalan di http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
