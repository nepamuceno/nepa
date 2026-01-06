package parser

import (
    "strings"
    "unicode"
)

//
// parser_bloques.go
//
// Este parser maneja los BLOQUES nativos de nepa.
// Reglas principales:
//   - Los bloques de código se inician con ":".
//   - Cada nivel de bloque se indenta con 4 espacios.
//   - Al entrar a un bloque hijo → +4 espacios.
//   - Al salir a un bloque padre → -4 espacios.
//   - Cuando se regresa a 0 espacios, el bloque termina.
//   - Los bloques vacíos son válidos.
//   - Los bloques anidados se soportan indefinidamente.
//   - Los bloques asociados a control de flujo (si_es, mientras, porcada) delegan aquí.
//   - Los {} se reservan SOLO para datos estilo JSON (matriz, lista, diccionario, estructura).
//
// Este parser debe:
//   - Detectar ":" como inicio de bloque.
//   - Contar espacios al inicio de cada línea.
//   - Agrupar instrucciones en Hijos según nivel de indentación.
//   - Validar que la indentación sea múltiplo de 4.
//   - Terminar el bloque cuando la indentación baja.
//   - Reportar errores claros si la indentación es irregular.
//

// parseBloque recibe una lista de líneas y construye un Nodo de tipo "bloque".
func parseBloque(lineas []string, nivelBase int) *Nodo {
    nodo := &Nodo{
        Tipo:  "bloque",
        Hijos: []Nodo{},
        Extra: map[string]interface{}{
            "delimitador":       ":",
            "nivelIndentacion":  nivelBase,
        },
    }

    for i := 0; i < len(lineas); i++ {
        linea := lineas[i]
        trim := strings.TrimSpace(linea)
        if trim == "" {
            continue
        }

        // contar espacios iniciales
        espacios := contarEspaciosIniciales(linea)
        if espacios%4 != 0 {
            // error: indentación irregular
            nodo.Hijos = append(nodo.Hijos, Nodo{
                Tipo:  "error",
                Valor: "Indentación inválida, debe ser múltiplo de 4",
            })
            continue
        }

        // determinar nivel actual
        nivel := espacios

        // si el nivel es menor al nivel base, significa que el bloque terminó
        if nivel < nivelBase {
            break
        }

        // si el nivel es exactamente el nivel base, parsear instrucción normal
        if nivel == nivelBase {
            child := parseLinea(trim)
            if child != nil {
                nodo.Hijos = append(nodo.Hijos, *child)
            }
            continue
        }

        // si el nivel es mayor al nivel base, significa que es un bloque hijo
        if nivel == nivelBase+4 {
            // recolectar todas las líneas del sub‑bloque
            subBloque := recolectarSubBloque(lineas[i:], nivel)
            child := parseBloque(subBloque, nivel)
            nodo.Hijos = append(nodo.Hijos, *child)
            // saltar las líneas ya consumidas
            i += len(subBloque) - 1
            continue
        }

        // si el nivel es mayor a +4 → error
        if nivel > nivelBase+4 {
            nodo.Hijos = append(nodo.Hijos, Nodo{
                Tipo:  "error",
                Valor: "Indentación inesperada, exceso de niveles",
            })
        }
    }

    return nodo
}

// contarEspaciosIniciales devuelve el número de espacios al inicio de la línea.
func contarEspaciosIniciales(linea string) int {
    count := 0
    for _, r := range linea {
        if r == ' ' {
            count++
        } else {
            break
        }
    }
    return count
}

// recolectarSubBloque toma las líneas desde un índice y devuelve las que pertenecen al sub‑bloque.
func recolectarSubBloque(lineas []string, nivel int) []string {
    var sub []string
    for _, l := range lineas {
        esp := contarEspaciosIniciales(l)
        if esp < nivel {
            break
        }
        sub = append(sub, l)
    }
    return sub
}

// parseLinea: función auxiliar que invoca el parser correspondiente para una línea.
// Usa el registro central Parsers.
func parseLinea(linea string) *Nodo {
    for _, f := range Parsers {
        n := f(linea)
        if n != nil {
            return n
        }
    }
    return &Nodo{Tipo: "expresion", Valor: linea}
}

// --- Registro en Parsers ---
func init() {
    Parsers["bloque"] = func(linea string) *Nodo {
        trim := strings.TrimSpace(linea)
        if !strings.HasSuffix(trim, ":") {
            return nil
        }
        // Evitar colisiones con palabras similares
        if len(trim) > len("bloque") && unicode.IsLetter(rune(trim[len("bloque")])) {
            return nil
        }
        // Un bloque se parsea a partir de las líneas siguientes, aquí devolvemos un nodo vacío
        return &Nodo{Tipo: "bloque", Hijos: []Nodo{}, Extra: map[string]interface{}{"delimitador": ":"}}
    }
}
