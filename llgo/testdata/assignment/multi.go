package main

func xyz() (int, int) {
	return 123, 456
}

func abc() (int, int) {
	var a, b = xyz()
	return a, b
}

type S struct {
	a int
	b int
}

func main() {
	a, b := xyz()
	println(a, b)
	b, a = abc()
	println(a, b)

	var s S
	s.a, s.b = a, b
	println(s.a, s.b)
}
