package matematicas

import (
	"fmt"
	"math"
	"nepa/desarrollo/interno/evaluador"
	"sort"
)

func inyectarEstadisticaGlobal() {

	// --- 1. MOMENTOS Y FORMA DE LA DISTRIBUCIÓN ---

	evaluador.Funciones["sesgo"] = func(args ...interface{}) (interface{}, error) {
		nums, err := validarN("sesgo", args)
		if err != nil { return nil, err }
		n := float64(len(nums))
		if n < 3 { return nil, fmt.Errorf("❌ ERROR: El sesgo requiere al menos 3 datos") }

		media, _ := calcularMedia(nums)
		var m3, m2 float64
		for _, v := range nums {
			m3 += math.Pow(v-media, 3)
			m2 += math.Pow(v-media, 2)
		}
		std := math.Sqrt(m2 / n)
		if std == 0 { return 0.0, nil }
		
		coef := (n / ((n - 1) * (n - 2)))
		return finalizar("sesgo", coef * (m3 / math.Pow(std, 3)))
	}

	evaluador.Funciones["curtosis"] = func(args ...interface{}) (interface{}, error) {
		nums, err := validarN("curtosis", args)
		if err != nil { return nil, err }
		n := float64(len(nums))
		if n < 4 { return nil, fmt.Errorf("❌ ERROR: La curtosis requiere al menos 4 datos") }

		media, _ := calcularMedia(nums)
		var m4, m2 float64
		for _, v := range nums {
			m4 += math.Pow(v-media, 4)
			m2 += math.Pow(v-media, 2)
		}
		varianza := m2 / n
		if varianza == 0 { return 0.0, nil }
		
		return finalizar("curtosis", (m4 / (n * varianza * varianza)) - 3)
	}

	// --- 2. RELACIÓN ENTRE VARIABLES (Bi-variada) ---

	evaluador.Funciones["correlacion_pearson"] = func(args ...interface{}) (interface{}, error) {
		if len(args) != 2 { return nil, fmt.Errorf("requiere dos matrices/listas (X e Y)") }
		
		x, errX := evaluador.ConvertirAListaReal(args[0])
		y, errY := evaluador.ConvertirAListaReal(args[1])
		if errX != nil || errY != nil { return nil, fmt.Errorf("error al procesar listas de entrada") }
		
		if len(x) != len(y) || len(x) == 0 { return nil, fmt.Errorf("listas deben tener igual tamaño") }
		
		mx, _ := calcularMedia(x)
		my, _ := calcularMedia(y)
		
		var num, denX, denY float64
		for i := 0; i < len(x); i++ {
			dx := x[i] - mx
			dy := y[i] - my
			num += dx * dy
			denX += dx * dx
			denY += dy * dy
		}
		res := num / math.Sqrt(denX*denY)
		return finalizar("correlacion_pearson", res)
	}

	evaluador.Funciones["regresion_lineal"] = func(args ...interface{}) (interface{}, error) {
		if len(args) != 2 { return nil, fmt.Errorf("requiere dos matrices/listas (X e Y)") }
		x, _ := evaluador.ConvertirAListaReal(args[0])
		y, _ := evaluador.ConvertirAListaReal(args[1])
		n := float64(len(x))
		if n == 0 { return nil, fmt.Errorf("las listas no pueden estar vacías") }
		
		var sumX, sumY, sumXY, sumX2 float64
		for i := 0; i < len(x); i++ {
			sumX += x[i]
			sumY += y[i]
			sumXY += x[i] * y[i]
			sumX2 += x[i] * x[i]
		}
		
		divisor := (n*sumX2 - sumX*sumX)
		if divisor == 0 { return nil, fmt.Errorf("no se puede calcular regresión (división por cero)") }
		
		pendiente := (n*sumXY - sumX*sumY) / divisor
		intercepto := (sumY - pendiente*sumX) / n
		
		return []float64{pendiente, intercepto}, nil
	}

	// --- 3. MEDIDAS DE POSICIÓN ---

	evaluador.Funciones["percentil"] = func(args ...interface{}) (interface{}, error) {
		if len(args) < 2 { return nil, fmt.Errorf("requiere matriz/lista y valor p (0-100)") }
		nums, err := evaluador.ConvertirAListaReal(args[0])
		if err != nil { return nil, err }
		p, _ := evaluador.ConvertirAReal(args[1])
		
		sort.Float64s(nums)
		idx := (p / 100) * float64(len(nums)-1)
		i := int(idx)
		frac := idx - float64(i)
		
		var res float64
		if i+1 < len(nums) {
			res = nums[i] + frac*(nums[i+1]-nums[i])
		} else {
			res = nums[i]
		}
		return finalizar("percentil", res)
	}

	// --- 4. LAS BÁSICAS MEJORADAS ---

	evaluador.Funciones["promedio"] = func(args ...interface{}) (interface{}, error) {
		nums, err := validarN("promedio", args)
		if err != nil { return nil, err }
		m, _ := calcularMedia(nums)
		return finalizar("promedio", m)
	}

	evaluador.Funciones["varianza_poblacional"] = func(args ...interface{}) (interface{}, error) {
		nums, err := validarN("varianza", args)
		if err != nil { return nil, err }
		media, _ := calcularMedia(nums)
		var sumaCuadrados float64
		for _, n := range nums { sumaCuadrados += math.Pow(n-media, 2) }
		return finalizar("varianza_poblacional", sumaCuadrados / float64(len(nums)))
	}
}

// Ayudante interno para no repetir código
func calcularMedia(nums []float64) (float64, error) {
	if len(nums) == 0 { return 0, nil }
	var s float64
	for _, v := range nums { s += v }
	return s / float64(len(nums)), nil
}
