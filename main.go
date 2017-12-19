package main

import (
	"fmt"
	"os"
	"time"
)

func playSoundAfter(t time.Duration) {
	time.Sleep(t)
	fmt.Printf("%v passed\n", t)
}

func parseTime(t string) time.Duration {
	parsed, err := time.ParseDuration(t)
	if err != nil {
		fmt.Printf("Invalid time format\n")
		os.Exit(1)
	}
	return parsed
}

func main() {
	// time is the first argument
	//e.g: 10h, 20s, 30m
	p := parseTime(os.Args[1])
	playSoundAfter(p)
}
