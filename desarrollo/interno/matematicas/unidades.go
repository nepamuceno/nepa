package matematicas

import (
	"fmt"
	"nepa/desarrollo/interno/evaluador"
)

func inyectarUnidadesGlobal() {

	// --- 1. LONGITUD Y DISTANCIA ---
	evaluador.Funciones["convertir_longitud"] = func(args ...interface{}) (interface{}, error) {
		v, u1, u2, err := extraerUnidades(args); if err != nil { return nil, err }
		factores := map[string]float64{
			"nm": 1e-9, "um": 1e-6, "mm": 0.001, "cm": 0.01, "m": 1, "km": 1000,
			"pulgada": 0.0254, "pie": 0.3048, "yarda": 0.9144, "milla": 1609.34, 
			"milla_nautica": 1852, "angstrom": 1e-10, "año_luz": 9.461e15, "parsec": 3.086e16,
		}
		return convertirGenerico(v, u1, u2, factores, "longitud")
	}

	// --- 2. ÁREA (Superficie) ---
	evaluador.Funciones["convertir_area"] = func(args ...interface{}) (interface{}, error) {
		v, u1, u2, err := extraerUnidades(args); if err != nil { return nil, err }
		factores := map[string]float64{
			"mm2": 1e-6, "cm2": 1e-4, "m2": 1, "km2": 1e6,
			"pulgada2": 0.00064516, "pie2": 0.092903, "acre": 4046.86, "hectarea": 10000,
		}
		return convertirGenerico(v, u1, u2, factores, "área")
	}

	// --- 3. VOLUMEN Y CAPACIDAD ---
	evaluador.Funciones["convertir_volumen"] = func(args ...interface{}) (interface{}, error) {
		v, u1, u2, err := extraerUnidades(args); if err != nil { return nil, err }
		factores := map[string]float64{
			"ml": 0.001, "litro": 1, "m3": 1000, "taza": 0.25, 
			"pinta": 0.473176, "galon": 3.78541, "barril": 158.987, "pie3": 28.3168,
		}
		return convertirGenerico(v, u1, u2, factores, "volumen")
	}

	// --- 4. MASA Y PESO ---
	evaluador.Funciones["convertir_masa"] = func(args ...interface{}) (interface{}, error) {
		v, u1, u2, err := extraerUnidades(args); if err != nil { return nil, err }
		factores := map[string]float64{
			"mg": 1e-6, "g": 0.001, "kg": 1, "tonelada": 1000,
			"onza": 0.0283495, "libra": 0.453592, "stone": 6.35029, "quintal": 100,
		}
		return convertirGenerico(v, u1, u2, factores, "masa")
	}

	// --- 5. TIEMPO ---
	evaluador.Funciones["convertir_tiempo"] = func(args ...interface{}) (interface{}, error) {
		v, u1, u2, err := extraerUnidades(args); if err != nil { return nil, err }
		factores := map[string]float64{
			"ns": 1e-9, "us": 1e-6, "ms": 0.001, "seg": 1, "min": 60, 
			"hora": 3600, "dia": 86400, "semana": 604800, "mes": 2629746, "año": 31556952,
		}
		return convertirGenerico(v, u1, u2, factores, "tiempo")
	}

	// --- 6. VELOCIDAD ---
	evaluador.Funciones["convertir_velocidad"] = func(args ...interface{}) (interface{}, error) {
		v, u1, u2, err := extraerUnidades(args); if err != nil { return nil, err }
		factores := map[string]float64{
			"m_s": 1, "km_h": 0.277778, "milla_h": 0.44704, "nudo": 0.514444, "mach": 343,
		}
		return convertirGenerico(v, u1, u2, factores, "velocidad")
	}

	// --- 7. ALMACENAMIENTO DIGITAL (Data) ---
	evaluador.Funciones["convertir_datos"] = func(args ...interface{}) (interface{}, error) {
		v, u1, u2, err := extraerUnidades(args); if err != nil { return nil, err }
		factores := map[string]float64{
			"bit": 0.125, "byte": 1, "kb": 1e3, "mb": 1e6, "gb": 1e9, "tb": 1e12,
			"kib": 1024, "mib": 1048576, "gib": 1073741824, "tib": 1099511627776,
		}
		return convertirGenerico(v, u1, u2, factores, "datos")
	}

	// --- 8. ENERGÍA Y POTENCIA ---
	evaluador.Funciones["convertir_energia"] = func(args ...interface{}) (interface{}, error) {
		v, u1, u2, err := extraerUnidades(args); if err != nil { return nil, err }
		factores := map[string]float64{
			"joule": 1, "caloria": 4.184, "kcal": 4184, "btu": 1055.06, "ev": 1.602e-19, "kwh": 3.6e6,
		}
		return convertirGenerico(v, u1, u2, factores, "energía")
	}

	evaluador.Funciones["convertir_potencia"] = func(args ...interface{}) (interface{}, error) {
		v, u1, u2, err := extraerUnidades(args); if err != nil { return nil, err }
		factores := map[string]float64{
			"watt": 1, "kw": 1000, "hp": 745.7, "cv": 735.5,
		}
		return convertirGenerico(v, u1, u2, factores, "potencia")
	}

	// --- 9. PRESIÓN ---
	evaluador.Funciones["convertir_presion"] = func(args ...interface{}) (interface{}, error) {
		v, u1, u2, err := extraerUnidades(args); if err != nil { return nil, err }
		factores := map[string]float64{
			"pascal": 1, "bar": 100000, "atm": 101325, "psi": 6894.76, "torr": 133.322,
		}
		return convertirGenerico(v, u1, u2, factores, "presión")
	}

	// --- 10. ÁNGULOS ---
	evaluador.Funciones["convertir_angulo"] = func(args ...interface{}) (interface{}, error) {
		v, u1, u2, err := extraerUnidades(args); if err != nil { return nil, err }
		factores := map[string]float64{
			"grado": 1, "radian": 57.2958, "gradian": 0.9, "arco_min": 0.0166667,
		}
		return convertirGenerico(v, u1, u2, factores, "ángulo")
	}

	// --- 11. TEMPERATURA (Lógica No-Lineal) ---
	evaluador.Funciones["convertir_temperatura"] = func(args ...interface{}) (interface{}, error) {
		v, u1, u2, err := extraerUnidades(args); if err != nil { return nil, err }
		var celsius float64
		switch u1 {
			case "c": celsius = v
			case "f": celsius = (v - 32) * 5 / 9
			case "k": celsius = v - 273.15
			default: return nil, fmt.Errorf("❌ Unidad origen '%s' no válida", u1)
		}
		switch u2 {
			case "c": return celsius, nil
			case "f": return (celsius * 9 / 5) + 32, nil
			case "k": return celsius + 273.15, nil
			default: return nil, fmt.Errorf("❌ Unidad destino '%s' no válida", u2)
		}
	}
}

// --- UTILIDADES INTERNAS DEL MOTOR ---

func extraerUnidades(args []interface{}) (float64, string, string, error) {
	if len(args) != 3 {
		return 0, "", "", fmt.Errorf("❌ ERROR: requiere (valor, 'origen', 'destino')")
	}
	v, err := evaluador.ConvertirAReal(args[0])
	if err != nil { return 0, "", "", err }
	return v, fmt.Sprintf("%v", args[1]), fmt.Sprintf("%v", args[2]), nil
}

func convertirGenerico(v float64, u1, u2 string, factores map[string]float64, tipo string) (float64, error) {
	f1, ok1 := factores[u1]
	f2, ok2 := factores[u2]
	if !ok1 || !ok2 {
		return 0, fmt.Errorf("❌ ERROR: Unidad en %s no reconocida. Disponibles: %v", tipo, obtenerLlaves(factores))
	}
	return (v * f1) / f2, nil
}

func obtenerLlaves(m map[string]float64) []string {
	keys := make([]string, 0, len(m))
	for k := range m { keys = append(keys, k) }
	return keys
}
