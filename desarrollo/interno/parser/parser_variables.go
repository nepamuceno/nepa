package parser

import (
	"strings"
)

func parseVariable(linea string) *Nodo {
	if !strings.HasPrefix(linea, "variable ") {
		return nil
	}

	contenido := strings.TrimSpace(linea[len("variable "):])
	var tipo, nombres string
	var valor interface{}

	if strings.Contains(contenido, ":=") {
		// Caso: variable bit a,b := 1
		partes := strings.SplitN(contenido, ":=", 2)
		izq := strings.TrimSpace(partes[0])
		der := strings.TrimSpace(partes[1])
		
		valor = parseValor(der)

		campos := strings.Fields(izq)
		if len(campos) >= 2 {
			tipo = campos[0]
			nombres = strings.Join(campos[1:], "")
		}
	} else {
		// Caso: variable bit a,b,c,d (Valores seguros)
		campos := strings.Fields(contenido)
		if len(campos) >= 2 {
			tipo = campos[0]
			nombres = strings.Join(campos[1:], "")
			valor = nil // IMPORTANTE: Enviamos nil para que CrearBit use su default
		}
	}

	if tipo != "" && nombres != "" {
		return &Nodo{
			Tipo:   "variable",
			Nombre: nombres, // "a,b,c,d"
			Valor:  valor,   // El valor parseado o nil
			Args:   []interface{}{tipo},
		}
	}
	return nil
}
