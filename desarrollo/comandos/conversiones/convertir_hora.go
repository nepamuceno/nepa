package conversiones

import (
    "time"
    "nepa/desarrollo/interno/evaluador"
)

const ayudaConvertirHora = `
convertir_hora(valor) → hora
Alias: a_hora(valor), convertir.hora(valor)

Convierte una cadena en formato HH:MM:SS a un objeto hora.
Ejemplo: convertir_hora("12:48:00")
`

func fnConvertirHora(args ...interface{}) interface{} {
    if len(args) < 1 {
        return evaluador.NuevaErrorConversion("convertir_hora", ayudaConvertirHora, nil)
    }
    s, ok := args[0].(string)
    if !ok {
        return evaluador.NuevaErrorConversion("convertir_hora", ayudaConvertirHora, args[0])
    }
    t, err := time.Parse("15:04:05", s)
    if err != nil {
        return evaluador.NuevaErrorConversion("convertir_hora", ayudaConvertirHora, s)
    }
    return t
}

func RegistrarConvertirHora(ctx *evaluador.Contexto) {
    evaluador.Funciones["convertir_hora"] = func(args ...interface{}) (interface{}, error) {
        r := fnConvertirHora(args...)
        if err, ok := r.(error); ok {
            return nil, err
        }
        return r, nil
    }
    evaluador.Funciones["a_hora"] = evaluador.Funciones["convertir_hora"]
    evaluador.Funciones["convertir.hora"] = evaluador.Funciones["convertir_hora"]
}
