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

// parseSiBloques: reconoce bloques si_es, pero_si, si_no y ayuda
// Soporta condiciones complejas con operadores, funciones externas, punteros, casts, etc.
func parseSiBloques(cuerpo []string) []Nodo {
    var resultado []Nodo
    i := 0
    n := len(cuerpo)

    for i < n {
        linea := strings.TrimSpace(cuerpo[i])
        if linea == "" {
            i++
            continue
        }

        // si_es
        if strings.HasPrefix(linea, "si_es") && strings.HasSuffix(linea, ":") {
            condicion := extraerCondicion(linea, "si_es")
            bloque, avanzados := recolectarBloqueIndentado(cuerpo[i+1:])
            i += 1 + avanzados
            resultado = append(resultado, Nodo{
                Tipo:      "si_es",
                Condicion: condicion,
                Valor:     bloque,
            })
            continue
        }

        // pero_si
        if strings.HasPrefix(linea, "pero_si") && strings.HasSuffix(linea, ":") {
            condicion := extraerCondicion(linea, "pero_si")
            bloque, avanzados := recolectarBloqueIndentado(cuerpo[i+1:])
            i += 1 + avanzados
            resultado = append(resultado, Nodo{
                Tipo:      "pero_si",
                Condicion: condicion,
                Valor:     bloque,
            })
            continue
        }

        // si_no
        if linea == "si_no:" {
            bloque, avanzados := recolectarBloqueIndentado(cuerpo[i+1:])
            i += 1 + avanzados
            resultado = append(resultado, Nodo{
                Tipo:  "si_no",
                Valor: bloque,
            })
            continue
        }

        // ayuda (bloque opcional de soporte/debug)
        if linea == "ayuda:" {
            bloque, avanzados := recolectarBloqueIndentado(cuerpo[i+1:])
            i += 1 + avanzados
            resultado = append(resultado, Nodo{
                Tipo:  "ayuda",
                Valor: bloque,
            })
            continue
        }

        i++
    }

    return resultado
}

// extraerCondicion: obtiene la condición de si_es/pero_si
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
// Aquí se integran todas las semánticas discutidas: operadores, funciones, punteros, casts, etc.
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
        case instr == "romper":
            bloque = append(bloque, Nodo{Tipo: "romper", Valor: "panico"})

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
    TiposControl["si_es"] = true
    TiposControl["pero_si"] = true
    TiposControl["si_no"] = true
    TiposControl["ayuda"] = true
}
