package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func renderHomePage(w http.ResponseWriter, r *http.Request) {

	pageData := struct {
		Devices []Device
	}{
		Devices: appData.Devices,
	}
	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, pageData)
	log.Println("Renedered the home page.")

}

func redirectToHomePage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusFound)
}

func checkHealth(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, "alive")

}
