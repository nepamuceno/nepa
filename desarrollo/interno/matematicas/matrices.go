package matematicas

import (
	"fmt"
	"nepa/desarrollo/interno/evaluador"
)

func inyectarMatricesGlobal() {

	// --- 1. DETERMINANTES ---

	// det2x2(a, b, c, d) o det2x2([[a,b],[c,d]])
	evaluador.Funciones["det2x2"] = func(args ...interface{}) (interface{}, error) {
		v, err := obtener4Valores(args)
		if err != nil { return nil, err }
		// | v0 v1 |  => (v0*v3) - (v1*v2)
		// | v2 v3 |
		return (v[0] * v[3]) - (v[1] * v[2]), nil
	}

	// det3x3 (Regla de Sarrus)
	// | a b c |
	// | d e f |
	// | g h i |
	evaluador.Funciones["det3x3"] = func(args ...interface{}) (interface{}, error) {
		v, err := obtener9Valores(args)
		if err != nil { return nil, err }
		
		pos := (v[0]*v[4]*v[8]) + (v[1]*v[5]*v[6]) + (v[2]*v[3]*v[7])
		neg := (v[2]*v[4]*v[6]) + (v[0]*v[5]*v[7]) + (v[1]*v[3]*v[8])
		
		return pos - neg, nil
	}

	// --- 2. OPERACIONES DINÁMICAS (N x M) ---

	// matriz_sumar(M1, M2)
	evaluador.Funciones["matriz_sumar"] = func(args ...interface{}) (interface{}, error) {
		m1, m2, err := validarDosMatrices(args)
		if err != nil { return nil, err }

		if len(m1) != len(m2) || len(m1[0]) != len(m2[0]) {
			return nil, fmt.Errorf("❌ ERROR: Las matrices deben tener las mismas dimensiones")
		}

		filas := len(m1)
		cols := len(m1[0])
		res := crearMatrizVacia(filas, cols)

		for i := 0; i < filas; i++ {
			for j := 0; j < cols; j++ {
				res[i][j] = m1[i][j] + m2[i][j]
			}
		}
		return res, nil
	}

	// matriz_multiplicar(M1, M2) -> El corazón de la computación
	evaluador.Funciones["matriz_multiplicar"] = func(args ...interface{}) (interface{}, error) {
		m1, m2, err := validarDosMatrices(args)
		if err != nil { return nil, err }

		// Validar: Columnas M1 == Filas M2
		if len(m1[0]) != len(m2) {
			return nil, fmt.Errorf("❌ ERROR: Columnas de M1 (%d) no coinciden con Filas de M2 (%d)", len(m1[0]), len(m2))
		}

		res := crearMatrizVacia(len(m1), len(m2[0]))

		for i := 0; i < len(m1); i++ {
			for j := 0; j < len(m2[0]); j++ {
				for k := 0; k < len(m2); k++ {
					res[i][j] += m1[i][k] * m2[k][j]
				}
			}
		}
		return res, nil
	}

	// matriz_transponer(M) -> Cambia filas por columnas
	evaluador.Funciones["matriz_transponer"] = func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 { return nil, fmt.Errorf("requiere una matriz") }
		m := args[0].([][]float64)
		
		filas := len(m)
		cols := len(m[0])
		res := crearMatrizVacia(cols, filas)

		for i := 0; i < filas; i++ {
			for j := 0; j < cols; j++ {
				res[j][i] = m[i][j]
			}
		}
		return res, nil
	}
}

// --- UTILERÍA PARA MATRICES ---

func crearMatrizVacia(f, c int) [][]float64 {
	m := make([][]float64, f)
	for i := range m { m[i] = make([]float64, c) }
	return m
}

func obtener4Valores(args []interface{}) ([]float64, error) {
	// Aquí podrías implementar lógica para aceptar tanto 4 args como una lista de 4
	res := make([]float64, 4)
	for i := 0; i < 4; i++ {
		v, _ := evaluador.ConvertirAReal(args[i])
		res[i] = v
	}
	return res, nil
}

func obtener9Valores(args []interface{}) ([]float64, error) {
	res := make([]float64, 9)
	if len(args) != 9 { return nil, fmt.Errorf("det3x3 requiere 9 valores") }
	for i := 0; i < 9; i++ {
		v, _ := evaluador.ConvertirAReal(args[i])
		res[i] = v
	}
	return res, nil
}

func validarDosMatrices(args []interface{}) ([][]float64, [][]float64, error) {
	if len(args) != 2 { return nil, nil, fmt.Errorf("se requieren 2 matrices") }
	m1, ok1 := args[0].([][]float64)
	m2, ok2 := args[1].([][]float64)
	if !ok1 || !ok2 { return nil, nil, fmt.Errorf("argumentos deben ser matrices [][]float64") }
	return m1, m2, nil
}
