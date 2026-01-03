package evaluador

import (
    "strings"
    "sync"

    "nepa/desarrollo/interno/core"
    "nepa/desarrollo/interno/parser"
)

type SolicitudEjecutar struct {
    Codigo     string
    Argumentos []interface{}
    Archivo    string
    Mensaje    string
}

func (s SolicitudEjecutar) Error() string {
    return s.Mensaje
}

type Handler func(parser.Nodo, *Contexto)

var manejadores = make(map[string]Handler)
var mu sync.RWMutex

func Registrar(tipo string, fn Handler) {
    mu.Lock()
    defer mu.Unlock()
    manejadores[tipo] = fn
}

func EjecutarConContexto(ast []parser.Nodo, args map[string]interface{},
    globales map[string]interface{}, constantes map[string]interface{},
    archivo string) (map[string]interface{}, error) {

    ctx := &Contexto{
        Variables:  map[string]interface{}{},
        Globales:   globales,
        Constantes: constantes,
        Funciones:  map[string]func(...interface{}) interface{}{},
    }

    for k, v := range args {
        ctx.Variables[k] = v
    }

    resultados := map[string]interface{}{}

    for i, nodo := range ast {
        linea := i + 1 // número de línea (1-based)

        if nodo.Tipo == "llamada" {
            nombre := strings.ToLower(nodo.Nombre)
            f, ok := Funciones[nombre]
            if ok {
                if _, err := f(nodo.Args...); err != nil {
                    return nil, core.ErrorEjecucion{
                        Archivo: archivo,
                        Linea:   linea,
                        Mensaje: "ejecución fallida en '" + nodo.Nombre + "' → " + err.Error(),
                    }
                }
            } else {
                return nil, core.ErrorEjecucion{
                    Archivo: archivo,
                    Linea:   linea,
                    Mensaje: "instrucción no reconocida '" + nodo.Nombre + "'",
                }
            }
            continue
        }

        mu.RLock()
        handler, ok := manejadores[nodo.Tipo]
        mu.RUnlock()
        if ok {
            handler(nodo, ctx)
        } else {
            return nil, core.ErrorEjecucion{
                Archivo: archivo,
                Linea:   linea,
                Mensaje: "tipo de nodo no reconocido '" + nodo.Tipo + "'",
            }
        }
    }

    for k, v := range ctx.Variables {
        resultados[k] = v
    }

    return resultados, nil
}
