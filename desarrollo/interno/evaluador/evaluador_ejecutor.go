package evaluador

import (
	"fmt"
	"strings"
	"sync"

	"nepa/desarrollo/interno/administrador"
	"nepa/desarrollo/interno/parser"
)

// SolicitudEjecutar representa la orden de trabajo para el intérprete.
type SolicitudEjecutar struct {
	Codigo     string
	Argumentos []interface{}
	Archivo    string
	Mensaje    string
}

func (s SolicitudEjecutar) Error() string {
	return s.Mensaje
}

// Manejador es la función que procesa un tipo de nodo específico (ej: asignar, si, mientras).
type Manejador func(parser.Nodo, *Contexto)

var (
	manejadores = make(map[string]Manejador)
	mu          sync.RWMutex
)

// Registrar vincula un tipo de nodo del parser con un manejador del evaluador.
func Registrar(tipo string, fn Manejador) {
	mu.Lock()
	defer mu.Unlock()
	manejadores[tipo] = fn
}

// EjecutarConContexto recorre el AST y ejecuta cada instrucción en orden.
func EjecutarConContexto(ast []parser.Nodo, args map[string]interface{},
	globales map[string]interface{}, constantes map[string]interface{},
	archivo string) (map[string]interface{}, error) {

	// 1. Inicializamos el contexto con la memoria compartida
	ctx := &Contexto{
		Variables:  map[string]interface{}{},
		Globales:   globales,
		Constantes: constantes,
		// Aquí podríamos usar RegistrarFuncionesEstandar(ctx) si fuera necesario
		Funciones:  map[string]func(...interface{}) interface{}{}, 
	}

	// 2. Cargamos los argumentos iniciales
	for k, v := range args {
		ctx.Variables[k] = v
	}

	// 3. Procesamiento línea por línea
	for i, nodo := range ast {
		linea := i + 1 

		// Si es una llamada directa (instrucción de Nepa)
		if nodo.Tipo == "llamada" {
			nombre := strings.ToLower(nodo.Nombre)
			
			// Buscamos en el registro global de funciones que pulimos antes
			f, existe := Funciones[nombre]
			if existe {
				// Convertimos los argumentos del nodo (que son strings/interfaces)
				// a valores reales antes de pasar a la función
				if _, err := f(nodo.Args...); err != nil {
					return nil, fmt.Errorf("%s:%d: fallo en '%s' → %v",
						archivo, linea, nodo.Nombre, err)
				}
			} else {
				return nil, fmt.Errorf("%s:%d: instrucción no reconocida '%s'",
					archivo, linea, nodo.Nombre)
			}
			continue
		}

		// Si es un nodo registrado (como una expresión compleja, un bucle, etc.)
		mu.RLock()
		manejador, ok := manejadores[nodo.Tipo]
		mu.RUnlock()

		if ok {
			manejador(nodo, ctx)
		} else {
			return nil, fmt.Errorf("%s:%d: tipo de instrucción no soportado '%s'",
				archivo, linea, nodo.Tipo)
		}
	}

	// 4. Recolección de resultados finales
	resultados := map[string]interface{}{}
	for k, v := range ctx.Variables {
		if varObj, ok := v.(administrador.Variable); ok {
			resultados[k] = varObj.Mostrar()
		} else {
			resultados[k] = v
		}
	}

	return resultados, nil
}
