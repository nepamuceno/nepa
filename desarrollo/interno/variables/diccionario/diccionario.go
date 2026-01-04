package diccionario

import (
	"fmt"
	"strings"
	"sync"

	"nepa/desarrollo/interno/administrador"
)

type Diccionario struct {
	mu     sync.RWMutex
	nombre string
	valor  map[string]interface{}
}

func CrearDiccionario(nombre string, v interface{}) (administrador.Variable, error) {
	d := &Diccionario{
		nombre: strings.TrimSpace(nombre),
		valor:  make(map[string]interface{}), // Blindaje: mapa inicializado, nunca nil
	}
	if v != nil {
		if err := d.AsignarDesdeInterface(v); err != nil {
			return nil, err
		}
	}
	return d, nil
}

func (d *Diccionario) Nombre() string { return d.nombre }
func (d *Diccionario) Tipo() string   { return "diccionario" }

func (d *Diccionario) Mostrar() string {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return fmt.Sprintf("%s:%s=%v", d.Tipo(), d.nombre, d.valor)
}

func (d *Diccionario) AsignarDesdeInterface(v interface{}) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if v == nil {
		d.valor = make(map[string]interface{})
		return nil
	}

	if val, ok := v.(map[string]interface{}); ok {
		d.valor = val
		return nil
	}
	return fmt.Errorf("❌ el valor debe ser un mapa para asignar a Diccionario")
}

func (d *Diccionario) ValorComoInterface() interface{} {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.valor
}

func (d *Diccionario) JSON() string {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return fmt.Sprintf(`{"tipo":"diccionario","nombre":"%s","elementos":%d}`, d.nombre, len(d.valor))
}

// Métodos de Interfaz Obligatorios
func (d *Diccionario) ABooleano() (bool, error) { return len(d.valor) > 0, nil }
func (d *Diccionario) AEntero() (int, error)    { return len(d.valor), nil }
func (d *Diccionario) AReal() (float64, error)  { return float64(len(d.valor)), nil }
