package fecha

import (
	"fmt"
	"strings"
	"sync"
	"time"
	"nepa/desarrollo/interno/administrador"
)

type Fecha struct {
	mu     sync.RWMutex
	nombre string
	valor  time.Time
}

func CrearFecha(nombre string, v interface{}) (administrador.Variable, error) {
	f := &Fecha{
		nombre: strings.TrimSpace(nombre),
		valor:  time.Now(),
	}
	if v != nil {
		f.AsignarDesdeInterface(v)
	}
	return f, nil
}

func (f *Fecha) Nombre() string { return f.nombre }
func (f *Fecha) Tipo() string   { return "fecha" }

func (f *Fecha) Mostrar() string {
	f.mu.RLock()
	defer f.mu.RUnlock()
	return fmt.Sprintf("%s:%s=%s", f.Tipo(), f.nombre, f.valor.Format("2006-01-02"))
}

func (f *Fecha) AsignarDesdeInterface(v interface{}) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if v == nil {
		f.valor = time.Now()
		return nil
	}
	if t, ok := v.(time.Time); ok {
		f.valor = t
		return nil
	}
	parsed, err := time.Parse("2006-01-02", fmt.Sprint(v))
	if err != nil { return err }
	f.valor = parsed
	return nil
}

func (f *Fecha) ValorComoInterface() interface{} {
	f.mu.RLock()
	defer f.mu.RUnlock()
	return f.valor
}

func (f *Fecha) JSON() string {
	return fmt.Sprintf(`{"tipo":"fecha","nombre":"%s","valor":"%s"}`, f.nombre, f.valor.Format("2006-01-02"))
}

func (f *Fecha) ABooleano() (bool, error) { return !f.valor.IsZero(), nil }
func (f *Fecha) AEntero() (int, error)    { return int(f.valor.Unix()), nil }
func (f *Fecha) AReal() (float64, error)  { return float64(f.valor.Unix()), nil }
