package bit

import (
    "encoding/json"
    "errors"
    "fmt"
    "strconv"
    "strings"
    "sync"

    "nepa/interno/administrador"
)

// Errores específicos del tipo Bit
var (
    ErrValorInvalido  = errors.New("valor inválido para Bit: solo se permite 0 o 1")
    ErrAsignacionNula = errors.New("asignación nula: se requiere un valor para Bit")
    ErrConversion     = errors.New("conversión inválida hacia/desde Bit")
    ErrNombreInvalido = errors.New("nombre inválido: vacío o con espacios")
)

// Bit representa un valor binario (0 o 1) con nombre y metadatos.
type Bit struct {
    mu        sync.RWMutex
    nombre    string
    valor     uint8
    etiquetas map[string]string
}

// CrearBit crea una instancia de Bit y devuelve administrador.Variable.
// Cumple con la firma esperada por el administrador.
func CrearBit(nombre string, v interface{}) (administrador.Variable, error) {
    nombre = strings.TrimSpace(nombre)
    if nombre == "" || strings.Contains(nombre, " ") {
        return nil, ErrNombreInvalido
    }

    b := &Bit{
        nombre:    nombre,
        valor:     0, // por defecto
        etiquetas: make(map[string]string),
    }

    if v != nil {
        if err := b.AsignarDesdeInterface(v); err != nil {
            return nil, err
        }
    }
    return b, nil
}

// Implementación de la interfaz Variable

func (b *Bit) Nombre() string {
    b.mu.RLock()
    defer b.mu.RUnlock()
    return b.nombre
}

func (b *Bit) Tipo() string { return "bit" }

func (b *Bit) Mostrar() string {
    b.mu.RLock()
    defer b.mu.RUnlock()
    return fmt.Sprintf("%s:%s=%d", b.Tipo(), b.nombre, b.valor)
}

func (b *Bit) JSON() string {
    b.mu.RLock()
    defer b.mu.RUnlock()
    payload := map[string]interface{}{
        "tipo":   b.Tipo(),
        "nombre": b.nombre,
        "valor":  b.valor,
    }
    data, _ := json.Marshal(payload)
    return string(data)
}

func (b *Bit) AsignarDesdeInterface(v interface{}) error {
    switch val := v.(type) {
    case nil:
        return ErrAsignacionNula
    case bool:
        if val {
            b.valor = 1
        } else {
            b.valor = 0
        }
        return nil
    case int:
        if val == 0 || val == 1 {
            b.valor = uint8(val)
            return nil
        }
        return ErrValorInvalido
    case string:
        s := strings.TrimSpace(strings.ToLower(val))
        switch s {
        case "0", "false", "falso":
            b.valor = 0
            return nil
        case "1", "true", "verdadero":
            b.valor = 1
            return nil
        default:
            if n, err := strconv.Atoi(s); err == nil {
                if n == 0 || n == 1 {
                    b.valor = uint8(n)
                    return nil
                }
            }
            return ErrValorInvalido
        }
    default:
        return ErrValorInvalido
    }
}

func (b *Bit) ValorComoInterface() interface{} {
    b.mu.RLock()
    defer b.mu.RUnlock()
    return b.valor
}

func (b *Bit) ABooleano() (bool, error) {
    switch b.valor {
    case 0:
        return false, nil
    case 1:
        return true, nil
    default:
        return false, ErrConversion
    }
}

func (b *Bit) AEntero() (int, error) {
    if b.valor == 0 || b.valor == 1 {
        return int(b.valor), nil
    }
    return 0, ErrConversion
}

func (b *Bit) AReal() (float64, error) {
    if b.valor == 0 || b.valor == 1 {
        return float64(b.valor), nil
    }
    return 0.0, ErrConversion
}

// Métodos adicionales propios del Bit

func (b *Bit) Invertir() {
    b.mu.Lock()
    if b.valor == 0 {
        b.valor = 1
    } else {
        b.valor = 0
    }
    b.mu.Unlock()
}

func (b *Bit) EsCero() bool {
    b.mu.RLock()
    defer b.mu.RUnlock()
    return b.valor == 0
}

func (b *Bit) EsUno() bool {
    b.mu.RLock()
    defer b.mu.RUnlock()
    return b.valor == 1
}

// Registro automático del constructor en el administrador
func init() {
    administrador.RegistrarConstructor("bit", CrearBit)
}
