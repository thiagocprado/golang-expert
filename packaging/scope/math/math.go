// O nome do pacote deve ser em minúsculo.
// Em Go, nomes que começam com letra maiúscula são exportados (visíveis fora do pacote),
// enquanto nomes com letra minúscula são privados ao pacote.
package math

// math (com letra minúscula) não é exportada, só pode ser usada dentro do pacote math.
type math struct {
	a int
	b int
}

// NewMath (com letra maiúscula) é exportada e pode ser usada em outros pacotes.
func NewMath(a, b int) math {
	return math{a: a, b: b}
}

// Add (com letra maiúscula) é exportada e pode ser usada em outros pacotes.
func (m *math) Add() int {
	return m.a + m.b
}

// Exemplo adicional: struct e função exportadas

// MathExported é exportada (letra maiúscula)
type MathExported struct {
	A int
	B int
}

// NewMathExported é exportada (letra maiúscula)
func NewMathExported(a, b int) *MathExported {
	return &MathExported{A: a, B: b}
}

// AddExported é exportada (letra maiúscula)
func (m *MathExported) AddExported() int {
	return m.A + m.B
}
