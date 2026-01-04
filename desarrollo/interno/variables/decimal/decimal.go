package decimal

import (
	"fmt"
	"strings"
	"sync"
	"nepa/desarrollo/interno/administrador"
	"nepa/desarrollo/interno/evaluador"
)

type Decimal struct {
	mu     sync.RWMutex
	nombre string
	valor  float64
}

func CrearDecimal(nombre string, v interface{}) (administrador.Variable, error) {
	d := &Decimal{
		nombre: strings.TrimSpace(nombre),
		valor:  0.00,
	}
	if v != nil {
		d.AsignarDesdeInterface(v)
	}
	return d, nil
}

func (d *Decimal) Nombre() string { return d.nombre }
func (d *Decimal) Tipo() string   { return "decimal" }

func (d *Decimal) Mostrar() string {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return fmt.Sprintf("%s:%s=%.2f", d.Tipo(), d.nombre, d.valor)
}

func (d *Decimal) AsignarDesdeInterface(v interface{}) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if v == nil {
		d.valor = 0.0
		return nil
	}
	val, err := evaluador.ConvertirAReal(v)
	if err != nil { return err }
	d.valor = val
	return nil
}

func (d *Decimal) ValorComoInterface() interface{} {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.valor
}

func (d *Decimal) JSON() string {
	return fmt.Sprintf(`{"tipo":"decimal","nombre":"%s","valor":%.2f}`, d.nombre, d.valor)
}

func (d *Decimal) ABooleano() (bool, error) { return d.valor != 0, nil }
func (d *Decimal) AEntero() (int, error)    { return int(d.valor), nil }
func (d *Decimal) AReal() (float64, error)  { return d.valor, nil }
