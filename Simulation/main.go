package main

import (
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var Ticker *time.Ticker

var Veheicles []*veheicle

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	Ticker = time.NewTicker(time.Millisecond * 100)

	log.Println("Hello World !")

	veheicles_com := make(chan string)

	// create veheicle object and append to the list
	for i := 0; i < 5; i++ {
		Veheicles = append(Veheicles, &veheicle{i, "car" + strconv.Itoa(i), veheicle_types.car, veheicle_types.car.priority, [2]int{rand.IntN(100), rand.IntN(100)}, [2]int{rand.IntN(100), rand.IntN(100)}, "0", veheicles_com})
	}

	// turn them on
	for i := 0; i < len(Veheicles); i++ {
		v := Veheicles[i]
		go v.Start()
	}

	// wait for com
	go func() {
		for {
			m := <-veheicles_com
			ms := strings.Split(m, "-")

			switch ms[0] {
			// if a arrived text
			case "arrived":
				log.Printf("Veheicle '%s' is arrived to destionation", ms[1])
				vhc_id, _ := strconv.Atoi(ms[1])

				// find it in list and remove
				for i, v := range Veheicles {
					if v.id == vhc_id {
						Veheicles = append(Veheicles[:i], Veheicles[i+1:]...)
					}
				}

			}
		}
	}()

	http.HandleFunc("/hi", hi)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/status", status)
	http.HandleFunc("/headers", headers)

	//http.ListenAndServe(":8090", nil)

	for { //this loop will spin, using 100% CPU (SA5002)
		//wait wait wait
	}

}
