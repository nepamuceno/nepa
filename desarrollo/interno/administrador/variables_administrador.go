// desarrollo/interno/administrador/variables_administrador.go
package administrador

import (
    "errors"
    "sync"
)

// Interfaz común que todos los tipos deben implementar.
// Ya la definimos en bit.go, pero lo ideal es centralizarla aquí
// para que todos los tipos la importen desde este paquete.
type Variable interface {
    Nombre() string
    Tipo() string
    Mostrar() string
    JSON() string

    AsignarDesdeInterface(v interface{}) error
    ValorComoInterface() interface{}

    // Conversiones básicas
    ABooleano() (bool, error)
    AEntero() (int, error)
    AReal() (float64, error)
}

// Errores comunes del administrador
var (
    ErrVariableNoEncontrada = errors.New("variable no encontrada")
    ErrVariableYaExiste     = errors.New("ya existe una variable con ese nombre")
    ErrConstructorNoExiste  = errors.New("no existe constructor para ese tipo")
)

// Tabla global de símbolos: nombre → Variable
var tabla = make(map[string]Variable)

// Mutex para concurrencia segura
var mu sync.RWMutex

// Registro de constructores: tipo → función(nombre, valor) → Variable
var Constructores = make(map[string]func(string, interface{}) (Variable, error))

// RegistrarConstructor permite añadir nuevos tipos al ecosistema.
// Ejemplo: administrador.RegistrarConstructor("bit", bit.CrearBit)
func RegistrarConstructor(tipo string, f func(string, interface{}) (Variable, error)) {
    Constructores[tipo] = f
}

// RegistrarVariable guarda una variable en la tabla global.
// Si ya existe una con el mismo nombre, retorna error.
func RegistrarVariable(nombre string, v Variable) error {
    mu.Lock()
    defer mu.Unlock()
    if _, existe := tabla[nombre]; existe {
        return ErrVariableYaExiste
    }
    tabla[nombre] = v
    return nil
}

// ObtenerVariable devuelve una variable por nombre.
func ObtenerVariable(nombre string) (Variable, error) {
    mu.RLock()
    defer mu.RUnlock()
    if v, ok := tabla[nombre]; ok {
        return v, nil
    }
    return nil, ErrVariableNoEncontrada
}

// ModificarVariable reemplaza el valor de una variable existente.
func ModificarVariable(nombre string, nuevoValor interface{}) error {
    mu.Lock()
    defer mu.Unlock()
    v, ok := tabla[nombre]
    if !ok {
        return ErrVariableNoEncontrada
    }
    return v.AsignarDesdeInterface(nuevoValor)
}

// BorrarVariable elimina una variable de la tabla global.
func BorrarVariable(nombre string) error {
    mu.Lock()
    defer mu.Unlock()
    if _, ok := tabla[nombre]; !ok {
        return ErrVariableNoEncontrada
    }
    delete(tabla, nombre)
    return nil
}

// MostrarVariable imprime la representación textual de una variable.
func MostrarVariable(nombre string) (string, error) {
    mu.RLock()
    defer mu.RUnlock()
    if v, ok := tabla[nombre]; ok {
        return v.Mostrar(), nil
    }
    return "", ErrVariableNoEncontrada
}

// ListarVariables devuelve todas las variables registradas.
func ListarVariables() []Variable {
    mu.RLock()
    defer mu.RUnlock()
    var lista []Variable
    for _, v := range tabla {
        lista = append(lista, v)
    }
    return lista
}
