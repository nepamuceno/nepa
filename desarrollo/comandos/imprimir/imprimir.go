package imprimir

import (
    "errors"
    "fmt"
    "reflect"
    "strings"

    "nepa/desarrollo/interno/evaluador"
)

var (
    ErrSintaxisInvalida = errors.New("sintaxis inválida: use 'imprimir()', 'imprimir \"texto\"', 'imprimir var' o 'imprimir expr'")
)

// imprimirValor convierte cualquier valor Go en una representación legible
func imprimirValor(v interface{}) string {
    if v == nil {
        return "nulo"
    }

    rv := reflect.ValueOf(v)
    switch rv.Kind() {
    case reflect.String:
        return rv.String()
    case reflect.Bool:
        if rv.Bool() {
            return "verdadero"
        }
        return "falso"
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        return fmt.Sprintf("%d", rv.Int())
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
        return fmt.Sprintf("%d", rv.Uint())
    case reflect.Float32, reflect.Float64:
        return fmt.Sprintf("%f", rv.Float())
    case reflect.Slice, reflect.Array:
        var partes []string
        for i := 0; i < rv.Len(); i++ {
            partes = append(partes, imprimirValor(rv.Index(i).Interface()))
        }
        return "[" + strings.Join(partes, ", ") + "]"
    case reflect.Map:
        var partes []string
        for _, key := range rv.MapKeys() {
            partes = append(partes,
                fmt.Sprintf("%s: %s", imprimirValor(key.Interface()), imprimirValor(rv.MapIndex(key).Interface())))
        }
        return "{" + strings.Join(partes, ", ") + "}"
    case reflect.Ptr:
        if rv.IsNil() {
            return "puntero(nil)"
        }
        return fmt.Sprintf("puntero→%s", imprimirValor(rv.Elem().Interface()))
    case reflect.Func:
        return "función"
    default:
        return fmt.Sprintf("%v", v)
    }
}

// Ejecutar mantiene compatibilidad con main.go (línea cruda)
func Ejecutar(linea string) error {
    linea = strings.TrimSpace(linea)

    if strings.HasPrefix(strings.ToLower(linea), "imprimir") {
        linea = strings.TrimSpace(linea[len("imprimir"):])
    }

    if linea == "" || linea == "()" || linea == `""` || linea == `''` {
        fmt.Println()
        return nil
    }

    // Literal entre comillas
    if (strings.HasPrefix(linea, "\"") && strings.HasSuffix(linea, "\"")) ||
        (strings.HasPrefix(linea, "'") && strings.HasSuffix(linea, "'")) {
        contenido := strings.Trim(linea, "\"'")
        fmt.Println(contenido)
        return nil
    }

    // Evaluar expresión completa
    resultado, err := evaluador.Eval(linea)
    if err != nil {
        return fmt.Errorf("❌ error evaluando expresión '%s': %v", linea, err)
    }

    fmt.Println(imprimirValor(resultado))
    return nil
}

// init registra este comando en el mapa global de funciones
func init() {
    evaluador.Funciones["imprimir"] = func(args ...interface{}) (interface{}, error) {
        if len(args) == 0 {
            fmt.Println()
            return nil, nil
        }
        for _, arg := range args {
            fmt.Println(imprimirValor(arg))
        }
        return nil, nil
    }
}
