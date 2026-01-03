package llamada

import (
    "fmt"

    "nepa/desarrollo/interno/evaluador"
    "nepa/desarrollo/interno/parser"
)

// init registra el handler para nodos tipo "llamada"
func init() {
    evaluador.Registrar("llamada", func(n parser.Nodo, ctx *evaluador.Contexto) {
        // Buscar la función en el contexto
        fn, ok := ctx.Funciones[n.Nombre]
        if !ok {
            fmt.Printf("⚠️ Función '%s' no existe\n", n.Nombre)
            return
        }

        // Preparar argumentos evaluados
        var args []interface{}
        for _, arg := range n.Args {
            // Si el argumento ya es un valor básico, úsalo directo
            switch v := arg.(type) {
            case string, int, int64, float64, bool:
                args = append(args, v)
            default:
                // Intentar evaluar como expresión
                res, err := evaluador.Eval(fmt.Sprintf("%v", v))
                if err != nil {
                    fmt.Printf("⚠️ Error evaluando argumento '%v': %v\n", v, err)
                    return
                }
                args = append(args, res)
            }
        }

        // Ejecutar la función con los argumentos del nodo
        resultado := fn(args...)

        // Mostrar el resultado de la llamada
        fmt.Printf("✔ Llamada a función '%s' → %v\n", n.Nombre, resultado)
    })
}
