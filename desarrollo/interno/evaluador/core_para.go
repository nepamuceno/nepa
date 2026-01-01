package evaluador

import (
	"bufio"
	"os"
)

// Iterador representa la interfaz que cualquier objeto complejo 
// (archivos, sockets, bases de datos) debe cumplir para ser usado en un 'para'
type Iterador interface {
	Siguiente() (interface{}, bool)
	Cerrar() error
}

// --- Soporte para Archivos ---

type ArchivoIterador struct {
	file    *os.File
	scanner *bufio.Scanner
}

func (ai *ArchivoIterador) Siguiente() (interface{}, bool) {
	if ai.scanner.Scan() {
		return ai.scanner.Text(), true
	}
	return nil, false
}

func (ai *ArchivoIterador) Cerrar() error {
	return ai.file.Close()
}

// --- Soporte para Rangos Infinitos o Generadores ---

type GeneradorRango struct {
	Actual float64
	Paso   float64
}

func (gr *GeneradorRango) Siguiente() (interface{}, bool) {
	val := gr.Actual
	gr.Actual += gr.Paso
	return val, true // Siempre hay un siguiente valor
}

// Para que GeneradorRango cumpla con la interfaz Iterador y no rompa el 'para'
func (gr *GeneradorRango) Cerrar() error {
	return nil
}

// --- Funciones de Utilidad Core ---

// ConvertirAIterador ahora usa *Interpretador para coincidir con tu evaluador.go
func (i *Interpretador) ConvertirAIterador(objeto interface{}) Iterador {
	switch v := objeto.(type) {
	case *os.File:
		return &ArchivoIterador{
			file:    v,
			scanner: bufio.NewScanner(v),
		}
	case Iterador:
		return v
	}
	return nil
}

// EsIterable verifica si el dato requiere un tratamiento especial de core
func EsIterable(objeto interface{}) bool {
	switch objeto.(type) {
	// Añadimos interfaces genéricas para que coincida con lo que el Evaluador produce
	// Se mantienen string, []interface{} (listas) y map[string]interface{} (objetos)
	case string, []interface{}, map[string]interface{}, Iterador, *os.File:
		return true
	}
	return false
}
