package parser

import (
    "strings"
)

// parseFuncion: Maneja la declaración de funciones estilo Python/nepa.
// - Sintaxis: funcion <nombre>([args]):
// - Bloque definido por indentación de 4 espacios.
// - Retorno por defecto: entero (0=éxito, 1=error).
// - Si se usa regresa_valor (<tipo>) <var>, se fuerza ese tipo.
// - Si se usa regresa <expresión>, se infiere el tipo de la expresión.
// - Detecta romper como salida de pánico.
func parseFuncion(linea string, cuerpo []string) *Nodo {
    if !strings.HasPrefix(linea, "funcion ") {
        return nil
    }

    def := strings.TrimSpace(strings.TrimPrefix(linea, "funcion "))
    if def == "" {
        return nil
    }

    // Separar nombre y parámetros
    partes := strings.SplitN(def, "(", 2)
    if len(partes) != 2 {
        return nil
    }

    nombre := strings.TrimSpace(partes[0])
    resto := partes[1]

    partes2 := strings.SplitN(resto, ")", 2)
    if len(partes2) != 2 {
        return nil
    }

    paramsParte := strings.TrimSpace(partes2[0])
    bloqueIndicador := strings.TrimSpace(partes2[1])

    // Validar que termina con ":"
    if !strings.HasSuffix(bloqueIndicador, ":") {
        return nil
    }

    // Parsear parámetros
    var params []Nodo
    if paramsParte != "" {
        listaParams := strings.Split(paramsParte, ",")
        for _, p := range listaParams {
            p = strings.TrimSpace(p)
            if p == "" {
                continue
            }
            campos := strings.Fields(p)
            tipoTokens, nextIdx := extraerTipoTokens(campos)
            if len(tipoTokens) == 0 || nextIdx >= len(campos) {
                continue
            }
            nombreParam := campos[nextIdx]
            params = append(params, Nodo{
                Tipo:   "parametro",
                Nombre: nombreParam,
                Args:   tipoTokens,
                Valor:  nil,
            })
        }
    }

    // Retorno por defecto: entero (0=éxito, 1=error)
    retorno := []interface{}{"entero"}
    retornoVar := ""
    retornoExpr := ""

    // Parsear cuerpo indentado (4 espacios)
    var cuerpoNodos []Nodo
    for _, lineaC := range cuerpo {
        if strings.HasPrefix(lineaC, "    ") { // 4 espacios
            instr := strings.TrimSpace(lineaC)

            // Detectar romper
            if instr == "romper" {
                cuerpoNodos = append(cuerpoNodos, Nodo{Tipo: "romper", Valor: "panico"})
                retorno = []interface{}{"entero"}
                retornoExpr = "romper"
                continue
            }

            // Detectar regresa_valor
            if strings.HasPrefix(instr, "regresa_valor") {
                campos := strings.Fields(strings.TrimPrefix(instr, "regresa_valor"))
                tipoTokens, nextIdx := extraerTipoTokens(campos)
                if len(tipoTokens) > 0 && nextIdx < len(campos) {
                    retorno = tipoTokens
                    retornoVar = campos[nextIdx]
                }
                cuerpoNodos = append(cuerpoNodos, Nodo{Tipo: "regresa_valor", Args: retorno, Nombre: retornoVar})
                continue
            }

            // Detectar regresa <expresión>
            if strings.HasPrefix(instr, "regresa ") {
                expr := strings.TrimSpace(strings.TrimPrefix(instr, "regresa"))
                retornoExpr = expr
                if strings.Contains(expr, ".") {
                    retorno = []interface{}{"real"}
                } else if strings.HasPrefix(expr, "\"") && strings.HasSuffix(expr, "\"") {
                    retorno = []interface{}{"texto"}
                } else if strings.Contains(expr, "[") && strings.Contains(expr, "]") {
                    retorno = []interface{}{"matriz"}
                } else {
                    retorno = []interface{}{"entero"}
                }
                cuerpoNodos = append(cuerpoNodos, Nodo{Tipo: "regresa", Valor: expr})
                continue
            }

            // Otras instrucciones
            if nodo := parseBloque(instr); nodo != nil {
                cuerpoNodos = append(cuerpoNodos, *nodo)
            }
        }
    }

    return &Nodo{
        Tipo:   "funcion",
        Nombre: nombre,
        Args:   retorno,   // tipo de retorno
        Valor:  append(params, cuerpoNodos...), // parámetros + cuerpo
        Extra: map[string]interface{}{
            "retornoVar":  retornoVar,
            "retornoExpr": retornoExpr,
        },
    }
}

// --- Registro en Parsers ---
func init() {
    Parsers["funcion"] = func(linea string) *Nodo {
        // Nota: funciones requieren cuerpo, así que aquí se devuelve nil.
        // El recolector de bloques es quien llama a parseFuncion con cuerpo.
        return nil
    }
}
