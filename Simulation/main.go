package main

import (
	"flag"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var Ticker *time.Ticker

var Veheicles []*veheicle

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	debug_file, _ := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	brodcast_file, _ := os.OpenFile("brodcast.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	debug = log.New(debug_file, "debug:", log.LstdFlags)
	broadcastLog = log.New(brodcast_file, "brodcast:", log.LstdFlags)
	debug.Println("Hi")
	Ticker = time.NewTicker(time.Millisecond * 100)
	veheicle_number := flag.Int("veheicle_number", 5, "Number of veheicles")

	flag.Parse()

	log.Println("Hello World !")

	veheicles_com := make(chan string)

	// create veheicle object and append to the list
	for i := 0; i < *veheicle_number; i++ {
		Veheicles = append(Veheicles, &veheicle{i, "car" + strconv.Itoa(i), veheicle_types.car, veheicle_types.car.priority, [2]int{rand.IntN(100), rand.IntN(100)}, [2]int{rand.IntN(100), rand.IntN(100)}, "0", veheicles_com})
	}

	// turn them on
	for i := 0; i < len(Veheicles); i++ {
		v := Veheicles[i]
		debug.Printf("Veheicle '%d' is starting to from %d, %d to %d, %d", v.id, v.position[0], v.position[1], v.destionation[0], v.destionation[1])
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
				debug.Printf("Veheicle '%s' is arrived to destionation", ms[1])
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
	http.HandleFunc("/ws", handleWebSocket)


	go periodicBroadcast()
	http.ListenAndServe(":8090", nil)
	}

}
