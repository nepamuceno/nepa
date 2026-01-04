package evaluador

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
)

// evaluarUnario maneja expresiones de un solo operando como -a, +5 o !booleano.
func evaluarUnario(n *ast.UnaryExpr, ctx *Contexto) (interface{}, error) {
	// Evaluamos lo que está a la derecha del operador (X)
	valor, err := evaluarNodo(n.X, ctx)
	if err != nil {
		return nil, err
	}

	switch n.Op {
	case token.ADD: // Caso: +x
		return ConvertirAReal(valor)

	case token.SUB: // Caso: -x
		f, err := ConvertirAReal(valor)
		if err != nil {
			return nil, err
		}
		return -f, nil

	case token.NOT: // Caso: !x (Negación lógica)
		b, ok := valor.(bool)
		if !ok {
			return nil, errors.New("❌ ERROR FATAL: el operador lógico '!' requiere un valor booleano")
		}
		return !b, nil

	default:
		return nil, fmt.Errorf("❌ ERROR FATAL: operador unario '%v' no soportado", n.Op)
	}
}
