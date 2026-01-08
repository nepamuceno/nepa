package nucleo

import "bytes"

// FUNCIONES DE UTILIDAD GENERALES PARA NEPA

// STRIP_BOM remueve el BOM de una cadena si existe
func STRIP_BOM(cadena string) string {
    return string(bytes.TrimPrefix([]byte(cadena), []byte{0xEF, 0xBB, 0xBF}))
}
