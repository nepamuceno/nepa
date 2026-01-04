package matematicas

import (
	"fmt"
	"math/bits"
	"nepa/desarrollo/interno/evaluador"
	"strconv"
	"strings"
)

func inyectarComputacionalGlobal() {

	// --- 1. CONVERSIÓN MAESTRA (AUTO-DETECCIÓN) ---

	// convertir_maestro(valor, [alfabeto_destino])
	// Si solo se pasa valor, lo convierte a decimal basándose en su propia anatomía.
	evaluador.Funciones["convertir_cualquier_base"] = func(args ...interface{}) (interface{}, error) {
		if len(args) < 1 {
			return nil, fmt.Errorf("❌ ERROR: falta el valor a convertir")
		}

		// Definir Alfabeto Destino (Default: Decimal)
		alfDestino := "0123456789"
		if len(args) >= 2 && args[1] != nil {
			alfDestino = fmt.Sprintf("%v", args[1])
		}
		if err := validarAlfabetoGuru(alfDestino); err != nil { return nil, err }

		var decimal int64
		entrada := args[0]

		// Lógica de detección de Origen
		if n, ok := entrada.(float64); ok {
			// Si ya es un número (ej: resultado de otra función), es decimal puro
			decimal = int64(n)
		} else {
			// Si es cadena, deducimos su alfabeto por las letras únicas que contiene
			strEntrada := fmt.Sprintf("%v", entrada)
			alfOrigen := ""
			visto := make(map[rune]bool)
			for _, r := range strEntrada {
				if !visto[r] {
					visto[r] = true
					alfOrigen += string(r)
				}
			}

			baseOrigen := int64(len(alfOrigen))
			if baseOrigen < 2 {
				// Si es un solo carácter repetido "aaaaa", no hay base base válida
				return nil, fmt.Errorf("❌ ERROR: no se puede deducir base de un solo carácter")
			}

			// Convertir a decimal usando el alfabeto auto-detectado
			for i := 0; i < len(strEntrada); i++ {
				idx := strings.IndexByte(alfOrigen, strEntrada[i])
				decimal = decimal*baseOrigen + int64(idx)
			}
		}

		// Convertir de Decimal al Alfabeto de Destino
		if decimal == 0 { return string(alfDestino[0]), nil }

		baseDest := int64(len(alfDestino))
		var res strings.Builder
		tempDec := decimal
		if tempDec < 0 { tempDec = -tempDec }

		for tempDec > 0 {
			res.WriteByte(alfDestino[tempDec%baseDest])
			tempDec /= baseDest
		}

		// Invertir cadena resultante
		runes := []rune(res.String())
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		
		final := string(runes)
		if decimal < 0 { final = "-" + final }
		return final, nil
	}

	// --- 2. WRAPPERS DE CONVERSIÓN RÁPIDA ---

	evaluador.Funciones["a_binario"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("a_binario", args); if err != nil { return nil, err }
		return strconv.FormatInt(int64(v), 2), nil
	}

	evaluador.Funciones["a_hex"] = func(args ...interface{}) (interface{}, error) {
		v, err := validar1("a_hex", args); if err != nil { return nil, err }
		return "0x" + strconv.FormatInt(int64(v), 16), nil
	}

	// --- 3. OPERACIONES LÓGICAS DE BITS (BITWISE) ---

	evaluador.Funciones["bit_and"] = func(args ...interface{}) (interface{}, error) {
		a, b, err := validar2("bit_and", args); if err != nil { return nil, err }
		return float64(int64(a) & int64(b)), nil
	}

	evaluador.Funciones["bit_or"] = func(args ...interface{}) (interface{}, error) {
		a, b, err := validar2("bit_or", args); if err != nil { return nil, err }
		return float64(int64(a) | int64(b)), nil
	}

	evaluador.Funciones["bit_xor"] = func(args ...interface{}) (interface{}, error) {
		a, b, err := validar2("bit_xor", args); if err != nil { return nil, err }
		return float64(int64(a) ^ int64(b)), nil
	}

	evaluador.Funciones["bit_not"] = func(args ...interface{}) (interface{}, error) {
		a, err := validar1("bit_not", args); if err != nil { return nil, err }
		return float64(^int64(a)), nil
	}

	// --- 4. DESPLAZAMIENTOS Y ROTACIONES ---

	evaluador.Funciones["desplazar_izq"] = func(args ...interface{}) (interface{}, error) {
		n, p, err := validar2("desplazar_izq", args); if err != nil { return nil, err }
		return float64(int64(n) << uint64(p)), nil
	}

	evaluador.Funciones["desplazar_der"] = func(args ...interface{}) (interface{}, error) {
		n, p, err := validar2("desplazar_der", args); if err != nil { return nil, err }
		return float64(int64(n) >> uint64(p)), nil
	}

	evaluador.Funciones["rotar_izq"] = func(args ...interface{}) (interface{}, error) {
		n, p, err := validar2("rotar_izq", args); if err != nil { return nil, err }
		return float64(bits.RotateLeft64(uint64(n), int(p))), nil
	}

	// --- 5. ANÁLISIS COMPUTACIONAL ---

	evaluador.Funciones["contar_bits_encendidos"] = func(args ...interface{}) (interface{}, error) {
		n, err := validar1("contar_bits_encendidos", args); if err != nil { return nil, err }
		return float64(bits.OnesCount64(uint64(n))), nil
	}

	evaluador.Funciones["paridad"] = func(args ...interface{}) (interface{}, error) {
		n, err := validar1("paridad", args); if err != nil { return nil, err }
		return float64(bits.OnesCount64(uint64(n)) % 2), nil
	}

	evaluador.Funciones["invertir_bytes"] = func(args ...interface{}) (interface{}, error) {
		n, err := validar1("invertir_bytes", args); if err != nil { return nil, err }
		return float64(bits.ReverseBytes64(uint64(n))), nil
	}

	evaluador.Funciones["es_potencia_de_dos"] = func(args ...interface{}) (interface{}, error) {
		n, err := validar1("es_potencia_de_dos", args); if err != nil { return nil, err }
		val := int64(n)
		if val <= 0 { return false, nil }
		return (val & (val - 1)) == 0, nil
	}
}
// Auxiliar para asegurar que el alfabeto del usuario sea válido
func validarAlfabetoGuru(alf string) error {
    if len(alf) < 2 { 
        return fmt.Errorf("❌ ERROR: alfabeto muy corto") 
    }
    visto := make(map[rune]bool)
    for _, r := range alf {
        if visto[r] { 
            return fmt.Errorf("❌ ERROR: carácter repetido en alfabeto: %c", r) 
        }
        visto[r] = true
    }
    return nil
}
