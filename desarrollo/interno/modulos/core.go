package modulos

import (
    "nepa/desarrollo/interno/evaluador"
    "nepa/desarrollo/interno/matematicas"
    // en el futuro: fechas, cadenas, ficheros, etc.
)

// CargarCore inyecta todas las librerías internas al contexto.
// Se llama una sola vez al inicio de Nepa.
func CargarCore(ctx *evaluador.Contexto) {
    if ctx.Funciones == nil {
        ctx.Funciones = map[string]func(...interface{}) interface{}{}
    }

    // Matemáticas básicas
    matematicas.InyectarBasicas(ctx)

    // Estadística
    matematicas.InyectarEstadistica(ctx)

    // Física
    matematicas.InyectarFisica(ctx)

    // Finanzas
    matematicas.InyectarFinanzas(ctx)

    // Álgebra
    matematicas.InyectarAlgebra(ctx)

    // Aquí puedes añadir otras librerías internas en el futuro
    // cadenas.InyectarCadenas(ctx)
    // fechas.InyectarFechas(ctx)
    // ficheros.InyectarFicheros(ctx)
}
