package main

import "fmt"

// 구조체 사람
type Person struct {
	name string
	age  int
}

// 기능 사람 이름 출력
func (p Person) PrintName() {
	fmt.Println(p.name)
}

// main
func main() {
	fmt.Println("go ---------------")

	var p Person
	p1 := Person{"개쏭이", 15}
	p2 := Person{name: "소똥이", age: 21}
	p3 := Person{name: "Carson"}
	p4 := Person{}

	fmt.Println(p, p1, p2, p3, p4)

	p.name = "기술자"
	p.age = 56

	fmt.Println(p)

	// 
	p.PrintName()

	fmt.Println("------------- lang")
}
