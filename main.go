package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func getParam(values url.Values, name string, missingValue int) (res int, err error) {
	if values.Has(name) {
		res, err = strconv.Atoi(values.Get(name))
		if err != nil {
			err = fmt.Errorf("Error in param %s: %s", name, err)
		}
		return
	}
	return missingValue, nil
}

func initThrottleOpts() (throttle bool, opts middleware.ThrottleOpts, err error) {
	throttle = false
	l := os.Getenv("THROTTLE_LIMIT")
	if l != "" {
		throttle = true
		opts.Limit, err = strconv.Atoi(l)
		if err != nil {
			return
		}
	}
	bl := os.Getenv("THROTTLE_BACKLOG_LIMIT")
	if bl != "" {
		throttle = true
		opts.BacklogLimit, err = strconv.Atoi(bl)
		if err != nil {
			return
		}
	}
	bt := os.Getenv("THROTTLE_BACKLOG_TIMEOUT")
	if bt != "" {
		throttle = true
		var backlogTimeout int
		backlogTimeout, err = strconv.Atoi(bt)
		if err != nil {
			return
		}
		opts.BacklogTimeout = time.Duration(backlogTimeout) * time.Second
	}
	return
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	throttle, throttleOpts, err := initThrottleOpts()
	if err != nil {
		panic(err)
		return
	}
	if throttle {
		fmt.Printf("Starting with Throttle: %#v", throttleOpts)
		r.Use(middleware.ThrottleWithOpts(throttleOpts))
	}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello. Call /wait?time=n to wait n seconds. Or use min and max to define range. If no time or range given, range 2-30 is used."))
	})
	r.Get("/wait", func(w http.ResponseWriter, r *http.Request) {
		delay, err := getParam(r.URL.Query(), "time", 0)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		min, err := getParam(r.URL.Query(), "min", 2)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		max, err := getParam(r.URL.Query(), "max", 30)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		if delay == 0 {
			delay = min + rand.Intn(max-min)
		}
		time.Sleep(time.Duration(delay) * time.Second)
		w.Write([]byte(fmt.Sprintf("waited %d seconds", delay)))
	})
	http.ListenAndServe(":3000", r)
}
