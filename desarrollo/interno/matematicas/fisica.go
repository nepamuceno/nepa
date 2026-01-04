package matematicas

import (
	"math"
	"fmt"
	"nepa/desarrollo/interno/evaluador"
)

// Constantes Universales para el "Guru"
const (
	G_Universal = 6.67430e-11 // Constante de Gravitación
	C_Luz       = 299792458   // Velocidad de la luz en m/s
	H_Planck    = 6.62607e-34 // Constante de Planck
)

func inyectarFisicaGlobal() {

	// --- 1. CINEMÁTICA (MOVIMIENTO) ---

	// velocidad(distancia, tiempo)
	evaluador.Funciones["velocidad"] = func(args ...interface{}) (interface{}, error) {
		d, t, err := validar2("velocidad", args); if err != nil { return nil, err }
		return finalizar("velocidad", d/t)
	}

	// posicion_mrua(posicion_inicial, velocidad_inicial, aceleracion, tiempo)
	// d = xi + vi*t + 0.5*a*t^2
	evaluador.Funciones["posicion_mrua"] = func(args ...interface{}) (interface{}, error) {
		xi, vi, a, t, err := validar4("posicion_mrua", args); if err != nil { return nil, err }
		res := xi + (vi * t) + (0.5 * a * math.Pow(t, 2))
		return finalizar("posicion_mrua", res)
	}

	// --- 2. DINÁMICA Y FUERZAS (NEWTON) ---

	// fuerza(masa, aceleracion) -> F = m * a
	evaluador.Funciones["fuerza"] = func(args ...interface{}) (interface{}, error) {
		m, a, err := validar2("fuerza", args); if err != nil { return nil, err }
		return finalizar("fuerza", m*a)
	}

	// peso(masa, gravedad) -> P = m * g
	evaluador.Funciones["peso"] = func(args ...interface{}) (interface{}, error) {
		m, g, err := validar2("peso", args); if err != nil { return nil, err }
		return finalizar("peso", m*g)
	}

	// --- 3. ENERGÍA Y TRABAJO ---

	// energia_cinetica(masa, velocidad) -> Ec = 0.5 * m * v^2
	evaluador.Funciones["energia_cinetica"] = func(args ...interface{}) (interface{}, error) {
		m, v, err := validar2("energia_cinetica", args); if err != nil { return nil, err }
		return finalizar("energia_cinetica", 0.5*m*math.Pow(v, 2))
	}

	// energia_potencial(masa, gravedad, altura) -> Ep = m * g * h
	evaluador.Funciones["energia_potencial"] = func(args ...interface{}) (interface{}, error) {
		m, g, h, err := validar3("energia_potencial", args); if err != nil { return nil, err }
		return finalizar("energia_potencial", m*g*h)
	}

	// energia_masa(masa) -> E = m * c^2 (Einstein)
	evaluador.Funciones["energia_masa"] = func(args ...interface{}) (interface{}, error) {
		m, err := validar1("energia_masa", args); if err != nil { return nil, err }
		return finalizar("energia_masa", m*math.Pow(C_Luz, 2))
	}

	// --- 4. ASTROFÍSICA (GRAVITACIÓN) ---

	// atraccion_gravitatoria(masa1, masa2, distancia)
	// F = G * (m1 * m2) / r^2
	evaluador.Funciones["atraccion_gravitatoria"] = func(args ...interface{}) (interface{}, error) {
		m1, m2, r, err := validar3("atraccion_gravitatoria", args); if err != nil { return nil, err }
		res := G_Universal * (m1 * m2) / math.Pow(r, 2)
		return finalizar("atraccion_gravitatoria", res)
	}

	// --- 5. MECÁNICA CUÁNTICA BÁSICA ---

	// energia_foton(frecuencia) -> E = h * f
	evaluador.Funciones["energia_foton"] = func(args ...interface{}) (interface{}, error) {
		f, err := validar1("energia_foton", args); if err != nil { return nil, err }
		return finalizar("energia_foton", H_Planck*f)
	}

	// --- 6. RELATIVIDAD ESPECIAL ---

	// dilatacion_tiempo(tiempo_propio, velocidad)
	// t = t0 / sqrt(1 - v^2/c^2)
	evaluador.Funciones["dilatacion_tiempo"] = func(args ...interface{}) (interface{}, error) {
		t0, v, err := validar2("dilatacion_tiempo", args); if err != nil { return nil, err }
		if v >= C_Luz { return nil, fmt.Errorf("❌ ERROR: La velocidad no puede ser mayor o igual a la de la luz") }
		factor := math.Sqrt(1 - math.Pow(v, 2)/math.Pow(C_Luz, 2))
		return finalizar("dilatacion_tiempo", t0/factor)
	}

	// --- 7. FLUIDOS Y TERMODINÁMICA ---

	// presion(fuerza, area)
	evaluador.Funciones["presion"] = func(args ...interface{}) (interface{}, error) {
		f, a, err := validar2("presion", args); if err != nil { return nil, err }
		return finalizar("presion", f/a)
	}

	// celsius_a_fahrenheit(c)
	evaluador.Funciones["celsius_a_fahrenheit"] = func(args ...interface{}) (interface{}, error) {
		c, err := validar1("celsius_a_fahrenheit", args); if err != nil { return nil, err }
		return finalizar("celsius_a_fahrenheit", (c*9/5)+32)
	}
	
	// celsius_a_kelvin(c)
	evaluador.Funciones["celsius_a_kelvin"] = func(args ...interface{}) (interface{}, error) {
		c, err := validar1("celsius_a_kelvin", args); if err != nil { return nil, err }
		return finalizar("celsius_a_kelvin", c + 273.15)
	}
}
