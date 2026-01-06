package conversiones

import (
    "time"
    "nepa/desarrollo/interno/evaluador"
)

// Ayuda integrada para el comando convertir_fecha
const ayudaConvertirFecha = `
convertir_fecha(valor) → fecha
Alias: a_fecha(valor), convertir.fecha(valor)

Convierte una cadena en formato de fecha a un objeto fecha.
Formatos aceptados:
- "AAAA-MM-DD" → fecha
- "AAAA-MM-DD HH:MM:SS" → fecha y hora

Ejemplo:
    convertir_fecha("2026-01-05") → 5 de enero de 2026
    convertir_fecha("2026-01-05 12:34:56") → 5 de enero de 2026, 12:34:56
`

// fnConvertirFecha realiza la conversión robusta
func fnConvertirFecha(args ...interface{}) interface{} {
    if len(args) < 1 {
        return evaluador.NuevaErrorConversion("convertir_fecha", ayudaConvertirFecha, nil)
    }

    s, ok := args[0].(string)
    if !ok {
        return evaluador.NuevaErrorConversion("convertir_fecha", ayudaConvertirFecha, args[0])
    }

    // Intentar formato completo con hora
    if t, err := time.Parse("2006-01-02 15:04:05", s); err == nil {
        return t
    }

    // Intentar solo fecha
    if t, err := time.Parse("2006-01-02", s); err == nil {
        return t
    }

    // Ningún formato válido
    return evaluador.NuevaErrorConversion("convertir_fecha", ayudaConvertirFecha, s)
}

// RegistrarConvertirFecha registra el comando con sus tres alias
func RegistrarConvertirFecha(ctx *evaluador.Contexto) {
    evaluador.Funciones["convertir_fecha"] = func(args ...interface{}) (interface{}, error) {
        r := fnConvertirFecha(args...)
        if err, ok := r.(error); ok {
            return nil, err
        }
        return r, nil
    }
    evaluador.Funciones["a_fecha"] = evaluador.Funciones["convertir_fecha"]
    evaluador.Funciones["convertir.fecha"] = evaluador.Funciones["convertir_fecha"]
}

