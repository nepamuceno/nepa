package texto

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"nepa/desarrollo/interno/administrador"
)

type Texto struct {
	mu     sync.RWMutex
	nombre string
	valor  string
}

func CrearTexto(nombre string, v interface{}) (administrador.Variable, error) {
	t := &Texto{
		nombre: strings.TrimSpace(nombre),
		valor:  "", // Default seguro
	}
	if v != nil {
		if err := t.AsignarDesdeInterface(v); err != nil {
			return nil, err
		}
	}
	return t, nil
}

func (t *Texto) Nombre() string { return t.nombre }
func (t *Texto) Tipo() string   { return "texto" }

func (t *Texto) Mostrar() string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	res := t.valor
	if len(res) > 50 {
		res = res[:47] + "..."
	}
	return fmt.Sprintf("%s:%s=\"%s\"", t.Tipo(), t.nombre, res)
}

func (t *Texto) AsignarDesdeInterface(v interface{}) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	if v == nil {
		t.valor = ""
		return nil
	}
	t.valor = fmt.Sprint(v)
	return nil
}

func (t *Texto) ValorComoInterface() interface{} {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.valor
}

func (t *Texto) JSON() string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	data, _ := json.Marshal(map[string]interface{}{
		"tipo": t.Tipo(), "nombre": t.nombre, "valor": t.valor,
	})
	return string(data)
}

func (t *Texto) ABooleano() (bool, error) { return len(t.valor) > 0, nil }
func (t *Texto) AEntero() (int, error)    { return 0, fmt.Errorf("❌ conversión no soportada") }
func (t *Texto) AReal() (float64, error)  { return 0.0, fmt.Errorf("❌ conversión no soportada") }
