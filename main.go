package main

import (
	"fmt"
	"time"
	"math"
	"os"
	"strconv"
)

func main() {
	threshold,_ := strconv.Atoi(os.Getenv("TH"))
	before := int64(0)
	c := 0
	for _ = range time.Tick(time.Second * 1) {
		//fmt.Println("Ticking every 3 seconds")
		now :=  time.Now().UnixNano()/1000 // microseconds
		jitter := math.Abs(float64(1000000-(now-before)))
		//fmt.Printf("%.0f\n", jitter )
		if (int(jitter) > threshold) {
			fmt.Printf(".")
		}
		before = now
		if (c > 60) {
			fmt.Printf("\n")
			c = 0
		}
		c++
	}
}
