package cadena

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"nepa/desarrollo/interno/administrador"
)

const MaxLenCadena = 255

type Cadena struct {
	mu     sync.RWMutex
	nombre string
	valor  string
}

func CrearCadena(nombre string, v interface{}) (administrador.Variable, error) {
	c := &Cadena{
		nombre: strings.TrimSpace(nombre),
		valor:  "", // Default seguro
	}
	if v != nil {
		if err := c.AsignarDesdeInterface(v); err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *Cadena) Nombre() string { return c.nombre }
func (c *Cadena) Tipo() string   { return "cadena" }

func (c *Cadena) Mostrar() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return fmt.Sprintf("%s:%s='%s'", c.Tipo(), c.nombre, c.valor)
}

func (c *Cadena) AsignarDesdeInterface(v interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if v == nil {
		c.valor = ""
		return nil
	}
	s := fmt.Sprint(v)
	if len(s) > MaxLenCadena {
		return fmt.Errorf("❌ cadena excede el límite de %d caracteres", MaxLenCadena)
	}
	c.valor = s
	return nil
}

func (c *Cadena) ValorComoInterface() interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.valor
}

func (c *Cadena) JSON() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	data, _ := json.Marshal(map[string]interface{}{
		"tipo": c.Tipo(), "nombre": c.nombre, "valor": c.valor,
	})
	return string(data)
}

func (c *Cadena) ABooleano() (bool, error) { return len(c.valor) > 0, nil }
func (c *Cadena) AEntero() (int, error)    { return 0, fmt.Errorf("❌ conversión no soportada") }
func (c *Cadena) AReal() (float64, error)  { return 0.0, fmt.Errorf("❌ conversión no soportada") }
