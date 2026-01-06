package conversiones

import (
    "fmt"
    "strconv"
    "nepa/desarrollo/interno/evaluador"
)

// Ayuda integrada para el comando convertir_hexadecimal
const ayudaConvertirHexadecimal = `
convertir_hexadecimal(valor) → cadena
Alias: a_hexadecimal(valor), convertir.hexadecimal(valor)

Convierte un número entero o real a su representación hexadecimal.
Ejemplo:
    convertir_hexadecimal(255) → "ff"
    convertir_hexadecimal("1024") → "400"
`

// fnConvertirHexadecimal realiza la conversión robusta
func fnConvertirHexadecimal(args ...interface{}) interface{} {
    if len(args) < 1 {
        return evaluador.NuevaErrorConversion("convertir_hexadecimal", ayudaConvertirHexadecimal, nil)
    }

    switch v := args[0].(type) {
    case int:
        return fmt.Sprintf("%x", v)

    case int64:
        return fmt.Sprintf("%x", v)

    case float64:
        return fmt.Sprintf("%x", int(v))

    case string:
        // Intentar convertir cadena a entero
        n, err := strconv.Atoi(v)
        if err != nil {
            f, err2 := strconv.ParseFloat(v, 64)
            if err2 != nil {
                return evaluador.NuevaErrorConversion("convertir_hexadecimal", ayudaConvertirHexadecimal, v)
            }
            return fmt.Sprintf("%x", int(f))
        }
        return fmt.Sprintf("%x", n)

    default:
        return evaluador.NuevaErrorConversion("convertir_hexadecimal", ayudaConvertirHexadecimal, v)
    }
}

// RegistrarConvertirHexadecimal registra el comando con sus tres alias
func RegistrarConvertirHexadecimal(ctx *evaluador.Contexto) {
    evaluador.Funciones["convertir_hexadecimal"] = func(args ...interface{}) (interface{}, error) {
        r := fnConvertirHexadecimal(args...)
        if err, ok := r.(error); ok {
            return nil, err
        }
        return r, nil
    }
    evaluador.Funciones["a_hexadecimal"] = evaluador.Funciones["convertir_hexadecimal"]
    evaluador.Funciones["convertir.hexadecimal"] = evaluador.Funciones["convertir_hexadecimal"]
}
