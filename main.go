package main

import (
//"github.com/pkg/term"
"fmt"
"log"
"os"
"net/http"
"github.com/stianeikeland/go-rpio"
"github.com/gorilla/mux"
)

var pin1, pin2, pin3, pin4, pin5, pin6, pin7, pin8 rpio.Pin
const interval = 30
var lit bool

func main(){
        rpio.Open()
        defer rpio.Close()
        Initiate()
        
        router := mux.NewRouter().StrictSlash(true)
	    router.HandleFunc("/", index)
		router.HandleFunc("/w", serve_forward)
		router.HandleFunc("/s", serve_backwards)
		router.HandleFunc("/a", serve_rotate_left)
		router.HandleFunc("/d", serve_rotate_right)
		router.HandleFunc("/lightson", serve_lightson)
		router.HandleFunc("/lightsoff", serve_lightsoff)
		router.HandleFunc("/stop", stop)

		log.Fatal(http.ListenAndServe(":8080", router))
}

func stop(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Stopping backend webserver")
    Lightsoff()
    os.Exit(0)
}

func serve_lightson(w http.ResponseWriter, r *http.Request) {
	Lightson()
    fmt.Fprintln(w, "Turned on lights")
}

func serve_lightsoff(w http.ResponseWriter, r *http.Request) {
	Lightsoff()
    fmt.Fprintln(w, "Turned off lights")
}

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Robotic platform backend")
}

func serve_forward(w http.ResponseWriter, r *http.Request) {
	Forward(interval)
    fmt.Fprintln(w, "Moved forward")
}

func serve_backwards(w http.ResponseWriter, r *http.Request) {
	Backwards(interval)
    fmt.Fprintln(w, "Moved backwards")
}

func serve_rotate_left(w http.ResponseWriter, r *http.Request) {
	Rotate_left(interval)
    fmt.Fprintln(w, "Moved Left")
}

func serve_rotate_right(w http.ResponseWriter, r *http.Request) {
	Rotate_right(interval)
    fmt.Fprintln(w, "Moved right")
}
