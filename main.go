package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/sleep", wait)
	log.Fatal(http.ListenAndServe(":4321", nil))

}

func wait(w http.ResponseWriter, r *http.Request) {
	seconds, _ := strconv.Atoi(r.URL.Query().Get("seconds"))
	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Fprintf(w, "Slept for %v seconds\n", seconds)

}
