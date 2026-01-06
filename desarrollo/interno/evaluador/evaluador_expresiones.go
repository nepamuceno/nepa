package evaluador

import (
    "fmt"
    "strings"
)

// ResolverExpresion ahora soporta:
// - &x  → referencia (puntero a x)
// - *p  → desreferencia (valor apuntado por p)
// - (tipo) expr → cast normal
// - (tipo*) expr / (puntero tipo) expr → cast a puntero del tipo
func ResolverExpresion(entrada string, ctx *Contexto) (interface{}, error) {
    if entrada == "" {
        return nil, fmt.Errorf("la expresión está vacía")
    }
    entrada = strings.TrimSpace(entrada)

    // &x → referencia
    if strings.HasPrefix(entrada, "&") {
        nombre := strings.TrimSpace(entrada[1:])
        valor, err := EvalConContexto(nombre, ctx)
        if err != nil {
            return nil, fmt.Errorf("error al resolver referencia '&%s': %v", nombre, err)
        }
        return NuevoPuntero(valor), nil
    }

    // *p → desreferencia
    if strings.HasPrefix(entrada, "*") {
        expr := strings.TrimSpace(entrada[1:])
        v, err := EvalConContexto(expr, ctx)
        if err != nil {
            return nil, fmt.Errorf("error al desreferenciar '*%s': %v", expr, err)
        }
        if p, ok := v.(Puntero); ok {
            return p.Valor, nil
        }
        return nil, fmt.Errorf("no se puede desreferenciar: %T no es puntero", v)
    }

    // (tipo) expr  y  (tipo*) expr  y  (puntero tipo) expr
    if strings.HasPrefix(entrada, "(") && strings.Contains(entrada, ")") {
        cierre := strings.Index(entrada, ")")
        if cierre > 1 {
            tipoBruto := strings.TrimSpace(entrada[1:cierre])
            valorExpr := strings.TrimSpace(entrada[cierre+1:])

            // Detectar forma (puntero tipo)
            esFormaPuntero := false
            tipoDestino := tipoBruto
            if strings.HasPrefix(strings.ToLower(tipoBruto), "puntero ") {
                esFormaPuntero = true
                tipoDestino = strings.TrimSpace(tipoBruto[len("puntero "):])
            }
            // Detectar sufijo * → (entero*)
            if strings.HasSuffix(tipoDestino, "*") {
                esFormaPuntero = true
                tipoDestino = strings.TrimSuffix(tipoDestino, "*")
                tipoDestino = strings.TrimSpace(tipoDestino)
            }

            // Evaluar el valor original
            valor, err := EvalConContexto(valorExpr, ctx)
            if err != nil {
                return nil, fmt.Errorf("error al resolver valor en cast: %v", err)
            }

            // Resolver conversión base
            var convertido interface{}
            switch tipoDestino {
            case "entero":
                convertido = ctx.Funciones["convertir_entero"](valor)
            case "real":
                convertido = ctx.Funciones["convertir_real"](valor)
            case "cadena":
                convertido = ctx.Funciones["convertir_cadena"](valor)
            case "booleano":
                convertido = ctx.Funciones["convertir_booleano"](valor)
            case "binario":
                convertido = ctx.Funciones["convertir_binario"](valor)
            case "hexadecimal", "hex":
                convertido = ctx.Funciones["convertir_hexadecimal"](valor)
            case "fecha":
                convertido = ctx.Funciones["convertir_fecha"](valor)
            case "hora":
                convertido = ctx.Funciones["convertir_hora"](valor)
            case "tiempo":
                convertido = ctx.Funciones["convertir_tiempo"](valor)
            case "matriz":
                convertido = ctx.Funciones["convertir_matriz"](valor)
            default:
                return nil, NuevaErrorConversion("conversion", "Tipo de conversión no soportado", tipoDestino)
            }

            // Si la conversión devolvió error, propagar
            if convErr, ok := convertido.(error); ok {
                return nil, convErr
            }

            // Si se pidió puntero, envolver
            if esFormaPuntero {
                return NuevoPuntero(convertido), nil
            }
            return convertido, nil
        }
    }

    // Caso normal
    resultado, err := EvalConContexto(entrada, ctx)
    if err != nil {
        return nil, fmt.Errorf("error al resolver expresión: %v", err)
    }
    return resultado, nil
}

// --- Funciones explícitas de puntero ---

// fnPuntero crea un puntero a un valor
func fnPuntero(args ...interface{}) interface{} {
    if len(args) < 1 {
        return NuevaErrorConversion("puntero", "Se esperaba un valor para crear puntero", nil)
    }
    return NuevoPuntero(args[0])
}

// fnDesreferenciar devuelve el valor apuntado por un puntero
func fnDesreferenciar(args ...interface{}) interface{} {
    if len(args) < 1 {
        return NuevaErrorConversion("desreferenciar", "Se esperaba un puntero para desreferenciar", nil)
    }
    v := args[0]
    if p, ok := v.(Puntero); ok {
        return p.Valor
    }
    return NuevaErrorConversion("desreferenciar", "El valor no es puntero", v)
}

// RegistrarFuncionesPuntero agrega las funciones de puntero al contexto
func RegistrarFuncionesPuntero(ctx *Contexto) {
    ctx.Funciones["puntero"] = fnPuntero
    ctx.Funciones["desreferenciar"] = fnDesreferenciar
}
