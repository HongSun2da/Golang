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
func (s Student) ViewSungjuk() {
	fmt.Println(s.sungjuk)
}

// 기능 성적 변경
func (s Student) InputSungjuk(name string, grade string) {
	s.sungjuk.name = name
	s.sungjuk.grade = grade
}

// 일반 기능
func ViewSungjuk(s Student) {
	fmt.Println(s.sungjuk)
}

// main
func main() {
	fmt.Println("go ------------------------------")

	//
	s := Student{name: "서니", class: 2, sungjuk: Sungjuk{name: "국어", grade: "A"}}

	s.ViewSungjuk()

	//
	var ss Student
	ss.name = "김철수"
	ss.class = 1
	ss.sungjuk.name = "영어"
	ss.sungjuk.grade = "C"

	ss.ViewSungjuk()
	ViewSungjuk(ss)

	// 성적변경
	ss.InputSungjuk("지리", "D")
	ss.ViewSungjuk()

	fmt.Println("---------------------------- lang")
}
