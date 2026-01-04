package lista

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"nepa/desarrollo/interno/administrador"
)

type Lista struct {
	mu     sync.RWMutex
	nombre string
	valor  []interface{}
}

func CrearLista(nombre string, v interface{}) (administrador.Variable, error) {
	l := &Lista{
		nombre: strings.TrimSpace(nombre),
		valor:  make([]interface{}, 0),
	}
	if v != nil {
		if err := l.AsignarDesdeInterface(v); err != nil {
			return nil, err
		}
	}
	return l, nil
}

func (l *Lista) Nombre() string { return l.nombre }
func (l *Lista) Tipo() string   { return "lista" }

func (l *Lista) Mostrar() string {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return fmt.Sprintf("%s:%s=%v", l.Tipo(), l.nombre, l.valor)
}

func (l *Lista) AsignarDesdeInterface(v interface{}) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if v == nil {
		l.valor = make([]interface{}, 0)
		return nil
	}

	// Usamos reflexión para manejar cualquier tipo de slice/array que venga de Go
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array {
		nuevoSlice := make([]interface{}, rv.Len())
		for i := 0; i < rv.Len(); i++ {
			nuevoSlice[i] = rv.Index(i).Interface()
		}
		l.valor = nuevoSlice
		return nil
	}

	// Si no es un slice, lo envolvemos como primer elemento de la lista
	l.valor = []interface{}{v}
	return nil
}

func (l *Lista) ValorComoInterface() interface{} {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.valor
}

func (l *Lista) JSON() string {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return fmt.Sprintf(`{"tipo":"lista","nombre":"%s","longitud":%d}`, l.nombre, len(l.valor))
}

func (l *Lista) ABooleano() (bool, error) { 
	l.mu.RLock()
	defer l.mu.RUnlock()
	return len(l.valor) > 0, nil 
}

func (l *Lista) AEntero() (int, error) { 
	l.mu.RLock()
	defer l.mu.RUnlock()
	return len(l.valor), nil 
}

func (l *Lista) AReal() (float64, error) { 
	l.mu.RLock()
	defer l.mu.RUnlock()
	return float64(len(l.valor)), nil 
}
