package imprimir

import (
    "errors"
    "fmt"
    "strings"

    "nepa/interno/evaluador"
)

var (
    ErrSintaxisInvalida = errors.New("sintaxis inválida: use 'imprimir()', 'imprimir \"texto\"', 'imprimir var' o 'imprimir expr'")
    ErrTipoNoImprimible = errors.New("tipo de dato no imprimible")
)

func Ejecutar(linea string) error {
    linea = strings.TrimSpace(linea)

    // Quitar palabra clave "imprimir"
    if strings.HasPrefix(strings.ToLower(linea), "imprimir") {
        linea = strings.TrimSpace(linea[len("imprimir"):])
    }

    // Caso: imprimir sin nada → error
    if linea == "" {
        return ErrSintaxisInvalida
    }

    // Caso: imprimir() → línea en blanco
    if linea == "()" {
        fmt.Println()
        return nil
    }

    // Caso: imprimir "" o '' → línea en blanco
    if linea == `""` || linea == `''` {
        fmt.Println()
        return nil
    }

    // Evaluar expresión completa (concatenaciones, funciones, variables, etc.)
    resultado, err := evaluador.Eval(linea)
    if err != nil {
        return fmt.Errorf("error evaluando expresión '%s': %v", linea, err)
    }

    // Convertir resultado a string imprimible
    switch v := resultado.(type) {
    case string:
        fmt.Println(v)
    case int, int64, float64, bool:
        fmt.Println(fmt.Sprintf("%v", v))
    default:
        return fmt.Errorf("%w: %T", ErrTipoNoImprimible, v)
    }

    return nil
}
