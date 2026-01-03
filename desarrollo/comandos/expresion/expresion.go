package expresion

import (
    "fmt"

    "nepa/desarrollo/interno/evaluador"
    "nepa/desarrollo/interno/parser"
)

// init registra el handler para nodos tipo "expresion"
func init() {
    evaluador.Registrar("expresion", func(n parser.Nodo, ctx *evaluador.Contexto) {
        switch v := n.Valor.(type) {
        case string:
            // Si es un literal de texto, imprimir directo
            fmt.Println(v)
        case int, int64, float64, bool:
            // Tipos básicos también se imprimen directo
            fmt.Println(v)
        default:
            // Para otros casos, intentar evaluar como expresión
            resultado, err := evaluador.Eval(fmt.Sprintf("%v", n.Valor))
            if err != nil {
                fmt.Printf("⚠️ Error evaluando expresión '%v': %v\n", n.Valor, err)
                return
            }
            fmt.Printf("✔ Expresión evaluada: %v → %v\n", n.Valor, resultado)
        }
    })
}
