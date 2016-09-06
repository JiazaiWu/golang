package main

import (
	"fmt"
)


type Base struct {
    Name string
}

func (base *Base) foo() {
    fmt.Println("this is Base!")
}

func (base *Base) getName() {
    fmt.Println("My name is" + base.Name)
}

type Foo struct {
    Base
}

func (foo *Foo) foo() {
    fmt.Println("this is Foo!")
	foo.Base.foo()
}

func main() {
   myFoo := &Foo{Name:"jiazai"}
   myFoo.foo()
   myFoo.getName()
}
