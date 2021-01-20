package test

import (
	"fmt"
	"testing"
)

type R struct {
	name string
	age  int
}

func TestEqual(t *testing.T) {
	//rs := make([] R, 5)
	//for _, r := range(rs) {
	//	r.name = "111"
	//	r.age = 11
	//}
	//fmt.Println(rs)
	r := new(R)
	r.age = 21
	r.name = "123"

	r1 := R{"jack", 12}
	r2 := r1
	fmt.Println(r)
	fmt.Println(r1)
	fmt.Println(r2)
}
