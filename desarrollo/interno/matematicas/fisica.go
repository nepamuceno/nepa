package matematicas

import (
    "fmt"
    "math"

    "nepa/desarrollo/interno/evaluador"
)

// InyectarFisica agrega funciones físicas al contexto
func InyectarFisica(ctx *evaluador.Contexto) {
    if ctx.Funciones == nil {
        ctx.Funciones = map[string]func(...interface{}) interface{}{}
    }

    reg := func(n string, f func(...interface{}) interface{}) {
        ctx.Funciones[n] = f
    }

    // --- Energía y movimiento ---
    reg("energia_relativista", func(args ...interface{}) interface{} {
        if len(args) != 1 {
            return "Error: energia_relativista requiere 1 argumento (masa)"
        }
        return toFloat(args[0]) * math.Pow(299792458.0, 2)
    })

    reg("energia_cinetica", func(args ...interface{}) interface{} {
        if len(args) != 2 {
            return "Error: energia_cinetica requiere 2 argumentos (masa, velocidad)"
        }
        return 0.5 * toFloat(args[0]) * math.Pow(toFloat(args[1]), 2)
    })

    reg("trabajo", func(args ...interface{}) interface{} {
        if len(args) != 3 {
            return "Error: trabajo requiere 3 argumentos (fuerza, distancia, ángulo)"
        }
        ang := toFloat(args[2]) * math.Pi / 180
        return toFloat(args[0]) * toFloat(args[1]) * math.Cos(ang)
    })

    // --- Proyectiles ---
    reg("caida_libre", func(args ...interface{}) interface{} {
        if len(args) != 1 {
            return "Error: caida_libre requiere 1 argumento (tiempo)"
        }
        return 0.5 * 9.80665 * math.Pow(toFloat(args[0]), 2)
    })

    reg("proyectil_pos", func(args ...interface{}) interface{} {
        if len(args) != 3 {
            return "Error: proyectil_pos requiere 3 argumentos (velocidad inicial, ángulo, tiempo)"
        }
        v0 := toFloat(args[0])
        ang := toFloat(args[1]) * math.Pi / 180
        t := toFloat(args[2])
        x := v0 * math.Cos(ang) * t
        y := (v0 * math.Sin(ang) * t) - (0.5 * 9.80665 * t * t)
        return fmt.Sprintf("%f,%f", x, y)
    })

    reg("proyectil_x", func(args ...interface{}) interface{} {
        if len(args) != 3 {
            return "Error: proyectil_x requiere 3 argumentos (velocidad inicial, ángulo, tiempo)"
        }
        return toFloat(args[0]) * math.Cos(toFloat(args[1])*math.Pi/180) * toFloat(args[2])
    })

    reg("proyectil_y", func(args ...interface{}) interface{} {
        if len(args) != 3 {
            return "Error: proyectil_y requiere 3 argumentos (velocidad inicial, ángulo, tiempo)"
        }
        v0, ang, t := toFloat(args[0]), toFloat(args[1])*math.Pi/180, toFloat(args[2])
        return (v0 * math.Sin(ang) * t) - (0.5 * 9.80665 * t * t)
    })

    // --- Física general ---
    reg("densidad", func(args ...interface{}) interface{} {
        if len(args) != 2 {
            return "Error: densidad requiere 2 argumentos (masa, volumen)"
        }
        if toFloat(args[1]) == 0 {
            return "Error: volumen no puede ser 0"
        }
        return toFloat(args[0]) / toFloat(args[1])
    })

    reg("presion_gas", func(args ...interface{}) interface{} {
        if len(args) != 3 {
            return "Error: presion_gas requiere 3 argumentos (n moles, temperatura, volumen)"
        }
        if toFloat(args[2]) == 0 {
            return "Error: volumen no puede ser 0"
        }
        return (toFloat(args[0]) * 8.314462 * toFloat(args[1])) / toFloat(args[2])
    })

    reg("fuerza_gravitatoria", func(args ...interface{}) interface{} {
        if len(args) != 3 {
            return "Error: fuerza_gravitatoria requiere 3 argumentos (masa1, masa2, distancia)"
        }
        if toFloat(args[2]) == 0 {
            return "Error: distancia no puede ser 0"
        }
        G := 6.67430e-11
        return G * (toFloat(args[0]) * toFloat(args[1])) / math.Pow(toFloat(args[2]), 2)
    })

    reg("ley_ohm_v", func(args ...interface{}) interface{} {
        if len(args) != 2 {
            return "Error: ley_ohm_v requiere 2 argumentos (corriente, resistencia)"
        }
        return toFloat(args[0]) * toFloat(args[1])
    })

    reg("magnitud_vector", func(args ...interface{}) interface{} {
        if len(args) != 3 {
            return "Error: magnitud_vector requiere 3 argumentos (x, y, z)"
        }
        return math.Sqrt(math.Pow(toFloat(args[0]), 2) +
            math.Pow(toFloat(args[1]), 2) +
            math.Pow(toFloat(args[2]), 2))
    })
}
