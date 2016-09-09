package main

import (
	"errors"
	"fmt"
	"golang/gobase/mylib"
	"strings"
	"time"
	"runtime"
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
	Op   string
	Path string
	Err  error
}

//指定struct PathError （别名pe），implement Error并实现
func (pe PathError) Error() string {
	return pe.Op + " " + pe.Path + ": " + pe.Err.Error()
}

//出错处理例子--

func main() {

	//ascii
	{
		var a = 'a'
        fmt.Println("a = ", a)
		b := a + 'Z' - 'z'
		fmt.Println("b = ", b)
	}
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
			fmt.Println(pos(i), neg(-2*i))
		}
	}

	//错误处理
	{
		err := &PathError{
			Op:   "err Op ",
			Path: "err Path",
			Err:  errors.New("test custom err"),
		}

		fmt.Println(err.Error())
	}

	//strings.HasRrefix  the same usage for HasSuffix
	{
		jiazaiName := "jiazai"
		fmt.Println("has prefix ? ", strings.HasPrefix(jiazaiName, "jia"))
		fmt.Printf("%t\n", strings.HasPrefix(jiazaiName, "jia"))
	}

	//strings.Contains
	{
		jiazaiName := "jiazai"
		fmt.Println("jiazai contain az? ", strings.Contains(jiazaiName, "az"))
	}

	//strings.Index
	{
		jiazaiName := "jiazai"
		fmt.Println("which index jiazai contain az? ", strings.Index(jiazaiName, "az"))
	}

	//strings.Fields will devide string to array by empty char(\n \t space,etc)
	//strings.Split will devide string to array by arguement
	//string.Join will link array to string by arguement
	{
		str := "The quick brown fox jumps over the lazy dog"
		sl := strings.Fields(str)
		fmt.Printf("Splitted in slice: %v\n", sl)
		for _, val := range sl {
			fmt.Printf("%s - ", val)
		}
		fmt.Println()
		str2 := "GO1|The ABC of Go|25"
		sl2 := strings.Split(str2, "|")
		fmt.Printf("Splitted in slice: %v\n", sl2)
		for _, val := range sl2 {
			fmt.Printf("%s - ", val)
		}
		fmt.Println()
		str3 := strings.Join(sl2, ";")
		fmt.Printf("sl2 joined by ;: %s\n", str3)
	}

    where := func() {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("get where %s:%d\n", file, line)
	}

	{
		fmt.Println("How to use time")
		var week time.Duration
		t := time.Now()
		where()
		fmt.Println(t) // e.g. Wed Dec 21 09:52:14 +0100 RST 2011
		fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
		// 21.12.2011
		t = time.Now().UTC()
		fmt.Println(t)          // Wed Dec 21 08:52:14 +0000 UTC 2011
		fmt.Println(time.Now()) // Wed Dec 21 09:52:14 +0100 RST 2011
		// calculating times:
		week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
		week_from_now := t.Add(week)
		fmt.Println(week_from_now) // Wed Dec 28 08:52:14 +0000 UTC 2011
		// formatting times:
		where()
		fmt.Println(t.Format(time.RFC822))         // 21 Dec 11 0852 UTC
		fmt.Println(t.Format(time.ANSIC))          // Wed Dec 21 08:56:34 2011
		fmt.Println(t.Format("02 Jan 2006 15:04")) // 21 Dec 2011 08:52
		s := t.Format("20060102")
		fmt.Println(t, "=>", s)
	}
}
