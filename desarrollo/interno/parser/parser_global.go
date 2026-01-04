package parser

import (
	"strings"
)

func parseGlobal(linea string) *Nodo {
	if strings.HasPrefix(linea, "global ") && strings.Contains(linea, "=") {
		partes := strings.SplitN(linea[len("global "):], "=", 2)
		if len(partes) == 2 {
			return &Nodo{
				Tipo:   "global",
				Nombre: strings.TrimSpace(partes[0]),
				Valor:  parseValor(strings.TrimSpace(partes[1])),
			}
		}
	}
	return nil
}
