package core

import "fmt"

type ErrorEjecucion struct {
    Archivo string
    Linea   int
    Mensaje string
}

func (e ErrorEjecucion) Error() string {
    rojoBold := "\033[1;31m"  // rojo + bold
    grisBold := "\033[1;90m"  // gris + bold
    reset := "\033[0m"        // reset

    return fmt.Sprintf("%serror%s %s%s:%d%s → %s",
        rojoBold, reset, grisBold, e.Archivo, e.Linea, reset, e.Mensaje)
}
