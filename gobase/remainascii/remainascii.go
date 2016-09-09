package main 

import (
	"fmt"
)

func isascii(c rune) bool {
	if c > 127 {
		return false
	}
	return true
}

func getnewstr(str string, f func(rune) bool) string{
	var ret string
    for _, ch := range str {
		if (f(ch)) {
            ret += string(ch) 
		} else {
			fmt.Println(ch, "is replaced")
			ret += " "
		}
	}

	return ret
}


func main() {

    str := "asSASA ddd dsjkdsjsこん dk"

	fmt.Println("len", len(str))

    str1 :=  getnewstr(str, isascii)

    fmt.Println(str1)
}
