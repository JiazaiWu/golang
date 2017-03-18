package main

import (
	"fmt"
	"strconv"
)

type Wing interface {
	Fly() bool
	Drop()
}

//class Bird start
type Bird struct {
	nTime int
}

func (b *Bird) Fly() bool {
	b.nTime--
	if b.nTime >= 0 {
		fmt.Println("Brid fly ok! rest " + strconv.Itoa(b.nTime) + " times")
		return true
	}

	fmt.Println("Bird fly bad! rest " + strconv.Itoa(b.nTime) + " times")
	return false
}

func (b Bird) Drop() {
	b.nTime++
	fmt.Println("Brid dropped")
}

//class Brid end

func main() {
	//var b1 = Bird{10}
	//var myWing1 Wing = b1
	var myWing2 Wing = new(Bird)

	//myWing1.Fly()
	myWing2.Fly()
}
