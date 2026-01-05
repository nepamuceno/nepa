package parser

import (
	"strings"
)

// SegmentarEstructura devuelve un []interface{} para que matriz.go lo reconozca como Slice
func SegmentarEstructura(raw string) interface{} {
	raw = strings.TrimSpace(raw)

	// Si no empieza con corchete, es un valor final (string)
	if !strings.HasPrefix(raw, "[") {
		return raw
	}

	// Quitamos solo los corchetes de este nivel: [[a],[b]] -> [a],[b]
	if strings.HasPrefix(raw, "[") && strings.HasSuffix(raw, "]") {
		raw = raw[1 : len(raw)-1]
	}

	var elementos []interface{}
	var buffer strings.Builder
	nivelCorchete := 0
	nivelParentesis := 0

	for _, char := range raw {
		switch char {
		case '[':
			nivelCorchete++
			buffer.WriteRune(char)
		case ']':
			nivelCorchete--
			buffer.WriteRune(char)
		case '(':
			nivelParentesis++
			buffer.WriteRune(char)
		case ')':
			nivelParentesis--
			buffer.WriteRune(char)
		case ',':
			// Solo cortamos si estamos en el nivel raíz de ESTA estructura
			if nivelCorchete == 0 && nivelParentesis == 0 {
				str := strings.TrimSpace(buffer.String())
				if str != "" {
					// RECURSIÓN: Esto crea la jerarquía de Slices que matriz.go ama
					elementos = append(elementos, SegmentarEstructura(str))
				}
				buffer.Reset()
				continue
			}
			buffer.WriteRune(char)
		default:
			buffer.WriteRune(char)
		}
	}

	finalStr := strings.TrimSpace(buffer.String())
	if finalStr != "" {
		elementos = append(elementos, SegmentarEstructura(finalStr))
	}

	return elementos
}
