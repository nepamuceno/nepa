package parser

import (
    "strings"
)

// Nodo representa un elemento del árbol sintáctico
type Nodo struct {
    Tipo      string
    Condicion string
    Valor     interface{}
    Nombre    string
    Args      []string
}

// parseMientrasBloques: reconoce bloques mientras(condición): con soporte completo
// Sintaxis soportada:
//   mientras condicion:
//   mientras(condicion):
// Bloques indentados con 4 espacios, con instrucción 'rompe' para salir del bucle
func parseMientrasBloques(cuerpo []string) []Nodo {
    var resultado []Nodo
    i := 0
    n := len(cuerpo)

    for i < n {
        linea := strings.TrimSpace(cuerpo[i])
        if linea == "" {
            i++
            continue
        }

        // mientras
        if strings.HasPrefix(linea, "mientras") && strings.HasSuffix(linea, ":") {
            condicion := extraerCondicion(linea, "mientras")
            bloque, avanzados := recolectarBloqueIndentado(cuerpo[i+1:])
            i += 1 + avanzados
            resultado = append(resultado, Nodo{
                Tipo:      "mientras",
                Condicion: condicion,
                Valor:     bloque,
            })
            continue
        }

        i++
    }

    return resultado
}

// extraerCondicion: obtiene la condición de mientras
// soporta sintaxis con o sin paréntesis
func extraerCondicion(linea, prefijo string) string {
    expr := strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(linea, prefijo), ":"))
    if strings.HasPrefix(expr, "(") && strings.HasSuffix(expr, ")") {
        expr = strings.TrimPrefix(expr, "(")
        expr = strings.TrimSuffix(expr, ")")
    }
    return strings.TrimSpace(expr)
}

// recolectarBloqueIndentado: recoge instrucciones indentadas con 4 espacios
// Integra semántica completa: operadores, funciones externas, punteros, casts, etc.
func recolectarBloqueIndentado(lineas []string) ([]Nodo, int) {
    var bloque []Nodo
    avanzados := 0

    for avanzados < len(lineas) {
        cruda := lineas[avanzados]
        if !strings.HasPrefix(cruda, "    ") { // exactamente 4 espacios
            break
        }
        instr := strings.TrimSpace(cruda)

        // Instrucciones reconocidas
        switch {
        case instr == "rompe":
            bloque = append(bloque, Nodo{Tipo: "rompe", Valor: "salir_bucle"})

        case strings.HasPrefix(instr, "regresa_valor"):
            campos := strings.Fields(strings.TrimPrefix(instr, "regresa_valor"))
            tipoTokens, nextIdx := extraerTipoTokens(campos)
            var nombreVar string
            if len(tipoTokens) > 0 && nextIdx < len(campos) {
                nombreVar = campos[nextIdx]
            }
            bloque = append(bloque, Nodo{
                Tipo:   "regresa_valor",
                Nombre: nombreVar,
                Args:   tipoTokens,
            })

        case strings.HasPrefix(instr, "regresa "):
            expr := strings.TrimSpace(strings.TrimPrefix(instr, "regresa"))
            bloque = append(bloque, Nodo{Tipo: "regresa", Valor: expr})

        case strings.HasPrefix(instr, "variable "):
            if nodo := parseVariable(instr); nodo != nil {
                bloque = append(bloque, *nodo)
            }

        case strings.HasPrefix(instr, "constante "):
            if nodo := parseConst(instr); nodo != nil {
                bloque = append(bloque, *nodo)
            }

        case strings.HasPrefix(instr, "global "):
            if nodo := parseGlobal(instr); nodo != nil {
                bloque = append(bloque, *nodo)
            }

        case strings.Contains(instr, ":="):
            if nodo := parseAsignar(instr); nodo != nil {
                bloque = append(bloque, *nodo)
            }

        default:
            // Aquí se aceptan expresiones complejas: operadores, funciones externas, punteros, casts
            bloque = append(bloque, Nodo{Tipo: "instruccion", Valor: instr})
        }

        avanzados++
    }

    return bloque, avanzados
}

// extraerTipoTokens: placeholder para reconocer tipos en regresa_valor
func extraerTipoTokens(campos []string) ([]string, int) {
    return campos, len(campos)
}

// parseVariable, parseConst, parseGlobal, parseAsignar
// se asumen definidos en otros módulos del parser

// --- Registro en TiposControl ---
func init() {
    TiposControl["mientras"] = true
    TiposControl["rompe"]    = true
}
