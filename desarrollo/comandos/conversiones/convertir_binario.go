package conversiones

import (
    "fmt"
    "strconv"
    "nepa/desarrollo/interno/evaluador"
)

// Ayuda integrada para el comando convertir_binario
const ayudaConvertirBinario = `
convertir_binario(valor) → cadena
Alias: a_binario(valor), convertir.binario(valor)

Convierte un número entero o real a su representación binaria.
Ejemplo:
    convertir_binario(255) → "11111111"
    convertir_binario("10") → "1010"
`

// fnConvertirBinario realiza la conversión robusta
func fnConvertirBinario(args ...interface{}) interface{} {
    if len(args) < 1 {
        return evaluador.NuevaErrorConversion("convertir_binario", ayudaConvertirBinario, nil)
    }

    switch v := args[0].(type) {
    case int:
        return fmt.Sprintf("%b", v)

    case int64:
        return fmt.Sprintf("%b", v)

    case float64:
        return fmt.Sprintf("%b", int(v))

    case string:
        // Intentar convertir cadena a entero
        n, err := strconv.Atoi(v)
        if err != nil {
            f, err2 := strconv.ParseFloat(v, 64)
            if err2 != nil {
                return evaluador.NuevaErrorConversion("convertir_binario", ayudaConvertirBinario, v)
            }
            return fmt.Sprintf("%b", int(f))
        }
        return fmt.Sprintf("%b", n)

    default:
        return evaluador.NuevaErrorConversion("convertir_binario", ayudaConvertirBinario, v)
    }
}

// RegistrarConvertirBinario registra el comando con sus tres alias
func RegistrarConvertirBinario(ctx *evaluador.Contexto) {
    evaluador.Funciones["convertir_binario"] = func(args ...interface{}) (interface{}, error) {
        r := fnConvertirBinario(args...)
        if err, ok := r.(error); ok {
            return nil, err
        }
        return r, nil
    }
    evaluador.Funciones["a_binario"] = evaluador.Funciones["convertir_binario"]
    evaluador.Funciones["convertir.binario"] = evaluador.Funciones["convertir_binario"]
}
