package parser

import (
    "strings"
)

// Tipos base disponibles según desarrollo/interno/variables/
var TiposBase = map[string]bool{
    "bit":         true,
    "booleano":    true,
    "cadena":      true,
    "caracter":    true,
    "complejo":    true,
    "decimal":     true,
    "diccionario": true,
    "entero":      true,
    "fecha":       true,
    "hora":        true,
    "lista":       true,
    "matriz":      true,
    "objeto":      true,
    "puntero":     true,
    "real":        true,
    "texto":       true,
    "tiempo":      true,
}
