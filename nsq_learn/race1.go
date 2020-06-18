package main

import (
	"log"
	"time"
)

func main() {
	a := 1
	go func() {
		a = 2
	}()
	a = 3
	log.Printf("a:%v", a)
	time.Sleep(time.Second)
}
