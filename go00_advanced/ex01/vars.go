package main

import "fmt"

func main() {
	a := "42"
	var b uint = 42
	c := 42
	var d uint8 = 42
	var e uint16 = 42
	var f uint32 = 42
	var g uint64 = 42
	h := false
	var i float32 = 42
	var j float64 = 42
	var k complex64 = 42
	var l complex128 = complex(42, 21)
	type FortyTwo struct {
	}
	m := FortyTwo{}
	n := [1]int{42}
	o := map[string]int{a: 42}
	var p *int
	q := make([]int, 0, 0)
	var r chan int
	var s interface{}

	fmt.Printf("%v is a %T.\n", a, a)
	fmt.Printf("%v is a %T.\n", b, b)
	fmt.Printf("%v is an %T.\n", c, c)
	fmt.Printf("%v is a %T.\n", d, d)
	fmt.Printf("%v is an %T.\n", e, e)
	fmt.Printf("%v is a %T.\n", f, f)
	fmt.Printf("%v is an %T.\n", g, g)
	fmt.Printf("%v is a %T.\n", h, h)
	fmt.Printf("%v is a %T.\n", i, i)
	fmt.Printf("%v is a %T.\n", j, j)
	fmt.Printf("%v is a %T.\n", k, k)
	fmt.Printf("%v is a %T.\n", l, l)
	fmt.Printf("%v is a %T.\n", m, m)
	fmt.Printf("%v is a %T.\n", n, n)
	fmt.Printf("%v is a %T.\n", o, o)
	fmt.Printf("%p is an %T.\n", p, p)
	fmt.Printf("%v is a %T.\n", q, q)
	fmt.Printf("%v is a %T.\n", r, r)
	fmt.Printf("%v is a %T.\n", s, s)
}
