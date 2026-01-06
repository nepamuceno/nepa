package parser

import (
    "strconv"
    "strings"
)

// Nodo representa un elemento del AST
type Nodo struct {
    Tipo   string        // variable, global, constante, llamada, bloque, expresion, asignar, conversion, lista, indice, funcion
    Nombre string        // nombre de variable, función u operador
    Valor  interface{}   // valor literal, expresión o cuerpo de bloque ([]Nodo)
    Args   []interface{} // argumentos de llamadas, índices, operadores o parámetros de función
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

        // --- Funciones estilo Python: funcion nombre(args): ---
        if strings.HasPrefix(linea, "funcion ") && strings.HasSuffix(linea, ":") {
            if nodo := parseFuncion(linea); nodo != nil {
                ast = append(ast, *nodo)
                continue
            }
        }

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

// --- Funciones ---

// funcion nombre(args):
func parseFuncion(linea string) *Nodo {
    def := strings.TrimSuffix(strings.TrimSpace(linea), ":")
    def = strings.TrimPrefix(def, "funcion ")
    nombre := def
    var params []interface{}

    if strings.Contains(def, "(") && strings.HasSuffix(def, ")") {
        nombre = strings.TrimSpace(def[:strings.Index(def, "(")])
        contenido := def[strings.Index(def, "(")+1 : strings.LastIndex(def, ")")]
        for _, p := range splitArgs(contenido) {
            p = strings.TrimSpace(p)
            if p != "" {
                params = append(params, p)
            }
        }
    }

    return &Nodo{
        Tipo:   "funcion",
        Nombre: nombre,
        Args:   params,
        Valor:  []Nodo(nil), // cuerpo se llena después si manejas indentación
    }
}

// --- Llamadas ---

// Llamadas con y sin paréntesis: imprimir("x") o imprimir "x"
func parseLlamada(linea string) *Nodo {
    linea = strings.TrimSpace(linea)

    // Con paréntesis
    if strings.Contains(linea, "(") && strings.HasSuffix(linea, ")") && !strings.HasPrefix(linea, "(") {
        nombre := strings.TrimSpace(strings.SplitN(linea, "(", 2)[0])
        return &Nodo{Tipo: "llamada", Nombre: nombre, Args: extraerArgs(linea)}
    }

    // Sin paréntesis: imprimir "hola"
    campos := strings.Fields(linea)
    if len(campos) >= 2 {
        nombre := campos[0]
        resto := strings.TrimSpace(linea[len(nombre):])
        var args []interface{}
        partes := splitArgs(resto)
        if len(partes) == 0 {
            partes = []string{resto}
        }
        for _, p := range partes {
            p = strings.TrimSpace(p)
            if p != "" {
                args = append(args, parseValor(p))
            }
        }
        return &Nodo{Tipo: "llamada", Nombre: nombre, Args: args}
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
    partes := splitArgs(contenido)
    var args []interface{}
    for _, p := range partes {
        arg := strings.TrimSpace(p)
        if arg != "" {
            args = append(args, parseValor(arg))
        }
    }
    return args
}

// splitArgs separa por comas respetando anidación de () y []
func splitArgs(s string) []string {
    var res []string
    nivelPar, nivelCor := 0, 0
    inicio := 0
    for i := 0; i < len(s); i++ {
        c := s[i]
        switch c {
        case '(':
            nivelPar++
        case ')':
            if nivelPar > 0 {
                nivelPar--
            }
        case '[':
            nivelCor++
        case ']':
            if nivelCor > 0 {
                nivelCor--
            }
        case ',':
            if nivelPar == 0 && nivelCor == 0 {
                res = append(res, strings.TrimSpace(s[inicio:i]))
                inicio = i + 1
            }
        }
    }
    final := strings.TrimSpace(s[inicio:])
    if final != "" {
        res = append(res, final)
    }
    return res
}

// --- Valores y expresiones ---

func parseValor(raw string) interface{} {
    raw = strings.TrimSpace(raw)

    // Literales básicos
    if raw == "verdadero" { return true }
    if raw == "falso" { return false }
    if raw == "nada" { return nil }

    // Números
    if i, err := strconv.Atoi(raw); err == nil { return i }
    if f, err := strconv.ParseFloat(raw, 64); err == nil { return f }

    // Cadenas
    if strings.HasPrefix(raw, "\"") && strings.HasSuffix(raw, "\"") {
        return strings.Trim(raw, "\"")
    }
    if strings.HasPrefix(raw, "'") && strings.HasSuffix(raw, "'") {
        contenido := strings.Trim(raw, "'")
        if len(contenido) == 1 { return rune(contenido[0]) }
        return contenido
    }

    // Listas / matrices
    if strings.HasPrefix(raw, "[") && strings.HasSuffix(raw, "]") {
        contenido := strings.Trim(raw, "[]")
        partes := splitArgs(contenido)
        var lista []interface{}
        for _, p := range partes {
            val := strings.TrimSpace(p)
            if val != "" {
                lista = append(lista, parseValor(val))
            }
        }
        return lista
    }

    // Conversión
    if strings.HasPrefix(raw, "(") && strings.Contains(raw, ")") {
        fin := strings.Index(raw, ")")
        if fin > 0 {
            tipo := strings.TrimSpace(raw[1:fin])
            resto := strings.TrimSpace(raw[fin+1:])
            return &Nodo{Tipo:"conversion", Nombre:tipo, Valor:parseValor(resto)}
        }
    }

    // Índices: a[1], matriz[i][j]
    if strings.Contains(raw, "[") && strings.HasSuffix(raw, "]") && !strings.HasPrefix(raw, "[") {
        nombre := strings.SplitN(raw, "[", 2)[0]
        indicesRaw := raw[len(nombre):]
        var indices []interface{}
        for strings.Contains(indicesRaw, "[") {
            ini := strings.Index(indicesRaw, "[")
            fin := strings.Index(indicesRaw, "]")
            if ini == -1 || fin == -1 || fin <= ini {
                break
            }
            idx := strings.TrimSpace(indicesRaw[ini+1 : fin])
            indices = append(indices, parseValor(idx))
            indicesRaw = indicesRaw[fin+1:]
        }
        return &Nodo{Tipo:"indice", Nombre:strings.TrimSpace(nombre), Args:indices}
    }

    // Llamada: nombre(args)
    if strings.Contains(raw, "(") && strings.HasSuffix(raw, ")") && !strings.HasPrefix(raw, "(") {
        nombre := strings.SplitN(raw, "(", 2)[0]
        return &Nodo{Tipo:"llamada", Nombre:strings.TrimSpace(nombre), Args:extraerArgs(raw)}
    }

    // Expresiones con operadores
    if contieneOperadores(raw) {
        return parseExpr(raw)
    }

    // Valor por defecto: variable o literal desconocido
    return raw
}

// --- Subparser con precedencia ---
// Orden: OR lógico (o) > AND lógico (y) > bit OR (|) > bit XOR (^) > bit AND (&) > SHIFT (<< >>) > POW (**) > ADD/SUB (+ -) > MUL/DIV (* /) > UNARIOS (! ~ ++ --) > COMPARACIONES (== != < > <= >=) > FACTOR

func parseExpr(expr string) interface{} {
    tokens := tokenize(expr)
    pos := 0
    return parseLogOr(tokens, &pos)
}

// Lógicos en español: o (OR), y (AND)
func parseLogOr(tokens []string, pos *int) interface{} {
    node := parseLogAnd(tokens, pos)
    for *pos < len(tokens) {
        op := tokens[*pos]
        if op != "o" {
            break
        }
        *pos++
        right := parseLogAnd(tokens, pos)
        node = &Nodo{Tipo:"expresion", Nombre:op, Args:[]interface{}{node, right}}
    }
    return node
}

func parseLogAnd(tokens []string, pos *int) interface{} {
    node := parseBitOr(tokens, pos)
    for *pos < len(tokens) {
        op := tokens[*pos]
        if op != "y" {
            break
        }
        *pos++
        right := parseBitOr(tokens, pos)
        node = &Nodo{Tipo:"expresion", Nombre:op, Args:[]interface{}{node, right}}
    }
    return node
}

func parseBitOr(tokens []string, pos *int) interface{} {
    node := parseBitXor(tokens, pos)
    for *pos < len(tokens) {
        op := tokens[*pos]
        if op != "|" {
            break
        }
        *pos++
        right := parseBitXor(tokens, pos)
        node = &Nodo{Tipo:"expresion", Nombre:op, Args:[]interface{}{node, right}}
    }
    return node
}

func parseBitXor(tokens []string, pos *int) interface{} {
    node := parseBitAnd(tokens, pos)
    for *pos < len(tokens) {
        op := tokens[*pos]
        if op != "^" {
            break
        }
        *pos++
        right := parseBitAnd(tokens, pos)
        node = &Nodo{Tipo:"expresion", Nombre:op, Args:[]interface{}{node, right}}
    }
    return node
}

func parseBitAnd(tokens []string, pos *int) interface{} {
    node := parseShift(tokens, pos)
    for *pos < len(tokens) {
        op := tokens[*pos]
        if op != "&" {
            break
        }
        *pos++
        right := parseShift(tokens, pos)
        node = &Nodo{Tipo:"expresion", Nombre:op, Args:[]interface{}{node, right}}
    }
    return node
}

func parseShift(tokens []string, pos *int) interface{} {
    node := parsePow(tokens, pos)
    for *pos < len(tokens) {
        op := tokens[*pos]
        if op != "<<" && op != ">>" {
            break
        }
        *pos++
        right := parsePow(tokens, pos)
        node = &Nodo{Tipo:"expresion", Nombre:op, Args:[]interface{}{node, right}}
    }
    return node
}

func parsePow(tokens []string, pos *int) interface{} {
    node := parseAddSub(tokens, pos)
    for *pos < len(tokens) {
        op := tokens[*pos]
        if op != "**" {
            break
        }
        *pos++
        right := parseAddSub(tokens, pos)
        node = &Nodo{Tipo:"expresion", Nombre:op, Args:[]interface{}{node, right}}
    }
    return node
}

func parseAddSub(tokens []string, pos *int) interface{} {
    node := parseMulDiv(tokens, pos)
    for *pos < len(tokens) {
        op := tokens[*pos]
        if op != "+" && op != "-" {
            break
        }
        *pos++
        right := parseMulDiv(tokens, pos)
        node = &Nodo{Tipo:"expresion", Nombre:op, Args:[]interface{}{node, right}}
    }
    return node
}

func parseMulDiv(tokens []string, pos *int) interface{} {
    node := parseUnary(tokens, pos)
    for *pos < len(tokens) {
        op := tokens[*pos]
        if op != "*" && op != "/" {
            break
        }
        *pos++
        right := parseUnary(tokens, pos)
        node = &Nodo{Tipo:"expresion", Nombre:op, Args:[]interface{}{node, right}}
    }
    return node
}

func parseUnary(tokens []string, pos *int) interface{} {
    if *pos >= len(tokens) {
        return nil
    }
    tok := tokens[*pos]

    // Unarios: ! (lógico), ~ (bit), ++, --
    if tok == "!" || tok == "~" || tok == "++" || tok == "--" {
        *pos++
        right := parseUnary(tokens, pos)
        return &Nodo{Tipo:"expresion", Nombre:tok, Args:[]interface{}{right}}
    }

    // Primario y comparaciones
    node := parsePrimary(tokens, pos)
    for *pos < len(tokens) {
        op := tokens[*pos]
        if op != "==" && op != "!=" && op != "<" && op != ">" && op != "<=" && op != ">=" {
            break
        }
        *pos++
        right := parsePrimary(tokens, pos)
        node = &Nodo{Tipo:"expresion", Nombre:op, Args:[]interface{}{node, right}}
    }
    return node
}

func parsePrimary(tokens []string, pos *int) interface{} {
    if *pos >= len(tokens) {
        return nil
    }
    tok := tokens[*pos]

    // Paréntesis
    if tok == "(" {
        *pos++
        node := parseLogOr(tokens, pos)
        if *pos < len(tokens) && tokens[*pos] == ")" {
            *pos++
        }
        return node
    }

    // Consumir token y delegar a parseValor para llamadas, índices, conversiones, literales y variables
    *pos++
    return parseValor(tok)
}

// --- Tokenizador y utilidades de operadores ---

func contieneOperadores(s string) bool {
    return strings.ContainsAny(s, "+-*/|&^~!()<>=") ||
        strings.Contains(s, "<<") || strings.Contains(s, ">>") ||
        strings.Contains(s, "**") ||
        strings.Contains(s, "==") || strings.Contains(s, "!=") ||
        strings.Contains(s, "<=") || strings.Contains(s, ">=")
}

func tokenize(expr string) []string {
    // Operadores de dos caracteres primero
    reemplazos2 := []string{"<<", ">>", "**", "==", "!=", "<=", ">="}
    for _, r := range reemplazos2 {
        expr = strings.ReplaceAll(expr, r, " "+r+" ")
    }

    // Unarios y binarios de un carácter
    reemplazos1 := []string{"(", ")", "+", "-", "*", "/", "|", "&", "^", "~", "!", "[", "]", ",", "<", ">"}
    for _, r := range reemplazos1 {
        expr = strings.ReplaceAll(expr, r, " "+r+" ")
    }

    // Incremento/decremento como tokens separados si aparecen pegados (var++ / var--)
    expr = strings.ReplaceAll(expr, "++", " ++ ")
    expr = strings.ReplaceAll(expr, "--", " -- ")

    // Compactar espacios
    campos := strings.Fields(expr)
    return campos
}
