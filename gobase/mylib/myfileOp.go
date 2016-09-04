package mylib

import (
	"os"//文件操作
	"fmt"
	"bufio"//字符串
	"io"//获取输入
	"strconv"//string操作
)

func ReadValues(infile string)(values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open the input file ", infile)
		return
	}
	defer file.Close()
	bufferReader := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix, err1 := bufferReader.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			return
		}
		str := string(line) // 转换字符数组为字符串
		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)
	}
	return
}
