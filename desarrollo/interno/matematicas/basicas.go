package matematicas

import (
	"fmt"
	"math"
	"nepa/desarrollo/interno/evaluador"
)

// --- INYECCIÓN TOTAL DE LIBRERÍA MATH ---

func inyectarBasicasGlobal() {
	// 1. ARITMÉTICA Y POTENCIAS
	evaluador.Funciones["absoluto"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("absoluto", args); if err != nil { return nil, err }
		return finalizar("absoluto", math.Abs(v))
	}
	evaluador.Funciones["raiz"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("raiz", args); if err != nil { return nil, err }
		if v < 0 { return nil, fmt.Errorf("❌ ERROR: raiz de numero negativo") }
		return finalizar("raiz", math.Sqrt(v))
	}
	evaluador.Funciones["raiz_cubica"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("raiz_cubica", args); if err != nil { return nil, err }
		return finalizar("raiz_cubica", math.Cbrt(v))
	}
	evaluador.Funciones["potencia"] = func(args ...interface{}) (interface{}, error) {
		b, e, err := validar2("potencia", args); if err != nil { return nil, err }
		return finalizar("potencia", math.Pow(b, e))
	}
	evaluador.Funciones["resto"] = func(args ...interface{}) (interface{}, error) {
		v1, v2, err := validar2("resto", args); if err != nil { return nil, err }
		return finalizar("resto", math.Mod(v1, v2))
	}
	evaluador.Funciones["hipotenusa"] = func(args ...interface{}) (interface{}, error) {
		v1, v2, err := validar2("hipotenusa", args); if err != nil { return nil, err }
		return finalizar("hipotenusa", math.Hypot(v1, v2))
	}

	// 2. TRIGONOMETRÍA BÁSICA E INVERSA
	evaluador.Funciones["seno"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("seno", args); if err != nil { return nil, err }
		return finalizar("seno", math.Sin(v))
	}
	evaluador.Funciones["coseno"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("coseno", args); if err != nil { return nil, err }
		return finalizar("coseno", math.Cos(v))
	}
	evaluador.Funciones["tangente"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("tangente", args); if err != nil { return nil, err }
		return finalizar("tangente", math.Tan(v))
	}
	evaluador.Funciones["arcoseno"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("arcoseno", args); if err != nil { return nil, err }
		return finalizar("arcoseno", math.Asin(v))
	}
	evaluador.Funciones["arcocoseno"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("arcocoseno", args); if err != nil { return nil, err }
		return finalizar("arcocoseno", math.Acos(v))
	}
	evaluador.Funciones["arcotangente"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("arcotangente", args); if err != nil { return nil, err }
		return finalizar("arcotangente", math.Atan(v))
	}
	evaluador.Funciones["arcotangente2"] = func(args ...interface{}) (interface{}, error) {
		y, x, err := validar2("arcotangente2", args); if err != nil { return nil, err }
		return finalizar("arcotangente2", math.Atan2(y, x))
	}

	// 3. TRIGONOMETRÍA HIPERBÓLICA
	evaluador.Funciones["seno_h"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("seno_h", args); if err != nil { return nil, err }
		return finalizar("seno_h", math.Sinh(v))
	}
	evaluador.Funciones["coseno_h"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("coseno_h", args); if err != nil { return nil, err }
		return finalizar("coseno_h", math.Cosh(v))
	}
	evaluador.Funciones["tangente_h"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("tangente_h", args); if err != nil { return nil, err }
		return finalizar("tangente_h", math.Tanh(v))
	}

	// 4. EXPONENCIALES Y LOGARITMOS
	evaluador.Funciones["exp"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("exp", args); if err != nil { return nil, err }
		return finalizar("exp", math.Exp(v))
	}
	evaluador.Funciones["exp2"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("exp2", args); if err != nil { return nil, err }
		return finalizar("exp2", math.Exp2(v))
	}
	evaluador.Funciones["logaritmo"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("logaritmo", args); if err != nil { return nil, err }
		return finalizar("logaritmo", math.Log(v))
	}
	evaluador.Funciones["logaritmo10"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("logaritmo10", args); if err != nil { return nil, err }
		return finalizar("logaritmo10", math.Log10(v))
	}
	evaluador.Funciones["logaritmo2"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("logaritmo2", args); if err != nil { return nil, err }
		return finalizar("logaritmo2", math.Log2(v))
	}

	// 5. REDONDEO Y CLASIFICACIÓN
	evaluador.Funciones["techo"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("techo", args); if err != nil { return nil, err }
		return finalizar("techo", math.Ceil(v))
	}
	evaluador.Funciones["piso"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("piso", args); if err != nil { return nil, err }
		return finalizar("piso", math.Floor(v))
	}
	evaluador.Funciones["truncar"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("truncar", args); if err != nil { return nil, err }
		return finalizar("truncar", math.Trunc(v))
	}
	evaluador.Funciones["redondear"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("redondear", args); if err != nil { return nil, err }
		return finalizar("redondear", math.Round(v))
	}

	// 6. MÁXIMOS, MÍNIMOS Y SIGNOS
	evaluador.Funciones["maximo"] = func(args ...interface{}) (interface{}, error) {
		v1, v2, err := validar2("maximo", args); if err != nil { return nil, err }
		return finalizar("maximo", math.Max(v1, v2))
	}
	evaluador.Funciones["minimo"] = func(args ...interface{}) (interface{}, error) {
		v1, v2, err := validar2("minimo", args); if err != nil { return nil, err }
		return finalizar("minimo", math.Min(v1, v2))
	}
	evaluador.Funciones["copiar_signo"] = func(args ...interface{}) (interface{}, error) {
		v1, v2, err := validar2("copiar_signo", args); if err != nil { return nil, err }
		return finalizar("copiar_signo", math.Copysign(v1, v2))
	}
	evaluador.Funciones["diferencia_positiva"] = func(args ...interface{}) (interface{}, error) {
		v1, v2, err := validar2("diferencia_positiva", args); if err != nil { return nil, err }
		return finalizar("diferencia_positiva", math.Dim(v1, v2))
	}
}
