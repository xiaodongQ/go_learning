package main

import (
	"log"
	"math"
	"os"
	"runtime/pprof"
	"time"
)

/*
 Get a iso time
  eg: 2018-03-16T18:02:48.284Z
*/
func IsoTime() string {
	utcTime := time.Now().UTC()
	iso := utcTime.String()
	isoBytes := []byte(iso)
	iso = string(isoBytes[:10]) + "T" + string(isoBytes[11:23]) + "Z"
	log.Printf("value:%f\n", math.Sqrt(200))
	return iso
}

func main() {
	log.Println("start profiling...")
	file, err := os.Create("./cpu.profile")
	if err != nil {
		log.Printf("create file err![%v]\n", err)
		return
	}
	defer file.Close()

	pprof.StartCPUProfile(file)
	defer pprof.StopCPUProfile()

	for i := 0; i < 1000; i++ {
		log.Printf("isotime:%v\n", IsoTime())
		log.Print("=============")
	}

	heapfile, err := os.Create("./mem.profile")
	if err != nil {
		log.Fatal("create error!\n")
	}
	defer heapfile.Close()
	pprof.WriteHeapProfile(heapfile)

}

func init() {
	// log.SetFlags(log.Lshortfile)
}
