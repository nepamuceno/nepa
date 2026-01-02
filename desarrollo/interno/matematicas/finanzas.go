package matematicas

import (
    "math"

    "nepa/desarrollo/interno/evaluador"
)

// InyectarFinanzas agrega funciones financieras al contexto
func InyectarFinanzas(ctx *evaluador.Contexto) {
    if ctx.Funciones == nil {
        ctx.Funciones = map[string]func(...interface{}) interface{}{}
    }

    reg := func(n string, f func(...interface{}) interface{}) {
        ctx.Funciones[n] = f
    }

    // --- Interés compuesto ---
    reg("interes_compuesto", func(args ...interface{}) interface{} {
        // c * (1+i)^t
        if len(args) != 3 {
            return "Error: interes_compuesto requiere 3 argumentos (capital, tasa, tiempo)"
        }
        return toFloat(args[0]) * math.Pow(1+toFloat(args[1]), toFloat(args[2]))
    })

    // --- Valor presente ---
    reg("valor_presente", func(args ...interface{}) interface{} {
        // vf / (1+i)^t
        if len(args) != 3 {
            return "Error: valor_presente requiere 3 argumentos (valor futuro, tasa, tiempo)"
        }
        return toFloat(args[0]) / math.Pow(1+toFloat(args[1]), toFloat(args[2]))
    })

    // --- Amortización ---
    reg("amortizacion", func(args ...interface{}) interface{} {
        // Fórmula de cuota fija: (P * i * (1+i)^n) / ((1+i)^n - 1)
        if len(args) != 3 {
            return "Error: amortizacion requiere 3 argumentos (principal, tasa anual, número de meses)"
        }
        p := toFloat(args[0])
        i := toFloat(args[1]) / 12 // tasa mensual
        n := toFloat(args[2])
        if n == 0 {
            return "Error: número de meses no puede ser 0"
        }
        return (p * i * math.Pow(1+i, n)) / (math.Pow(1+i, n) - 1)
    })

    // --- Tasa de crecimiento ---
    reg("tasa_crecimiento", func(args ...interface{}) interface{} {
        // (final/inicial)^(1/t) - 1
        if len(args) != 3 {
            return "Error: tasa_crecimiento requiere 3 argumentos (valor final, valor inicial, tiempo)"
        }
        if toFloat(args[1]) == 0 {
            return "Error: valor inicial no puede ser 0"
        }
        return math.Pow(toFloat(args[0])/toFloat(args[1]), 1/toFloat(args[2])) - 1
    })
}
