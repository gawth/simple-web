package main

import (
    "net/http"
)

func imageHandler(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Cache-Control", "max-age=0, no-cache")
    http.ServeFile(w, r, "image.jpg")
}

func main() {
    http.Handle("/", http.FileServer(http.Dir("./")))
    http.HandleFunc("/img", imageHandler)
    http.ListenAndServe(":8080", nil)
}
