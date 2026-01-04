package matematicas

import (
	"math"
	"math/rand"
	"nepa/desarrollo/interno/evaluador"
	"time"
)

func inyectarProbabilidadGlobal() {
	rand.Seed(time.Now().UnixNano())

	// --- 1. COMBINATORIA ---
	// combinaciones(n, r) -> n! / (r!(n-r)!)
	evaluador.Funciones["combinaciones"] = func(args ...interface{}) (interface{}, error) {
		n, r, err := validar2("combinaciones", args); if err != nil { return nil, err }
		return finalizar("combinaciones", math.Gamma(n+1)/(math.Gamma(r+1)*math.Gamma(n-r+1)))
	}

	// --- 2. DISTRIBUCIONES (La Campana de Gauss) ---
	// normal_pdf(x, media, desviacion)
	evaluador.Funciones["distribucion_normal"] = func(args ...interface{}) (interface{}, error) {
		x, m, d, err := validar3("distribucion_normal", args); if err != nil { return nil, err }
		exponente := math.Pow(x-m, 2) / (2 * math.Pow(d, 2))
		coeficiente := 1 / (d * math.Sqrt(2*math.Pi))
		return finalizar("distribucion_normal", coeficiente*math.Exp(-exponente))
	}

	// --- 3. GENERADORES ---
	evaluador.Funciones["aleatorio_rango"] = func(args ...interface{}) (interface{}, error) {
		min, max, err := validar2("aleatorio_rango", args); if err != nil { return nil, err }
		return min + rand.Float64()*(max-min), nil
	}
}
