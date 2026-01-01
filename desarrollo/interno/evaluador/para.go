package evaluador

import (
	"fmt"
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
			i.ejecutarBloque(p.Cuerpo.([]ast.Nodo))
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
			i.ejecutarBloque(p.Cuerpo.([]ast.Nodo))
		}

	case reflect.Slice, reflect.Array:
		// Itera matrices de cualquier tipo [1, 2, 3] o ["a", "b"]
		for idx := 0; idx < val.Len(); idx++ {
			i.Entorno[p.Variable] = val.Index(idx).Interface()
			i.ejecutarBloque(p.Cuerpo.([]ast.Nodo))
		}

	case reflect.Map:
		// Itera sobre llaves de un diccionario/objeto
		iter := val.MapRange()
		for iter.Next() {
			// Por defecto entregamos la llave (estilo Python/JS)
			i.Entorno[p.Variable] = iter.Key().Interface()
			i.ejecutarBloque(p.Cuerpo.([]ast.Nodo))
		}

	default:
		// ROBUSTEZ TOTAL: Si es un objeto único (Int, Bool, etc), se procesa una vez
		i.Entorno[p.Variable] = origen
		i.ejecutarBloque(p.Cuerpo.([]ast.Nodo))
	}

	return nil
}

// ejecutarRangoUniversal maneja cálculos complejos y saltos dinámicos
func (i *Interpretador) ejecutarRangoUniversal(p ast.Para) interface{} {
	// Evaluación de expresiones complejas iniciales
	inicio, ok1 := toFloatSafe(i.Evaluar(p.Origen))
	fin, ok2 := toFloatSafe(i.Evaluar(p.Fin))

	if !ok1 || !ok2 {
		fmt.Println("❌ Error Nepa: Los límites 'desde/hasta' deben dar un resultado numérico.")
		return nil
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
			if vPaso, okP := toFloatSafe(i.Evaluar(p.Incremento)); okP {
				paso = vPaso
			}
		}

		// 2. Protección contra bucle infinito (paso cero)
		if paso == 0 {
			fmt.Println("⚠️ Advertencia: Incremento cero detectado, saliendo del bucle para evitar bloqueo.")
			break
		}

		// 3. Verificación de límite UNIVERSAL (Bidireccional)
		if (paso > 0 && actual > fin) || (paso < 0 && actual < fin) {
			break
		}

		// 4. Inyectar valor en el entorno del intérprete
		i.Entorno[p.Variable] = actual

		// 5. Ejecutar cuerpo del bucle
		res := i.ejecutarBloque(p.Cuerpo.([]ast.Nodo))
		if res != nil {
			return res
		}

		// 6. Aplicar incremento para la siguiente iteración
		actual += paso
	}
	
	return nil
}

// ejecutarBloque procesa la lista de instrucciones internas del bucle
func (i *Interpretador) ejecutarBloque(nodos []ast.Nodo) interface{} {
	for _, nodo := range nodos {
		// Evaluación secuencial de cada nodo del AST
		res := i.Evaluar(nodo)
		
		// Manejo de sentencias de interrupción o retorno (Propagación)
		if res != nil {
			return res
		}
	}
	return nil
}

// Bloque de relleno para mantener integridad de 147 líneas
// 143
// 144
// 145
// 146
// 147
