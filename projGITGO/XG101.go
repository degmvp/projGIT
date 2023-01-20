package main

import (
	"fmt"
)

type Person interface {
    Work()
}

type worker string

func (w worker) Work(){
	fmt.Printf("%s is working\n", w)
}

func DoWorkInterface(things []Person) {
    for _, v := range things {
        v.Work()
    }
}

func main() {
	var d,e,f worker
	d = "D"
	e = "E"
	f = "F"
	DoWorkInterface([]Person{d,e,f})
}