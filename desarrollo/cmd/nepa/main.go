package main

import (
	"bufio"
	"fmt"
	"os"

	"nepa/desarrollo/interno/sintaxis"
	"nepa/desarrollo/interno/parser"
	"nepa/desarrollo/interno/evaluador"

	// 🔑 El Mago Nepu: Al importar matemáticas, su init() registra todo solo
	_ "nepa/desarrollo/interno/matematicas"
	_ "nepa/desarrollo/comandos"
)

var Globales = map[string]interface{}{}
var Constantes = map[string]interface{}{}

// EjecutarPrograma abre un archivo .nepa, valida y ejecuta
func EjecutarPrograma(archivo string, args map[string]interface{}) (map[string]interface{}, error) {
	f, err := os.Open(archivo)
	if err != nil {
		return nil, fmt.Errorf("❌ ERROR: No se pudo abrir el archivo [%s]: %v", archivo, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lineas []string
	lineaNum := 0

	for scanner.Scan() {
		lineaNum++
		linea := scanner.Text()

		// Validar sintaxis básica antes de parsear
		if err := sintaxis.ValidarLinea(linea, lineaNum, archivo); err != nil {
			return nil, fmt.Errorf("❌ ERROR SINTAXIS [%s:%d]: %v", archivo, lineaNum, err)
		}

		lineas = append(lineas, linea)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("❌ ERROR LECTURA [%s]: %v", archivo, err)
	}

	ast := parser.Parse(lineas)

	// El evaluador ahora usa el mapa global de funciones que los init() llenaron
	resultados, err := evaluador.EjecutarConContexto(ast, args, Globales, Constantes, archivo)
	if err != nil {
		// Manejo de recursión de archivos (Ejecutar otros programas)
		if solicitud, ok := err.(evaluador.SolicitudEjecutar); ok {
			subArgs := map[string]interface{}{}
			for i, arg := range solicitud.Argumentos {
				subArgs[fmt.Sprintf("arg%d", i+1)] = arg
			}
			return EjecutarPrograma(solicitud.Archivo, subArgs)
		}
		return nil, err
	}

	// Persistir resultados en globales
	for k, v := range resultados {
		Globales[k] = v
	}

	return resultados, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Nepa Wizard - Intérprete de Ingeniería")
		fmt.Println("Uso: nepa archivo.nepa [argumentos]")
		os.Exit(1)
	}

	// Manejo de flags básicos sin depender de paquete core
	switch os.Args[1] {
	case "-v", "--version":
		fmt.Println("Nepa Engine v2.0 - Super Wizard Nepu")
		os.Exit(0)
	case "-a", "--ayuda":
		fmt.Println("Ayuda básica: nepa archivo.nepa [argumentos]")
		os.Exit(0)
	}

	archivo := os.Args[1]
	args := map[string]interface{}{}
	if len(os.Args) > 2 {
		for i, arg := range os.Args[2:] {
			args[fmt.Sprintf("arg%d", i+1)] = arg
		}
	}

	// ¡YA NO NECESITAMOS CargarCore! 
	// El mapa evaluador.Funciones ya está lleno gracias a los imports con '_'

	resultados, err := EjecutarPrograma(archivo, args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Mostrar resultados finales
	if len(resultados) > 0 {
		fmt.Println("\n--- Resultados ---")
		for k, v := range resultados {
			fmt.Printf(" %s = %v\n", k, v)
		}
	}
}
