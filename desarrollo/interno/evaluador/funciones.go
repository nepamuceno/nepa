package evaluador

import (
    "errors"
    "strings"
)

// Funciones internas y externas registradas.
// La clave es el nombre en minúsculas de la función o "tipo.metodo" para métodos.
var Funciones = map[string]func(args ...interface{}) (interface{}, error){
    // --- Funciones matemáticas ---
    "suma": func(args ...interface{}) (interface{}, error) {
        var total float64
        for _, a := range args {
            f, err := ConvertirAReal(a)
            if err != nil {
                return nil, err
            }
            total += f
        }
        return total, nil
    },

    "max": func(args ...interface{}) (interface{}, error) {
        if len(args) == 0 {
            return nil, errors.New("max requiere al menos un argumento")
        }
        m, err := ConvertirAReal(args[0])
        if err != nil {
            return nil, err
        }
        for _, a := range args[1:] {
            f, err := ConvertirAReal(a)
            if err != nil {
                return nil, err
            }
            if f > m {
                m = f
            }
        }
        return m, nil
    },

    // --- Métodos sobre cadenas ---
    "cadena.convertir_caracter": func(args ...interface{}) (interface{}, error) {
        if len(args) < 1 {
            return nil, errors.New("convertir_caracter requiere al menos una cadena")
        }
        s, ok := args[0].(string)
        if !ok {
            return nil, errors.New("primer argumento no es una cadena")
        }
        // Ejemplo simple: convertir a mayúsculas
        return strings.ToUpper(s), nil
    },

    "cadena.longitud": func(args ...interface{}) (interface{}, error) {
        if len(args) < 1 {
            return nil, errors.New("longitud requiere al menos una cadena")
        }
        s, ok := args[0].(string)
        if !ok {
            return nil, errors.New("primer argumento no es una cadena")
        }
        return len(s), nil
    },
}
