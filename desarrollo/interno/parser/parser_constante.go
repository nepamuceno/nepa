package parser

import (
    "strings"
)

// parseConst: Maneja declaraciones de constantes.
// Siempre requiere un valor inicial. Soporta tipos compuestos y valores complejos.
// Usa TiposBase centralizado desde parser_tipo_nativo.go.
func parseConst(linea string) *Nodo {
    if !strings.HasPrefix(linea, "constante ") {
        return nil
    }

    def := strings.TrimSpace(strings.TrimPrefix(linea, "constante "))
    if def == "" {
        return nil
    }

    campos := strings.Fields(def)
    if len(campos) < 3 {
        return nil
    }

    // Extraer tokens de tipo (puntero, matriz, etc.)
    tipoTokens, nextIdx := extraerTipoTokens(campos)
    if len(tipoTokens) == 0 {
        return nil
    }

    resto := strings.TrimSpace(strings.Join(campos[nextIdx:], " "))

    // Constantes deben tener asignación obligatoria
    if !strings.Contains(resto, ":=") {
        return nil
    }

    partes := strings.SplitN(resto, ":=", 2)
    nombresParte := strings.TrimSpace(partes[0])
    valorParte := strings.TrimSpace(partes[1])

    if nombresParte == "" || valorParte == "" {
        return nil
    }

    valor := parseValor(valorParte)

    // Separar nombres por coma
    nombres := strings.Split(nombresParte, ",")
    var nodos []Nodo
    for _, n := range nombres {
        nombre := strings.TrimSpace(n)
        if nombre == "" {
            continue
        }
        nodos = append(nodos, Nodo{
            Tipo:   "constante",
            Nombre: nombre,
            Args:   tipoTokens,
            Valor:  valor,
        })
    }

    if len(nodos) == 1 {
        return &nodos[0]
    }

    // Para múltiples constantes, devolvemos un nodo agrupador con hijos
    return &Nodo{
        Tipo:   "constantes",
        Nombre: "",
        Args:   tipoTokens,
        Valor:  nodos,
    }
}

// --- Registro en Parsers ---
func init() {
    Parsers["constante"] = func(linea string) *Nodo {
        return parseConst(linea)
    }
}
