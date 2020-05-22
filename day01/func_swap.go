package main

import "fmt"
/*
交换值
 */
func swap(a int, b int) (int, int) {

	tmp := a
	a = b
	b = tmp
	return a, b
}

func main() {
	a := 1
	b := 2
	a, b = swap(a, b)
	fmt.Println(a, b)

	c := 10
	d := 20
	fmt.Println(c, d)
	c, d = d, c
	fmt.Println(c, d)
}
