package tiempo

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"nepa/desarrollo/interno/administrador"
)

type Tiempo struct {
	mu     sync.RWMutex
	nombre string
	valor  time.Time
}

func CrearTiempo(nombre string, v interface{}) (administrador.Variable, error) {
	t := &Tiempo{
		nombre: strings.TrimSpace(nombre),
		valor:  time.Now(), // Blindaje: Si no hay valor, usamos el "Ahora"
	}
	if v != nil {
		if err := t.AsignarDesdeInterface(v); err != nil {
			return nil, err
		}
	}
	return t, nil
}

func (t *Tiempo) Nombre() string { return t.nombre }
func (t *Tiempo) Tipo() string   { return "tiempo" }

func (t *Tiempo) Mostrar() string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return fmt.Sprintf("%s:%s=%s", t.Tipo(), t.nombre, t.valor.Format("2006-01-02 15:04:05"))
}

func (t *Tiempo) AsignarDesdeInterface(v interface{}) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if v == nil {
		t.valor = time.Now()
		return nil
	}

	switch val := v.(type) {
	case time.Time:
		t.valor = val
	case string:
		// Intentamos parsear formatos comunes de ingeniería
		parsed, err := time.Parse("2006-01-02", val)
		if err != nil {
			parsed, err = time.Parse("2006-01-02 15:04:05", val)
		}
		if err != nil {
			return fmt.Errorf("❌ formato de tiempo inválido. Use AAAA-MM-DD o AAAA-MM-DD HH:MM:SS")
		}
		t.valor = parsed
	default:
		return fmt.Errorf("❌ no se puede convertir %T a tiempo", v)
	}
	return nil
}

func (t *Tiempo) ValorComoInterface() interface{} {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.valor
}

func (t *Tiempo) JSON() string {
	return fmt.Sprintf(`{"tipo":"tiempo","nombre":"%s","valor":"%s"}`, t.nombre, t.valor.String())
}

// Métodos de Interfaz Obligatorios
func (t *Tiempo) ABooleano() (bool, error) { return !t.valor.IsZero(), nil }
func (t *Tiempo) AEntero() (int, error)    { return int(t.valor.Unix()), nil } // Unix timestamp
func (t *Tiempo) AReal() (float64, error)  { return float64(t.valor.UnixNano()), nil }
