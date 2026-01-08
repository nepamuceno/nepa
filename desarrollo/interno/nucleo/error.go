package nucleo

import (
    "fmt"
    "os"
)

// TIPOS DE MENSAJE
const (
    _FATAL       = "❌ FATAL"
    _ADVERTENCIA = "⚠️ ADVERTENCIA"
    _INFO        = "ℹ️ INFO"
)

// EMITIR ERROR
func EmitirError(tipo string, archivo string, linea int, codigo int, args ...interface{}) {
    plantilla, ok := MENSAJES_ERROR[codigo]
    if !ok {
        plantilla = MENSAJES_ERROR[9999] // Error desconocido centralizado
    }
    mensaje := fmt.Sprintf(plantilla, archivo, linea, codigo, args...)
    fmt.Fprintf(os.Stderr, "%s %s\n", tipo, mensaje)
}

// EMITIR DETALLE
func EmitirDetalle(nivel int, archivo string, linea int, codigo int, args ...interface{}) {
    if nivel <= _DETALLE {
        plantilla, ok := MENSAJES_DETALLE[codigo]
        if !ok {
            plantilla = MENSAJES_DETALLE[6999] // Detalle desconocido centralizado
        }
        mensaje := fmt.Sprintf(plantilla, args...)
        fmt.Fprintf(os.Stdout, "[DETALLE-%d] %s\n", nivel, mensaje)
    }
}

// EMITIR DEPURACION
func EmitirDepuracion(nivel int, archivo string, linea int, codigo int, args ...interface{}) {
    if nivel <= _DEPURACION {
        plantilla, ok := MENSAJES_DEPURACION[codigo]
        if !ok {
            plantilla = MENSAJES_DEPURACION[6099] // Depuración desconocida centralizado
        }
        mensaje := fmt.Sprintf(plantilla, args...)
        fmt.Fprintf(os.Stdout, "[DEPURACION-%d] %s\n", nivel, mensaje)
    }
}
