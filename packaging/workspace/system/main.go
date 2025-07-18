package main

import "github.com/thiagocprado/golang-expert/packaging/replace/math"

func main() {
	m := math.NewMath(3, 4)
	println(m.Add())
}
