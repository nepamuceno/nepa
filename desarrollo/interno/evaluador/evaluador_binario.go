package evaluador

import (
    "errors"
    "fmt"
    "go/ast"
    "go/token"
)

// evalBinario maneja operaciones binarias (+, -, *, /, %, concatenación, comparaciones, lógicas).
func evalBinario(n *ast.BinaryExpr) (interface{}, error) {
    izquierda, err := evalNode(n.X)
    if err != nil {
        return nil, err
    }
    derecha, err := evalNode(n.Y)
    if err != nil {
        return nil, err
    }
    return aplicarOperacion(n.Op, izquierda, derecha)
}

// aplicarOperacion ejecuta la operación según el tipo de dato.
func aplicarOperacion(op token.Token, izquierda, derecha interface{}) (interface{}, error) {
    switch op {
    case token.ADD:
        // Concatenación si alguno es cadena
        if ls, ok := izquierda.(string); ok {
            return ls + fmt.Sprint(derecha), nil
        }
        if rs, ok := derecha.(string); ok {
            return fmt.Sprint(izquierda) + rs, nil
        }
        // Suma numérica
        return operarNumeros(izquierda, derecha, func(a, b float64) float64 { return a + b })

    case token.SUB:
        return operarNumeros(izquierda, derecha, func(a, b float64) float64 { return a - b })

    case token.MUL:
        return operarNumeros(izquierda, derecha, func(a, b float64) float64 { return a * b })

    case token.QUO:
        return operarNumerosConValidacion(izquierda, derecha, func(a, b float64) (float64, error) {
            if b == 0 {
                return 0, errors.New("❌ ERROR FATAL: división por cero")
            }
            return a / b, nil
        })

    case token.REM:
        li, err := ConvertirAReal(izquierda)
        if err != nil {
            return nil, err
        }
        ri, err := ConvertirAReal(derecha)
        if err != nil {
            return nil, err
        }
        if int(ri) == 0 {
            return nil, errors.New("❌ ERROR FATAL: módulo por cero")
        }
        return float64(int(li) % int(ri)), nil

    // Comparaciones
    case token.EQL:
        return izquierda == derecha, nil
    case token.NEQ:
        return izquierda != derecha, nil
    case token.LSS:
        return compararNumeros(izquierda, derecha, func(a, b float64) bool { return a < b })
    case token.GTR:
        return compararNumeros(izquierda, derecha, func(a, b float64) bool { return a > b })
    case token.LEQ:
        return compararNumeros(izquierda, derecha, func(a, b float64) bool { return a <= b })
    case token.GEQ:
        return compararNumeros(izquierda, derecha, func(a, b float64) bool { return a >= b })

    // Operadores lógicos
    case token.LAND:
        li, lok := izquierda.(bool)
        ri, rok := derecha.(bool)
        if !lok || !rok {
            return nil, errors.New("❌ ERROR FATAL: operador lógico requiere valores booleanos")
        }
        return li && ri, nil
    case token.LOR:
        li, lok := izquierda.(bool)
        ri, rok := derecha.(bool)
        if !lok || !rok {
            return nil, errors.New("❌ ERROR FATAL: operador lógico requiere valores booleanos")
        }
        return li || ri, nil

    default:
        return nil, errors.New("❌ ERROR FATAL: expresión binaria inválida")
    }
}

// Helpers para evitar duplicación
func operarNumeros(izquierda, derecha interface{}, op func(a, b float64) float64) (float64, error) {
    li, err := ConvertirAReal(izquierda)
    if err != nil {
        return 0, err
    }
    ri, err := ConvertirAReal(derecha)
    if err != nil {
        return 0, err
    }
    return op(li, ri), nil
}

func operarNumerosConValidacion(izquierda, derecha interface{}, op func(a, b float64) (float64, error)) (float64, error) {
    li, err := ConvertirAReal(izquierda)
    if err != nil {
        return 0, err
    }
    ri, err := ConvertirAReal(derecha)
    if err != nil {
        return 0, err
    }
    return op(li, ri)
}

func compararNumeros(izquierda, derecha interface{}, cmp func(a, b float64) bool) (bool, error) {
    li, err := ConvertirAReal(izquierda)
    if err != nil {
        return false, err
    }
    ri, err := ConvertirAReal(derecha)
    if err != nil {
        return false, err
    }
    return cmp(li, ri), nil
}
