package parser

import (
    "strings"
    "unicode"
)

//
// parseEstructura: Parser robusto para estructuras/objetos estilo JSON,
// definiciones con campos y tipos, accesos anidados y asignaciones.
//
// Sintaxis soportada:
//   - Literal JSON-like:
//       estructura { "nombre":"Juan", "edad":30 }
//       estructura { "persona": { "nombre":"Ana", "hijos":[{"nombre":"Luis"}] } }
//   - Definición con campos y tipos:
//       estructura Persona { texto nombre; entero edad; }
//       estructura Registro { bit activo; real saldo; puntero ref; }
//     (Opcional: prefijo de tipo compuesto para la estructura, vía extraerTipoTokens)
//   - Acceso a campos (anidados con . y []):
//       estructura Persona.nombre
//       estructura Persona.direccion.calle
//       estructura Persona.hijos[0].nombre
//   - Asignación opcional:
//       estructura Persona.nombre := "Juan"
//       estructura Persona.edad := 30
//       estructura Persona.hijos[0].nombre := "Luis"
//   - Valores complejos:
//       estructura Persona.saldo := sin(x)*sqrt(y)
//       estructura Registro.ref := obtenerReferencia()
//
// Reglas:
//   - Soporta llaves {} y corchetes [] anidados sin límite (JSON-like).
//   - Soporta definición de campos con tipos explícitos dentro de {}.
//   - Soporta accesos con '.' y con índices [ ... ] anidados.
//   - Soporta asignación con := a nivel toplevel.
//   - Integra tipos nativos y compuestos vía extraerTipoTokens.
//
// Devuelve:
//   Nodo{
//     Tipo:   "estructura",
//     Nombre: opcional (identificador de la estructura),
//     Args:   []interface{}{ tipoTokens?, campos? },
//     Valor:  literal/expresión/asignación opcional,
//     Extra:  map[string]interface{}{
//       "modo": "literal|definicion|acceso|expresion|vacio|desconocido",
//       "campos": []interface{} opcional (definición),
//       "acceso": []interface{} opcional (ruta de acceso con segmentos y/o índices),
//     },
//   }
func parseEstructura(linea string) *Nodo {
    trim := strings.TrimSpace(linea)
    if !strings.HasPrefix(trim, "estructura") {
        return nil
    }

    resto := strings.TrimSpace(strings.TrimPrefix(trim, "estructura"))
    if resto == "" {
        return &Nodo{Tipo: "estructura", Args: []interface{}{}, Extra: map[string]interface{}{"modo": "vacio"}}
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

    // Caso 1: literal JSON-like si comienza con '{'
    if strings.HasPrefix(restoIzq, "{") {
        lit := strings.TrimSpace(restoIzq)
        if !llavesBalanceadas(lit) {
            return nil
        }
        n := &Nodo{
            Tipo:  "estructura",
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

    // Caso 2: definición con identificador y bloque { ... }
    id, defStr := splitIdentDefBlock(restoIzq)
    if id != "" && defStr != "" {
        camposDef := parseCamposDefinicion(defStr)
        n := &Nodo{
            Tipo:   "estructura",
            Nombre: id,
            Args:   nil,
            Extra:  map[string]interface{}{"modo": "definicion", "campos": camposDef},
        }
        if len(tipoTokens) > 0 {
            n.Args = append(n.Args, tipoTokens...)
        }
        if der != "" {
            n.Valor = parseValor(der)
        }
        return n
    }

    // Caso 3: acceso a campos anidados (ruta con '.' y posibles índices [ ... ])
    idAcceso, ruta := splitAccessChain(restoIzq)
    if idAcceso != "" && len(ruta) > 0 {
        n := &Nodo{
            Tipo:   "estructura",
            Nombre: idAcceso,
            Args:   nil,
            Extra:  map[string]interface{}{"modo": "acceso", "acceso": ruta},
        }
        if len(tipoTokens) > 0 {
            n.Args = append(n.Args, tipoTokens...)
        }
        if der != "" {
            n.Valor = parseValor(der)
        }
        return n
    }

    // Caso 4: expresión genérica (si queda algo)
    if restoIzq != "" {
        n := &Nodo{
            Tipo:  "estructura",
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
    return &Nodo{Tipo: "estructura", Args: tipoTokens, Extra: map[string]interface{}{"modo": "desconocido"}}
}

// --- Utilidades internas ---

// splitTopLevelAsignacion: separa izquierda y derecha de := a nivel toplevel (respeta comillas, paréntesis, corchetes y llaves)
func splitTopLevelAsignacion(s string) (string, string) {
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

// llavesBalanceadas: valida balance de {} a nivel toplevel (permite anidación)
func llavesBalanceadas(s string) bool {
    quote := rune(0)
    curly := 0
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
            case '{':
                curly++
            case '}':
                curly--
                if curly < 0 {
                    return false
                }
            }
        }
    }
    return curly == 0 && quote == 0
}

// splitIdentDefBlock: separa identificador y bloque de definición { ... }
func splitIdentDefBlock(s string) (string, string) {
    s = strings.TrimSpace(s)
    if s == "" {
        return "", ""
    }
    // identificador al inicio
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
    if !strings.HasPrefix(rest, "{") {
        return "", ""
    }
    block := extractTopLevelBraces(rest)
    return id.String(), block
}

// extractTopLevelBraces: extrae el contenido del primer bloque { ... } a nivel toplevel
func extractTopLevelBraces(s string) string {
    quote := rune(0)
    curly := 0
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
            case '{':
                if curly == 0 {
                    start = i
                }
                curly++
            case '}':
                curly--
                if curly == 0 {
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

// parseCamposDefinicion: parsea campos dentro de { ... } en forma "tipo nombre;" por línea
// Devuelve una lista de descriptores (strings o estructuras simples) manteniendo el orden.
func parseCamposDefinicion(def string) []interface{} {
    inner := strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(strings.TrimSpace(def), "{"), "}"))
    if inner == "" {
        return []interface{}{}
    }
    // dividir por ';' a nivel toplevel (respetando comillas y llaves/corchetes/paréntesis)
    var res []interface{}
    var buf strings.Builder
    quote := rune(0)
    par := 0
    br := 0
    curly := 0

    flush := func() {
        if buf.Len() == 0 {
            return
        }
        line := strings.TrimSpace(buf.String())
        buf.Reset()
        if line == "" {
            return
        }
        // cada línea puede ser: "<tipo> <nombre>" (opcionalmente con más tokens)
        fields := strings.Fields(line)
        if len(fields) == 0 {
            return
        }
        // extraer tipo tokens al inicio
        tokens, next := extraerTipoTokens(fields)
        nombre := strings.TrimSpace(strings.Join(fields[next:], " "))
        if nombre != "" {
            res = append(res, map[string]interface{}{
                "tipo":   tokens,
                "nombre": nombre,
            })
        } else {
            // si no hay nombre, guardar la línea cruda
            res = append(res, line)
        }
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
            case '(':
                par++
                buf.WriteRune(r)
            case ')':
                if par > 0 {
                    par--
                }
                buf.WriteRune(r)
            case '[':
                br++
                buf.WriteRune(r)
            case ']':
                if br > 0 {
                    br--
                }
                buf.WriteRune(r)
            case '{':
                curly++
                buf.WriteRune(r)
            case '}':
                if curly > 0 {
                    curly--
                }
                buf.WriteRune(r)
            case ';':
                if par == 0 && br == 0 && curly == 0 {
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

// splitAccessChain: separa identificador base y ruta de acceso anidada con '.' y [ ... ]
// Ejemplos:
//   "Persona.nombre"                -> id="Persona", ruta=["nombre"]
//   "Persona.hijos[0].nombre"       -> id="Persona", ruta=["hijos", [0], "nombre"]
//   "Registro.ref->campo"           -> id="Registro", ruta=["ref->campo"] (se deja crudo si hay '->')
func splitAccessChain(s string) (string, []interface{}) {
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
    if rest == "" {
        return id.String(), nil
    }

    // tokenizar ruta respetando comillas, paréntesis, corchetes y llaves
    var ruta []interface{}
    var buf strings.Builder
    quote := rune(0)
    par := 0
    br := 0
    curly := 0

    flush := func() {
        if buf.Len() == 0 {
            return
        }
        token := strings.TrimSpace(buf.String())
        buf.Reset()
        if token == "" {
            return
        }
        // si el token comienza con '[' y termina con ']', parsear índices
        if strings.HasPrefix(token, "[") && strings.HasSuffix(token, "]") {
            ruta = append(ruta, parseIndicesTopLevel(token))
            return
        }
        // si contiene '->' (puntero), dejar crudo como segmento
        ruta = append(ruta, token)
    }

    for i := 0; i < len(rest); i++ {
        r := rune(rest[i])
        switch {
        case quote != 0:
            buf.WriteRune(r)
            if r == quote && (i == 0 || rest[i-1] != '\\') {
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
            case '[':
                // si el buffer tiene un nombre previo, flush antes de capturar el bloque
                if buf.Len() > 0 {
                    flush()
                }
                // extraer bloque completo [ ... ]
                block := extractTopLevelBracket(rest[i:])
                if block != "" {
                    ruta = append(ruta, parseIndicesTopLevel(block))
                    i += len(block) - 1
                }
            case ']':
                // no debería ocurrir sin '[' previo; lo tratamos como texto
                buf.WriteRune(r)
            case '{':
                curly++
                buf.WriteRune(r)
            case '}':
                if curly > 0 {
                    curly--
                }
                buf.WriteRune(r)
            case '.':
                if par == 0 && br == 0 && curly == 0 {
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
    return id.String(), ruta
}

// --- Registro en Parsers ---
func init() {
    Parsers["estructura"] = func(linea string) *Nodo {
        trim := strings.TrimSpace(linea)
        if !strings.HasPrefix(trim, "estructura") {
            return nil
        }
        // Evitar colisiones con prefijos (p.ej., "estructurax")
        if len(trim) > len("estructura") && unicode.IsLetter(rune(trim[len("estructura")])) {
            return nil
        }
        return parseEstructura(trim)
    }
}
