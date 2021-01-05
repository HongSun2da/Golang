package main

import "fmt"

func main() {
	var s []int
	var t []int
	var u []int

	s = make([]int, 3)
	s[0] = 100
	s[1] = 200
	s[2] = 300
	fmt.Println(s, len(s), cap(s))

	s = append(s, 400)
	fmt.Println(s, len(s), cap(s))

	t = append(s, 500)
	fmt.Println(t, len(t), cap(t))

	u = append(t, 600)
	fmt.Println(u, len(u), cap(u))

	fmt.Println("----------------")

	u[0] = 9999
	fmt.Println(s, len(s), cap(s))
	fmt.Println(t, len(t), cap(t))
	fmt.Println(u, len(u), cap(u))
}
