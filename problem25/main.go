package main

import (
	"fmt"
	"time"
)

/*
	Реализовать собственную функцию sleep.
*/
//done

func mySleep(duration time.Duration) {
	<-time.After(duration)
}

func mySleep2(duration time.Duration) {
	start := time.Now()
	for time.Since(start) < duration {
	}
}

func main() {
	fmt.Println("start")
	mySleep(1 * time.Second)
	fmt.Println("=1=")
	mySleep2(1 * time.Second)
	fmt.Println("finish")
}
