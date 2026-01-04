package puntero

import (
	"fmt"
	"strings"
	"sync"
	"nepa/desarrollo/interno/administrador"
)

type Puntero struct {
	mu       sync.RWMutex
	nombre   string
	objetivo administrador.Variable
}

func CrearPuntero(nombre string, v interface{}) (administrador.Variable, error) {
	p := &Puntero{
		nombre: strings.TrimSpace(nombre),
	}
	if v != nil {
		p.AsignarDesdeInterface(v)
	}
	return p, nil
}

func (p *Puntero) Nombre() string { return p.nombre }
func (p *Puntero) Tipo() string   { return "puntero" }

func (p *Puntero) Mostrar() string {
	p.mu.RLock()
	defer p.mu.RUnlock()
	if p.objetivo == nil {
		return fmt.Sprintf("%s:%s=nulo", p.Tipo(), p.nombre)
	}
	return fmt.Sprintf("%s:%s→(%s)", p.Tipo(), p.nombre, p.objetivo.Nombre())
}

func (p *Puntero) AsignarDesdeInterface(v interface{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if v == nil {
		p.objetivo = nil
		return nil
	}
	if target, ok := v.(administrador.Variable); ok {
		p.objetivo = target
		return nil
	}
	return fmt.Errorf("❌ un puntero solo puede apuntar a otra Variable")
}

func (p *Puntero) ValorComoInterface() interface{} {
	p.mu.RLock()
	defer p.mu.RUnlock()
	if p.objetivo == nil { return nil }
	return p.objetivo.ValorComoInterface()
}

func (p *Puntero) JSON() string {
	return fmt.Sprintf(`{"tipo":"puntero","nombre":"%s"}`, p.nombre)
}

func (p *Puntero) ABooleano() (bool, error) { return p.objetivo != nil, nil }
func (p *Puntero) AEntero() (int, error)    { if p.objetivo == nil {return 0,nil}; return p.objetivo.AEntero() }
func (p *Puntero) AReal() (float64, error)  { if p.objetivo == nil {return 0,nil}; return p.objetivo.AReal() }
