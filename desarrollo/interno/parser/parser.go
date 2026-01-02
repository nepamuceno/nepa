package parser

import (
    "strconv"
    "strings"
)

// Nodo representa un elemento del AST
type Nodo struct {
    Tipo   string        // variable, global, constante, llamada, imprimir, bloque, expresion
    Nombre string        // nombre de variable o función
    Valor  interface{}   // valor literal o expresión
    Args   []interface{} // argumentos de llamadas
}

// Parse convierte líneas validadas en un AST
func Parse(lineas []string) []Nodo {
    var ast []Nodo

    for _, linea := range lineas {
        linea = strings.TrimSpace(linea)
        if linea == "" || strings.HasPrefix(linea, "#") {
            continue
        }

        // Bloques: si:, mientras:, para:, funcion:, etc.
        if strings.HasSuffix(linea, ":") {
            nombre := strings.TrimSuffix(linea, ":")
            ast = append(ast, Nodo{Tipo: "bloque", Nombre: nombre})
            continue
        }

        // --- Asignaciones globales y constantes ---
        if strings.HasPrefix(linea, "global ") && strings.Contains(linea, "=") {
            partes := strings.SplitN(linea[len("global "):], "=", 2)
            if len(partes) == 2 {
                nombre := strings.TrimSpace(partes[0])
                valor := parseValor(strings.TrimSpace(partes[1]))
                ast = append(ast, Nodo{Tipo: "global", Nombre: nombre, Valor: valor})
                continue
            }
        }

        if strings.HasPrefix(linea, "const ") && strings.Contains(linea, "=") {
            partes := strings.SplitN(linea[len("const "):], "=", 2)
            if len(partes) == 2 {
                nombre := strings.TrimSpace(partes[0])
                valor := parseValor(strings.TrimSpace(partes[1]))
                ast = append(ast, Nodo{Tipo: "constante", Nombre: nombre, Valor: valor})
                continue
            }
        }

        // --- Asignación local ---
        if strings.Contains(linea, "=") && !strings.Contains(linea, "==") {
            partes := strings.SplitN(linea, "=", 2)
            if len(partes) == 2 {
                nombre := strings.TrimSpace(partes[0])
                valor := parseValor(strings.TrimSpace(partes[1]))
                ast = append(ast, Nodo{Tipo: "variable", Nombre: nombre, Valor: valor})
                continue
            }
        }

        // --- Llamadas a funciones ---
        if strings.Contains(linea, "(") && strings.HasSuffix(linea, ")") {
            fn := linea[:strings.Index(linea, "(")]
            fn = strings.TrimSpace(fn)
            args := extraerArgs(linea)

            if fn == "imprimir" {
                ast = append(ast, Nodo{Tipo: "imprimir", Nombre: fn, Args: args})
            } else {
                ast = append(ast, Nodo{Tipo: "llamada", Nombre: fn, Args: args})
            }
            continue
        }

        // --- Expresión suelta ---
        ast = append(ast, Nodo{Tipo: "expresion", Valor: parseValor(linea)})
    }

    return ast
}

// extraerArgs obtiene los argumentos de una llamada tipo funcion(arg1, arg2)
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

// parseValor convierte un literal string en su tipo correcto
func parseValor(raw string) interface{} {
    // Booleanos
    if raw == "true" {
        return true
    }
    if raw == "false" {
        return false
    }

    // Nulo
    if raw == "nulo" {
        return nil
    }

    // Enteros
    if i, err := strconv.Atoi(raw); err == nil {
        return i
    }

    // Flotantes
    if f, err := strconv.ParseFloat(raw, 64); err == nil {
        return f
    }

    // Strings con comillas
    if strings.HasPrefix(raw, "\"") && strings.HasSuffix(raw, "\"") {
        return strings.Trim(raw, "\"")
    }

    // Listas/matrices simples: [1,2,3]
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

    // Si no coincide con nada, devolver como string
    return raw
}
