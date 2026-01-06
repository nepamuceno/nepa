package evaluador

import "fmt"

// Puntero representa una referencia a un valor.
// Puede usarse para simular &variable y *puntero en Nepa.
type Puntero struct {
    Valor interface{}
}

// String devuelve una representación legible del puntero.
func (p Puntero) String() string {
    return fmt.Sprintf("&%v", p.Valor)
}

// NuevoPuntero crea un puntero a un valor dado.
func NuevoPuntero(v interface{}) Puntero {
    return Puntero{Valor: v}
}

// EsPuntero indica si un valor es de tipo Puntero.
func EsPuntero(v interface{}) bool {
    _, ok := v.(Puntero)
    return ok
}

// Desreferenciar devuelve el valor apuntado por un Puntero.
// Si el valor no es puntero, devuelve un error claro.
func Desreferenciar(v interface{}) (interface{}, error) {
    if p, ok := v.(Puntero); ok {
        return p.Valor, nil
    }
    return nil, fmt.Errorf("❌ error: no se puede desreferenciar, %T no es puntero", v)
}
