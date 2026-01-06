package parser

import (
    "strconv"
    "strings"
    "unicode"
)

//
// parseMatriz: Parser robusto para matrices y accesos anidados.
//
// Sintaxis soportada (flexible y sin límites):
//   - Constructor literal estilo JSON:
//       matriz [[1,2],[3,4]]
//       matriz [[["a","b"],["c","d"]],[["e","f"],["g","h"]]]
//   - Con tipo de elementos y/o dimensiones:
//       matriz entero 2x2
//       matriz real 3x4 := [[1.1,2.2,3.3],[4.4,5.5,6.6],[7.7,8.8,9.9]]
//       matriz texto NxN
//   - Accesos y multi-índices:
//       matriz A[2]
//       matriz A[2,[x,y]]
//       matriz A[ i , j , k ]
//       matriz A[ i , [ j , [ k , l ] ] ]
//   - Mixto (tipo + acceso + asignación):
//       matriz entero A[ i , j ] := 42
//
// Reglas:
//   - Soporta índices anidados y listas de índices con comas.
//   - Soporta dimensiones con notación NxM, 2x2, 3x3, 4x4, NxN (N puede ser número o identificador).
//   - Soporta literales JSON-like con corchetes anidados sin límite.
//   - Soporta asignación opcional con := <valor> (literal, expresión o llamada).
//   - El tipo de elementos es opcional y se parsea con extraerTipoTokens (bit, entero, real, texto, puntero, etc.).
//
// Devuelve:
//   Nodo{
//     Tipo:   "matriz",
//     Nombre: opcional (identificador de la matriz en accesos),
//     Args:   []interface{}{ tipoTokens?, dims?, indices? },
//     Valor:  literal/expresión/asignación opcional,
//     Extra:  map[string]interface{}{
//       "modo": "literal|acceso|dimensiones",
//       "dims": []string{"N","M"} opcional,
//       "indices": []interface{} opcional,
//     },
//   }
func parseMatriz(linea string) *Nodo {
    trim := strings.TrimSpace(linea)
    if !strings.HasPrefix(trim, "matriz") {
        return nil
    }

    resto := strings.TrimSpace(strings.TrimPrefix(trim, "matriz"))
    if resto == "" {
        // matriz sin contenido: nodo vacío con modo desconocido
        return &Nodo{Tipo: "matriz", Args: []interface{}{}, Extra: map[string]interface{}{"modo": "vacio"}}
    }

    // Separar posible asignación := a nivel toplevel
    izq, der := splitTopLevelAsignacion(resto)

    // Parsear parte izquierda: puede contener tipo, dimensiones, identificador+indices o literal [[...]]
    campos := strings.Fields(izq)

    // Intentar extraer tipo de elementos al inicio
    var tipoTokens []interface{}
    nextIdx := 0
    if len(campos) > 0 {
        tipoTokens, nextIdx = extraerTipoTokens(campos)
    }

    // Reconstruir el resto tras tipo
    restoIzq := strings.TrimSpace(strings.Join(campos[nextIdx:], " "))

    // Caso 1: literal JSON-like si comienza con '['
    if strings.HasPrefix(restoIzq, "[") {
        lit := strings.TrimSpace(restoIzq)
        // Validar balance de corchetes a nivel toplevel
        if !corchetesBalanceados(lit) {
            return nil
        }
        n := &Nodo{
            Tipo:  "matriz",
            Args:  nil,
            Valor: parseValor(lit),
            Extra: map[string]interface{}{"modo": "literal"},
        }
        if len(tipoTokens) > 0 {
            n.Args = append(n.Args, tipoTokens...)
        }
        // Si hay asignación explícita (der), la preferimos como Valor
        if der != "" {
            n.Valor = parseValor(der)
        }
        return n
    }

    // Caso 2: dimensiones explícitas (token con 'x', p.ej. 2x2, 3x4, NxN)
    // Buscamos el primer token con 'x' fuera de comillas
    dimsTok := firstDimsToken(restoIzq)
    if dimsTok != "" {
        dims := splitDims(dimsTok)
        n := &Nodo{
            Tipo: "matriz",
            Args: nil,
            Extra: map[string]interface{}{
                "modo": "dimensiones",
                "dims": dims,
            },
        }
        if len(tipoTokens) > 0 {
            n.Args = append(n.Args, tipoTokens...)
        }
        // Si hay asignación, parsear como literal/expresión
        if der != "" {
            n.Valor = parseValor(der)
        }
        return n
    }

    // Caso 3: acceso con identificador y corchetes: A[...]
    // Extraemos identificador y bloque de índices
    id, indicesStr := splitIdentIndices(restoIzq)
    if id != "" && indicesStr != "" {
        indices := parseIndicesTopLevel(indicesStr)
        n := &Nodo{
            Tipo:   "matriz",
            Nombre: id,
            Args:   nil,
            Extra: map[string]interface{}{
                "modo":    "acceso",
                "indices": indices,
            },
        }
        if len(tipoTokens) > 0 {
            n.Args = append(n.Args, tipoTokens...)
        }
        if der != "" {
            n.Valor = parseValor(der)
        }
        return n
    }

    // Si no encaja en ninguno, intentamos tratar el resto como expresión/literal genérico de matriz
    if restoIzq != "" {
        n := &Nodo{
            Tipo:  "matriz",
            Args:  tipoTokens,
            Valor: parseValor(restoIzq),
            Extra: map[string]interface{}{"modo": "expresion"},
        }
        if der != "" {
            n.Valor = parseValor(der)
        }
        return n
    }

    // Fallback: nodo vacío con tipo si lo hubo
    return &Nodo{Tipo: "matriz", Args: tipoTokens, Extra: map[string]interface{}{"modo": "desconocido"}}
}

// --- Utilidades internas ---

// splitTopLevelAsignacion: separa izquierda y derecha de := a nivel toplevel (respeta comillas, paréntesis y corchetes)
func splitTopLevelAsignacion(s string) (string, string) {
    var left strings.Builder
    var right strings.Builder
    quote := rune(0)
    par := 0
    br := 0
    found := false

    for i, r := range s {
        switch {
        case quote != 0:
            if !found {
                left.WriteRune(r)
            } else {
                right.WriteRune(r)
            }
            if r == quote && (i == 0 || s[i-1] != '\\') {
                quote = 0
            }
        default:
            switch r {
            case '"', '\'':
                quote = r
                if !found {
                    left.WriteRune(r)
                } else {
                    right.WriteRune(r)
                }
            case '(':
                par++
                if !found {
                    left.WriteRune(r)
                } else {
                    right.WriteRune(r)
                }
            case ')':
                if par > 0 {
                    par--
                }
                if !found {
                    left.WriteRune(r)
                } else {
                    right.WriteRune(r)
                }
            case '[':
                br++
                if !found {
                    left.WriteRune(r)
                } else {
                    right.WriteRune(r)
                }
            case ']':
                if br > 0 {
                    br--
                }
                if !found {
                    left.WriteRune(r)
                } else {
                    right.WriteRune(r)
                }
            case ':':
                // posible := (mirar siguiente)
                if !found && i+1 < len(s) && s[i+1] == '=' && par == 0 && br == 0 {
                    found = true
                    // saltar '=' en el siguiente ciclo
                    continue
                }
                if !found {
                    left.WriteRune(r)
                } else {
                    right.WriteRune(r)
                }
            case '=':
                if !found {
                    left.WriteRune(r)
                } else {
                    right.WriteRune(r)
                }
            default:
                if !found {
                    left.WriteRune(r)
                } else {
                    right.WriteRune(r)
                }
            }
        }
    }
    return strings.TrimSpace(left.String()), strings.TrimSpace(right.String())
}

// corchetesBalanceados: valida balance de corchetes a nivel toplevel (permite anidación)
func corchetesBalanceados(s string) bool {
    quote := rune(0)
    br := 0
    for i, r := range s {
        switch {
        case quote != 0:
            if r == quote && (i == 0 || s[i-1] != '\\') {
                quote = 0
            }
        default:
            switch r {
            case '"', '\'':
                quote = r
            case '[':
                br++
            case ']':
                br--
                if br < 0 {
                    return false
                }
            }
        }
    }
    return br == 0 && quote == 0
}

// firstDimsToken: devuelve el primer token con 'x' (p.ej. 2x2, 3x4, NxN) fuera de comillas
func firstDimsToken(s string) string {
    fields := strings.Fields(s)
    for _, f := range fields {
        if strings.Contains(f, "x") && !strings.ContainsAny(f, "\"'") {
            // validar que sea algo como NxM (N y M pueden ser números o identificadores)
            parts := strings.Split(f, "x")
            if len(parts) == 2 && parts[0] != "" && parts[1] != "" {
                return f
            }
        }
    }
    return ""
}

// splitDims: separa "NxM" en []string{"N","M"} (N y M pueden ser números o identificadores)
func splitDims(tok string) []string {
    parts := strings.Split(tok, "x")
    return []string{parts[0], parts[1]}
}

// splitIdentIndices: separa identificador y bloque de índices A[...]
func splitIdentIndices(s string) (string, string) {
    s = strings.TrimSpace(s)
    if s == "" {
        return "", ""
    }
    // identificador: letras, números y guiones bajos al inicio
    var id strings.Builder
    i := 0
    for i < len(s) {
        r := rune(s[i])
        if i == 0 {
            if !(unicode.IsLetter(r) || r == '_') {
                break
            }
            id.WriteRune(r)
            i++
            continue
        }
        if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
            id.WriteRune(r)
            i++
            continue
        }
        break
    }
    if id.Len() == 0 {
        return "", ""
    }
    rest := strings.TrimSpace(s[i:])
    if !strings.HasPrefix(rest, "[") {
        return "", ""
    }
    // extraer bloque de índices hasta el corchete de cierre toplevel
    idx := extractTopLevelBracket(rest)
    return id.String(), idx
}

// extractTopLevelBracket: extrae el contenido del primer bloque [ ... ] a nivel toplevel
func extractTopLevelBracket(s string) string {
    quote := rune(0)
    br := 0
    start := -1
    end := -1
    for i, r := range s {
        switch {
        case quote != 0:
            if r == quote && (i == 0 || s[i-1] != '\\') {
                quote = 0
            }
        default:
            switch r {
            case '"', '\'':
                quote = r
            case '[':
                if br == 0 {
                    start = i
                }
                br++
            case ']':
                br--
                if br == 0 {
                    end = i
                    goto done
                }
            }
        }
    }
done:
    if start >= 0 && end > start {
        return s[start : end+1]
    }
    return ""
}

// parseIndicesTopLevel: parsea índices dentro de [ ... ], soporta listas y anidación
// Ejemplos:
//   [2]                -> [2]
//   [ i , j , k ]      -> [i, j, k]
//   [ 2 , [x,y] ]      -> [2, [x,y]]
func parseIndicesTopLevel(idx string) []interface{} {
    inner := strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(strings.TrimSpace(idx), "["), "]"))
    if inner == "" {
        return []interface{}{}
    }
    // separar por comas a nivel toplevel respetando corchetes y comillas
    var res []interface{}
    var buf strings.Builder
    quote := rune(0)
    br := 0
    par := 0

    flush := func() {
        if buf.Len() == 0 {
            return
        }
        token := strings.TrimSpace(buf.String())
        buf.Reset()
        if token == "" {
            return
        }
        // si el token comienza con '[' y termina con ']', es un sub-bloque
        if strings.HasPrefix(token, "[") && strings.HasSuffix(token, "]") {
            res = append(res, parseIndicesTopLevel(token))
            return
        }
        // si es número, guardarlo como número; si no, parseValor
        if n, err := strconv.ParseFloat(token, 64); err == nil {
            res = append(res, n)
            return
        }
        res = append(res, parseValor(token))
    }

    for i, r := range inner {
        switch {
        case quote != 0:
            buf.WriteRune(r)
            if r == quote && (i == 0 || inner[i-1] != '\\') {
                quote = 0
            }
        default:
            switch r {
            case '"', '\'':
                quote = r
                buf.WriteRune(r)
            case '[':
                br++
                buf.WriteRune(r)
            case ']':
                if br > 0 {
                    br--
                }
                buf.WriteRune(r)
            case '(':
                par++
                buf.WriteRune(r)
            case ')':
                if par > 0 {
                    par--
                }
                buf.WriteRune(r)
            case ',':
                if br == 0 && par == 0 {
                    flush()
                } else {
                    buf.WriteRune(r)
                }
            default:
                buf.WriteRune(r)
            }
        }
    }
    flush()
    return res
}

// --- Registro en Parsers ---
func init() {
    Parsers["matriz"] = func(linea string) *Nodo {
        trim := strings.TrimSpace(linea)
        if !strings.HasPrefix(trim, "matriz") {
            return nil
        }
        // Evitar colisiones con prefijos (p.ej., "matrizX")
        if len(trim) > len("matriz") && unicode.IsLetter(rune(trim[len("matriz")])) {
            return nil
        }
        return parseMatriz(trim)
    }
}
