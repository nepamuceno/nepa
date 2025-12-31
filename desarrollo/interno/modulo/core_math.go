package modulo

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

func toFloat(v interface{}) float64 {
	switch t := v.(type) {
	case float64: return t
	case int:     return float64(t)
	case int64:   return float64(t)
	case float32: return float64(t)
	default:      return 0.0
	}
}

func InyectarMatematicas(lib *LibreriaNepa) {
	ayudaInfo := make(map[string]string)

	rand.Seed(time.Now().UnixNano())

	reg := func(n string, desc string, f func(...interface{}) interface{}) {
		lib.Funciones[n] = f
		ayudaInfo[n] = desc
	}

	// --- CONSTANTES ---
	lib.Variables["pi"] = math.Pi
	lib.Variables["e"] = math.E
	lib.Variables["gravedad"] = 9.80665
	lib.Variables["phi"] = 1.618033988749895
	lib.Variables["luz"] = 299792458.0 
	lib.Variables["infinito"] = math.Inf(1)
	lib.Variables["cero_absoluto"] = -273.15
	lib.Variables["planck"] = 6.62607015e-34
	lib.Variables["stefan_boltzmann"] = 5.670374e-8
	lib.Variables["gas_ideal"] = 8.314462 

	// --- SISTEMA DE AYUDA ---
	reg("ayuda", "Muestra ayuda: ayuda('funcion') o ayuda() para lista", func(args ...interface{}) interface{} {
		if len(args) > 0 {
			nom := fmt.Sprintf("%v", args[0])
			if val, ok := ayudaInfo[nom]; ok {
				return fmt.Sprintf("Sintaxis [%s]: %s", nom, val)
			}
			return "Función no encontrada."
		}
		var lista string
		for k := range ayudaInfo { lista += k + ", " }
		return "Funciones disponibles: " + lista
	})

	// --- BÁSICAS Y TRIGONOMETRÍA ---
	reg("raiz", "raiz(n) -> Raíz cuadrada", func(args ...interface{}) interface{} { return math.Sqrt(toFloat(args[0])) })
	reg("raiz_n", "raiz_n(v, n) -> Raíz enésima", func(args ...interface{}) interface{} { return math.Pow(toFloat(args[0]), 1/toFloat(args[1])) })
	reg("potencia", "potencia(base, exp)", func(args ...interface{}) interface{} { return math.Pow(toFloat(args[0]), toFloat(args[1])) })
	reg("abs", "abs(n) -> Valor absoluto", func(args ...interface{}) interface{} { return math.Abs(toFloat(args[0])) })
	reg("log", "log(n) -> Logaritmo natural", func(args ...interface{}) interface{} { return math.Log(toFloat(args[0])) })
	reg("log10", "log10(n) -> Logaritmo base 10", func(args ...interface{}) interface{} { return math.Log10(toFloat(args[0])) })
	reg("signo", "signo(n) -> Retorna -1, 0 o 1", func(args ...interface{}) interface{} {
		val := toFloat(args[0])
		if val > 0 { return 1.0 } else if val < 0 { return -1.0 }
		return 0.0
	})
	
	reg("seno", "seno(rad)", func(args ...interface{}) interface{} { return math.Sin(toFloat(args[0])) })
	reg("coseno", "coseno(rad)", func(args ...interface{}) interface{} { return math.Cos(toFloat(args[0])) })
	reg("tangente", "tangente(rad)", func(args ...interface{}) interface{} { return math.Tan(toFloat(args[0])) })
	reg("secante", "secante(rad)", func(args ...interface{}) interface{} { return 1 / math.Cos(toFloat(args[0])) })
	reg("cosecante", "cosecante(rad)", func(args ...interface{}) interface{} { return 1 / math.Sin(toFloat(args[0])) })
	reg("cotangente", "cotangente(rad)", func(args ...interface{}) interface{} { return 1 / math.Tan(toFloat(args[0])) })
	
	reg("senoh", "senoh(n)", func(args ...interface{}) interface{} { return math.Sinh(toFloat(args[0])) })
	reg("cosenoh", "cosenoh(n)", func(args ...interface{}) interface{} { return math.Cosh(toFloat(args[0])) })
	reg("tangenteh", "tangenteh(n)", func(args ...interface{}) interface{} { return math.Tanh(toFloat(args[0])) })
	reg("aseno", "aseno(n) -> Arcoseno", func(args ...interface{}) interface{} { return math.Asin(toFloat(args[0])) })
	reg("acoseno", "acoseno(n) -> Arcocoseno", func(args ...interface{}) interface{} { return math.Acos(toFloat(args[0])) })
	reg("atangente", "atangente(n) -> Arcotangente", func(args ...interface{}) interface{} { return math.Atan(toFloat(args[0])) })
	reg("a_radianes", "a_radianes(grados)", func(args ...interface{}) interface{} { return toFloat(args[0]) * (math.Pi / 180) })
	reg("a_grados", "a_grados(rad)", func(args ...interface{}) interface{} { return toFloat(args[0]) * (180 / math.Pi) })

	// --- GEOMETRÍA ANALÍTICA Y ÁREAS (2D) ---
	reg("distancia2d", "distancia2d(x1, y1, x2, y2)", func(args ...interface{}) interface{} {
		return math.Hypot(toFloat(args[2])-toFloat(args[0]), toFloat(args[3])-toFloat(args[1]))
	})
	reg("punto_medio", "punto_medio(x1, y1, x2, y2) -> string 'x,y'", func(args ...interface{}) interface{} {
		return fmt.Sprintf("%f,%f", (toFloat(args[0])+toFloat(args[2]))/2, (toFloat(args[1])+toFloat(args[3]))/2)
	})
	reg("hipotenusa", "hipotenusa(cat1, cat2)", func(args ...interface{}) interface{} { return math.Hypot(toFloat(args[0]), toFloat(args[1])) })
	
	reg("area_circulo", "area_circulo(radio)", func(args ...interface{}) interface{} { return math.Pi * math.Pow(toFloat(args[0]), 2) })
	reg("area_rectangulo", "area_rectangulo(b, a)", func(args ...interface{}) interface{} { return toFloat(args[0]) * toFloat(args[1]) })
	reg("area_triangulo", "area_triangulo(b, a)", func(args ...interface{}) interface{} { return (toFloat(args[0]) * toFloat(args[1])) / 2 })
	reg("area_elipse", "area_elipse(eje_a, eje_b)", func(args ...interface{}) interface{} { return math.Pi * toFloat(args[0]) * toFloat(args[1]) })
	reg("area_trapecio", "area_trapecio(b_menor, b_mayor, h)", func(args ...interface{}) interface{} { return ((toFloat(args[0]) + toFloat(args[1])) * toFloat(args[2])) / 2 })
	reg("area_poligono", "area_poligono(n_lados, lado, apotema)", func(args ...interface{}) interface{} { return (toFloat(args[0]) * toFloat(args[1]) * toFloat(args[2])) / 2 })

	// --- GEOMETRÍA 3D (VOLÚMENES Y SUPERFICIES) ---
	reg("vol_esfera", "vol_esfera(radio)", func(args ...interface{}) interface{} { return (4.0/3.0) * math.Pi * math.Pow(toFloat(args[0]), 3) })
	reg("vol_cubo", "vol_cubo(lado)", func(args ...interface{}) interface{} { return math.Pow(toFloat(args[0]), 3) })
	reg("vol_cilindro", "vol_cilindro(radio, h)", func(args ...interface{}) interface{} { return math.Pi * math.Pow(toFloat(args[0]), 2) * toFloat(args[1]) })
	reg("vol_cono", "vol_cono(radio, h)", func(args ...interface{}) interface{} { return (1.0/3.0) * math.Pi * math.Pow(toFloat(args[0]), 2) * toFloat(args[1]) })
	reg("vol_piramide", "vol_piramide(area_base, h)", func(args ...interface{}) interface{} { return (1.0/3.0) * toFloat(args[0]) * toFloat(args[1]) })
	reg("area_sup_esfera", "area_sup_esfera(r)", func(args ...interface{}) interface{} { return 4 * math.Pi * math.Pow(toFloat(args[0]), 2) })
	reg("area_sup_cilindro", "area_sup_cilindro(r, h)", func(args ...interface{}) interface{} {
		r, h := toFloat(args[0]), toFloat(args[1])
		return 2 * math.Pi * r * (r + h)
	})

	// --- ESTADÍSTICA Y PROBABILIDAD ---
	reg("media", "media(n1, n2, ...)", func(args ...interface{}) interface{} {
		sum := 0.0; for _, v := range args { sum += toFloat(v) }
		return sum / float64(len(args))
	})
	reg("mediana", "mediana(n1, n2, ...)", func(args ...interface{}) interface{} {
		vals := []float64{}; for _, v := range args { vals = append(vals, toFloat(v)) }
		sort.Float64s(vals); l := len(vals)
		if l%2 == 0 { return (vals[l/2-1] + vals[l/2]) / 2 }
		return vals[l/2]
	})
	reg("varianza", "varianza(n1, n2, ...)", func(args ...interface{}) interface{} {
		var sum, sumSq float64; for _, v := range args { val := toFloat(v); sum += val; sumSq += val * val }
		n := float64(len(args)); return (sumSq / n) - (math.Pow(sum/n, 2))
	})
	reg("desviacion_est", "desviacion_est(n1, n2...)", func(args ...interface{}) interface{} {
		var sum, sumSq float64; for _, v := range args { val := toFloat(v); sum += val; sumSq += val * val }
		n := float64(len(args)); return math.Sqrt((sumSq / n) - (math.Pow(sum/n, 2)))
	})
	reg("rango", "rango(n1, n2...)", func(args ...interface{}) interface{} {
		vals := []float64{}; for _, v := range args { vals = append(vals, toFloat(v)) }
		sort.Float64s(vals); return vals[len(vals)-1] - vals[0]
	})

	// --- AZAR Y SIMULACIÓN ALEATORIA ---
	reg("azar", "azar() -> float [0, 1)", func(args ...interface{}) interface{} { return rand.Float64() })
	reg("azar_rango", "azar_rango(min, max)", func(args ...interface{}) interface{} { return toFloat(args[0]) + rand.Float64()*(toFloat(args[1])-toFloat(args[0])) })
	reg("azar_int", "azar_int(min, max)", func(args ...interface{}) interface{} { 
		min := int(toFloat(args[0])); max := int(toFloat(args[1]))
		return float64(rand.Intn(max-min+1) + min) 
	})
	reg("azar_dado", "azar_dado(caras)", func(args ...interface{}) interface{} { return float64(rand.Intn(int(toFloat(args[0]))) + 1) })

	// --- COMBINATORIA Y NÚMEROS ---
	reg("factorial", "factorial(n)", func(args ...interface{}) interface{} {
		n := int(toFloat(args[0])); res := 1.0; for i := 1; i <= n; i++ { res *= float64(i) }
		return res
	})
	reg("combinacion", "combinacion(n, k) -> nCr", func(args ...interface{}) interface{} {
		n := toFloat(args[0]); k := toFloat(args[1]); res := 1.0
		for i := 1.0; i <= k; i++ { res = res * (n - i + 1) / i }; return math.Round(res)
	})
	reg("permutacion", "permutacion(n, k) -> nPr", func(args ...interface{}) interface{} {
		n := toFloat(args[0]); k := toFloat(args[1]); res := 1.0
		for i := 0.0; i < k; i++ { res *= (n - i) }; return res
	})
	reg("es_primo", "es_primo(n) -> bool", func(args ...interface{}) interface{} {
		n := int(toFloat(args[0])); if n < 2 { return false }
		for i := 2; i*i <= n; i++ { if n%i == 0 { return false } }; return true
	})
	reg("mcd", "mcd(a, b) -> MCD", func(args ...interface{}) interface{} {
		a, b := int(toFloat(args[0])), int(toFloat(args[1]))
		for b != 0 { a, b = b, a%b }; return float64(a)
	})
	reg("mcm", "mcm(a, b) -> MCM", func(args ...interface{}) interface{} {
		a, b := toFloat(args[0]), toFloat(args[1]); tempA, tempB := int(a), int(b)
		for tempB != 0 { tempA, tempB = tempB, tempA%tempB }; return math.Abs(a*b) / float64(tempA)
	})

	// --- ÁLGEBRA LINEAL ---
	reg("det2x2", "det2x2(a,b,c,d)", func(args ...interface{}) interface{} {
		return (toFloat(args[0]) * toFloat(args[3])) - (toFloat(args[1]) * toFloat(args[2]))
	})
	reg("det3x3", "det3x3(a11,a12,a13,a21,a22,a23,a31,a32,a33)", func(args ...interface{}) interface{} {
		a, b, c := toFloat(args[0]), toFloat(args[1]), toFloat(args[2])
		d, e, f := toFloat(args[3]), toFloat(args[4]), toFloat(args[5])
		g, h, i := toFloat(args[6]), toFloat(args[7]), toFloat(args[8])
		return a*(e*i-f*h) - b*(d*i-f*g) + c*(d*h-e*g)
	})
	reg("producto_punto", "producto_punto(x1,y1,z1, x2,y2,z2)", func(args ...interface{}) interface{} {
		return (toFloat(args[0]) * toFloat(args[3])) + (toFloat(args[1]) * toFloat(args[4])) + (toFloat(args[2]) * toFloat(args[5]))
	})

	// --- REDONDEOS Y FORMATO ---
	reg("redondear", "redondear(n)", func(args ...interface{}) interface{} { return math.Round(toFloat(args[0])) })
	reg("piso", "piso(n) -> Floor", func(args ...interface{}) interface{} { return math.Floor(toFloat(args[0])) })
	reg("techo", "techo(n) -> Ceil", func(args ...interface{}) interface{} { return math.Ceil(toFloat(args[0])) })
	reg("formatear", "formatear(n, decimales)", func(args ...interface{}) interface{} {
		val := toFloat(args[0]); prec := int(toFloat(args[1]))
		res, _ := strconv.ParseFloat(fmt.Sprintf("%.*f", prec, val), 64)
		return res
	})

	// --- BITWISE Y BASES ---
	reg("bit_and", "bit_and(a, b)", func(args ...interface{}) interface{} { return float64(int64(toFloat(args[0])) & int64(toFloat(args[1]))) })
	reg("bit_or", "bit_or(a, b)", func(args ...interface{}) interface{} { return float64(int64(toFloat(args[0])) | int64(toFloat(args[1]))) })
	reg("bit_xor", "bit_xor(a, b)", func(args ...interface{}) interface{} { return float64(int64(toFloat(args[0])) ^ int64(toFloat(args[1]))) })
	reg("binario", "binario(n) -> string", func(args ...interface{}) interface{} { return strconv.FormatInt(int64(toFloat(args[0])), 2) })
	reg("hex", "hex(n) -> string", func(args ...interface{}) interface{} { return strconv.FormatInt(int64(toFloat(args[0])), 16) })

	// --- FÍSICA Y SIMULACIÓN ---
	reg("energia_relativista", "energia_relativista(masa) -> E=mc^2", func(args ...interface{}) interface{} {
		return toFloat(args[0]) * math.Pow(299792458.0, 2)
	})
	reg("caida_libre", "caida_libre(tiempo) -> d=0.5*g*t^2", func(args ...interface{}) interface{} {
		return 0.5 * 9.80665 * math.Pow(toFloat(args[0]), 2)
	})
	reg("proyectil_pos", "proyectil_pos(v0, ang, t) -> string 'x,y'", func(args ...interface{}) interface{} {
		v0 := toFloat(args[0]); ang := toFloat(args[1]) * math.Pi / 180; t := toFloat(args[2])
		x := v0 * math.Cos(ang) * t
		y := (v0 * math.Sin(ang) * t) - (0.5 * 9.80665 * t * t)
		return fmt.Sprintf("%f,%f", x, y)
	})
	// AGREGADAS PARA EVITAR PANIC EN CONCATENACIÓN
	reg("proyectil_x", "proyectil_x(v0, ang, t)", func(args ...interface{}) interface{} {
		return toFloat(args[0]) * math.Cos(toFloat(args[1])*math.Pi/180) * toFloat(args[2])
	})
	reg("proyectil_y", "proyectil_y(v0, ang, t)", func(args ...interface{}) interface{} {
		v0, ang, t := toFloat(args[0]), toFloat(args[1])*math.Pi/180, toFloat(args[2])
		return (v0 * math.Sin(ang) * t) - (0.5 * 9.80665 * t * t)
	})

	reg("presion_gas", "presion_gas(n, t, v) -> P=nRT/V", func(args ...interface{}) interface{} {
		return (toFloat(args[0]) * 8.314462 * toFloat(args[1])) / toFloat(args[2])
	})

	// --- FINANZAS AVANZADAS ---
	reg("interes_compuesto", "interes_compuesto(c, i, t)", func(args ...interface{}) interface{} {
		return toFloat(args[0]) * math.Pow(1+toFloat(args[1]), toFloat(args[2]))
	})
	reg("valor_presente", "valor_presente(vf, i, t)", func(args ...interface{}) interface{} {
		return toFloat(args[0]) / math.Pow(1+toFloat(args[1]), toFloat(args[2]))
	})
	reg("amortizacion", "amortizacion(capital, tasa, meses)", func(args ...interface{}) interface{} {
		p := toFloat(args[0]); i := toFloat(args[1]) / 12; n := toFloat(args[2])
		return (p * i * math.Pow(1+i, n)) / (math.Pow(1+i, n) - 1)
	})
	reg("tasa_crecimiento", "tasa_crecimiento(final, inicial, t)", func(args ...interface{}) interface{} {
		return math.Pow(toFloat(args[0])/toFloat(args[1]), 1/toFloat(args[2])) - 1
	})

	// --- FÍSICA Y MATEMÁTICA AVANZADA ---
	reg("fuerza_gravitatoria", "fuerza_grav(m1, m2, r)", func(args ...interface{}) interface{} {
		G := 6.67430e-11
		return G * (toFloat(args[0]) * toFloat(args[1])) / math.Pow(toFloat(args[2]), 2)
	})
	reg("energia_cinetica", "energia_cinetica(m, v)", func(args ...interface{}) interface{} {
		return 0.5 * toFloat(args[0]) * math.Pow(toFloat(args[1]), 2)
	})
	reg("trabajo", "trabajo(f, d, ang_deg)", func(args ...interface{}) interface{} {
		ang := toFloat(args[2]) * math.Pi / 180
		return toFloat(args[0]) * toFloat(args[1]) * math.Cos(ang)
	})
	reg("densidad", "densidad(masa, vol)", func(args ...interface{}) interface{} {
		return toFloat(args[0]) / toFloat(args[1])
	})
	reg("ley_ohm_v", "ley_ohm_v(i, r) -> V", func(args ...interface{}) interface{} { return toFloat(args[0]) * toFloat(args[1]) })
	reg("magnitud_vector", "magnitud_vector(x, y, z)", func(args ...interface{}) interface{} {
		return math.Sqrt(math.Pow(toFloat(args[0]), 2) + math.Pow(toFloat(args[1]), 2) + math.Pow(toFloat(args[2]), 2))
	})
}




