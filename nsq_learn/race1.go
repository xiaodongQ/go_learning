package main

import (
	"log"
)

func main() {
	a := 1
	go func() {
		a = 2
	}()
	a = 3
	log.Printf("a:%v", a)
	// time.Sleep(time.Second)
	log.Printf("=============")
	hashtable()
}

func hashtable() {
	m := make(map[int]string)
	m[1] = "123"
	m[4] = "xxx"
	log.Printf("m:%v", m)
}
