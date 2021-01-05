package main

import "fmt"

type Bread struct {
	name string // 이름
}

func (b *Bread) PutJam(jam *Jam) {
	b.name += " + " + jam.GetVal()
}

func (b *Bread) String() string {
	return b.name
}

type Jam struct {
	name string
	val  string
}

func (j *Jam) GetVal() string {
	return j.val
}

// main
func main() {
	fmt.Println("go ---------------")

	bread := &Bread{name: "bread"}
	jam := &Jam{name: "jam", val: "jam_val"}

	bread.PutJam(jam)
	fmt.Println(bread)

	fmt.Println("------------- lang")
}
