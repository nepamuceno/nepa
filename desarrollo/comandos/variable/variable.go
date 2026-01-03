package variable

import (
    "errors"
    "fmt"
    "strings"

    "nepa/interno/administrador"
)

// Errores específicos del comando variable.
var (
    ErrSintaxisInvalida = errors.New("sintaxis inválida para comando 'variable'")
    ErrTipoDesconocido  = errors.New("tipo de variable desconocido")
)

// Ejecutar interpreta el comando 'variable'.
// Sintaxis esperada: variable <tipo> <nombre> [valor_opcional]
func Ejecutar(args []string) error {
    if len(args) < 2 {
        return ErrSintaxisInvalida
    }

    tipo := strings.ToLower(args[0])
    nombre := args[1]
    var valor interface{} = nil
    if len(args) > 2 {
        valor = args[2]
    }

    constructor, ok := administrador.Constructores[tipo]
    if !ok {
        return fmt.Errorf("%w: %s", ErrTipoDesconocido, tipo)
    }

    v, err := constructor(nombre, valor)
    if err != nil {
        return fmt.Errorf("error creando variable '%s' de tipo '%s': %w", nombre, tipo, err)
    }

    administrador.RegistrarVariable(nombre, v)
    fmt.Printf("✔ Variable creada: %s\n", v.Mostrar())
    return nil
}
