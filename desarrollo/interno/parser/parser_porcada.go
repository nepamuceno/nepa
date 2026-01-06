package parser

import (
    "strings"
)

// Nodo representa un elemento del árbol sintáctico
type Nodo struct {
    Tipo      string
    Condicion string
    Init      string
    Post      string
    Valor     interface{}
    Nombre    string
    Args      []string
}

// parsePorcadaBloques: reconoce bloques porcada(condición, init, post): con soporte completo
// Sintaxis soportada:
//   porcada condicion:
//   porcada(condicion):
//   porcada(condicion, init, post):
// Bloques indentados con 4 espacios, con instrucción 'rompe' para salir del bucle
func parsePorcadaBloques(cuerpo []string) []Nodo {
    var resultado []Nodo
    i := 0
    n := len(cuerpo)

    for i < n {
        linea := strings.TrimSpace(cuerpo[i])
        if linea == "" {
            i++
            continue
        }

        // porcada
        if strings.HasPrefix(linea, "porcada") && strings.HasSuffix(linea, ":") {
            condicion, init, post := extraerPorcada(linea)
            bloque, avanzados := recolectarBloqueIndentado(cuerpo[i+1:])
            i += 1 + avanzados
            resultado = append(resultado, Nodo{
                Tipo:      "porcada",
                Condicion: condicion,
                Init:      init,
                Post:      post,
                Valor:     bloque,
            })
            continue
        }

        i++
    }

    return resultado
}

// extraerPorcada: obtiene la condición, inicialización y post-expresión de porcada
func extraerPorcada(linea string) (string, string, string) {
    expr := strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(linea, "porcada"), ":"))
    if strings.HasPrefix(expr, "(") && strings.HasSuffix(expr, ")") {
        expr = strings.TrimPrefix(expr, "(")
        expr = strings.TrimSuffix(expr, ")")
    }
    partes := strings.Split(expr, ",")
    condicion, init, post := "", "", ""
    if len(partes) > 0 {
        condicion = strings.TrimSpace(partes[0])
    }
    if len(partes) > 1 {
        init = strings.TrimSpace(partes[1])
    }
    if len(partes) > 2 {
        post = strings.TrimSpace(partes[2])
    }
    return condicion, init, post
}

// recolectarBloqueIndentado: recoge instrucciones indentadas con 4 espacios
func recolectarBloqueIndentado(lineas []string) ([]Nodo, int) {
    var bloque []Nodo
    avanzados := 0

    for avanzados < len(lineas) {
        cruda := lineas[avanzados]
        if !strings.HasPrefix(cruda, "    ") {
            break
        }
        instr := strings.TrimSpace(cruda)

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
    TiposControl["porcada"] = true
    TiposControl["rompe"]   = true
}
