package administrador

import (
    "errors"
    "fmt"
    "sync"
)

// Interfaz común que todos los tipos deben implementar.
// Todos los paquetes de variables deben importar esta interfaz.
type Variable interface {
    Nombre() string
    Tipo() string
    Mostrar() string
    JSON() string

    AsignarDesdeInterface(v interface{}) error
    ValorComoInterface() interface{}

    // Conversiones básicas universales
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

// RegistrarConstructor permite añadir un nuevo tipo al ecosistema.
// Ejemplo: administrador.RegistrarConstructor("bit", bit.CrearBit)
func RegistrarConstructor(tipo string, f func(string, interface{}) (Variable, error)) {
    Constructores[tipo] = f
}

// RegistrarConstructores permite inyectar múltiples tipos de una sola vez.
// Ejemplo:
// administrador.RegistrarConstructores(map[string]func(string, interface{}) (Variable, error){
//     "bit":    bit.CrearBit,
//     "entero": entero.CrearEntero,
// })
func RegistrarConstructores(mapa map[string]func(string, interface{}) (Variable, error)) {
    for tipo, f := range mapa {
        Constructores[tipo] = f
    }
}

// CrearVariableUniversal crea una variable de cualquier tipo registrado.
// Si el constructor no existe, retorna ErrConstructorNoExiste.
func CrearVariableUniversal(tipo, nombre string, valor interface{}) (Variable, error) {
    mu.Lock()
    defer mu.Unlock()
    if _, existe := tabla[nombre]; existe {
        return nil, ErrVariableYaExiste
    }
    constructor, ok := Constructores[tipo]
    if !ok {
        return nil, ErrConstructorNoExiste
    }
    v, err := constructor(nombre, valor)
    if err != nil {
        return nil, fmt.Errorf("error creando variable '%s' de tipo '%s': %w", nombre, tipo, err)
    }
    tabla[nombre] = v
    return v, nil
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
