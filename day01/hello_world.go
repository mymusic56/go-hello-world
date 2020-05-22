package main

import (
	"fmt"
	"strconv"
)

//定义常量，不能被修改
const LENGTH = 10
const LIMIT_AGE int = 100
//iota: 在第一行会被置为0， 新增一行iota值加一
const (
	A1 = 1
	A2 = iota
	A3
	A4
	B1 = "zhangsn"
	B2 = iota
	B3
)

//代码格式化： ctrl + alt + L

func main() {
	fmt.Print("hell world\r\n")

	//变量声明， 声明时赋值，类型可选
	var a string = "zhangsan"
	var b = "12345"
	var c bool
	var d int
	var e string
	fmt.Println("B： " + b)
	b = "lisi"
	//使用 := 进行变量声明， 声明并赋值，无需使用关键字var、类型， 系统自动推断类型
	f := "wangwu"
	//多变量声明
	var g, h, i int
	var j, k int = 1, 2
	var l, m = 123, "hello"
	n, o := "hello world", 12

	fmt.Println("A： " + a)
	fmt.Println("B： " + b)
	//boole转字符串
	fmt.Println("C： " + strconv.FormatBool(c))
	//数字转字符串
	fmt.Println("D： " + fmt.Sprintf("%d", d))
	fmt.Println("f： " + f)
	//字符串拼接
	fmt.Println("e:" + e)
	d = 10
	fmt.Println("a + d = " + a + strconv.Itoa(d))
	fmt.Print("a的地址：")
	fmt.Println(&a)
	a2 := a
	fmt.Println(&a2)

	g = 20
	h = 20
	i = 20
	fmt.Print(g + h + i)
	fmt.Print(j + k)
	fmt.Print(l, m, n, o)

	fmt.Println(LENGTH)
	fmt.Println(LIMIT_AGE)
	fmt.Println("--------------------")

	fmt.Print(A1, A2, A3, A4, B1, B2, B3)
}
