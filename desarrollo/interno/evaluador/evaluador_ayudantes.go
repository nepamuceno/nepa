package evaluador

import (
	"errors"
	"strconv"
)

// Eliminamos la línea de ErrTipoNoSoportado porque ya existe en evaluador.go

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
			return 0, errors.New("❌ ERROR FATAL: no se pudo convertir la cadena a número → " + x)
		}
		return f, nil
	case bool:
		if x {
			return 1.0, nil
		}
		return 0.0, nil
	default:
		// Aquí usamos la variable que ya está declarada en evaluador.go
		return 0, ErrTipoNoSoportado 
	}
}

// ConvertirAListaReal transforma una lista genérica en un slice de float64.
func ConvertirAListaReal(entrada interface{}) ([]float64, error) {
	if lista, ok := entrada.([]float64); ok {
		return lista, nil
	}

	if reflejo, ok := entrada.([]interface{}); ok {
		res := make([]float64, len(reflejo))
		for i, v := range reflejo {
			num, err := ConvertirAReal(v)
			if err != nil {
				return nil, errors.New("❌ ERROR: elemento en posición " + strconv.Itoa(i) + " no es numérico")
			}
			res[i] = num
		}
		return res, nil
	}

	return nil, errors.New("❌ ERROR: se esperaba una lista o arreglo de números")
}
