package evaluador

import (
	"errors"
	"go/ast"
	"go/token"
	"strconv"
	"strings"
)

// evaluarLiteral maneja valores escritos directamente: 10, 3.14, "hola".
func evaluarLiteral(n *ast.BasicLit) (interface{}, error) {
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
		// Eliminamos las comillas externas
		crudo := strings.Trim(n.Value, `"`)
		return desescaparCadena(crudo), nil

	case token.CHAR:
		crudo := strings.TrimSpace(strings.Trim(n.Value, `'`))
		if len(crudo) != 1 {
			return nil, errors.New("❌ ERROR FATAL: carácter inválido → " + n.Value)
		}
		return rune(crudo[0]), nil

	default:
		return nil, errors.New(ErrTipoNoSoportado.Error() + " → " + n.Kind.String())
	}
}

// desescaparCadena procesa secuencias de escape comunes en cadenas.
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
