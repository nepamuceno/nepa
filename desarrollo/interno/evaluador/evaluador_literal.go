package evaluador

import (
    "errors"
    "go/ast"
    "go/token"
    "strconv"
    "strings"
)

// evalLiteral maneja literales básicos: enteros, reales, cadenas y caracteres.
// Se invoca desde el despachador en evaluador.go cuando el nodo es *ast.BasicLit.
func evalLiteral(n *ast.BasicLit) (interface{}, error) {
    switch n.Kind {
    case token.INT:
        valor, err := strconv.Atoi(n.Value)
        if err != nil {
            return nil, errors.New("❌ ERROR FATAL: entero inválido → " + n.Value)
        }
        return valor, nil

    case token.FLOAT:
        valor, err := strconv.ParseFloat(n.Value, 64)
        if err != nil {
            return nil, errors.New("❌ ERROR FATAL: número real inválido → " + n.Value)
        }
        return valor, nil

    case token.STRING:
        raw := strings.Trim(n.Value, `"`)
        return desescaparCadena(raw), nil

    case token.CHAR:
        raw := strings.TrimSpace(strings.Trim(n.Value, `'`))
        if len(raw) != 1 {
            return nil, errors.New("❌ ERROR FATAL: carácter inválido → " + n.Value)
        }
        return rune(raw[0]), nil

    default:
        // Aquí usamos ErrTipoNoSoportado definido en evaluador.go
        return nil, errors.New(ErrTipoNoSoportado.Error() + " → " + n.Kind.String())
    }
}

// desescaparCadena procesa secuencias de escape comunes en cadenas
func desescaparCadena(s string) string {
    reemplazos := map[string]string{
        `\n`: "\n",
        `\t`: "\t",
        `\"`: `"`,
        `\\`: `\`,
        `\r`: "\r",
        `\b`: "\b",
        `\f`: "\f",
    }
    resultado := s
    for k, v := range reemplazos {
        resultado = strings.ReplaceAll(resultado, k, v)
    }
    return resultado
}
