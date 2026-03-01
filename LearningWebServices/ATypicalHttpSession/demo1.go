package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("new request received:")
	fmt.Println("Method:", r.Method)
	fmt.Println("URL:", r.URL.Path)
	fmt.Println("Proto:", r.Proto)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello from GO HTTP server")
}

// POST /contact
func contactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	name := r.FormValue("name")

	fmt.Println("Form name:", name)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Received form from %s\n", name)
}

// 404 handler
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/notfound", notFoundHandler)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
