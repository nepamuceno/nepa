package entero

import (
	"fmt"
	"strings"
	"sync"

	"nepa/desarrollo/interno/administrador"
	"nepa/desarrollo/interno/evaluador"
)

type Entero struct {
	mu     sync.RWMutex
	nombre string
	valor  int64
}

func CrearEntero(nombre string, v interface{}) (administrador.Variable, error) {
	e := &Entero{
		nombre: strings.TrimSpace(nombre),
		valor:  0,
	}
	if v != nil {
		if err := e.AsignarDesdeInterface(v); err != nil {
			return nil, err
		}
	}
	return e, nil
}

func (e *Entero) Nombre() string { return e.nombre }
func (e *Entero) Tipo() string   { return "entero" }

func (e *Entero) Mostrar() string {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return fmt.Sprintf("%s:%s=%d", e.Tipo(), e.nombre, e.valor)
}

func (e *Entero) AsignarDesdeInterface(v interface{}) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	if v == nil {
		e.valor = 0
		return nil
	}
	val, err := evaluador.ConvertirAReal(v)
	if err != nil {
		return fmt.Errorf("❌ valor inválido para Entero: %v", v)
	}
	e.valor = int64(val)
	return nil
}

func (e *Entero) ValorComoInterface() interface{} {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.valor
}

func (e *Entero) JSON() string {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return fmt.Sprintf(`{"tipo":"entero","nombre":"%s","valor":%d}`, e.nombre, e.valor)
}

func (e *Entero) ABooleano() (bool, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.valor != 0, nil
}

func (e *Entero) AEntero() (int, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return int(e.valor), nil
}

func (e *Entero) AReal() (float64, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return float64(e.valor), nil
}
