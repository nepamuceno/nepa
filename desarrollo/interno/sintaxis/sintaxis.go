package sintaxis

import (
    "fmt"
    "strings"

    "nepa/desarrollo/interno/bloque"
)

// ValidarLinea revisa una línea de código y devuelve error si hay problema
func ValidarLinea(linea string, num int, archivo string) error {
    l := strings.TrimSpace(linea)
    if l == "" {
        return nil
    }

    // Validar que no use palabras reservadas como variables
    for _, palabra := range bloque.PalabrasReservadas {
        if strings.HasPrefix(l, palabra+" ") || strings.HasPrefix(l, palabra+"=") {
            return fmt.Errorf("Uso inválido de palabra reservada '%s' en %s línea %d", palabra, archivo, num)
        }
    }

    // Validar bloques reservados (ejemplo: "si", "para", etc.)
    for _, b := range bloque.BloquesReservados {
        if strings.HasPrefix(l, b) && !strings.HasSuffix(l, ":") {
            return fmt.Errorf("Bloque '%s' en %s línea %d requiere ':' al final", b, archivo, num)
        }
    }

    // Validar paréntesis balanceados
    if strings.Count(l, "(") != strings.Count(l, ")") {
        return fmt.Errorf("Paréntesis desbalanceados en %s línea %d", archivo, num)
    }

    // Validar comillas dobles balanceadas
    if strings.Count(l, "\"")%2 != 0 {
        return fmt.Errorf("Comillas dobles desbalanceadas en %s línea %d", archivo, num)
    }

    // Validar comillas simples balanceadas
    if strings.Count(l, "'")%2 != 0 {
        return fmt.Errorf("Comillas simples desbalanceadas en %s línea %d", archivo, num)
    }

    return nil
}

// EsReservada indica si una palabra es reservada en Nepa
func EsReservada(p string) bool {
    for _, r := range bloque.PalabrasReservadas {
        if p == r {
            return true
        }
    }
    for _, r := range bloque.BloquesReservados {
        if p == r {
            return true
        }
    }
    return false
}
