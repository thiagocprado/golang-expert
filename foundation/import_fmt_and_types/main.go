package main

import "fmt"

// podemos criar tipos personalizados
type ID int

var (
	b bool = true
	i int  = 1
	s string
	f ID = 10
)

// verbos de formatação (format verbs)
func main() {
	fmt.Printf("o tipo de B é %T \n", b)
	fmt.Printf("o tipo de I é %T", i)
	fmt.Printf("o tipo de S é %T", s)
	fmt.Printf("o tipo de F é %T", f)
}

// | Verbo | Significado                                        |
// | ----- | -------------------------------------------------- |
// | `%s`  | String (texto)                                     |
// | `%d`  | Inteiro decimal (`int`, `int64`, etc.)             |
// | `%f`  | Número de ponto flutuante (float)                  |
// | `%v`  | Valor "padrão" (útil para structs, genérico)       |
// | `%+v` | Valor com nomes de campos (para structs)           |
// | `%#v` | Representação Go válida do valor (útil para debug) |
// | `%T`  | Tipo do valor                                      |
// | `%t`  | Booleano                                           |
// | `%q`  | String com aspas                                   |
// | `%x`  | Hexadecimal (inteiros ou bytes)                    |
