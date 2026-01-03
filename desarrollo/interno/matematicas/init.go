package matematicas

import (
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

    // Sistemas de numeración y operaciones bit a bit
    InyectarBases(ctx)
}
