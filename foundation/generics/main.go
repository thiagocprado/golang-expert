package main

// generics
func sum[T int | float64](m map[string]T) T {
	var sum T

	for _, v := range m {
		sum += v
	}

	return sum
}

type MyNumber int

type Number interface {
	// com o til, conseguimos abrir uma exceção para "outros tipos" de int, como é o caso do MyNumber
	~int | ~float64
}

// também podemos utilizar generics em structs
type Client[T Number] struct {
	CPF  T
	Name string
}

// constraints são utilizadas para limitar os tipos que podem ser utilizados com generics
func sum2[T Number](m map[string]T) T {
	var sum T

	for _, v := range m {
		sum += v
	}

	return sum
}

// funciona apenas com comparações
func compare[T comparable](a T, b T) bool {
	return a == b
}

func main() {
	m := map[string]int{
		"one":   100,
		"two":   200,
		"three": 300,
	}

	m2 := map[string]float64{
		"one":   100.10,
		"two":   200.20,
		"three": 300.30,
	}

	m3 := map[string]MyNumber{
		"one":   100,
		"two":   200,
		"three": 300,
	}

	// o uso do generics é implicito para manter a retrocompatibilidade
	println(sum(m))
	println(sum2(m2))
	println(sum2(m3))
	println(compare(10, 10.0))
}
