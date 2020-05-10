package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

/*
// net/http/pprof包的init已经定义了几个http endpoint页面
func init() {
	http.HandleFunc("/debug/pprof/", Index)
	http.HandleFunc("/debug/pprof/cmdline", Cmdline)
	http.HandleFunc("/debug/pprof/profile", Profile)
	http.HandleFunc("/debug/pprof/symbol", Symbol)
	http.HandleFunc("/debug/pprof/trace", Trace)
}
*/

var datas []string

func Add(str string) string {
	data := []byte(str)
	sData := string(data)
	datas = append(datas, sData)

	return sData
}

func main() {
	// 注意起了一个协程一直写全局，消耗内存会比较多
	go func() {
		for {
			log.Println(Add("https://github.com/EDDYCJY"))
		}
	}()

	// 实现一个http web真容易。。。
	http.ListenAndServe("0.0.0.0:6060", nil)
}
