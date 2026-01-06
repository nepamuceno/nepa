package parser

import (
    "strings"
    "unicode"
)

//
// parseDiccionario: Parser robusto para diccionarios/mapas estilo JSON + golang + python,
// con literales, anidación, accesos por clave, asignaciones y operaciones estándar.
//
// Sintaxis soportada:
//   - Literal JSON-like (anidado sin límite):
//       diccionario { "a":1, "b":"dos", "c":[1,2], "d":{"x":9} }
//       diccionario { "persona": { "nombre":"Ana", "edad":30 }, "activo":true }
//   - Definición con tipos explícitos (opcional, flexible):
//       diccionario texto->entero
//       diccionario clave texto, valor real
//     (Se integra vía extraerTipoTokens; se guarda como tokens en Args)
//   - Acceso a claves (anidado con [] y encadenado):
//       diccionario D["clave1"]
//       diccionario D["persona"]["nombre"]
//       diccionario D["lista"][0]                 // mezcla de clave y índice
//       diccionario D["a"][i]["b"][j,[k,l]]       // anidación profunda
//   - Asignación opcional:
//       diccionario D["clave1"] := 42
//       diccionario D["persona"]["edad"] := 31
//       diccionario D["lista"][0] := {"x":1}
//   - Operaciones estilo Go/Python (se parsean como expresiones):
//       diccionario len(D)
//       diccionario delete(D,"clave1")
//       diccionario keys(D)
//       diccionario values(D)
//       diccionario setdefault(D,"k",valor)
//   - Mezcla de tipos (heterogéneo, estilo JSON):
//       diccionario { "a":1, "b":"dos", "c":true, "d":[1,2], "e":{"x":9} }
//
// Reglas:
//   - Soporta llaves {} y corchetes [] anidados sin límite (JSON-like).
//   - Soporta accesos con claves string, variables y expresiones dentro de [ ... ].
//   - Soporta asignación con := a nivel toplevel.
//   - Integra tipos nativos y compuestos vía extraerTipoTokens si se usan como prefijo.
//   - Permite mezcla de tipos en valores (heterogéneo), coherente con JSON.
//
// Devuelve:
//   Nodo{
//     Tipo:   "diccionario",
//     Nombre: opcional (identificador del diccionario),
//     Args:   []interface{}{ tipoTokens? },
//     Valor:  literal/expresión/asignación opcional,
//     Extra:  map[string]interface{}{
//       "modo": "literal|definicion|acceso|expresion|vacio|desconocido",
//       "acceso": []interface{} opcional (ruta de acceso con segmentos: claves/índices),
//     },
//   }
func parseDiccionario(linea string) *Nodo {
    trim := strings.TrimSpace(linea)
    if !strings.HasPrefix(trim, "diccionario") {
        return nil
    }

    resto := strings.TrimSpace(strings.TrimPrefix(trim, "diccionario"))
    if resto == "" {
        return &Nodo{Tipo: "diccionario", Args: []interface{}{}, Extra: map[string]interface{}{"modo": "vacio"}}
    }

    // Separar posible asignación := a nivel toplevel
    izq, der := splitTopLevelAsignacionDic(resto)

    // Intentar extraer tipo compuesto al inicio (opcional)
    campos := strings.Fields(izq)
    var tipoTokens []interface{}
    nextIdx := 0
    if len(campos) > 0 {
        tipoTokens, nextIdx = extraerTipoTokens(campos)
    }
    restoIzq := strings.TrimSpace(strings.Join(campos[nextIdx:], " "))

    // Caso 1: literal JSON-like si comienza con '{'
    if strings.HasPrefix(restoIzq, "{") {
        lit := strings.TrimSpace(restoIzq)
        if !llavesBalanceadas(lit) {
            return nil
        }
        n := &Nodo{
            Tipo:  "diccionario",
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

    // Caso 2: acceso con identificador y cadena de claves/índices: D["k"][i]...
    id, acceso := splitDictAccessChain(restoIzq)
    if id != "" && len(acceso) > 0 {
        n := &Nodo{
            Tipo:   "diccionario",
            Nombre: id,
            Args:   nil,
            Extra:  map[string]interface{}{"modo": "acceso", "acceso": acceso},
        }
        if len(tipoTokens) > 0 {
            n.Args = append(n.Args, tipoTokens...)
        }
        if der != "" {
            n.Valor = parseValor(der)
        }
        return n
    }

    // Caso 3: definición con tipos explícitos (si quedaron tokens y no hay literal/acceso)
    if len(tipoTokens) > 0 && restoIzq == "" {
        return &Nodo{
            Tipo:  "diccionario",
            Args:  tipoTokens,
            Extra: map[string]interface{}{"modo": "definicion"},
        }
    }

    // Caso 4: operaciones estilo Go/Python o expresiones genéricas
    if restoIzq != "" {
        n := &Nodo{
            Tipo:  "diccionario",
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
    return &Nodo{Tipo: "diccionario", Args: tipoTokens, Extra: map[string]interface{}{"modo": "desconocido"}}
}

// --- Utilidades internas ---

// splitTopLevelAsignacionDic: separa izquierda y derecha de := a nivel toplevel (respeta comillas, paréntesis, corchetes y llaves)
func splitTopLevelAsignacionDic(s string) (string, string) {
    var left strings.Builder
    var right strings.Builder
    quote := rune(0)
    par := 0
    br := 0
    curly := 0
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
            case '{':
                curly++
                if !found {
                    left.WriteRune(r)
                } else {
                    right.WriteRune(r)
                }
            case '}':
                if curly > 0 {
                    curly--
                }
                if !found {
                    left.WriteRune(r)
                } else {
                    right.WriteRune(r)
                }
            case ':':
                // posible := (mirar siguiente)
                if !found && i+1 < len(s) && s[i+1] == '=' && par == 0 && br == 0 && curly == 0 {
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

// splitDictAccessChain: separa identificador base y ruta de acceso anidada con [ ... ] encadenado.
// Acepta claves string ("clave"), variables, expresiones y también índices numéricos.
// Ejemplos:
//   D["a"]["b"]           -> id="D", acceso=["a","b"]
//   D["lista"][0]         -> id="D", acceso=["lista", 0]
//   D[k][i,[j,l]]         -> id="D", acceso=[k, [i,[j,l]]]
func splitDictAccessChain(s string) (string, []interface{}) {
    s = strings.TrimSpace(s)
    if s == "" {
        return "", nil
    }
    // identificador base
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
        return "", nil
    }
    rest := strings.TrimSpace(s[i:])
    if rest == "" || rest[0] != '[' {
        return "", nil
    }

    // recorrer bloques [ ... ] consecutivos
    var acceso []interface{}
    for i < len(s) {
        if rest == "" || rest[0] != '[' {
            break
        }
        block := extractTopLevelBracket(rest)
        if block == "" {
            break
        }
        // parsear contenido del bloque como clave/índice:
        // - si es string literal "clave" -> clave
        // - si es lista [x,y] -> lista de índices/claves
        // - si es expresión/variable -> parseValor(token)
        parsed := parseDictKeyOrIndices(block)
        acceso = append(acceso, parsed)
        // avanzar
        rest = strings.TrimSpace(rest[len(block):])
        if len(rest) > 0 && rest[0] == '[' {
            continue
        }
        break
    }

    return id.String(), acceso
}

// parseDictKeyOrIndices: interpreta el contenido de [ ... ] para diccionarios.
// - Si el contenido es una sola clave string: ["clave"] -> "clave"
// - Si es una lista separada por comas: [a,b] -> []interface{}{a,b} (cada elemento parseado)
// - Si es un único token no string: [expr] -> parseValor(expr)
func parseDictKeyOrIndices(block string) interface{} {
    inner := strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(strings.TrimSpace(block), "["), "]"))
    if inner == "" {
        return []interface{}{}
    }
    // caso clave string única: "clave"
    if (strings.HasPrefix(inner, "\"") && strings.HasSuffix(inner, "\"")) ||
        (strings.HasPrefix(inner, "'") && strings.HasSuffix(inner, "'")) {
        // devolver literal tal cual (parseValor mantiene el literal)
        return parseValor(inner)
    }

    // si contiene coma a toplevel, es lista de claves/índices
    parts := splitTopLevelCommas(inner)
    if len(parts) > 1 {
        var res []interface{}
        for _, p := range parts {
            p = strings.TrimSpace(p)
            if p == "" {
                continue
            }
            // si es sub-bloque [ ... ], parsear recursivo como índices anidados
            if strings.HasPrefix(p, "[") && strings.HasSuffix(p, "]") {
                res = append(res, parseIndicesTopLevel(p))
                continue
            }
            res = append(res, parseValor(p))
        }
        return res
    }

    // único token: puede ser número, variable o expresión
    return parseValor(inner)
}

// --- Registro en Parsers ---
func init() {
    Parsers["diccionario"] = func(linea string) *Nodo {
        trim := strings.TrimSpace(linea)
        if !strings.HasPrefix(trim, "diccionario") {
            return nil
        }
        // Evitar colisiones con prefijos (p.ej., "diccionarios")
        if len(trim) > len("diccionario") && unicode.IsLetter(rune(trim[len("diccionario")])) {
            return nil
        }
        return parseDiccionario(trim)
    }
}
