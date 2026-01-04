package caracter

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"nepa/desarrollo/interno/administrador"
)

type Caracter struct {
	mu     sync.RWMutex
	nombre string
	valor  rune
}

func CrearCaracter(nombre string, v interface{}) (administrador.Variable, error) {
	ca := &Caracter{
		nombre: strings.TrimSpace(nombre),
		valor:  0, // Caracter nulo por defecto
	}
	if v != nil {
		if err := ca.AsignarDesdeInterface(v); err != nil {
			return nil, err
		}
	}
	return ca, nil
}

func (ca *Caracter) Nombre() string { return ca.nombre }
func (ca *Caracter) Tipo() string   { return "caracter" }

func (ca *Caracter) Mostrar() string {
	ca.mu.RLock()
	defer ca.mu.RUnlock()
	return fmt.Sprintf("%s:%s='%c'", ca.Tipo(), ca.nombre, ca.valor)
}

func (ca *Caracter) AsignarDesdeInterface(v interface{}) error {
	ca.mu.Lock()
	defer ca.mu.Unlock()
	if v == nil {
		ca.valor = 0
		return nil
	}
	s := fmt.Sprint(v)
	runes := []rune(s)
	if len(runes) == 0 {
		ca.valor = 0
		return nil
	}
	ca.valor = runes[0]
	return nil
}

func (ca *Caracter) ValorComoInterface() interface{} {
	ca.mu.RLock()
	defer ca.mu.RUnlock()
	return ca.valor
}

func (ca *Caracter) JSON() string {
	ca.mu.RLock()
	defer ca.mu.RUnlock()
	data, _ := json.Marshal(map[string]interface{}{
		"tipo": ca.Tipo(), "nombre": ca.nombre, "valor": string(ca.valor),
	})
	return string(data)
}

func (ca *Caracter) ABooleano() (bool, error) { return ca.valor != 0, nil }
func (ca *Caracter) AEntero() (int, error)    { return int(ca.valor), nil }
func (ca *Caracter) AReal() (float64, error)  { return float64(ca.valor), nil }
