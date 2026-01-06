package conversiones

import (
    "nepa/desarrollo/interno/evaluador"
)

// Ayuda integrada para el comando convertir_puntero
const ayudaConvertirPuntero = `
convertir_puntero(valor) → puntero
Alias: a_puntero(valor), convertir.puntero(valor)

Convierte cualquier valor en un puntero que puede ser desreferenciado.
Ejemplo:
    variable entero e := 42
    variable puntero p := convertir_puntero(e)
    variable entero e2 := desreferenciar(p)  → 42
`

// fnConvertirPuntero crea un puntero al valor dado
func fnConvertirPuntero(args ...interface{}) interface{} {
    if len(args) < 1 {
        return evaluador.NuevaErrorConversion("convertir_puntero", ayudaConvertirPuntero, nil)
    }
    return evaluador.Puntero{Valor: args[0]}
}

// RegistrarConvertirPuntero registra el comando con sus tres alias
func RegistrarConvertirPuntero(ctx *evaluador.Contexto) {
    evaluador.Funciones["convertir_puntero"] = func(args ...interface{}) (interface{}, error) {
        r := fnConvertirPuntero(args...)
        if err, ok := r.(error); ok {
            return nil, err
        }
        return r, nil
    }
    evaluador.Funciones["a_puntero"] = evaluador.Funciones["convertir_puntero"]
    evaluador.Funciones["convertir.puntero"] = evaluador.Funciones["convertir_puntero"]

    evaluador.Funciones["desreferenciar"] = func(args ...interface{}) (interface{}, error) {
        if len(args) < 1 {
            return nil, evaluador.NuevaErrorConversion("desreferenciar", "Se esperaba un puntero válido", nil)
        }
        p, ok := args[0].(evaluador.Puntero)
        if !ok {
            return nil, evaluador.NuevaErrorConversion("desreferenciar", "Se esperaba un puntero válido", args[0])
        }
        return p.Valor, nil
    }
}
