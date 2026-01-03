package evaluador

import "fmt"

// FormatearValor convierte un valor a cadena para mostrarlo en main.go
func FormatearValor(v interface{}) string {
    return fmt.Sprintf("%v", v)
}
