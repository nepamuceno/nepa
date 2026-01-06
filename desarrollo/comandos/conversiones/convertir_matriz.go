package conversiones

import (
    "encoding/json"
    "nepa/desarrollo/interno/evaluador"
)

// Ayuda integrada para el comando convertir_matriz
const ayudaConvertirMatriz = `
convertir_matriz(valor) → matriz
Alias: a_matriz(valor), convertir.matriz(valor)

Convierte una cadena en formato JSON a una matriz.
Formatos aceptados:
- "[[1,2],[3,4]]"
- "[[true,false],[false,true]]"
- "[["a","b"],["c","d"]]"

Ejemplo:
    convertir_matriz("[[1,2],[3,4]]") → matriz con dos filas y dos columnas
`

// fnConvertirMatriz realiza la conversión robusta
func fnConvertirMatriz(args ...interface{}) interface{} {
    if len(args) < 1 {
        return evaluador.NuevaErrorConversion("convertir_matriz", ayudaConvertirMatriz, nil)
    }

    s, ok := args[0].(string)
    if !ok {
        return evaluador.NuevaErrorConversion("convertir_matriz", ayudaConvertirMatriz, args[0])
    }

    var m [][]interface{}
    if err := json.Unmarshal([]byte(s), &m); err != nil {
        return evaluador.NuevaErrorConversion("convertir_matriz", ayudaConvertirMatriz, s)
    }

    return m
}

// RegistrarConvertirMatriz registra el comando con sus tres alias
func RegistrarConvertirMatriz(ctx *evaluador.Contexto) {
    evaluador.Funciones["convertir_matriz"] = func(args ...interface{}) (interface{}, error) {
        r := fnConvertirMatriz(args...)
        if err, ok := r.(error); ok {
            return nil, err
        }
        return r, nil
    }
    evaluador.Funciones["a_matriz"] = evaluador.Funciones["convertir_matriz"]
    evaluador.Funciones["convertir.matriz"] = evaluador.Funciones["convertir_matriz"]
}
