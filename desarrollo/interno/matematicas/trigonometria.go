package matematicas

import (
	"math"
	"nepa/desarrollo/interno/evaluador"
)

func inyectarTrigonometriaGlobal() {
	// --- 1. HIPERBÓLICAS (Catenarias y Relatividad) ---
	evaluador.Funciones["sinh"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("sinh", args); if err != nil { return nil, err }
		return finalizar("sinh", math.Sinh(v))
	}
	evaluador.Funciones["cosh"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("cosh", args); if err != nil { return nil, err }
		return finalizar("cosh", math.Cosh(v))
	}

	// --- 2. RESOLUCIÓN DE TRIÁNGULOS (Ley de Cosenos) ---
	// lado_faltante(ladoA, ladoB, anguloGrados)
	evaluador.Funciones["triangulo_lado_c"] = func(args ...interface{}) (interface{}, error) {
		a, b, ang, err := validar3("triangulo_lado_c", args); if err != nil { return nil, err }
		rad := ang * (math.Pi / 180)
		// c² = a² + b² - 2ab * cos(C)
		res := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2) - 2*a*b*math.Cos(rad))
		return finalizar("triangulo_lado_c", res)
	}

	// --- 3. COORDENADAS (Sistemas de Posicionamiento) ---
	// cartesianas_a_polar(x, y) -> entrega [radio, angulo_grados]
	evaluador.Funciones["a_polar"] = func(args ...interface{}) (interface{}, error) {
		x, y, err := validar2("a_polar", args); if err != nil { return nil, err }
		radio := math.Hypot(x, y)
		angulo := math.Atan2(y, x) * (180 / math.Pi)
		return []float64{radio, angulo}, nil
	}
}
