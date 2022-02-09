package main

import (
	"flag"
	"fmt"
	"time"

	//"github.com/stianeikeland/go-rpio"
	"github.com/drahcirennobran/go-rpio-mock"
)

const (
	pin = rpio.Pin(13) //33
)

func main() {
	delay := flag.Int("delay", 100, "delay in ms")
	rep := flag.Uint64("rep", 11, "repetitions")
	flag.Parse()
	fmt.Println(*delay)
	fmt.Println(*rep)
	for i := uint64(0); i < *rep; i++ {
		fmt.Printf(".")
		pin.High()
		time.Sleep(time.Microsecond * time.Duration(*delay))
		pin.Low()
		time.Sleep(time.Microsecond * time.Duration(*delay))
	}
}
