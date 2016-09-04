package mylib

import "fmt"

//this func is used for get type UNFINISHED!!!!
func GetType(args ...interface{})  {
	for _, v := range args {
		switch v.(type) {
		case int32:
			fmt.Println("This is a int32")
		case int64:
			fmt.Println("This is a int64")
		case float32:
			fmt.Println("This is a float32")
		case float64:
			fmt.Println("This is a float64")
		case string:
			fmt.Println("This is a string")
		default:
			fmt.Println("unknown type")
		}
	}
}
