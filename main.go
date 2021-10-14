package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Global variables
var appConfig AppConfig
var appData AppData

func main() {
	setWorkingDir()
	loadData()
	setupWebServer()
}

func setWorkingDir() {
	thisApp, err := os.Executable()
	if err != nil {
		log.Fatalf("Error determining the directory. \"%s\"", err)
	}
	appPath := filepath.Dir(thisApp)
	os.Chdir(appPath)
	log.Printf("Set working directory: %s", appPath)
}

func setupWebServer() {

	// Init HTTP Router - mux
	router := mux.NewRouter()

	// map directory to server static files
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Define Home Route
	router.HandleFunc("/", renderHomePage).Methods("GET")

	// Define Wakeup functions with a Device Name
	router.HandleFunc("/wake/{deviceName}", wakeUpWithDeviceName).Methods("GET")
	router.HandleFunc("/wake/{deviceName}/", wakeUpWithDeviceName).Methods("GET")

	// Define Data save Api function
	router.HandleFunc("/data/save", saveData).Methods("POST")

	// Define Data get Api function
	router.HandleFunc("/data/get", getData).Methods("GET")

	// Define health check function
	router.HandleFunc("/health", checkHealth).Methods("GET")

	// Setup Webserver
	httpListen := ":8089"
	log.Printf("Startup Webserver on \"%s\"", httpListen)

	srv := &http.Server{
		Handler: handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(router),
		Addr:    httpListen,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
