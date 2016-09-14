package main

import "fmt"

func insertStringSlice(des, src []int, index int) []int{
	if len(des) < index {
         return nil
	}

	if index < 0 {
		return nil
	}

    s1 := des[:index-1]
	s2 := des[index-1:]
	
	return append(s1, append(src, s2...)...)
}

func main() {
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)

	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)

	s1 := []int{1, 2, 5, 6}
	s2 := []int{3, 4}
    s3 :=insertStringSlice(s1, s2, 3)
	
	fmt.Println(s3)
}
