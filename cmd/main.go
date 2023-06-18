package main

import (
	"go_learning/sometest"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

type IntSlice []int

func (a IntSlice) Len() int {
	return len(a)
}

func (a IntSlice) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a IntSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	//sometest.Neo4jTest()
	//sometest.DataBendTest()
	sometest.MapTest()
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	runtime.Gosched()
	select {}
}
