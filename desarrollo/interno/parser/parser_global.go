package parser

import (
    "strings"
)

// parseGlobal: Maneja declaraciones de variables globales.
// Puede tener o no valor inicial. Soporta tipos compuestos y múltiples nombres.
// Usa TiposBase centralizado desde parser_tipo_nativo.go.
func parseGlobal(linea string) *Nodo {
    if !strings.HasPrefix(linea, "global ") {
        return nil
    }

    def := strings.TrimSpace(strings.TrimPrefix(linea, "global "))
    if def == "" {
        return nil
    }

    campos := strings.Fields(def)
    if len(campos) < 2 {
        return nil
    }

    // Extraer tokens de tipo (puntero, matriz, etc.)
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

    // Separar nombres por coma
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
            Tipo:   "global",
            Nombre: nombre,
            Args:   tipoTokens,
            Valor:  valor,
        })
    }

    if len(nodos) == 1 {
        return &nodos[0]
    }

    // Para múltiples globales, devolvemos un nodo agrupador con hijos
    return &Nodo{
        Tipo:   "globales",
        Nombre: "",
        Args:   tipoTokens,
        Valor:  nodos,
    }
}

// --- Registro en Parsers ---
func init() {
    Parsers["global"] = func(linea string) *Nodo {
        return parseGlobal(linea)
    }
}
