package evaluador

import (
    "nepa/desarrollo/interno/parser"
)

// SolicitudEjecutar representa una petición de ejecución desde main.go.
// Implementa error para poder usarse en type assertions.
type SolicitudEjecutar struct {
    Codigo     string
    Argumentos []interface{}
    Archivo    string
    Mensaje    string
}

func (s SolicitudEjecutar) Error() string {
    return s.Mensaje
}

// EjecutarConContexto ejecuta código dentro de un Contexto dado.
// Ajustado para recibir []parser.Nodo como espera main.go.
func EjecutarConContexto(nodos []parser.Nodo, globales, constantes, funciones map[string]interface{}) (map[string]interface{}, error) {
    // TODO: implementar ejecución real con nodos del parser
    resultados := make(map[string]interface{})
    resultados["estado"] = "ejecución no implementada"
    return resultados, nil
}
