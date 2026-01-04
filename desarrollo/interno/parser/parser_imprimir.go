package parser

import (
	"strings"
)

// parseLlamada maneja funciones como imprimir(...) o llamadas directas
func parseLlamada(linea string) *Nodo {
	if strings.Contains(linea, "(") && strings.HasSuffix(linea, ")") {
		fn := strings.TrimSpace(linea[:strings.Index(linea, "(")])
		args := extraerArgs(linea)
		return &Nodo{Tipo: "llamada", Nombre: fn, Args: args}
	}
	if strings.Contains(linea, " ") {
		partes := strings.SplitN(linea, " ", 2)
		fn := strings.TrimSpace(partes[0])
		resto := strings.TrimSpace(partes[1])
		args := []interface{}{}
		for _, p := range strings.Split(resto, ",") {
			arg := strings.TrimSpace(p)
			if arg != "" {
				args = append(args, parseValor(arg))
			}
		}
		return &Nodo{Tipo: "llamada", Nombre: fn, Args: args}
	}
	return nil
}
