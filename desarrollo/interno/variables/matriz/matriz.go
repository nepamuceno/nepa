package matriz

import (
	"fmt"
	"strings"
	"sync"

	"nepa/desarrollo/interno/administrador"
)

type Matriz struct {
	mu     sync.RWMutex
	nombre string
	valor  [][]float64
}

func CrearMatriz(nombre string, v interface{}) (administrador.Variable, error) {
	m := &Matriz{
		nombre: strings.TrimSpace(nombre),
		valor:  [][]float64{}, // Blindaje: slice vacío, no nil
	}
	if v != nil {
		if err := m.AsignarDesdeInterface(v); err != nil {
			return nil, err
		}
	}
	return m, nil
}

func (m *Matriz) Nombre() string { return m.nombre }
func (m *Matriz) Tipo() string   { return "matriz" }

func (m *Matriz) Mostrar() string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	filas := len(m.valor)
	cols := 0
	if filas > 0 {
		cols = len(m.valor[0])
	}
	return fmt.Sprintf("%s:%s=[%dx%d]", m.Tipo(), m.nombre, filas, cols)
}

func (m *Matriz) AsignarDesdeInterface(v interface{}) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if v == nil {
		m.valor = [][]float64{}
		return nil
	}

	if val, ok := v.([][]float64); ok {
		m.valor = val
		return nil
	}
	return fmt.Errorf("❌ formato incompatible para Matriz")
}

func (m *Matriz) ValorComoInterface() interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.valor
}

func (m *Matriz) JSON() string {
	return fmt.Sprintf(`{"tipo":"matriz","nombre":"%s"}`, m.nombre)
}

// Métodos de Interfaz Obligatorios
func (m *Matriz) ABooleano() (bool, error) { return len(m.valor) > 0, nil }
func (m *Matriz) AEntero() (int, error)    { return len(m.valor), nil }
func (m *Matriz) AReal() (float64, error)  { return float64(len(m.valor)), nil }
