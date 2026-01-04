package booleano

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"nepa/desarrollo/interno/administrador"
)

type Booleano struct {
	mu     sync.RWMutex
	nombre string
	valor  bool
}

func CrearBooleano(nombre string, v interface{}) (administrador.Variable, error) {
	b := &Booleano{nombre: strings.TrimSpace(nombre)}
	if v != nil {
		if err := b.AsignarDesdeInterface(v); err != nil {
			return nil, err
		}
	}
	return b, nil
}

func (b *Booleano) Nombre() string { return b.nombre }
func (b *Booleano) Tipo() string   { return "booleano" }

func (b *Booleano) Mostrar() string {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return fmt.Sprintf("%s:%s=%t", b.Tipo(), b.nombre, b.valor)
}

func (b *Booleano) AsignarDesdeInterface(v interface{}) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	switch val := v.(type) {
	case bool:
		b.valor = val
	case int, float64, int64:
		b.valor = val != 0
	case string:
		s := strings.ToLower(strings.TrimSpace(val))
		b.valor = (s == "true" || s == "verdadero" || s == "1")
	default:
		return fmt.Errorf("❌ no se puede convertir a booleano")
	}
	return nil
}

func (b *Booleano) ValorComoInterface() interface{} {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.valor
}

func (b *Booleano) JSON() string {
	b.mu.RLock()
	defer b.mu.RUnlock()
	data, _ := json.Marshal(map[string]interface{}{
		"tipo": b.Tipo(), "nombre": b.nombre, "valor": b.valor,
	})
	return string(data)
}

// --- Métodos de Interfaz Obligatorios ---

func (b *Booleano) ABooleano() (bool, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.valor, nil
}

func (b *Booleano) AEntero() (int, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	if b.valor {
		return 1, nil
	}
	return 0, nil
}

func (b *Booleano) AReal() (float64, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	if b.valor {
		return 1.0, nil
	}
	return 0.0, nil
}
