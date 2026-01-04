package matematicas

import (
	"fmt"
	"math"
	"nepa/desarrollo/interno/evaluador"
)

func inyectarAlgebraGlobal() {

	// --- 1. RESOLUCIÓN DE ECUACIONES Y RAÍCES ---

	// resolver_cuadratica(a, b, c) -> Retorna [x1, x2] usando la fórmula general
	evaluador.Funciones["resolver_cuadratica"] = func(args ...interface{}) (interface{}, error) {
		a, b, c, err := validar3("resolver_cuadratica", args); if err != nil { return nil, err }
		disc := (b * b) - (4 * a * c)
		if disc < 0 {
			return nil, fmt.Errorf("❌ ERROR: Discriminante negativo (%f). Raíces imaginarias no soportadas", disc)
		}
		x1 := (-b + math.Sqrt(disc)) / (2 * a)
		x2 := (-b - math.Sqrt(disc)) / (2 * a)
		return []float64{x1, x2}, nil
	}

	// discriminante(a, b, c) -> b² - 4ac
	evaluador.Funciones["discriminante"] = func(args ...interface{}) (interface{}, error) {
		a, b, c, err := validar3("discriminante", args); if err != nil { return nil, err }
		return (b * b) - (4 * a * c), nil
	}

	// --- 2. TEORÍA DE NÚMEROS (Criptografía y Algoritmia) ---

	// mcd(a, b) -> Máximo Común Divisor (Algoritmo de Euclides)
	evaluador.Funciones["mcd"] = func(args ...interface{}) (interface{}, error) {
		a, b, err := validar2("mcd", args); if err != nil { return nil, err }
		ia, ib := int64(math.Abs(a)), int64(math.Abs(b))
		for ib != 0 {
			ia, ib = ib, ia%ib
		}
		return float64(ia), nil
	}

	// mcm(a, b) -> Mínimo Común Múltiplo
	evaluador.Funciones["mcm"] = func(args ...interface{}) (interface{}, error) {
		a, b, err := validar2("mcm", args); if err != nil { return nil, err }
		if a == 0 || b == 0 { return 0.0, nil }
		ia, ib := int64(math.Abs(a)), int64(math.Abs(b))
		// MCM(a,b) = |a*b| / MCD(a,b)
		tempA, tempB := ia, ib
		for tempB != 0 {
			tempA, tempB = tempB, tempA%tempB
		}
		return float64((ia * ib) / tempA), nil
	}

	// es_primo(n) -> Test de primalidad optimizado
	evaluador.Funciones["es_primo"] = func(args ...interface{}) (interface{}, error) {
		n, err := validar1("es_primo", args); if err != nil { return nil, err }
		num := int64(n)
		if num <= 1 { return false, nil }
		if num <= 3 { return true, nil }
		if num%2 == 0 || num%3 == 0 { return false, nil }
		for i := int64(5); i*i <= num; i += 6 {
			if num%i == 0 || num%(i+2) == 0 { return false, nil }
		}
		return true, nil
	}

	// --- 3. FUNCIONES ESPECIALES Y DE PRECISIÓN (Tu Aporte + Mejoras) ---

	evaluador.Funciones["gamma"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("gamma", args); if err != nil { return nil, err }
		return finalizar("gamma", math.Gamma(v))
	}

	evaluador.Funciones["log_gamma"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("log_gamma", args); if err != nil { return nil, err }
		res, _ := math.Lgamma(v)
		return finalizar("log_gamma", res)
	}

	evaluador.Funciones["error_mat"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("error_mat", args); if err != nil { return nil, err }
		return finalizar("error_mat", math.Erf(v))
	}

	evaluador.Funciones["error_mat_complementario"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("error_mat_complementario", args); if err != nil { return nil, err }
		return finalizar("error_mat_complementario", math.Erfc(v))
	}

	evaluador.Funciones["logaritmo_b"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("logaritmo_b", args); if err != nil { return nil, err }
		return finalizar("logaritmo_b", math.Logb(v))
	}

	evaluador.Funciones["logaritmo_1p"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("logaritmo_1p", args); if err != nil { return nil, err }
		return finalizar("logaritmo_1p", math.Log1p(v))
	}

	evaluador.Funciones["exp_m1"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("exp_m1", args); if err != nil { return nil, err }
		return finalizar("exp_m1", math.Expm1(v))
	}

	evaluador.Funciones["escalar_binario"] = func(args ...interface{}) (interface{}, error) {
		x, n, err := validar2("escalar_binario", args); if err != nil { return nil, err }
		return finalizar("escalar_binario", math.Ldexp(x, int(n)))
	}

	// --- 4. FUNCIONES DE BESSEL (Física de Ondas) ---

	evaluador.Funciones["bessel_j0"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("bessel_j0", args); if err != nil { return nil, err }
		return finalizar("bessel_j0", math.J0(v))
	}
	evaluador.Funciones["bessel_j1"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("bessel_j1", args); if err != nil { return nil, err }
		return finalizar("bessel_j1", math.J1(v))
	}
	evaluador.Funciones["bessel_y0"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("bessel_y0", args); if err != nil { return nil, err }
		return finalizar("bessel_y0", math.Y0(v))
	}
	evaluador.Funciones["bessel_y1"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("bessel_y1", args); if err != nil { return nil, err }
		return finalizar("bessel_y1", math.Y1(v))
	}

	// --- 5. EVALUACIÓN POLINÓMICA (Algoritmo de Horner) ---

	// poli_evaluar(x, c0, c1, c2...) -> evalúa c0 + c1*x + c2*x^2...
	evaluador.Funciones["poli_evaluar"] = func(args ...interface{}) (interface{}, error) {
		if len(args) < 2 { return nil, fmt.Errorf("❌ ERROR: poli_evaluar(x, coeficientes...)") }
		x, _ := evaluador.ConvertirAReal(args[0])
		var res float64
		// Iteramos desde el último coeficiente (grado más alto) hacia atrás
		for i := len(args) - 1; i >= 1; i-- {
			c, _ := evaluador.ConvertirAReal(args[i])
			res = res*x + c
		}
		return res, nil
	}
}
