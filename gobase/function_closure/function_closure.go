package main

import "fmt"

func fibonacci() func(int) int {
	var base int = 0

	
    return func(x int) int {

		defer func() {
            base = x
		}()

		return base + x
	}

}


func main() {
    fi := fibonacci()
    var current = 1

	for i:= 0; i < 20; i++ {
        fmt.Printf("%d ", current)
		current = fi(current)
	}
	fmt.Println()
}
