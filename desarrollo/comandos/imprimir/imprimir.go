package imprimir

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"nepa/desarrollo/interno/administrador" // Importante: acceso a las variables
	"nepa/desarrollo/interno/evaluador"
)

var (
	ErrSintaxisInvalida = errors.New("sintaxis inválida: use 'imprimir()', 'imprimir \"texto\"', 'imprimir var' o 'imprimir expr'")
)

func imprimirValor(v interface{}) string {
	if v == nil {
		return "nulo"
	}

	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.String:
		// Verificamos si el string es el nombre de una variable existente
		nombreVar := rv.String()
		if variable, err := administrador.ObtenerVariable(nombreVar); err == nil {
			// Si existe, imprimimos su valor real
			return fmt.Sprintf("%v", variable.ValorComoInterface())
		}
		return nombreVar
	case reflect.Bool:
		if rv.Bool() { return "verdadero" }
		return "falso"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", rv.Int())
	case reflect.Float32, reflect.Float64:
		// Usamos %g para que no imprima ceros innecesarios (3.14 en vez de 3.140000)
		return fmt.Sprintf("%g", rv.Float())
	case reflect.Slice, reflect.Array:
		var partes []string
		for i := 0; i < rv.Len(); i++ {
			partes = append(partes, imprimirValor(rv.Index(i).Interface()))
		}
		return "[" + strings.Join(partes, ", ") + "]"
	default:
		return fmt.Sprintf("%v", v)
	}
}

func Ejecutar(linea string) error {
	linea = strings.TrimSpace(linea)

	if strings.HasPrefix(strings.ToLower(linea), "imprimir") {
		// Manejo de paréntesis: imprimir(edad) -> edad
		linea = strings.TrimSpace(linea[len("imprimir"):])
		linea = strings.TrimPrefix(linea, "(")
		linea = strings.TrimSuffix(linea, ")")
	}

	if linea == "" || linea == `""` || linea == `''` {
		fmt.Println()
		return nil
	}

	// 1. Caso: Literal entre comillas (texto puro)
	if (strings.HasPrefix(linea, "\"") && strings.HasSuffix(linea, "\"")) ||
		(strings.HasPrefix(linea, "'") && strings.HasSuffix(linea, "'")) {
		fmt.Println(strings.Trim(linea, "\"'"))
		return nil
	}

	// 2. Caso: Es una variable conocida
	if v, err := administrador.ObtenerVariable(linea); err == nil {
		fmt.Println(imprimirValor(v.ValorComoInterface()))
		return nil
	}

	// 3. Caso: Evaluar como expresión (matemática, etc.)
	resultado, err := evaluador.Eval(linea)
	if err != nil {
		// Si no es variable ni expresión válida, imprimimos el texto original
		fmt.Println(linea)
		return nil
	}

	fmt.Println(imprimirValor(resultado))
	return nil
}

func init() {
	evaluador.Funciones["imprimir"] = func(args ...interface{}) (interface{}, error) {
		for i, arg := range args {
			fmt.Print(imprimirValor(arg))
			if i < len(args)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
		return nil, nil
	}
}
