package evaluador

import (
	"fmt"
)

// ResolverExpresion es el punto de entrada que conecta el texto plano con
// el motor de evaluación basado en AST.
func ResolverExpresion(entrada string, ctx *Contexto) (interface{}, error) {
	if entrada == "" {
		return nil, fmt.Errorf("la expresión está vacía")
	}

	// CAMBIO AQUÍ: Llamamos a EvalConContexto para poder usar el ctx
	// y resolver variables como 'a', 'x', 'y' en la expresión de la presa.
	resultado, err := EvalConContexto(entrada, ctx)
	if err != nil {
		return nil, fmt.Errorf("error al resolver expresión: %v", err)
	}

	return resultado, nil
}
