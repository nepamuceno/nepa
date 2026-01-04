package parser

import (
    "strconv"
    "strings"
)

// Nodo representa un elemento del AST
type Nodo struct {
    Tipo   string        // variable, global, constante, llamada, bloque, expresion, asignar
    Nombre string        // nombre de variable o función
    Valor  interface{}   // valor literal o expresión
    Args   []interface{} // argumentos de llamadas o metadatos (ej. tipo de variable)
}

// Parse convierte líneas validadas en un AST
func Parse(lineas []string) []Nodo {
    var ast []Nodo

    for _, linea := range lineas {
        linea = strings.TrimSpace(linea)
        if linea == "" || strings.HasPrefix(linea, "#") {
            continue
        }

        tokens := strings.Fields(linea)
        if len(tokens) == 0 {
            continue
        }
        token := tokens[0]

        // --- Caso especial: si_es ---
        if token == "si_es" {
            if strings.HasSuffix(linea, ":") {
                ast = append(ast, Nodo{Tipo: "bloque", Nombre: strings.TrimSuffix(linea, ":")})
            } else {
                ast = append(ast, Nodo{Tipo: "expresion", Valor: parseValor(linea)})
            }
            continue
        }

        // --- Bloques generales ---
        if strings.HasSuffix(linea, ":") {
            ast = append(ast, Nodo{Tipo: "bloque", Nombre: strings.TrimSuffix(linea, ":")})
            continue
        }

        // --- Globales ---
        if nodo := parseGlobal(linea); nodo != nil {
            ast = append(ast, *nodo)
            continue
        }

        // --- Constantes ---
        if nodo := parseConst(linea); nodo != nil {
            ast = append(ast, *nodo)
            continue
        }

        // --- Variables ---
        if nodo := parseVariable(linea); nodo != nil {
            ast = append(ast, *nodo)
            continue
        }

        // --- Asignaciones ---
        if nodo := parseAsignar(linea); nodo != nil {
            ast = append(ast, *nodo)
            continue
        }

        // --- Llamadas ---
        if nodo := parseLlamada(linea); nodo != nil {
            ast = append(ast, *nodo)
            continue
        }

        // --- Expresión suelta ---
        ast = append(ast, Nodo{Tipo: "expresion", Valor: parseValor(linea)})
    }

    return ast
}

// --- Funciones auxiliares ---

func parseGlobal(linea string) *Nodo {
    if strings.HasPrefix(linea, "global ") && strings.Contains(linea, "=") {
        partes := strings.SplitN(linea[len("global "):], "=", 2)
        if len(partes) == 2 {
            return &Nodo{Tipo: "global", Nombre: strings.TrimSpace(partes[0]), Valor: parseValor(strings.TrimSpace(partes[1]))}
        }
    }
    return nil
}

func parseConst(linea string) *Nodo {
    if strings.HasPrefix(linea, "const ") && strings.Contains(linea, "=") {
        partes := strings.SplitN(linea[len("const "):], "=", 2)
        if len(partes) == 2 {
            return &Nodo{Tipo: "constante", Nombre: strings.TrimSpace(partes[0]), Valor: parseValor(strings.TrimSpace(partes[1]))}
        }
    }
    return nil
}

func parseLlamada(linea string) *Nodo {
    if strings.Contains(linea, "(") && strings.HasSuffix(linea, ")") {
        fn := strings.TrimSpace(linea[:strings.Index(linea, "(")])
        args := extraerArgs(linea)
        return &Nodo{Tipo: "llamada", Nombre: fn, Args: args}
    }
    if strings.Contains(linea, " ") {
        partes := strings.SplitN(linea, " ", 2)
        fn := strings.TrimSpace(partes[0])
        resto := strings.TrimSpace(partes[1])
        args := []interface{}{}
        for _, p := range strings.Split(resto, ",") {
            arg := strings.TrimSpace(p)
            if arg != "" {
                args = append(args, parseValor(arg))
            }
        }
        return &Nodo{Tipo: "llamada", Nombre: fn, Args: args}
    }
    return nil
}

// --- Utilidades ---

func extraerArgs(linea string) []interface{} {
    ini := strings.Index(linea, "(")
    fin := strings.LastIndex(linea, ")")
    if ini == -1 || fin == -1 || fin <= ini {
        return nil
    }
    contenido := linea[ini+1 : fin]
    partes := strings.Split(contenido, ",")
    var args []interface{}
    for _, p := range partes {
        arg := strings.TrimSpace(p)
        if arg != "" {
            args = append(args, parseValor(arg))
        }
    }
    return args
}

func parseValor(raw string) interface{} {
    if raw == "true" {
        return true
    }
    if raw == "false" {
        return false
    }
    if raw == "nulo" {
        return nil
    }
    if i, err := strconv.Atoi(raw); err == nil {
        return i
    }
    if f, err := strconv.ParseFloat(raw, 64); err == nil {
        return f
    }
    if strings.HasPrefix(raw, "\"") && strings.HasSuffix(raw, "\"") {
        return strings.Trim(raw, "\"")
    }
    if strings.HasPrefix(raw, "'") && strings.HasSuffix(raw, "'") {
        contenido := strings.Trim(raw, "'")
        if len(contenido) == 1 {
            return rune(contenido[0])
        }
        return contenido
    }
    if strings.HasPrefix(raw, "[") && strings.HasSuffix(raw, "]") {
        contenido := strings.Trim(raw, "[]")
        partes := strings.Split(contenido, ",")
        var lista []interface{}
        for _, p := range partes {
            val := strings.TrimSpace(p)
            if val != "" {
                lista = append(lista, parseValor(val))
            }
        }
        return lista
    }
    // Si no coincide con literal, se devuelve tal cual (posible nombre de variable)
    return raw
}
