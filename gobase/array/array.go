package main 

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4}
	arr2 :=  new([4]int)

	fmt.Println(arr2)

	arr3 := arr1

	fmt.Println(arr3)

	arr3[1] = 5

	fmt.Println("arr2 = ", arr3)
	fmt.Println("arr1 = ", arr1)

	fmt.Printf("%p\n", arr3)
	fmt.Printf("%p\n", arr1)
}
