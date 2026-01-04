package evaluador

import "fmt"

// Contexto representa el entorno de ejecución del intérprete.
type Contexto struct {
	Variables  map[string]interface{}                       // Variables locales
	Globales   map[string]interface{}                       // Variables globales
	Constantes map[string]interface{}                       // Constantes definidas
	Funciones  map[string]func(...interface{}) interface{}  // Funciones registradas
}

// ObtenerVariable busca un valor en el orden: Constantes -> Locales -> Globales
func (ctx *Contexto) ObtenerVariable(nombre string) (interface{}, error) {
	// 1. Buscar en Constantes
	if valor, existe := ctx.Constantes[nombre]; existe {
		return valor, nil
	}

	// 2. Buscar en Locales
	if valor, existe := ctx.Variables[nombre]; existe {
		return valor, nil
	}

	// 3. Buscar en Globales
	if valor, existe := ctx.Globales[nombre]; existe {
		return valor, nil
	}

	return nil, fmt.Errorf("la variable o constante '%s' no está definida", nombre)
}

// ObtenerFuncion busca una función registrada en el contexto
func (ctx *Contexto) ObtenerFuncion(nombre string) (func(...interface{}) interface{}, error) {
	if fn, existe := ctx.Funciones[nombre]; existe {
		return fn, nil
	}
	return nil, fmt.Errorf("la función '%s' no existe", nombre)
}
