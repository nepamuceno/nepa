package matematicas

import (
    "math"

    "nepa/desarrollo/interno/evaluador"
)

// InyectarBasicas agrega funciones matemáticas básicas al contexto
func InyectarBasicas(ctx *evaluador.Contexto) {
    if ctx.Funciones == nil {
        ctx.Funciones = map[string]func(...interface{}) interface{}{}
    }

    reg := func(n string, f func(...interface{}) interface{}) {
        ctx.Funciones[n] = f
    }

    // --- Raíz cuadrada ---
    reg("raiz", func(args ...interface{}) interface{} {
        if len(args) != 1 {
            return "Error: raiz requiere 1 argumento"
        }
        return math.Sqrt(toFloat(args[0]))
    })

    // --- Potencia ---
    reg("potencia", func(args ...interface{}) interface{} {
        if len(args) != 2 {
            return "Error: potencia requiere 2 argumentos"
        }
        return math.Pow(toFloat(args[0]), toFloat(args[1]))
    })

    // --- Seno ---
    reg("seno", func(args ...interface{}) interface{} {
        if len(args) != 1 {
            return "Error: seno requiere 1 argumento"
        }
        return math.Sin(toFloat(args[0]))
    })

    // --- Coseno ---
    reg("coseno", func(args ...interface{}) interface{} {
        if len(args) != 1 {
            return "Error: coseno requiere 1 argumento"
        }
        return math.Cos(toFloat(args[0]))
    })

    // --- Tangente ---
    reg("tangente", func(args ...interface{}) interface{} {
        if len(args) != 1 {
            return "Error: tangente requiere 1 argumento"
        }
        return math.Tan(toFloat(args[0]))
    })

    // --- Logaritmo natural ---
    reg("log", func(args ...interface{}) interface{} {
        if len(args) != 1 {
            return "Error: log requiere 1 argumento"
        }
        return math.Log(toFloat(args[0]))
    })

    // --- Valor absoluto ---
    reg("abs", func(args ...interface{}) interface{} {
        if len(args) != 1 {
            return "Error: abs requiere 1 argumento"
        }
        return math.Abs(toFloat(args[0]))
    })
}
