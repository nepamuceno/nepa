package matematicas

import (
    "math"

    "nepa/desarrollo/interno/evaluador"
)

// InyectarAlgebra agrega funciones de álgebra lineal al contexto
func InyectarAlgebra(ctx *evaluador.Contexto) {
    if ctx.Funciones == nil {
        ctx.Funciones = map[string]func(...interface{}) interface{}{}
    }

    reg := func(n string, f func(...interface{}) interface{}) {
        ctx.Funciones[n] = f
    }

    // --- Determinantes ---
    reg("det2x2", func(args ...interface{}) interface{} {
        // |a b|
        // |c d| = ad - bc
        if len(args) != 4 {
            return "Error: det2x2 requiere 4 argumentos"
        }
        return (toFloat(args[0]) * toFloat(args[3])) - (toFloat(args[1]) * toFloat(args[2]))
    })

    reg("det3x3", func(args ...interface{}) interface{} {
        // determinante de matriz 3x3
        if len(args) != 9 {
            return "Error: det3x3 requiere 9 argumentos"
        }
        a, b, c := toFloat(args[0]), toFloat(args[1]), toFloat(args[2])
        d, e, f := toFloat(args[3]), toFloat(args[4]), toFloat(args[5])
        g, h, i := toFloat(args[6]), toFloat(args[7]), toFloat(args[8])
        return a*(e*i-f*h) - b*(d*i-f*g) + c*(d*h-e*g)
    })

    // --- Producto punto ---
    reg("producto_punto", func(args ...interface{}) interface{} {
        // (x1,y1,z1) · (x2,y2,z2)
        if len(args) != 6 {
            return "Error: producto_punto requiere 6 argumentos"
        }
        return (toFloat(args[0]) * toFloat(args[3])) +
            (toFloat(args[1]) * toFloat(args[4])) +
            (toFloat(args[2]) * toFloat(args[5]))
    })

    // --- Magnitud de vector ---
    reg("magnitud_vector", func(args ...interface{}) interface{} {
        // ||v|| = sqrt(x^2 + y^2 + z^2)
        if len(args) != 3 {
            return "Error: magnitud_vector requiere 3 argumentos"
        }
        return math.Sqrt(math.Pow(toFloat(args[0]), 2) +
            math.Pow(toFloat(args[1]), 2) +
            math.Pow(toFloat(args[2]), 2))
    })
}
