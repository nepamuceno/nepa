package conversiones

import (
    "time"
    "nepa/desarrollo/interno/evaluador"
)

// Ayuda integrada para el comando convertir_tiempo
const ayudaConvertirTiempo = `
convertir_tiempo(valor) → tiempo
Alias: a_tiempo(valor), convertir.tiempo(valor)

Convierte una cadena en formato de duración a un objeto tiempo.
Formatos aceptados (según Go):
- "2h30m" → 2 horas 30 minutos
- "45s"   → 45 segundos
- "1h15m30s" → 1 hora 15 minutos 30 segundos

Ejemplo:
    convertir_tiempo("2h30m") → duración de 2 horas y 30 minutos
    convertir_tiempo("90s")   → duración de 90 segundos
`

// fnConvertirTiempo realiza la conversión robusta
func fnConvertirTiempo(args ...interface{}) interface{} {
    if len(args) < 1 {
        return evaluador.NuevaErrorConversion("convertir_tiempo", ayudaConvertirTiempo, nil)
    }

    s, ok := args[0].(string)
    if !ok {
        return evaluador.NuevaErrorConversion("convertir_tiempo", ayudaConvertirTiempo, args[0])
    }

    // Usar el parser de duraciones de Go
    d, err := time.ParseDuration(s)
    if err != nil {
        return evaluador.NuevaErrorConversion("convertir_tiempo", ayudaConvertirTiempo, s)
    }

    return d
}

// RegistrarConvertirTiempo registra el comando con sus tres alias
func RegistrarConvertirTiempo(ctx *evaluador.Contexto) {
    evaluador.Funciones["convertir_tiempo"] = func(args ...interface{}) (interface{}, error) {
        r := fnConvertirTiempo(args...)
        if err, ok := r.(error); ok {
            return nil, err
        }
        return r, nil
    }
    evaluador.Funciones["a_tiempo"] = evaluador.Funciones["convertir_tiempo"]
    evaluador.Funciones["convertir.tiempo"] = evaluador.Funciones["convertir_tiempo"]
}

