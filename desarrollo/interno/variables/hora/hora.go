package hora

import (
	"fmt"
	"strings"
	"sync"
	"time"
	"nepa/desarrollo/interno/administrador"
)

type Hora struct {
	mu     sync.RWMutex
	nombre string
	valor  time.Time
}

func CrearHora(nombre string, v interface{}) (administrador.Variable, error) {
	h := &Hora{
		nombre: strings.TrimSpace(nombre),
		valor:  time.Now(),
	}
	if v != nil {
		h.AsignarDesdeInterface(v)
	}
	return h, nil
}

func (h *Hora) Nombre() string { return h.nombre }
func (h *Hora) Tipo() string   { return "hora" }

func (h *Hora) Mostrar() string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return fmt.Sprintf("%s:%s=%s", h.Tipo(), h.nombre, h.valor.Format("15:04:05"))
}

func (h *Hora) AsignarDesdeInterface(v interface{}) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	if v == nil {
		h.valor = time.Now()
		return nil
	}
	s := fmt.Sprint(v)
	parsed, err := time.Parse("15:04:05", s)
	if err != nil {
		return fmt.Errorf("❌ formato de hora inválido: use HH:MM:SS")
	}
	h.valor = parsed
	return nil
}

func (h *Hora) ValorComoInterface() interface{} {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.valor
}

func (h *Hora) JSON() string {
	return fmt.Sprintf(`{"tipo":"hora","nombre":"%s","valor":"%s"}`, h.nombre, h.valor.Format("15:04:05"))
}

func (h *Hora) ABooleano() (bool, error) { return true, nil }
func (h *Hora) AEntero() (int, error)    { return h.valor.Hour()*3600 + h.valor.Minute()*60 + h.valor.Second(), nil }
func (h *Hora) AReal() (float64, error)  { e, _ := h.AEntero(); return float64(e), nil }
