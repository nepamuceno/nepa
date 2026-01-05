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
		// Caso: variable matriz m := [[1,2],[3,4]]
		partes := strings.SplitN(contenido, ":=", 2)
		izq := strings.TrimSpace(partes[0])
		der := strings.TrimSpace(partes[1])

		// Extraemos tipo y nombre antes de procesar el valor
		campos := strings.Fields(izq)
		if len(campos) >= 2 {
			tipo = campos[0]
			nombres = strings.Join(campos[1:], "")
		}

		// 🛡️ PROTECCIÓN DE ESTRUCTURAS COMPLEJAS
		// Si es matriz, enviamos la cadena 'der' completa al nodo.
		// Esto evita que parseValor rompa funciones como potencia(2, 3) por sus espacios.
		if tipo == "matriz" {
			valor = der
		} else {
			valor = parseValor(der)
		}

	} else {
		// Caso: variable entero a,b,c (Sin asignación inmediata)
		campos := strings.Fields(contenido)
		if len(campos) >= 2 {
			tipo = campos[0]
			nombres = strings.Join(campos[1:], "")
			valor = nil // IMPORTANTE: Enviamos nil para que Crear use su default
		}
	}

	if tipo != "" && nombres != "" {
		return &Nodo{
			Tipo:   "variable",
			Nombre: nombres, // "m_funciones" o "a,b,c"
			Valor:  valor,   // La cadena original para matrices o el valor parseado
			Args:   []interface{}{tipo},
		}
	}
	return nil
}
