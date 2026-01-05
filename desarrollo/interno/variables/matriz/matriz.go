package matriz

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"nepa/desarrollo/interno/administrador"
	"nepa/desarrollo/interno/evaluador"
)

type Matriz struct {
	mu     sync.RWMutex
	nombre string
	valor  [][]float64
}

func CrearMatriz(nombre string, v interface{}) (administrador.Variable, error) {
	m := &Matriz{
		nombre: strings.TrimSpace(nombre),
		valor:  [][]float64{},
	}
	if v != nil {
		if err := m.AsignarDesdeInterface(v); err != nil {
			return nil, err
		}
	}
	return m, nil
}

func (m *Matriz) Nombre() string { return m.nombre }
func (m *Matriz) Tipo() string   { return "matriz" }

func (m *Matriz) Mostrar() string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return fmt.Sprintf("%v", m.valor)
}

func (m *Matriz) AsignarDesdeInterface(v interface{}) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	datosPlanos, err := m.aplanarRecursivo(v)
	if err != nil {
		return err
	}

	if len(datosPlanos) == 0 {
		m.valor = [][]float64{}
		return nil
	}

	// Lógica de organización de dimensiones
	if len(datosPlanos) == 4 {
		m.valor = [][]float64{datosPlanos[:2], datosPlanos[2:]}
	} else {
		m.valor = [][]float64{datosPlanos}
	}

	return nil
}

func (m *Matriz) aplanarRecursivo(v interface{}) ([]float64, error) {
	var resultado []float64

	rv := reflect.ValueOf(v)
	if !rv.IsValid() {
		return resultado, nil
	}

	switch rv.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < rv.Len(); i++ {
			res, err := m.aplanarRecursivo(rv.Index(i).Interface())
			if err != nil {
				return nil, err
			}
			resultado = append(resultado, res...)
		}

	case reflect.String:
		raw := rv.String()
		// 1. Limpieza de escombros: removemos los corchetes que separan filas
		// Convertimos "[[a,b],[c,d]]" en "a,b,c,d" protegiendo funciones
		elementos := m.segmentarLimpiandoCorchetes(raw)
		
		for _, s := range elementos {
			s = strings.TrimSpace(s)
			if s == "" {
				continue
			}

			res := evaluador.ResolverEstructuraRecursiva(s, nil)
			val, err := evaluador.ConvertirAReal(res)
			if err != nil {
				return nil, fmt.Errorf("❌ ERROR FATAL: no se pudo convertir la cadena a número → %s", s)
			}
			resultado = append(resultado, val)
		}

	default:
		val, err := evaluador.ConvertirAReal(v)
		if err != nil {
			return nil, err
		}
		resultado = append(resultado, val)
	}

	return resultado, nil
}

// segmentarLimpiandoCorchetes ignora los corchetes al separar por comas
func (m *Matriz) segmentarLimpiandoCorchetes(raw string) []string {
	var pars []string
	var buf strings.Builder
	nivelParentesis := 0

	for _, r := range raw {
		switch r {
		case '(':
			nivelParentesis++
			buf.WriteRune(r)
		case ')':
			nivelParentesis--
			buf.WriteRune(r)
		case '[', ']':
			// Simplemente ignoramos los corchetes, no los metemos al buffer
			continue
		case ',':
			if nivelParentesis == 0 {
				pars = append(pars, buf.String())
				buf.Reset()
				continue
			}
			buf.WriteRune(r)
		default:
			buf.WriteRune(r)
		}
	}
	if buf.Len() > 0 {
		pars = append(pars, buf.String())
	}
	return pars
}

func (m *Matriz) ValorComoInterface() interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.valor
}

func (m *Matriz) JSON() string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return fmt.Sprintf(`{"tipo":"matriz","nombre":"%s","datos":%v}`, m.nombre, m.valor)
}

func (m *Matriz) ABooleano() (bool, error) { return len(m.valor) > 0, nil }
func (m *Matriz) AEntero() (int, error)    { return len(m.valor), nil }
func (m *Matriz) AReal() (float64, error)  { return float64(len(m.valor)), nil }
