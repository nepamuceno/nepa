package objeto

import (
	"fmt"
	"strings"
	"sync"
	"nepa/desarrollo/interno/administrador"
)

type Objeto struct {
	mu     sync.RWMutex
	nombre string
	valor  interface{}
}

func CrearObjeto(nombre string, v interface{}) (administrador.Variable, error) {
	o := &Objeto{
		nombre: strings.TrimSpace(nombre),
		valor:  make(map[string]interface{}),
	}
	if v != nil {
		o.AsignarDesdeInterface(v)
	}
	return o, nil
}

func (o *Objeto) Nombre() string { return o.nombre }
func (o *Objeto) Tipo() string   { return "objeto" }

func (o *Objeto) Mostrar() string {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return fmt.Sprintf("%s:%s=%v", o.Tipo(), o.nombre, o.valor)
}

func (o *Objeto) AsignarDesdeInterface(v interface{}) error {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.valor = v // Acepta cualquier estructura
	return nil
}

func (o *Objeto) ValorComoInterface() interface{} {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return o.valor
}

func (o *Objeto) JSON() string {
	return fmt.Sprintf(`{"tipo":"objeto","nombre":"%s"}`, o.nombre)
}

func (o *Objeto) ABooleano() (bool, error) { return o.valor != nil, nil }
func (o *Objeto) AEntero() (int, error)    { return 0, nil }
func (o *Objeto) AReal() (float64, error)  { return 0, nil }
