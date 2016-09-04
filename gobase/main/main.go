package main

import (
	"fmt"
	"golang/gobase/mylib"
	"errors"
)
//闭包例子++ addr 函数名，()参数，func(int) int 返回值；也就是说会返回一个func(int) int 的function。。。
func adder() func(int) int {
	sum := 0
	innerfunc := func(x int) int {
		//当调用处得到innerfunc，并使用时，sum会变化
		sum += x
		return sum
	}
	return innerfunc
}
//闭包例子--

//出错处理例子++
type PathError struct {
	Op string
	Path string
	Err error
}

//指定struct PathError （别名pe），implement Error并实现
func (pe PathError) Error() string {
	return pe.Op + " " + pe.Path + ": " + pe.Err.Error()
}
//出错处理例子--

func main(){

	//org 中元素发生改变 arrayint中的元素也会发生改变===>go的数组切片其实是共享内存
	{
		var org [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		var arrayint []int = org[:9]
		org[1] = 0
		for _, j := range arrayint {
			fmt.Print(j, " ")
		}
		fmt.Println(org)
	}

	//测试args ...interface{}
	{
		var v1 int32 = 1
		var v2 int64 = 2
		var v3 float32 = 1.1
		var v4 float64 = 2.2
		var v5 string = "test"
		mylib.GetType(v1, v2, v3, v4, v5)
	}

	//函数闭包
	{
		pos, neg := adder(), adder()
		for i := 0; i < 10; i++ {
			fmt.Println(pos(i), neg(-2 * i))
		}
	}

	//错误处理
	{
		err := &PathError{
			Op: "err Op ",
			Path: "err Path",
			Err:   errors.New("test custom err"),
		}

		fmt.Println(err.Error())
	}


}
