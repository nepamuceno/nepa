package evaluador

import (
    "fmt"
    "go/ast"
    "strings"
)

// evalLlamada maneja llamadas a funciones internas/externas y métodos.
func evalLlamada(n *ast.CallExpr) (interface{}, error) {
    switch fn := n.Fun.(type) {
    case *ast.Ident:
        nombreFuncion := strings.ToLower(fn.Name)
        argumentos, err := evalArgs(n.Args)
        if err != nil {
            return nil, err
        }
        f, ok := Funciones[nombreFuncion]
        if !ok {
            return nil, fmt.Errorf("%w → %s", ErrFuncionNoExiste, nombreFuncion)
        }
        return f(argumentos...)

    case *ast.SelectorExpr:
        objeto, err := evalNode(fn.X)
        if err != nil {
            return nil, err
        }
        nombreMetodo := strings.ToLower(fn.Sel.Name)
        argumentos, err := evalArgs(n.Args)
        if err != nil {
            return nil, err
        }
        tipo := tipoEnEspañol(objeto)
        clave := fmt.Sprintf("%s.%s", tipo, nombreMetodo)
        if f, ok := Funciones[clave]; ok {
            return f(append([]interface{}{objeto}, argumentos...)...)
        }
        return nil, fmt.Errorf("%w → método %s no existe para tipo %s", ErrFuncionNoExiste, nombreMetodo, tipo)

    default:
        return nil, ErrFuncionNoExiste
    }
}

// evalArgs evalúa todos los argumentos de una llamada.
func evalArgs(args []ast.Expr) ([]interface{}, error) {
    var valores []interface{}
    for _, a := range args {
        valor, err := evalNode(a)
        if err != nil {
            return nil, err
        }
        valores = append(valores, valor)
    }
    return valores, nil
}

// tipoEnEspañol traduce tipos básicos a nombres en español.
func tipoEnEspañol(v interface{}) string {
    switch v.(type) {
    case string:
        return "cadena"
    case int, int64:
        return "entero"
    case float64:
        return "real"
    case bool:
        return "booleano"
    case rune:
        return "caracter"
    case []interface{}:
        return "lista"
    default:
        return "objeto"
    }
}
