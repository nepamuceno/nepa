package parser

import (
    "strings"
)

// parseAsignar: Maneja declaraciones de variables con o sin valor inicial.
// Sintaxis: asignar <tipo> <var[,var2,...]> [:= <valor>]
// Soporta múltiples nombres separados por coma, punteros y matrices (incluye notación [][]).
func parseAsignar(linea string) *Nodo {
    if !strings.HasPrefix(linea, "asignar ") {
        return nil
    }

    def := strings.TrimSpace(strings.TrimPrefix(linea, "asignar "))
    if def == "" {
        return nil
    }

    campos := strings.Fields(def)
    if len(campos) < 2 {
        return nil
    }

    // Extraer tokens de tipo
    tipoTokens, nextIdx := extraerTipoTokens(campos)
    if len(tipoTokens) == 0 {
        return nil
    }

    resto := strings.TrimSpace(strings.Join(campos[nextIdx:], " "))

    // Ver si hay asignación
    var nombresParte, valorParte string
    if strings.Contains(resto, ":=") {
        partes := strings.SplitN(resto, ":=", 2)
        nombresParte = strings.TrimSpace(partes[0])
        valorParte = strings.TrimSpace(partes[1])
    } else {
        nombresParte = strings.TrimSpace(resto)
        valorParte = ""
    }

    nombres := strings.Split(nombresParte, ",")
    var nodos []Nodo
    for _, n := range nombres {
        nombre := strings.TrimSpace(n)
        if nombre == "" {
            continue
        }
        var valor interface{}
        if valorParte != "" {
            valor = parseValor(valorParte)
        } else {
            valor = nil
        }
        nodos = append(nodos, Nodo{
            Tipo:   "asignar",
            Nombre: nombre,
            Args:   tipoTokens,
            Valor:  valor,
        })
    }

    if len(nodos) == 1 {
        return &nodos[0]
    }

    return &Nodo{
        Tipo:   "asignaciones",
        Nombre: "",
        Args:   tipoTokens,
        Valor:  nodos,
    }
}

// extraerTipoTokens toma una secuencia de tokens y devuelve los que forman el tipo compuesto.
// Ejemplos válidos:
// - ["entero", "x"] -> ["entero"]
// - ["puntero", "entero", "x"] -> ["puntero","entero"]
// - ["matriz", "entero", "m"] -> ["matriz","entero"]
// - ["matriz[][]", "real", "tabla"] -> ["matriz[][]","real"]
// - ["puntero", "matriz", "caracter", "pm"] -> ["puntero","matriz","caracter"]
func extraerTipoTokens(campos []string) ([]interface{}, int) {
    var tokens []interface{}
    i := 0

    for i < len(campos) {
        tok := campos[i]

        // Notación de matriz con dimensiones explícitas ([][], [][][], etc.)
        if strings.HasPrefix(tok, "[]") {
            tokens = append(tokens, tok)
            i++
            continue
        }

        // Palabras clave de tipo compuesto
        if tok == "puntero" || tok == "matriz" || tok == "diccionario" || tok == "objeto" || tok == "lista" {
            tokens = append(tokens, tok)
            i++
            continue
        }

        // Tipo base (validado contra TiposBase centralizado)
        if TiposBase[tok] {
            tokens = append(tokens, tok)
            i++
            break
        }

        break
    }

    return tokens, i
}

// --- Registro en Parsers ---
func init() {
    Parsers["asignar"] = func(linea string) *Nodo {
        return parseAsignar(linea)
    }
}
