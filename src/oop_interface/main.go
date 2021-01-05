package main

import "fmt"

type SpoonOfJam interface {
	String() string
}

// interface 선언 //////////////////////////////////////////////////////////
type Jam interface {
	GetOneSpoon() SpoonOfJam
}

// Bread struct ////////////////////////////////////////////////////////////
type Bread struct {
	name string
}

func (b *Bread) PutJam(jam Jam) {
	spoon := jam.GetOneSpoon()
	b.name += " + " + spoon.String()
}

func (b *Bread) String() string {
	return b.name
}

// StrawberryJam struct ////////////////////////////////////////////////////
type StrawberryJam struct {
	name string
}

func (j *StrawberryJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfStrawberryJam{name: "Spoon Of Strawberry Jam"}
}

// SpoonOfStrawberryJam struct ////////////////////////////////////////////
type SpoonOfStrawberryJam struct {
	name string
}

func (s *SpoonOfStrawberryJam) String() string {
	return s.name
}

// OrangeJam struct ////////////////////////////////////////////////////
type OrangeJam struct {
	name string
}

func (j *OrangeJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfOrangeJam{name: "Spoon Of Orange Jam"}
}

// SpoonOfStrawberryJam struct ////////////////////////////////////////////
type SpoonOfOrangeJam struct {
	name string
}

func (s *SpoonOfOrangeJam) String() string {
	return s.name
}

// AppleJam struct ////////////////////////////////////////////////////
type AppleJam struct {
	name string
}

func (j *AppleJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfAppleJam{name: "Spoon Of Apple Jam"}
}

// SpoonOfAppleJam struct ////////////////////////////////////////////
type SpoonOfAppleJam struct {
	name string
}

func (s *SpoonOfAppleJam) String() string {
	return s.name
}

// main
func main() {
	fmt.Println("go ---------------")

	bread := &Bread{name: "Bread 1"}
	jam := &StrawberryJam{name: "Strawberry Jam 1"}
	bread.PutJam(jam)
	fmt.Println(bread)

	bread2 := &Bread{name: "Bread 2"}
	jam2 := &OrangeJam{name: "Orange Jam 1"}
	bread2.PutJam(jam2)
	fmt.Println(bread2)

	bread3 := &Bread{name: "Bread 3"}
	jam3 := &AppleJam{name: "Apple Jam 1"}
	bread3.PutJam(jam3)
	fmt.Println(bread3)

	fmt.Println("------------- lang")
}
