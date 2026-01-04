package evaluador

import (
	"fmt"
	"strconv"
)

// FormatearValor convierte cualquier valor interno de Nepa a una cadena legible.
// Se asegura de que los tipos se muestren de forma amigable para el usuario.
func FormatearValor(v interface{}) string {
	switch x := v.(type) {
	case bool:
		if x {
			return "verdadero"
		}
		return "falso"

	case float64:
		// Formateamos para que si es un entero (ej: 5.0) no muestre el .0
		// Pero si tiene decimales, los muestre todos sin notación científica extraña.
		return strconv.FormatFloat(x, 'f', -1, 64)

	case string:
		return x

	case nil:
		return "nulo"

	case []interface{}:
		// Formateo para listas: [1, 2, 3]
		res := "["
		for i, elem := range x {
			res += FormatearValor(elem)
			if i < len(x)-1 {
				res += ", "
			}
		}
		res += "]"
		return res

	default:
		// Para tipos inyectados (como bit) o estructuras complejas
		return fmt.Sprintf("%v", x)
	}
}
