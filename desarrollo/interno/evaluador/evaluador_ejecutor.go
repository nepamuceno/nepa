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

	// 1. Inicializamos el contexto
	ctx := &Contexto{
		Variables:  map[string]interface{}{},
		Globales:   globales,
		Constantes: constantes,
		Funciones:  map[string]func(...interface{}) interface{}{},
	}

	for k, v := range args {
		ctx.Variables[k] = v
	}

	// 2. Procesamiento línea por línea
	for i := range ast {
		linea := i + 1
		nodo := &ast[i] 

		// --- CASO A: LLAMADAS DIRECTAS ---
		if nodo.Tipo == "llamada" {
			nombre := strings.ToLower(nodo.Nombre)
			f, existe := Funciones[nombre]
			if existe {
				argsResueltos := make([]interface{}, len(nodo.Args))
				for idx, argRaw := range nodo.Args {
					argsResueltos[idx] = ResolverEstructuraRecursiva(argRaw, ctx)
				}

				if _, err := f(argsResueltos...); err != nil {
					return nil, fmt.Errorf("%s:%d: fallo en '%s' → %v",
						archivo, linea, nodo.Nombre, err)
				}
			} else {
				return nil, fmt.Errorf("%s:%d: instrucción no reconocida '%s'",
					archivo, linea, nodo.Nombre)
			}
			continue
		}

		// --- CASO B: NODOS REGISTRADOS ---
		mu.RLock()
		manejador, ok := manejadores[nodo.Tipo]
		mu.RUnlock()

		if ok {
			// --- RESOLUCIÓN PREVENTIVA ---
			// Extraemos el valor como string para analizarlo
			if nodo.Valor != nil {
				valorStr := fmt.Sprintf("%v", nodo.Valor)

				// Si contiene una llamada a función como "promedio(" o "binario("
				if strings.Contains(valorStr, "(") {
					res := ResolverEstructuraRecursiva(nodo.Valor, ctx)
					
					// Actualizamos el nodo con el resultado real (número o matriz)
					nodo.Valor = res

					// Inyectamos en el administrador para saltar errores de tipo string
					if nodo.Nombre != "" {
						if v, err := administrador.CrearVariableUniversal("", nodo.Nombre, res); err == nil {
							_ = administrador.RegistrarVariable(nodo.Nombre, v)
							ctx.Variables[nodo.Nombre] = v
						}
					}
				}
			}

			manejador(*nodo, ctx)
		} else {
			return nil, fmt.Errorf("%s:%d: tipo de instrucción no soportado '%s'",
				archivo, linea, nodo.Tipo)
		}
	}

	// 3. Recolección de resultados
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
