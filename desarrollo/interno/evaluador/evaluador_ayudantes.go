package evaluador

import (
	"fmt"
	"strconv"
)

// ConvertirAReal convierte cualquier valor a float64 para cálculos universales.
func ConvertirAReal(v interface{}) (float64, error) {
	switch x := v.(type) {
	case int:
		return float64(x), nil
	case int32:
		return float64(x), nil
	case int64:
		return float64(x), nil
	case uint8:
		return float64(x), nil
	case uint:
		return float64(x), nil
	case uint32:
		return float64(x), nil
	case uint64:
		return float64(x), nil
	case float32:
		return float64(x), nil
	case float64:
		return x, nil
	case string:
		f, err := strconv.ParseFloat(x, 64)
		if err != nil {
			return 0, fmt.Errorf("❌ ERROR FATAL: no se pudo convertir la cadena a número → %s", x)
		}
		return f, nil
	case bool:
		if x {
			return 1.0, nil
		}
		return 0.0, nil
	default:
		return 0, fmt.Errorf("tipo no soportado")
	}
}

// ConvertirAListaReal adaptado: resuelve y valida cada elemento.
func ConvertirAListaReal(entrada interface{}) ([]float64, error) {
	// 1. Resolvemos variables/funciones primero (Contexto global nil)
	resuelto := ResolverEstructuraRecursiva(entrada, nil)

	if reflejo, ok := resuelto.([]interface{}); ok {
		res := make([]float64, len(reflejo))
		for i, v := range reflejo {
			num, err := ConvertirAReal(v)
			if err != nil {
				// Aquí capturamos el fallo (ej. potencia) y cortamos la ejecución
				return nil, fmt.Errorf("❌ ERROR: el elemento '%v' en la posición %d no es válido", v, i)
			}
			res[i] = num
		}
		return res, nil
	}

	// Caso de un solo valor que no viene en lista
	num, err := ConvertirAReal(resuelto)
	if err != nil {
		return nil, fmt.Errorf("❌ ERROR: se esperaba un valor numérico")
	}
	return []float64{num}, nil
}

// ResolverEstructuraRecursiva recorre la estructura y resuelve variables/funciones.
func ResolverEstructuraRecursiva(v interface{}, ctx *Contexto) interface{} {
	if ctx == nil {
		ctx = &Contexto{
			Variables: make(map[string]interface{}),
			Funciones: make(map[string]func(...interface{}) interface{}),
		}
	}

	switch x := v.(type) {
	case string:
		res, err := EvalConContexto(x, ctx)
		if err == nil {
			return res
		}
		return x
	case []interface{}:
		res := make([]interface{}, len(x))
		for i, elem := range x {
			res[i] = ResolverEstructuraRecursiva(elem, ctx)
		}
		return res
	default:
		return v
	}
}
