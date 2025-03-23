package handler

import (
	"embed"
	"fmt"
	"net/http"
)

var StaticFiles embed.FS

func Custom404Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<html><head><title>404 not found</title></head<body><h1>404 - Page Not Found</h1><p>Sorry, we couldn't find the page you're looking for.</p></body></html>")
}

func GetFavicon(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/favicon.ico" {
		Custom404Handler(w, r)
		return
	}
	data, _ := StaticFiles.ReadFile("static/favicon.ico")
	w.Header().Set("Content-Type", "image/x-icon")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
