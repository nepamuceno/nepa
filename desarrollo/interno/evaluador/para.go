package evaluador

import (
	"fmt"
	"os"
	"nepa/desarrollo/interno/ast"
	"reflect"
)

// EjecutarPara es el motor universal de iteración de Nepa.
// Soporta: Rangos (C++), Colecciones (Python), Cadenas (Bash) y Objetos.
func (i *Interpretador) EjecutarPara(p ast.Para) interface{} {
	
	// CASO 1: Rango Numérico (Desde/Hasta/Incremento) - Estilo C++/Fortran
	if p.Fin != nil {
		return i.ejecutarRangoUniversal(p)
	}

	// CASO 2: Iteración de Contenido (En) - Estilo Python/Bash
	origen := i.Evaluar(p.Origen)
	if origen == nil {
		return nil
	}

	// --- NUEVO: SOPORTE PARA CORE_PARA (Iteradores de Archivos/Sistemas) ---
	if it, ok := origen.(Iterador); ok {
		// defer it.Cerrar() // Opcional: Descomenta si quieres cierre automático
		for {
			valor, siguiente := it.Siguiente()
			if !siguiente {
				break
			}
			i.Entorno[p.Variable] = valor
			res := i.ejecutarCuerpoUniversal(p.Cuerpo)
			// Solo propagamos si es un NepaRetorno real
			if _, ok := res.(NepaRetorno); ok { return res } 
		}
		return nil
	}

	// ROBUSTEZ: Indirect resuelve punteros o interfaces para llegar al valor real
	val := reflect.Indirect(reflect.ValueOf(origen))

	switch val.Kind() {
	case reflect.String:
		// Itera letra por letra
		for _, char := range val.String() {
			i.Entorno[p.Variable] = string(char)
			res := i.ejecutarCuerpoUniversal(p.Cuerpo)
			if _, ok := res.(NepaRetorno); ok { return res } 
		}

	case reflect.Slice, reflect.Array:
		// Itera matrices de cualquier tipo [1, 2, 3] o ["a", "b"]
		for idx := 0; idx < val.Len(); idx++ {
			i.Entorno[p.Variable] = val.Index(idx).Interface()
			res := i.ejecutarCuerpoUniversal(p.Cuerpo)
			if _, ok := res.(NepaRetorno); ok { return res } 
		}

	case reflect.Map:
		// Itera sobre llaves de un diccionario/objeto
		iter := val.MapRange()
		for iter.Next() {
			// Por defecto entregamos la llave (estilo Python/JS)
			i.Entorno[p.Variable] = iter.Key().Interface()
			res := i.ejecutarCuerpoUniversal(p.Cuerpo)
			if _, ok := res.(NepaRetorno); ok { return res } 
		}

	default:
		// ROBUSTEZ TOTAL: Si es un objeto único (Int, Bool, etc), se procesa una vez
		i.Entorno[p.Variable] = origen
		res := i.ejecutarCuerpoUniversal(p.Cuerpo)
		if _, ok := res.(NepaRetorno); ok { return res }
		return nil
	}

	return nil
}

// ejecutarRangoUniversal maneja cálculos complejos y saltos dinámicos
func (i *Interpretador) ejecutarRangoUniversal(p ast.Para) interface{} {
	// Evaluación de expresiones complejas iniciales
	inicio, ok1 := toFloatSafe(i.Evaluar(p.Origen))
	fin, ok2 := toFloatSafe(i.Evaluar(p.Fin))

	// ESTRICTO: Si los límites no son números, abortamos inmediatamente
	if !ok1 || !ok2 {
		fmt.Printf("❌ Error de Ejecución Nepa: Los límites 'desde/hasta' deben ser numéricos. (Línea aproximada: %d)\n", i.Entorno["__linea__"])
		os.Exit(1)
	}

	actual := inicio
	paso := 1.0 // Paso por defecto
	if inicio > fin { paso = -1.0 } // Auto-detección de dirección regresiva

	for {
		// 1. Recalcular paso y meta (Dinamismo Total en cada vuelta)
		if vFin, okF := toFloatSafe(i.Evaluar(p.Fin)); okF {
			fin = vFin
		}
		
		if p.Incremento != nil {
			vPaso, okP := toFloatSafe(i.Evaluar(p.Incremento))
			if !okP {
				fmt.Println("❌ Error de Ejecución Nepa: El incremento del bucle 'para' debe ser un número.")
				os.Exit(1)
			}
			paso = vPaso
		}

		// 2. Protección contra bucle infinito (paso cero)
		// ESTRICTO: En Nepa, un incremento de cero es un error fatal, no una advertencia.
		if paso == 0 {
			fmt.Println("❌ Error de Ejecución Nepa: Incremento cero detectado. Bucle infinito prohibido.")
			os.Exit(1)
		}

		// 3. Verificación de límite UNIVERSAL (Bidireccional)
		if (paso > 0 && actual > fin) || (paso < 0 && actual < fin) {
			break
		}

		// 4. Inyectar valor en el entorno del intérprete
		// Importante: p.Variable debe existir y ser un identificador válido
		i.Entorno[p.Variable] = actual

		// 5. Ejecutar cuerpo del bucle
		res := i.ejecutarCuerpoUniversal(p.Cuerpo)
		
		// Propagación estricta de retornos
		if _, ok := res.(NepaRetorno); ok {
			return res
		}

		// 6. Aplicar incremento para la siguiente iteración
		actual += paso
	}
	
	return nil
}

// ejecutarCuerpoUniversal es una función de apoyo para manejar cualquier tipo de nodo en el cuerpo
func (i *Interpretador) ejecutarCuerpoUniversal(cuerpo interface{}) interface{} {
	switch c := cuerpo.(type) {
	case []ast.Nodo:
		return i.ejecutarBloque(c)
	case ast.Nodo:
		// En Nepa con indentación, el cuerpo suele ser []ast.Nodo, 
		// pero mantenemos esto por compatibilidad con expresiones simples.
		return i.Evaluar(c)
	default:
		return nil
	}
}

// ejecutarBloque procesa la lista de instrucciones internas del bucle
func (i *Interpretador) ejecutarBloque(nodos []ast.Nodo) interface{} {
	var ultimo interface{}
	for _, nodo := range nodos {
		// Evaluación secuencial de cada nodo del AST
		ultimo = i.Evaluar(nodo)
		
		// Manejo de sentencias de interrupción o retorno (Propagación)
		if _, ok := ultimo.(NepaRetorno); ok {
			return ultimo
		}
	}
	return ultimo
}
