package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	port := "5432"
	if p, ok := os.LookupEnv("PORT"); ok {
		port = p
	}
	log.Printf("Starting at http://:%v?s=1", port)
	http.HandleFunc("/", wait)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func wait(w http.ResponseWriter, r *http.Request) {
	seconds, _ := strconv.Atoi(r.URL.Query().Get("s"))
	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Fprintf(w, "Slept for %v seconds\n", seconds)
}
