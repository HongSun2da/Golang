package main

import "fmt"

// main
func main() {
	fmt.Println("go ------------------------------")

	//
	var a int
	fmt.Println("[ var a int ]--------------------")
	fmt.Printf("var a  val : ")
	fmt.Println(a)
	fmt.Printf("var &a val : ")
	fmt.Println(&a)

	var p *int
	fmt.Println("[ var p *int ]--------------------")
	fmt.Printf("var p  val : ")
	fmt.Println(p)
	fmt.Printf("var &p val : ")
	fmt.Println(&p)
	fmt.Printf("var *p val : ")
	fmt.Println("invalid memory address or nil pointer dereference")
	fmt.Println("---------------------------------")

	a = 3
	fmt.Println("[ a = 3 ]--------------------")
	fmt.Printf("var a  val : ")
	fmt.Println(a)
	fmt.Printf("var &a val : ")
	fmt.Println(&a)

	fmt.Printf("var p  val : ")
	fmt.Println(p)
	fmt.Printf("var &p val : ")
	fmt.Println(&p)
	fmt.Printf("var *p val : ")
	fmt.Println("invalid memory address or nil pointer dereference")
	fmt.Println("---------------------------------")

	p = &a
	fmt.Println("[ p = &a ]--------------------")
	fmt.Printf("var a  val : ")
	fmt.Println(a)
	fmt.Printf("var &a val : ")
	fmt.Println(&a)

	fmt.Printf("var p  val : ")
	fmt.Println(p)
	fmt.Printf("var &p val : ")
	fmt.Println(&p)
	fmt.Printf("var *p val : ")
	fmt.Println(*p)
	fmt.Println("---------------------------------")

	var b int
	b = 2
	fmt.Println("[ b = 2 ]--------------------")
	fmt.Printf("var a  val : ")
	fmt.Println(a)
	fmt.Printf("var &a val : ")
	fmt.Println(&a)

	fmt.Printf("var b  val : ")
	fmt.Println(b)
	fmt.Printf("var &b val : ")
	fmt.Println(&b)

	fmt.Printf("var p  val : ")
	fmt.Println(p)
	fmt.Printf("var &p val : ")
	fmt.Println(&p)
	fmt.Printf("var *p val : ")
	fmt.Println(*p)
	fmt.Println("---------------------------------")

	p = &b
	fmt.Println("[ p = &b ]--------------------")
	fmt.Printf("var a  val : ")
	fmt.Println(a)
	fmt.Printf("var &a val : ")
	fmt.Println(&a)

	fmt.Printf("var b  val : ")
	fmt.Println(b)
	fmt.Printf("var &b val : ")
	fmt.Println(&b)

	fmt.Printf("var p  val : ")
	fmt.Println(p)
	fmt.Printf("var &p val : ")
	fmt.Println(&p)
	fmt.Printf("var *p val : ")
	fmt.Println(*p)
	fmt.Println("---------------------------------")

	fmt.Println("---------------------------- lang")
}
