package main

func main() {

	s := make([]int, 3)

	noItems := s[:0]

	println(s, noItems)

	a := 4
	b := 2
	c := 3

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

	if a > b {
		println("a > b")
	}

	if a > b && c > a {
		println("a > b && c > a")
	}

	if a > b || c > a {
		println("a > b || c > a")
	}

	if a > b {
		println(a)
	} else {
		println(b)
	}
}
