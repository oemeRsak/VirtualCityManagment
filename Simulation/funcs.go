package main

import (
	"log"
	"math"
	"strconv"
)

//"log"

func (vhc *veheicle) Start() {
	directions := [2]bool{vhc.destionation[0] > vhc.position[0], vhc.destionation[1] > vhc.position[1]}
	distances := [2]int{}

	if directions[0] {
		distances[0] = vhc.destionation[0] - vhc.position[0]

	} else {
		distances[0] = vhc.position[0] - vhc.destionation[0]
	}

	if directions[1] {
		distances[1] = vhc.destionation[1] - vhc.position[1]

	} else {
		distances[1] = vhc.position[1] - vhc.destionation[1]
	}

	if math.Abs(float64(distances[0])) > math.Abs(float64(distances[1])) {
		vhc.direction = "x"

	} else {
		vhc.direction = "y"
	}

	for {
		select {
		case <-Ticker.C:
			if vhc.direction == "x" {
				if directions[0] {
					vhc.position[0] = vhc.position[0] + 1
				} else {
					vhc.position[0] = vhc.position[0] - 1
				}

			}

			if vhc.direction == "y" {
				if directions[1] {
					vhc.position[1] = vhc.position[1] + 1
				} else {
					vhc.position[1] = vhc.position[1] - 1
				}

			}

			if vhc.position[0] == vhc.destionation[0] {
				vhc.direction = "y"
			}
			if vhc.position[1] == vhc.destionation[1] {
				vhc.direction = "x"
			}

			if vhc.position == vhc.destionation {
				vhc.communication <- "arrived-" + strconv.Itoa(vhc.id)
				return
			}
			log.Println(vhc.id, vhc.position, vhc.destionation)
		}
	}
	/*
		i := 0
		//log.Printf("Veheicle '%d' started to drive from %d,%d", vhc.id, vhc.position[0], vhc.position[1])
		for {
			select {
			case <-Ticker.C:
				//log.Printf("Veheicle '%d' is driving currently", vhc.id)
				i++
				if i == 10 {
					vhc.communication <- "arrived-" + strconv.Itoa(vhc.id)
					return
				}
			}
		}
	*/
}
