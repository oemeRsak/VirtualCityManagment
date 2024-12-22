package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func hi(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hi\n")
}

func hello(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():

		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func status(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		if name == "For" {
			for _, h := range headers {
				switch h {

				//curl 127.0.0.1:8090/status -H "For:veheicles"
				case "veheicles":
					log.Println("Requested status ")
					fmt.Fprint(w, Veheicles)
					fmt.Fprint(w, strconv.Itoa(len(Veheicles)))
				}
			}
		}
	}
}
