package evaluador

import (
    "errors"
    "strconv"
)

// ConvertirAReal convierte cualquier valor a float64 para cálculos universales.
// Se usa en todo el intérprete para evitar duplicación de lógica (sustituye a matematicas/toFloat).
func ConvertirAReal(v interface{}) (float64, error) {
    switch x := v.(type) {
    case int:
        return float64(x), nil
    case int32:
        return float64(x), nil
    case int64:
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
            return 1, nil
        }
        return 0, nil
    default:
        return 0, ErrTipoNoSoportado
    }
}
