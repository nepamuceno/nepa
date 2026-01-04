package matematicas

import (
	"math"
	"fmt"
	"nepa/desarrollo/interno/evaluador"
)

func inyectarFinanzasGlobal() {

	// --- 1. INTERÉS Y VALOR TEMPORAL ---

	// interes_simple(capital, tasa, tiempo)
	evaluador.Funciones["interes_simple"] = func(args ...interface{}) (interface{}, error) {
		c, i, t, err := validar3("interes_simple", args); if err != nil { return nil, err }
		return finalizar("interes_simple", c * i * t)
	}

	// interes_compuesto(capital, tasa, tiempo) -> Monto Final
	evaluador.Funciones["interes_compuesto"] = func(args ...interface{}) (interface{}, error) {
		c, i, t, err := validar3("interes_compuesto", args); if err != nil { return nil, err }
		return finalizar("interes_compuesto", c * math.Pow(1+i, t))
	}

	// valor_presente(monto_futuro, tasa, tiempo)
	evaluador.Funciones["valor_presente"] = func(args ...interface{}) (interface{}, error) {
		m, i, t, err := validar3("valor_presente", args); if err != nil { return nil, err }
		return finalizar("valor_presente", m / math.Pow(1+i, t))
	}

	// --- 2. PRÉSTAMOS Y ANUALIDADES (SISTEMA FRANCÉS) ---

	// cuota_prestamo(capital, tasa_periodo, num_periodos)
	evaluador.Funciones["cuota_prestamo"] = func(args ...interface{}) (interface{}, error) {
		p, i, n, err := validar3("cuota_prestamo", args); if err != nil { return nil, err }
		if i == 0 { return finalizar("cuota_prestamo", p/n) }
		numerador := i * math.Pow(1+i, n)
		denominador := math.Pow(1+i, n) - 1
		return finalizar("cuota_prestamo", p * (numerador / denominador))
	}

	// total_pagado(cuota, num_periodos)
	evaluador.Funciones["total_pagado"] = func(args ...interface{}) (interface{}, error) {
		c, n, err := validar2("total_pagado", args); if err != nil { return nil, err }
		return finalizar("total_pagado", c * n)
	}

	// --- 3. INVERSIÓN Y RENTABILIDAD ---

	// roi(ganancia, inversion) -> Retorno en %
	evaluador.Funciones["roi"] = func(args ...interface{}) (interface{}, error) {
		g, i, err := validar2("roi", args); if err != nil { return nil, err }
		return finalizar("roi", (g/i)*100)
	}

	// van(inversion_inicial, tasa_descuento, flujo1, flujo2...)
	// Calcula el Valor Actual Neto de una serie de flujos de caja.
	evaluador.Funciones["van"] = func(args ...interface{}) (interface{}, error) {
		nums, err := validarN("van", args); if err != nil { return nil, err }
		if len(nums) < 3 { return nil, fmt.Errorf("❌ 'van' requiere: inversion, tasa y al menos 1 flujo") }
		
		inversion := nums[0] // Generalmente negativo
		tasa := nums[1]
		flujos := nums[2:]
		
		sumaPresente := inversion
		for t, flujo := range flujos {
			sumaPresente += flujo / math.Pow(1+tasa, float64(t+1))
		}
		return finalizar("van", sumaPresente)
	}

	// --- 4. NEGOCIOS Y PRECIOS ---

	// margen_ganancia(precio_venta, costo) -> %
	evaluador.Funciones["margen_ganancia"] = func(args ...interface{}) (interface{}, error) {
		v, c, err := validar2("margen_ganancia", args); if err != nil { return nil, err }
		return finalizar("margen_ganancia", ((v-c)/v)*100)
	}

	// punto_equilibrio(costos_fijos, precio_venta, costo_variable)
	// Cuántas unidades vender para no ganar ni perder.
	evaluador.Funciones["punto_equilibrio"] = func(args ...interface{}) (interface{}, error) {
		fijos, precio, variable, err := validar3("punto_equilibrio", args); if err != nil { return nil, err }
		if precio <= variable { return nil, fmt.Errorf("❌ ERROR: El precio debe ser mayor al costo variable") }
		return finalizar("punto_equilibrio", fijos / (precio - variable))
	}

	// --- 5. ECONOMÍA REAL ---

	// poder_adquisitivo(monto, inflacion_anual, años)
	// Cuánto valdrá ese dinero en el futuro ajustado por inflación.
	evaluador.Funciones["poder_adquisitivo"] = func(args ...interface{}) (interface{}, error) {
		m, i, t, err := validar3("poder_adquisitivo", args); if err != nil { return nil, err }
		return finalizar("poder_adquisitivo", m / math.Pow(1+i, t))
	}

	// tasa_real(tasa_nominal, inflacion)
	// La ganancia real de una inversión descontando la inflación.
	evaluador.Funciones["tasa_real"] = func(args ...interface{}) (interface{}, error) {
		n, inf, err := validar2("tasa_real", args); if err != nil { return nil, err }
		// Formula de Fisher: [(1 + nominal) / (1 + inflacion)] - 1
		return finalizar("tasa_real", ((1+n)/(1+inf)) - 1)
	}
}
