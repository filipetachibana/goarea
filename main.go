package goarea

import "math"

// Pi é uma proporção númerica definida pela relação entre
// o perímetro de uma circunferência e seu diâmetro
const Pi float64 = 3.141516

// Circ é a circunferência do circulo
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é responsável por calcular a área de um retângulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não é visível
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
