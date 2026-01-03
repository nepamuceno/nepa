package evaluador

import (
    "errors"
    "go/ast"
    "go/token"
)

// evalUnario maneja expresiones unarias:
// - +x → devuelve el valor numérico positivo
// - -x → devuelve el valor numérico negado
// - !x → devuelve la negación lógica booleana
func evalUnario(n *ast.UnaryExpr) (interface{}, error) {
    valor, err := evalNode(n.X)
    if err != nil {
        return nil, err
    }

    switch n.Op {
    case token.ADD: // +x
        return ConvertirAReal(valor)

    case token.SUB: // -x
        f, err := ConvertirAReal(valor)
        if err != nil {
            return nil, err
        }
        return -f, nil

    case token.NOT: // !x
        b, ok := valor.(bool)
        if !ok {
            return nil, errors.New("❌ ERROR FATAL: operador lógico '!' requiere un valor booleano")
        }
        return !b, nil

    default:
        return nil, errors.New("❌ ERROR FATAL: expresión unaria inválida")
    }
}
