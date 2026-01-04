package parser

import (
	"strings"
)

func parseConst(linea string) *Nodo {
	// Cambiado de "const " a "constante " según la nueva regla
	if strings.HasPrefix(linea, "constante ") && strings.Contains(linea, "=") {
		partes := strings.SplitN(linea[len("constante "):], "=", 2)
		if len(partes) == 2 {
			return &Nodo{
				Tipo:   "constante",
				Nombre: strings.TrimSpace(partes[0]),
				Valor:  parseValor(strings.TrimSpace(partes[1])),
			}
		}
	}
	return nil
}
