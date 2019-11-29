package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"
)

func main() {
	http.Handle("/assets/",
		http.StripPrefix("/assets/",
			http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", Index)
	fmt.Println("server started at localhost:5000")
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	filepath := path.Join("views", "index.html")
	tmpl, err := template.ParseFiles(filepath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
