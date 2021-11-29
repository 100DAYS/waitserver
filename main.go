package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello. Call /wait?time=n to wait n seconds. If no time given, random is used."))
	})
	r.Get("/wait", func(w http.ResponseWriter, r *http.Request) {
		t := r.URL.Query().Get("time")
		var delay int
		if t == "" {
			delay = 2 + rand.Intn(30)
		} else {
			var err error
			delay, err = strconv.Atoi(t)
			if err != nil {
				w.Write([]byte("could not parse time: " + t))
			}
		}
		time.Sleep(time.Duration(delay) * time.Second)
		w.Write([]byte(fmt.Sprintf("waited %d seconds", delay)))
	})
	http.ListenAndServe(":3000", r)
}
