package evaluador

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"strings"
)

// Errores originales para que los comandos externos no fallen
var (
	ErrExpresionInvalida = errors.New("expresión inválida")
	ErrFuncionNoExiste   = errors.New("función no registrada")
	ErrTipoNoSoportado   = errors.New("tipo no soportado en evaluación")
	ErrConcatenacion     = errors.New("error de tipo en concatenación")
)

// Eval es la función que buscan 'asignar', 'imprimir', etc.
// Mantenemos EXACTAMENTE 1 argumento para que compile el resto del proyecto.
func Eval(expr string) (interface{}, error) {
	// Creamos un contexto vacío al vuelo. 
	// Así, el resto del sistema no tiene que saber que ahora usamos contextos.
	ctx := &Contexto{
		Variables: make(map[string]interface{}),
		Funciones: make(map[string]func(...interface{}) interface{}),
	}
	
	return EvalConContexto(expr, ctx)
}

// EvalConContexto es la versión que usaremos internamente o en el futuro 
// cuando queramos pasar variables de forma explícita.
func EvalConContexto(expr string, ctx *Contexto) (interface{}, error) {
	expr = strings.TrimSpace(expr)
	if expr == "" {
		return nil, ErrExpresionInvalida
	}

	node, err := parser.ParseExpr(expr)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrExpresionInvalida, err)
	}

	return evaluarNodo(node, ctx)
}

// evaluarNodo es el despachador interno que ya conoce el contexto.
func evaluarNodo(node ast.Expr, ctx *Contexto) (interface{}, error) {
	switch n := node.(type) {
	case *ast.BasicLit:
		return evaluarLiteral(n)
	case *ast.Ident:
		return evaluarIdentificador(n, ctx)
	case *ast.UnaryExpr:
		return evaluarUnario(n, ctx)
	case *ast.BinaryExpr:
		return evaluarBinario(n, ctx)
	case *ast.CallExpr:
		return evaluarLlamada(n, ctx)
	case *ast.ParenExpr:
		// Soporte crítico para la expresión de la presa: ( ... )
		return evaluarNodo(n.X, ctx)
	default:
		return nil, ErrExpresionInvalida
	}
}
