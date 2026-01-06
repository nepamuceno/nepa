package parser

import (
    "strings"
    "unicode"
)

//
// parseLista: Parser robusto para listas estilo JSON + golang (slices),
// con literales, anidación, accesos, asignaciones y mezcla de tipos.
//
// Sintaxis soportada:
//   - Literal JSON-like (anidado sin límite):
//       lista [1,2,3]
//       lista [[1,2],[3,4]]
//       lista ["a", {"k":"v"}, [true, 3.14]]
//   - Accesos y multi-índices (golang-style):
//       lista L[0]
//       lista L[i]
//       lista L[i,[j,k]]        // anidado de índices (para listas de listas)
//   - Asignación opcional:
//       lista L[0] := "hola"
//       lista L[i] := sin(x)*sqrt(y)
//   - Operaciones estilo golang (se parsean como expresiones):
//       lista append(L, valor)
//       lista copy(dst, src)
//       lista len(L)
//       lista cap(L)
//     (Estas se tratan como expresiones/llamadas vía parseValor)
//   - Mezcla de tipos (heterogéneo, estilo JSON):
//       lista [1, "dos", verdadero, {"a":1}, [3,4]]
//
// Reglas:
//   - Soporta corchetes [] anidados sin límite (JSON-like).
//   - Soporta accesos con índices simples y listas de índices anidadas.
//   - Soporta asignación con := a nivel toplevel.
//   - Integra tipos nativos y compuestos vía extraerTipoTokens si se usan como prefijo.
//
// Devuelve:
//   Nodo{
//     Tipo:   "lista",
//     Nombre: opcional (identificador de la lista en accesos),
//     Args:   []interface{}{ tipoTokens? },
//     Valor:  literal/expresión/asignación opcional,
//     Extra:  map[string]interface{}{
//       "modo": "literal|acceso|expresion|vacio|desconocido",
//       "indices": []interface{} opcional,
//     },
//   }
func parseLista(linea string) *Nodo {
    trim := strings.TrimSpace(linea)
    if !strings.HasPrefix(trim, "lista") {
        return nil
    }

    resto := strings.TrimSpace(strings.TrimPrefix(trim, "lista"))
    if resto == "" {
        return &Nodo{Tipo: "lista", Args: []interface{}{}, Extra: map[string]interface{}{"modo": "vacio"}}
    }

    // Separar posible asignación := a nivel toplevel
    izq, der := splitTopLevelAsignacion(resto)

    // Intentar extraer tipo compuesto al inicio (opcional)
    campos := strings.Fields(izq)
    var tipoTokens []interface{}
    nextIdx := 0
    if len(campos) > 0 {
        tipoTokens, nextIdx = extraerTipoTokens(campos)
    }
    restoIzq := strings.TrimSpace(strings.Join(campos[nextIdx:], " "))

    // Caso 1: literal JSON-like si comienza con '['
    if strings.HasPrefix(restoIzq, "[") {
        lit := strings.TrimSpace(restoIzq)
        if !corchetesBalanceados(lit) {
            return nil
        }
        n := &Nodo{
            Tipo:  "lista",
            Args:  nil,
            Valor: parseValor(lit),
            Extra: map[string]interface{}{"modo": "literal"},
        }
        if len(tipoTokens) > 0 {
            n.Args = append(n.Args, tipoTokens...)
        }
        if der != "" {
            n.Valor = parseValor(der)
        }
        return n
    }

    // Caso 2: acceso con identificador y corchetes: L[...]
    id, indicesStr := splitIdentIndices(restoIzq)
    if id != "" && indicesStr != "" {
        indices := parseIndicesTopLevel(indicesStr)
        n := &Nodo{
            Tipo:   "lista",
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

    // Caso 3: operaciones estilo golang o expresiones genéricas (append, copy, len, cap, llamadas)
    if restoIzq != "" {
        n := &Nodo{
            Tipo:  "lista",
            Args:  tipoTokens,
            Valor: parseValor(restoIzq),
            Extra: map[string]interface{}{"modo": "expresion"},
        }
        if der != "" {
            n.Valor = parseValor(der)
        }
        return n
    }

    // Fallback
    return &Nodo{Tipo: "lista", Args: tipoTokens, Extra: map[string]interface{}{"modo": "desconocido"}}
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

// corchetesBalanceados: valida balance de [] a nivel toplevel (permite anidación)
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

// splitIdentIndices: separa identificador y bloque de índices L[...]
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
        // si es número, dejar que parseValor lo trate (para mantener tipos homogéneos con el resto)
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
    Parsers["lista"] = func(linea string) *Nodo {
        trim := strings.TrimSpace(linea)
        if !strings.HasPrefix(trim, "lista") {
            return nil
        }
        // Evitar colisiones con prefijos (p.ej., "listado")
        if len(trim) > len("lista") && unicode.IsLetter(rune(trim[len("lista")])) {
            return nil
        }
        return parseLista(trim)
    }
}
