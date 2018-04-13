package main

import (
	"log"
	"net/http"
)

var cache map[string]string

func main() {
	cache = make(map[string]string, 4096)

	http.HandleFunc("/add", addWebHandler)
	http.HandleFunc("/query", queryWebHandler)
	http.HandleFunc("/remove", removeWebHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func addWebHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")
	if len(key) == 0 || len(value) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
	}
	cache[key] = value
	w.WriteHeader(http.StatusOK)
}

func queryWebHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if len(key) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
	}
	if val, ok := cache[key]; ok {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(val))
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func removeWebHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if len(key) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
	}
	if _, ok := cache[key]; ok {
		delete(cache, key)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
