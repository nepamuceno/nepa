package variable

import (
	"fmt"
	"regexp"
	"strings"

	"nepa/desarrollo/interno/administrador"
	"nepa/desarrollo/interno/bloque"
	"nepa/desarrollo/interno/evaluador"
	"nepa/desarrollo/interno/parser"
)

// esNombreValido asegura que el nombre no sea una palabra reservada ni tenga caracteres raros
func esNombreValido(nombre string) bool {
	re := regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*$`)
	if !re.MatchString(nombre) {
		return false
	}
	for _, r := range bloque.PalabrasReservadas {
		if nombre == r {
			return false
		}
	}
	return true
}

func init() {
	evaluador.Registrar("variable", func(n parser.Nodo, ctx *evaluador.Contexto) {
		// 1. Obtener tipo desde Args
		var tipo string
		if len(n.Args) > 0 {
			if t, ok := n.Args[0].(string); ok {
				tipo = strings.ToLower(strings.TrimSpace(t))
			}
		}

		// 2. Buscar el constructor (CrearEntero, CrearReal, etc.)
		constructor, ok := administrador.Constructores[tipo]
		if !ok {
			fmt.Printf("❌ Error: tipo de variable '%s' no reconocido\n", tipo)
			return
		}

		// 3. Evaluar el valor (Resuelve expresiones como base + ajuste)
		var valorFinal interface{} = n.Valor
		if strValor, ok := n.Valor.(string); ok && strValor != "" {
			// Intentamos calcular el resultado
			res, err := evaluador.Eval(strValor)
			if err == nil {
				valorFinal = res
			} else {
				// Si no es una expresión matemática, se queda como literal
				valorFinal = strValor
			}
		}

		// 4. Crear cada variable (soporta comas: a, b, c)
		nombres := strings.Split(n.Nombre, ",")
		for _, nombre := range nombres {
			nombre = strings.TrimSpace(nombre)
			if !esNombreValido(nombre) {
				fmt.Printf("❌ Error: nombre de variable '%s' inválido\n", nombre)
				continue
			}

			// Invocamos al constructor con el valor ya evaluado
			v, err := constructor(nombre, valorFinal)
			if err != nil {
				fmt.Printf("❌ Error creando %s (%s): %v\n", nombre, tipo, err)
				continue
			}

			// Registro en el administrador y en el contexto de ejecución
			administrador.RegistrarVariable(nombre, v)
			if ctx != nil && ctx.Variables != nil {
				ctx.Variables[nombre] = v
			}
			fmt.Printf("✔ Variable creada: %s\n", v.Mostrar())
		}
	})
}
