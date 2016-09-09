package main

import "fmt"
import "strings"

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
	    if !strings.HasSuffix(name, suffix) {
	            return name + suffix
	    }
	    return name
	}
}

func main() {
    addBmp := MakeAddSuffix(".bmp")
    addJpeg := MakeAddSuffix(".jpeg")
	str1 := addBmp("file") // returns: file.bmp
	str2 := addJpeg("file") // returns: file.jpeg

	fmt.Println(str1)
	fmt.Println(str2)
}
