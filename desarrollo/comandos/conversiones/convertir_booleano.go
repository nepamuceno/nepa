package conversiones

import (
    "strings"
    "nepa/desarrollo/interno/evaluador"
)

// Ayuda integrada para el comando convertir_booleano
const ayudaConvertirBooleano = `
convertir_booleano(valor) → booleano
Alias: a_booleano(valor), convertir.booleano(valor)

Se espera:
- cadena "verdadero"/"falso" o "true"/"false" → booleano
- entero (0 → falso, distinto de 0 → verdadero)
- real (0.0 → falso, distinto de 0.0 → verdadero)
- booleano (ya es booleano) → se devuelve tal cual

Devuelve un booleano. Si el valor no puede convertirse, se muestra este mensaje de ayuda.
`

// fnConvertirBooleano realiza la conversión robusta
func fnConvertirBooleano(args ...interface{}) interface{} {
    if len(args) < 1 {
        return evaluador.NuevaErrorConversion("convertir_booleano", ayudaConvertirBooleano, nil)
    }

    switch v := args[0].(type) {
    case bool:
        return v

    case int:
        return v != 0

    case int64:
        return v != 0

    case float64:
        return v != 0.0

    case string:
        s := strings.ToLower(strings.TrimSpace(v))
        switch s {
        case "true", "verdadero", "sí", "si":
            return true
        case "false", "falso", "no":
            return false
        default:
            return evaluador.NuevaErrorConversion("convertir_booleano", ayudaConvertirBooleano, v)
        }

    default:
        return evaluador.NuevaErrorConversion("convertir_booleano", ayudaConvertirBooleano, v)
    }
}

// RegistrarConvertirBooleano registra el comando con sus tres alias
func RegistrarConvertirBooleano(ctx *evaluador.Contexto) {
    evaluador.Funciones["convertir_booleano"] = func(args ...interface{}) (interface{}, error) {
        r := fnConvertirBooleano(args...)
        if err, ok := r.(error); ok {
            return nil, err
        }
        return r, nil
    }
    evaluador.Funciones["a_booleano"] = evaluador.Funciones["convertir_booleano"]
    evaluador.Funciones["convertir.booleano"] = evaluador.Funciones["convertir_booleano"]
}

