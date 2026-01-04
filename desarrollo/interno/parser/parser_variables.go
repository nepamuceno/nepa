package parser

import (
    "strings"
)

// parseVariable: universal para cualquier tipo
// Ejemplos:
//   variable bit a
//   variable bit a := 1
//   variable caracter c
//   variable entero x, y, z := 0
func parseVariable(linea string) *Nodo {
    if !strings.HasPrefix(linea, "variable ") {
        return nil
    }

    partes := strings.Fields(linea)
    if len(partes) < 3 {
        return nil
    }

    tipo := partes[1]
    nombres := partes[2]
    valor := ""

    if strings.Contains(linea, ":=") {
        idx := strings.Index(linea, ":=")
        nombres = strings.TrimSpace(linea[len("variable ")+len(tipo) : idx])
        valor = strings.TrimSpace(linea[idx+2:])
    }

    return &Nodo{
        Tipo:   "variable",
        Nombre: nombres,
        Valor:  parseValor(valor), // nil → constructor aplica default
        Args:   []interface{}{tipo},
    }
}
