package evaluador

import (
    "errors"
    "fmt"
    "go/ast"
    "go/parser"
    "strings"
)

var (
    ErrExpresionInvalida = errors.New("expresión inválida")
    ErrFuncionNoExiste   = errors.New("función no registrada")
    ErrTipoNoSoportado   = errors.New("tipo no soportado en evaluación")
    ErrConcatenacion     = errors.New("error de tipo en concatenación")
)

// Eval evalúa una expresión textual y devuelve un resultado como interface{}.
func Eval(expr string) (interface{}, error) {
    expr = strings.TrimSpace(expr)
    if expr == "" {
        return nil, ErrExpresionInvalida
    }
    node, err := parser.ParseExpr(expr)
    if err != nil {
        return nil, fmt.Errorf("%w: %v", ErrExpresionInvalida, err)
    }
    return evalNode(node)
}

// Dispatcher central: delega según el tipo de nodo AST
func evalNode(node ast.Expr) (interface{}, error) {
    switch n := node.(type) {
    case *ast.BasicLit:
        return evalLiteral(n)
    case *ast.Ident:
        return evalIdent(n)
    case *ast.UnaryExpr:
        return evalUnario(n)   // ✅ nombre en español
    case *ast.BinaryExpr:
        return evalBinario(n)  // ✅ nombre en español
    case *ast.CallExpr:
        return evalLlamada(n)  // ✅ nombre en español
    default:
        return nil, ErrExpresionInvalida
    }
}
