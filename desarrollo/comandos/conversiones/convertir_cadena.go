package conversiones

import (
    "fmt"
    "nepa/desarrollo/interno/evaluador"
)

// Ayuda integrada para el comando convertir_cadena
const ayudaConvertirCadena = `
convertir_cadena(valor) → cadena
Alias: a_cadena(valor), convertir.cadena(valor)

Se espera:
- entero (123) → "123"
- real (3.14) → "3.14"
- booleano (verdadero/falso) → "verdadero"/"falso"
- cadena (ya es cadena) → se devuelve tal cual
- puntero → representación textual del puntero

Devuelve una cadena. Si el valor no puede convertirse, se muestra este mensaje de ayuda.
`

// fnConvertirCadena realiza la conversión robusta
func fnConvertirCadena(args ...interface{}) interface{} {
    if len(args) < 1 {
        return evaluador.NuevaErrorConversion("convertir_cadena", ayudaConvertirCadena, nil)
    }

    switch v := args[0].(type) {
    case string:
        // Ya es cadena
        return v

    case int:
        return fmt.Sprintf("%d", v)

    case int64:
        return fmt.Sprintf("%d", v)

    case float64:
        return fmt.Sprintf("%g", v)

    case bool:
        if v {
            return "verdadero"
        }
        return "falso"

    case evaluador.Puntero:
        // Representación textual de puntero
        return fmt.Sprintf("&%v", v.Valor)

    default:
        // Tipo no soportado
        return evaluador.NuevaErrorConversion("convertir_cadena", ayudaConvertirCadena, v)
    }
}

// RegistrarConvertirCadena registra el comando con sus tres alias
func RegistrarConvertirCadena(ctx *evaluador.Contexto) {
    evaluador.Funciones["convertir_cadena"] = func(args ...interface{}) (interface{}, error) {
        r := fnConvertirCadena(args...)
        if err, ok := r.(error); ok {
            return nil, err
        }
        return r, nil
    }
    evaluador.Funciones["a_cadena"] = evaluador.Funciones["convertir_cadena"]
    evaluador.Funciones["convertir.cadena"] = evaluador.Funciones["convertir_cadena"]
}
