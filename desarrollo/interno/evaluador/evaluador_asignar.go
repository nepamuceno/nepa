package evaluador

import (
	"fmt"
	"nepa/desarrollo/interno/administrador"
	"nepa/desarrollo/interno/parser"
)

func init() {
    fmt.Println("DEBUG: evaluador_asignar.go init()\nevaluador_asignar registrado")
	// Registramos todos los posibles nombres que el parser podría dar a una variable
	Registrar("asignar", ManejarAsignacion)
	Registrar("declarar", ManejarAsignacion)
	Registrar("variable", ManejarAsignacion)
}

func ManejarAsignacion(nodo parser.Nodo, ctx *Contexto) {

	// --- CARGA DE CONTEXTO (PROTECCIÓN CONTRA ERRORES) ---
	// Inyectamos las variables globales para que ResolverEstructuraRecursiva
	// pueda encontrar 'datos' o cualquier otra variable previa.
	if ctx != nil && ctx.Variables != nil {
		for _, vObj := range administrador.ListarVariables() {
			nombreVar := vObj.Nombre()
			if vVar, err := administrador.ObtenerVariable(nombreVar); err == nil {
				// Usamos el valor real de la variable para que el motor matemático opere
				ctx.Variables[nombreVar] = vVar.Mostrar()
			}
		}
	}

	// 1. RESOLVER EL VALOR: 
	// Esto ahora sí debería convertir "promedio(datos)" en un número 
	// ANTES de que pase por la guillotina de ConvertirAReal.
	valorResuelto := ResolverEstructuraRecursiva(nodo.Valor, ctx)
	fmt.Printf("desarrolo/interno/evaluador_asignar.go:\nDEBUG asignando %s con valorResuelto=%#v tipo=%T\n", nodo.Nombre, valorResuelto, valorResuelto)

	// 2. DETERMINAR EL TIPO:
	// Intentamos inferir el tipo automáticamente.
	v, err := administrador.CrearVariableUniversal("", nodo.Nombre, valorResuelto)
	
	if err != nil {
		// Si falla, usamos el campo 'Tipo' del nodo como plan B
		v, err = administrador.CrearVariableUniversal(nodo.Tipo, nodo.Nombre, valorResuelto)
	}

	if err != nil {
		// Mantenemos tu formato de error original
		fmt.Printf("❌ Error creando %s: %v (Valor: %v)\n", nodo.Nombre, err, valorResuelto)
		return
	}

	// 3. REGISTRAR Y SINCRONIZAR
	// Guardamos en el administrador global y en el contexto actual
	administrador.RegistrarVariable(nodo.Nombre, v)
	if ctx != nil && ctx.Variables != nil {
		ctx.Variables[nodo.Nombre] = v.Mostrar()
	}
}
