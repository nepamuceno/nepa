package complejo

import (
	"fmt"
	"strings"
	"sync"

	"nepa/desarrollo/interno/administrador"
)

type Complejo struct {
	mu     sync.RWMutex
	nombre string
	valor  complex128
}

func CrearComplejo(nombre string, v interface{}) (administrador.Variable, error) {
	c := &Complejo{
		nombre: strings.TrimSpace(nombre),
		valor:  0 + 0i,
	}
	if v != nil {
		if err := c.AsignarDesdeInterface(v); err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *Complejo) Nombre() string { return c.nombre }
func (c *Complejo) Tipo() string   { return "complejo" }

func (c *Complejo) Mostrar() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return fmt.Sprintf("%s:%s=%v", c.Tipo(), c.nombre, c.valor)
}

func (c *Complejo) AsignarDesdeInterface(v interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock() // Corregido de 'd' a 'c'
	if v == nil {
		c.valor = 0 + 0i
		return nil
	}
	if val, ok := v.(complex128); ok {
		c.valor = val
		return nil
	}
	return fmt.Errorf("❌ valor no compatible con tipo complejo")
}

func (c *Complejo) ValorComoInterface() interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.valor
}

func (c *Complejo) JSON() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return fmt.Sprintf(`{"tipo":"complejo","nombre":"%s","real":%g,"imag":%g}`, c.nombre, real(c.valor), imag(c.valor))
}

func (c *Complejo) ABooleano() (bool, error) { 
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.valor != 0, nil 
}

func (c *Complejo) AEntero() (int, error) { 
	c.mu.RLock()
	defer c.mu.RUnlock()
	return int(real(c.valor)), nil 
}

func (c *Complejo) AReal() (float64, error) { 
	c.mu.RLock()
	defer c.mu.RUnlock()
	return real(c.valor), nil 
}
