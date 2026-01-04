package evaluador

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
)

// evaluarBinario maneja operaciones entre dos valores (+, -, *, /, %, etc.).
func evaluarBinario(n *ast.BinaryExpr, ctx *Contexto) (interface{}, error) {
	// Evaluamos el lado izquierdo y derecho recursivamente
	// Al evaluar el nodo, si es un identificador (base, ajuste), 
	// evaluador_ident.go ya nos traerá su valor numérico real.
	izquierda, err := evaluarNodo(n.X, ctx)
	if err != nil {
		return nil, err
	}
	derecha, err := evaluarNodo(n.Y, ctx)
	if err != nil {
		return nil, err
	}
	return aplicarOperacion(n.Op, izquierda, derecha)
}

// aplicarOperacion ejecuta la lógica matemática o lógica según el operador.
func aplicarOperacion(op token.Token, izquierda, derecha interface{}) (interface{}, error) {
	switch op {
	case token.ADD:
		// MEJORA DE INTEROPERABILIDAD:
		// Si cualquiera de los dos lados es una cadena de texto, 
		// realizamos una concatenación en lugar de suma numérica.
		_, esIzqString := izquierda.(string)
		_, esDerString := derecha.(string)

		if esIzqString || esDerString {
			return fmt.Sprint(izquierda) + fmt.Sprint(derecha), nil
		}

		// Si no hay strings, procedemos a la suma numérica universal
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

	// Comparaciones universales
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

	// Operadores lógicos (requieren booleanos)
	case token.LAND:
		li, lok := izquierda.(bool)
		ri, rok := derecha.(bool)
		if !lok || !rok {
			return nil, errors.New("❌ ERROR FATAL: el operador 'Y' (&&) requiere valores booleanos")
		}
		return li && ri, nil
	case token.LOR:
		li, lok := izquierda.(bool)
		ri, rok := derecha.(bool)
		if !lok || !rok {
			return nil, errors.New("❌ ERROR FATAL: el operador 'O' (||) requiere valores booleanos")
		}
		return li || ri, nil

	default:
		return nil, fmt.Errorf("❌ ERROR FATAL: operador binario no soportado: %v", op)
	}
}

// Auxiliares para cálculo numérico
func operarNumeros(izq, der interface{}, operacion func(a, b float64) float64) (float64, error) {
	valIzq, err := ConvertirAReal(izq)
	if err != nil {
		return 0, err
	}
	valDer, err := ConvertirAReal(der)
	if err != nil {
		return 0, err
	}
	return operacion(valIzq, valDer), nil
}

func operarNumerosConValidacion(izq, der interface{}, operacion func(a, b float64) (float64, error)) (float64, error) {
	valIzq, err := ConvertirAReal(izq)
	if err != nil {
		return 0, err
	}
	valDer, err := ConvertirAReal(der)
	if err != nil {
		return 0, err
	}
	return operacion(valIzq, valDer)
}

func compararNumeros(izq, der interface{}, comparacion func(a, b float64) bool) (bool, error) {
	valIzq, err := ConvertirAReal(izq)
	if err != nil {
		return false, err
	}
	valDer, err := ConvertirAReal(der)
	if err != nil {
		return false, err
	}
	return comparacion(valIzq, valDer), nil
}
