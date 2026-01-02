package matematicas

import (
    "fmt"
    "strconv"

    "nepa/desarrollo/interno/evaluador"
)

// InyectarTodo carga todas las funciones matemáticas en el contexto
func InyectarTodo(ctx *evaluador.Contexto) {
    if ctx.Funciones == nil {
        ctx.Funciones = map[string]func(...interface{}) interface{}{}
    }

    // Matemáticas básicas
    InyectarBasicas(ctx)

    // Estadística
    InyectarEstadistica(ctx)

    // Física
    InyectarFisica(ctx)

    // Finanzas
    InyectarFinanzas(ctx)

    // Álgebra
    InyectarAlgebra(ctx)
}

// toFloat convierte cualquier tipo soportado a float64
func toFloat(v interface{}) float64 {
    switch val := v.(type) {
    case int:
        return float64(val)
    case int32:
        return float64(val)
    case int64:
        return float64(val)
    case float32:
        return float64(val)
    case float64:
        return val
    case string:
        f, err := strconv.ParseFloat(val, 64)
        if err != nil {
            fmt.Printf("Error: no se pudo convertir '%s' a número\n", val)
            return 0
        }
        return f
    default:
        fmt.Printf("Error: tipo no soportado (%T)\n", v)
        return 0
    }
}
