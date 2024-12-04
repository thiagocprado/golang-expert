package main

func main() {
	a := 1
	b := 2
	c := 3

	if a > b {
		println(a)
	} else if a < b {
		println(b)
	} else {
		println("equal")
	}

	if a > b && a > c {
		println(a)
	} else if b > a && b > c {
		println(b)
	} else {
		println(c)
	}

	switch a {
	case 1:
		println("a")
	case 2:
		println("b")
	case 3:
		println("c")
	default:
		println("d")
	}
}
