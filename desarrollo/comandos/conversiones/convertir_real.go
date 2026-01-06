package conversiones

import (
    "strconv"
    "nepa/desarrollo/interno/evaluador"
)

// Ayuda integrada para el comando convertir_real
const ayudaConvertirReal = `
convertir_real(valor) → real
Alias: a_real(valor), convertir.real(valor)

Se espera:
- cadena numérica ("3.14") → real 3.14
- entero (123) → real 123.0
- booleano (verdadero/falso) → 1.0/0.0
- real (ya es real) → se devuelve tal cual

Devuelve un número real (float64). Si el valor no puede convertirse, se muestra este mensaje de ayuda.
`

// fnConvertirReal realiza la conversión robusta
func fnConvertirReal(args ...interface{}) interface{} {
    if len(args) < 1 {
        return evaluador.NuevaErrorConversion("convertir_real", ayudaConvertirReal, nil)
    }

    switch v := args[0].(type) {
    case float64:
        // Ya es real
        return v

    case int:
        return float64(v)

    case int64:
        return float64(v)

    case bool:
        if v {
            return 1.0
        }
        return 0.0

    case string:
        // Intentar convertir cadena a real
        n, err := strconv.ParseFloat(v, 64)
        if err != nil {
            return evaluador.NuevaErrorConversion("convertir_real", ayudaConvertirReal, v)
        }
        return n

    default:
        // Tipo no soportado
        return evaluador.NuevaErrorConversion("convertir_real", ayudaConvertirReal, v)
    }
}

// RegistrarConvertirReal registra el comando con sus tres alias
func RegistrarConvertirReal(ctx *evaluador.Contexto) {
    evaluador.Funciones["convertir_real"] = func(args ...interface{}) (interface{}, error) {
        r := fnConvertirReal(args...)
        if err, ok := r.(error); ok {
            return nil, err
        }
        return r, nil
    }
    evaluador.Funciones["a_real"] = evaluador.Funciones["convertir_real"]
    evaluador.Funciones["convertir.real"] = evaluador.Funciones["convertir_real"]
}
