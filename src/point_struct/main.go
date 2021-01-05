package main

import "fmt"

// 구조체 학생
type Student struct {
	name  string
	class int

	sungjuk Sungjuk
}

// 구조체 성적
type Sungjuk struct {
	name  string
	grade string
}

// 기능 성적 출력
func (s *Student) ViewSungjuk() {
	fmt.Println(s.sungjuk)
}

// 기능 성적 변경
func (s *Student) InputSungjuk(name string, grade string) {
	s.sungjuk.name = name
	s.sungjuk.grade = grade

	fmt.Println("[ (name string, grade string) ]--------------------")
	fmt.Printf("var name val : ")
	fmt.Println(name)
	fmt.Printf("var &name val : ")
	fmt.Println(&name)
	fmt.Printf("var grade val : ")
	fmt.Println(grade)
	fmt.Printf("var &grade val : ")
	fmt.Println(&grade)

	fmt.Printf("var s.sungjuk.name val : ")
	fmt.Println(s.sungjuk.name)
	fmt.Printf("var &s.sungjuk.name val : ")
	fmt.Println(&s.sungjuk.name)
	fmt.Printf("var s.sungjuk.grade val : ")
	fmt.Println(s.sungjuk.grade)
	fmt.Printf("var &s.sungjuk.grade val : ")
	fmt.Println(&s.sungjuk.grade)

}

// main
func main() {
	fmt.Println("go ------------------------------")

	//
	var s Student
	fmt.Println("[ var s Student ]--------------------")
	fmt.Printf("var s  val : ")
	fmt.Println(s)
	fmt.Printf("var &s val : ")
	fmt.Println(&s)
	fmt.Printf("var &s.name val : ")
	fmt.Println(&s.name)
	fmt.Printf("var &s.class val : ")
	fmt.Println(&s.class)
	fmt.Printf("var &s.sungjuk.name val : ")
	fmt.Println(&s.sungjuk.name)
	fmt.Printf("var &s.sungjuk.grade val : ")
	fmt.Println(&s.sungjuk.grade)

	s.name = "김철수"
	s.class = 1
	s.sungjuk.name = "영어"
	s.sungjuk.grade = "C"

	fmt.Println("[ var s Student 속성 지정 ]-------------")
	fmt.Printf("var s  val : ")
	fmt.Println(s)
	fmt.Printf("var &s val : ")
	fmt.Println(&s)
	fmt.Printf("var &s.name val : ")
	fmt.Println(&s.name)
	fmt.Printf("var &s.class val : ")
	fmt.Println(&s.class)
	fmt.Printf("var &s.sungjuk.name val : ")
	fmt.Println(&s.sungjuk.name)
	fmt.Printf("var &s.sungjuk.grade val : ")
	fmt.Println(&s.sungjuk.grade)

	s.InputSungjuk("수학", "A")
	s.ViewSungjuk()

	fmt.Println("---------------------------- lang")
}
