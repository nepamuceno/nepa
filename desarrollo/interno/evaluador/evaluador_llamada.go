package evaluador

import (
	"fmt"
	"go/ast"
	"strings"
)

// evaluarLlamada maneja llamadas a funciones (ej: seno(x)) y métodos (ej: lista.limpiar()).
func evaluarLlamada(n *ast.CallExpr, ctx *Contexto) (interface{}, error) {
	switch fn := n.Fun.(type) {
	case *ast.Ident:
		nombreFuncion := strings.ToLower(fn.Name)
		
		argumentos, err := evaluarArgumentos(n.Args, ctx)
		if err != nil {
			return nil, err
		}

		f, ok := Funciones[nombreFuncion]
		if !ok {
			return nil, fmt.Errorf("%w → %s", ErrFuncionNoExiste, nombreFuncion)
		}
		
		return f(argumentos...)

	case *ast.SelectorExpr:
		objeto, err := evaluarNodo(fn.X, ctx)
		if err != nil {
			return nil, err
		}
		
		nombreMetodo := strings.ToLower(fn.Sel.Name)
		argumentos, err := evaluarArgumentos(n.Args, ctx)
		if err != nil {
			return nil, err
		}

		tipo := obtenerTipoEnEspañol(objeto)
		clave := fmt.Sprintf("%s.%s", tipo, nombreMetodo)
		
		f, ok := Funciones[clave]
		if !ok {
			return nil, fmt.Errorf("%w → método %s no existe para tipo %s", ErrFuncionNoExiste, nombreMetodo, tipo)
		}
		
		return f(append([]interface{}{objeto}, argumentos...)...)

	default:
		return nil, ErrFuncionNoExiste
	}
}

// evaluarArgumentos evalúa cada expresión pasada como parámetro.
func evaluarArgumentos(args []ast.Expr, ctx *Contexto) ([]interface{}, error) {
	var valores []interface{}
	for _, a := range args {
		valor, err := evaluarNodo(a, ctx)
		if err != nil {
			return nil, err
		}
		valores = append(valores, valor)
	}
	return valores, nil
}

// obtenerTipoEnEspañol traduce los tipos internos para la resolución de métodos.
func obtenerTipoEnEspañol(v interface{}) string {
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
