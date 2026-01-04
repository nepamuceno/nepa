package real

import (
	"fmt"
	"strings"
	"sync"

	"nepa/desarrollo/interno/administrador"
	"nepa/desarrollo/interno/evaluador"
)

type Real struct {
	mu     sync.RWMutex
	nombre string
	valor  float64
}

func CrearReal(nombre string, v interface{}) (administrador.Variable, error) {
	r := &Real{
		nombre: strings.TrimSpace(nombre),
		valor:  0.0,
	}
	if v != nil {
		if err := r.AsignarDesdeInterface(v); err != nil {
			return nil, err
		}
	}
	return r, nil
}

func (r *Real) Nombre() string { return r.nombre }
func (r *Real) Tipo() string   { return "real" }

func (r *Real) Mostrar() string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return fmt.Sprintf("%s:%s=%g", r.Tipo(), r.nombre, r.valor)
}

func (r *Real) AsignarDesdeInterface(v interface{}) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if v == nil {
		r.valor = 0.0
		return nil
	}
	val, err := evaluador.ConvertirAReal(v)
	if err != nil {
		return err
	}
	r.valor = val
	return nil
}

func (r *Real) ValorComoInterface() interface{} {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.valor
}

func (r *Real) JSON() string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return fmt.Sprintf(`{"tipo":"real","nombre":"%s","valor":%g}`, r.nombre, r.valor)
}

func (r *Real) ABooleano() (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.valor != 0, nil
}

func (r *Real) AEntero() (int, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return int(r.valor), nil
}

func (r *Real) AReal() (float64, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.valor, nil
}
