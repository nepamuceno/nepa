package core

import (
    "bytes"
    "fmt"
    "os"
)

// Versión global del intérprete
const Version = "0.0002"

// NepaError es el formato estándar de error en Nepa
type NepaError struct {
    Archivo string
    Linea   int
    Mensaje string
}

func (e NepaError) Error() string {
    if e.Archivo != "" && e.Linea > 0 {
        return fmt.Sprintf("Error en %s, línea %d: %s", e.Archivo, e.Linea, e.Mensaje)
    }
    return fmt.Sprintf("Error: %s", e.Mensaje)
}

// NewError crea un NepaError de forma sencilla
func NewError(archivo string, linea int, mensaje string) error {
    return NepaError{Archivo: archivo, Linea: linea, Mensaje: mensaje}
}

// StripBOM remueve BOM de una cadena si existe
func StripBOM(s string) string {
    return string(bytes.TrimPrefix([]byte(s), []byte{0xEF, 0xBB, 0xBF}))
}

// --- Logging helpers ---
var DebugMode = false

// Debug imprime mensajes solo si está activado el modo debug
func Debug(msg string) {
    if DebugMode {
        fmt.Fprintln(os.Stderr, "[DEPURAR]", msg)
    }
}

// Info imprime mensajes informativos
func Info(msg string) {
    fmt.Fprintln(os.Stderr, "[INFO]", msg)
}

// Warn imprime advertencias
func Warn(msg string) {
    fmt.Fprintln(os.Stderr, "[ADVERTENCIA]", msg)
}

// Error imprime errores
func Error(msg string) {
    fmt.Fprintln(os.Stderr, "[ERROR]", msg)
}
