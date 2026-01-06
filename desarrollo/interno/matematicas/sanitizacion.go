package matematicas

import (
	"fmt"
	"math"
	"nepa/desarrollo/interno/evaluador"
)

// --- CAPA DE SEGURIDAD Y CONTROL DE ERRORES (CORE) ---

func finalizar(nombre string, resultado float64) (interface{}, error) {
	if math.IsNaN(resultado) {
		return nil, fmt.Errorf("❌ ERROR MATEMÁTICO en '%s': resultado indefinido (NaN)", nombre)
	}
	if math.IsInf(resultado, 0) {
		return nil, fmt.Errorf("❌ ERROR MATEMÁTICO en '%s': el resultado es infinito", nombre)
	}
	return resultado, nil
}

// ... validar1, validar2, validar3, validar4 se mantienen igual ...

func validar1(nombre string, args []interface{}) (float64, error) {
	if len(args) != 1 {
		return 0, fmt.Errorf("❌ ERROR: '%s' requiere 1 argumento, recibiste %d", nombre, len(args))
	}
	return evaluador.ConvertirAReal(args[0])
}

func validar2(nombre string, args []interface{}) (float64, float64, error) {
	if len(args) != 2 {
		return 0, 0, fmt.Errorf("❌ ERROR: '%s' requiere 2 argumentos", nombre)
	}
	v1, err1 := evaluador.ConvertirAReal(args[0])
	v2, err2 := evaluador.ConvertirAReal(args[1])
	if err1 != nil || err2 != nil { return 0, 0, fmt.Errorf("❌ ERROR: argumentos de '%s' deben ser números", nombre) }
	return v1, v2, nil
}

func validar3(nombre string, args []interface{}) (float64, float64, float64, error) {
	if len(args) != 3 {
		return 0, 0, 0, fmt.Errorf("❌ ERROR: '%s' requiere 3 argumentos", nombre)
	}
	v1, err1 := evaluador.ConvertirAReal(args[0])
	v2, err2 := evaluador.ConvertirAReal(args[1])
	v3, err3 := evaluador.ConvertirAReal(args[2])
	if err1 != nil || err2 != nil || err3 != nil {
		return 0, 0, 0, fmt.Errorf("❌ ERROR: los 3 argumentos de '%s' deben ser números", nombre)
	}
	return v1, v2, v3, nil
}

func validar4(nombre string, args []interface{}) (float64, float64, float64, float64, error) {
	if len(args) != 4 {
		return 0, 0, 0, 0, fmt.Errorf("❌ ERROR: '%s' requiere 4 argumentos", nombre)
	}
	v1, _ := evaluador.ConvertirAReal(args[0])
	v2, _ := evaluador.ConvertirAReal(args[1])
	v3, _ := evaluador.ConvertirAReal(args[2])
	v4, _ := evaluador.ConvertirAReal(args[3])
	return v1, v2, v3, v4, nil
}

// validarN ACTUALIZADO: Ahora soporta matrices y listas
func validarN(nombre string, args []interface{}) ([]float64, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("❌ ERROR: '%s' requiere al menos 1 valor", nombre)
	}

	// Si el primer argumento es una matriz o lista, usamos ConvertirAListaReal
	// Esto permite hacer promedio(matriz)
	if len(args) == 1 {
		nums, err := evaluador.ConvertirAListaReal(args[0])
		if err == nil {
			return nums, nil
		}
	}

	// Si son varios argumentos (ej: promedio(1, 2, 3)), procesamos la lista
	nums := make([]float64, len(args))
	for i, arg := range args {
		v, err := evaluador.ConvertirAReal(arg)
		if err != nil { 
			return nil, fmt.Errorf("❌ ERROR en '%s' (posicion %d): %v", nombre, i+1, err) 
		}
		nums[i] = v
	}
	return nums, nil
}
