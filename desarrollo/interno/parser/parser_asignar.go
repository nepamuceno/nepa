package parser

import (
	"strings"
)

// parseAsignar: Maneja la asignación y reasignación. 
// Movido aquí para desaturar el core.
func parseAsignar(linea string) *Nodo {
	if !strings.Contains(linea, ":=") {
		return nil
	}

	// Si es declaración, que lo maneje parser_variables.go
	if strings.HasPrefix(linea, "variable ") {
		return nil
	}

	linea = strings.TrimPrefix(linea, "asignar ")
	partes := strings.SplitN(linea, ":=", 2)
	if len(partes) != 2 {
		return nil
	}

	izq := strings.TrimSpace(partes[0])
	der := strings.TrimSpace(partes[1])

	// Limpieza: si viene "bit a,b", mandamos solo "a,b"
	campos := strings.Fields(izq)
	nombreLimpio := campos[len(campos)-1]

	return &Nodo{
		Tipo:   "asignar",
		Nombre: nombreLimpio,
		Valor:  parseValor(der),
	}
}
