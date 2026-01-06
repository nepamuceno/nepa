package parser

import (
    "strings"
    "unicode"
)

// parseImprimir: Parser universal para la instrucción `imprimir`.
//
// Sintaxis soportada:
//   - imprimir(<expr>)
//   - imprimir <expr>[, <expr2>, ...]
//
// Donde <expr> puede ser:
//   - Literales con " o ' (ej. "hola 'sr.' juan", 'Hola "mundo"')
//   - Variables de cualquier tipo (bit, entero, real, texto, puntero, matriz, objeto)
//   - Expresiones complejas y anidadas (sin(x)*((a)/(b)*sqrt(c - otraFuncion(y,z))))
//   - Concatenaciones con . o + (ej. "Hola: ".nombre." saludos", "Hola " + nombre)
//   - Caracteres de control (\n, \t, \b, etc.)
//   - Funciones internas/externas, llamadas anidadas
//   - Matrices y punteros (ej. matriz[2][3], puntero->campo)
//   - Modo pipe: imprimir <expr> | otraFuncion(...)
//
// Devuelve un Nodo{Tipo:"imprimir", Args:[...], Extra:{ opcional: "pipe": Nodo{...}, "modo": "parentesis|inline" }}
func parseImprimir(linea string) *Nodo {
    if !strings.HasPrefix(strings.TrimSpace(linea), "imprimir") {
        return nil
    }

    resto := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(linea), "imprimir"))
    if resto == "" {
        // si no hay nada, devolvemos un nodo vacío
        return &Nodo{Tipo: "imprimir", Args: []interface{}{}}
    }

    // Detectar modo pipe a nivel toplevel (no dentro de comillas ni paréntesis)
    parte, pipe := splitTopLevelPipe(resto)

    // Determinar modo: paréntesis o inline
    modo := "inline"
    argsTexto := parte
    if strings.HasPrefix(parte, "(") && strings.HasSuffix(parte, ")") {
        modo = "parentesis"
        argsTexto = strings.TrimSpace(parte[1 : len(parte)-1])
    } else {
        argsTexto = strings.TrimSpace(argsTexto)
    }

    // Separar argumentos por comas a nivel toplevel
    argStrs := splitTopLevelCommas(argsTexto)

    // Parsear cada argumento, soportando concatenación con . y +
    var args []interface{}
    for _, raw := range argStrs {
        raw = strings.TrimSpace(raw)
        if raw == "" {
            continue
        }
        // Si el argumento contiene concatenación, tokenizar y convertir a lista de fragmentos
        frags := splitConcatenacionTopLevel(raw)
        if len(frags) == 1 {
            args = append(args, parseValor(frags[0]))
        } else {
            for _, f := range frags {
                f = strings.TrimSpace(f)
                if f == "" {
                    continue
                }
                args = append(args, parseValor(f))
            }
        }
    }

    // Construir nodo base
    n := &Nodo{
        Tipo:  "imprimir",
        Args:  args,
        Extra: map[string]interface{}{"modo": modo},
    }

    // Si hay pipe, parsearlo como llamada y guardarlo en Extra["pipe"]
    if pipe != "" {
        pipe = strings.TrimSpace(pipe)
        if nodoCall := parseLlamada(pipe); nodoCall != nil {
            n.Extra["pipe"] = *nodoCall
        } else {
            n.Extra["pipe"] = Nodo{Tipo: "expresion", Valor: pipe}
        }
    }

    return n
}

// --- Utilidades internas ---

// splitTopLevelCommas: separa por comas a nivel toplevel (respeta comillas y paréntesis)
func splitTopLevelCommas(s string) []string {
    var res []string
    var buf strings.Builder
    quote := rune(0)
    par := 0

    for i, r := range s {
        switch {
        case quote != 0:
            buf.WriteRune(r)
            if r == quote && (i == 0 || s[i-1] != '\\') {
                quote = 0
            }
        default:
            switch r {
            case '"', '\'':
                quote = r
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
                if par == 0 {
                    res = append(res, strings.TrimSpace(buf.String()))
                    buf.Reset()
                } else {
                    buf.WriteRune(r)
                }
            default:
                buf.WriteRune(r)
            }
        }
    }
    if buf.Len() > 0 {
        res = append(res, strings.TrimSpace(buf.String()))
    }
    return res
}

// splitConcatenacionTopLevel: separa por operadores . y + a nivel toplevel (respeta comillas y paréntesis)
func splitConcatenacionTopLevel(s string) []string {
    var res []string
    var buf strings.Builder
    quote := rune(0)
    par := 0

    flush := func() {
        if buf.Len() > 0 {
            res = append(res, strings.TrimSpace(buf.String()))
            buf.Reset()
        }
    }

    for i, r := range s {
        switch {
        case quote != 0:
            buf.WriteRune(r)
            if r == quote && (i == 0 || s[i-1] != '\\') {
                quote = 0
            }
        default:
            switch r {
            case '"', '\'':
                quote = r
                buf.WriteRune(r)
            case '(':
                par++
                buf.WriteRune(r)
            case ')':
                if par > 0 {
                    par--
                }
                buf.WriteRune(r)
            case '.', '+':
                if par == 0 {
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

// splitTopLevelPipe: separa la parte principal y el pipe a nivel toplevel (respeta comillas y paréntesis)
func splitTopLevelPipe(s string) (string, string) {
    var left strings.Builder
    var right strings.Builder
    quote := rune(0)
    par := 0
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
            case '|':
                if par == 0 && !found {
                    found = true
                } else {
                    if !found {
                        left.WriteRune(r)
                    } else {
                        right.WriteRune(r)
                    }
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

// --- Registro en Parsers ---
func init() {
    Parsers["imprimir"] = func(linea string) *Nodo {
        trim := strings.TrimSpace(linea)
        if !strings.HasPrefix(trim, "imprimir") {
            return nil
        }
        // Evitar confundir con otras palabras que comiencen igual (ej. "imprimirX")
        if len(trim) > len("imprimir") && unicode.IsLetter(rune(trim[len("imprimir")])) {
            return nil
        }
        return parseImprimir(trim)
    }
}
