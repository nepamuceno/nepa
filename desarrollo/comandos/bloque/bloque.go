package bloque

import (
    "fmt"

    "nepa/desarrollo/interno/evaluador"
    "nepa/desarrollo/interno/parser"
)

// init registra el handler para nodos tipo "bloque"
func init() {
    evaluador.Registrar("bloque", func(n parser.Nodo, ctx *evaluador.Contexto) {
        // Un bloque normalmente contiene subnodos en n.Args
        count := 0
        for i, arg := range n.Args {
            // Verificamos si el argumento es un Nodo
            if hijo, ok := arg.(parser.Nodo); ok {
                _, err := evaluador.EjecutarConContexto(
                    []parser.Nodo{hijo},
                    ctx.Variables,
                    ctx.Globales,
                    ctx.Constantes,
                    fmt.Sprintf("bloque:%d", i+1), // identificador de archivo/posición
                )
                if err != nil {
                    fmt.Printf("Error: bloque:%d: %v\n", i+1, err)
                    return // detener ejecución del bloque en caso de error fatal
                }
                count++
            }
        }
        fmt.Printf("✔ Bloque ejecutado (%d nodos)\n", count)
    })
}
