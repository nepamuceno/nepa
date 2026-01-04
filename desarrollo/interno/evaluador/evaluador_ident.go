package evaluador

import (
	"errors"
	"fmt"
	"go/ast"
	"strings"
)

// Error específico para identificadores inexistentes
var ErrIdentificadorNoExiste = errors.New("❌ ERROR FATAL: el identificador no existe")

// evaluarIdentificador maneja:
// - verdadero / falso → booleanos
// - nombres de variables buscando en el Contexto (Locales, Globales o Constantes)
func evaluarIdentificador(n *ast.Ident, ctx *Contexto) (interface{}, error) {
	nombre := strings.ToLower(strings.TrimSpace(n.Name))

	switch nombre {
	case "verdadero":
		return true, nil
	case "falso":
		return false, nil
	default:
		// Primero intentamos buscar en el contexto que trae la ejecución
		v, err := ctx.ObtenerVariable(nombre)
		if err != nil {
			return nil, fmt.Errorf("%w → %s", ErrIdentificadorNoExiste, nombre)
		}

		// Si el valor es una estructura de Nepa (como Bit), extraemos su valor básico
		// Aquí asumimos que lo que devuelve ObtenerVariable puede ser un objeto con ValorComoInterface()
		// o el valor directamente.
		if interfaz, ok := v.(interface{ ValorComoInterface() interface{} }); ok {
			return interfaz.ValorComoInterface(), nil
		}

		return v, nil
	}
}
