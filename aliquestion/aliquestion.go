package main

import (
	"fmt"
	"time"
	"math/rand"
)

const (
	member int = 20
	question int = 20
)

func initquestions(quesformem *[member][question]int, quesnum int) {
    
    r :=  rand.New(rand.NewSource(time.Now().UnixNano())) //time as seed to get random

	for i:= 0; i < member; i++ {
		for j:= 0; j < question; j++ {

			for {
                flag := true
	            var temp int = r.Intn(quesnum)

			    for k:= 0; k < j; k++ {
					if quesformem[i][k] == temp {
                        flag = false
					}
				}
                

				if flag {
					quesformem[i][j] = temp
					break
				}
			}
		}
	}
}

func checkquestion(question1, question2 [question]int) bool {
	var count int = 0

	for i:= 0; i < question; i++ {
		for j:= 0; j < question; j++ {
			if question1[i] == question2[j] {
				count++

		        if count > 3 {
				    return true
				}

				break
			}
		}
	}

	return false
}

func checkquestions(quesformem [member][question]int) bool {

	for i := 0; i < member; i++ {
		for j := i+1; j < member; j++ {
            if checkquestion(quesformem[i], quesformem[j]) {
				return false
			}
		}
	}

    return true
}


func main() {

	//set array 20 members of 20 questions
	var quesformem [member][question]int
    
    // questions add by 50
	for i:= 500; i < 2000; i += 50 {

		fmt.Println("question number is ", i)

		var count int = 0
	    for j := 0; j < 1000; j++ {

            initquestions(&quesformem, i)
			 
			if checkquestions(quesformem) {
				count++
			}

		}

		fmt.Println("count is", count)

	}
}
