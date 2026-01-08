package main

import (
    "bufio"
    "fmt"
    "os"

    "nepa/desarrollo/interno/sintaxis"
    "nepa/desarrollo/interno/parser"
    "nepa/desarrollo/interno/evaluador"
    "nepa/desarrollo/interno/core"

    // 🔑 El Mago Nepu: Al importar matemáticas, su init() registra todo solo
    _ "nepa/desarrollo/interno/matematicas"
    _ "nepa/desarrollo/comandos"
)

// Variables de entorno de Nepa (prefijo _MAYÚSCULAS en español)
var _GLOBALES = map[string]interface{}{}   // Variables globales
var _CONSTANTES = map[string]interface{}{} // Constantes
var _DEPURACION = 0                        // Nivel de depuración
var _DETALLE = 0                           // Nivel de detalle (verbose)
var _CONFIGURACION = "nepa.conf"           // Archivo de configuración
var _VERDADERO = true                      // Valor lógico verdadero
var _FALSO = false                         // Valor lógico falso

// EjecutarPrograma abre un archivo .nepa, valida y ejecuta
func EjecutarPrograma(archivo string, args map[string]interface{}) (map[string]interface{}, error) {
    f, err := os.Open(archivo)
    if err != nil {
        core.EmitirError(core._FATAL, archivo, 0, 1000, archivo) // Error: archivo no encontrado
        return nil, err
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
            core.EmitirError(core._FATAL, archivo, lineaNum, 2000, linea) // Error sintaxis inválida
            return nil, err
        }

        lineas = append(lineas, linea)
    }

    if err := scanner.Err(); err != nil {
        core.EmitirError(core._FATAL, archivo, 0, 1001, err.Error()) // Error lectura
        return nil, err
    }

    ast := parser.Parse(lineas)

    // Evaluador con entorno global
    resultados, err := evaluador.EjecutarConContexto(ast, args, _GLOBALES, _CONSTANTES, archivo)
    if err != nil {
        // Manejo de recursión de archivos (Ejecutar otros programas)
        if solicitud, ok := err.(evaluador.SolicitudEjecutar); ok {
            subArgs := map[string]interface{}{}
            for i, arg := range solicitud.Argumentos {
                subArgs[fmt.Sprintf("arg%d", i+1)] = arg
            }
            return EjecutarPrograma(solicitud.Archivo, subArgs)
        }
        core.EmitirError(core._FATAL, archivo, 0, 5000, err.Error()) // Error en evaluación
        return nil, err
    }

    // Persistir resultados en globales
    for k, v := range resultados {
        _GLOBALES[k] = v
    }

    return resultados, nil
}

func main() {
    if len(os.Args) < 2 {
        core.EmitirError(core._ADVERTENCIA, "main", 0, 9000, "Uso: nepa [opciones] <programa.nepa>")
        os.Exit(1)
    }

    // Manejo de opciones profesionales
    switch os.Args[1] {
    case "--version":
        fmt.Println("Nepa Engine v2.0 - Súper Mago Nepu")
        os.Exit(0)
    case "--creditos":
        fmt.Println("Autor: Nepamuceno Bartolo")
        fmt.Println("GitHub: https://github.com/nepamuceno/nepa")
        fmt.Println("Correo: zzerver@gmail.com")
        fmt.Println("Compañía: zSoft Software")
        os.Exit(0)
    case "--ayuda", "-a":
        fmt.Println("Uso: nepa [opciones] <programa.nepa> [argumentos]")
        fmt.Println("Opciones disponibles:")
        fmt.Println("  --version               Muestra la versión actual")
        fmt.Println("  --creditos              Muestra créditos del autor y compañía")
        fmt.Println("  --ayuda, -a             Muestra esta ayuda detallada")
        fmt.Println("  --v, --vv, --vvv, --vvvv Control de detalle (1 a 4 niveles)")
        fmt.Println("  --c, --configuracion <archivo.conf> Carga configuración (default: nepa.conf)")
        os.Exit(0)
    case "--v":
        _DETALLE = 1
    case "--vv":
        _DETALLE = 2
    case "--vvv":
        _DETALLE = 3
    case "--vvvv":
        _DETALLE = 4
    case "--c", "--configuracion":
        if len(os.Args) > 2 {
            _CONFIGURACION = os.Args[2]
        }
        // TODO: cargar archivo de configuración si existe
    }

    archivo := os.Args[len(os.Args)-1]
    args := map[string]interface{}{}
    if len(os.Args) > 2 {
        for i, arg := range os.Args[2:] {
            args[fmt.Sprintf("arg%d", i+1)] = arg
        }
    }

    resultados, err := EjecutarPrograma(archivo, args)
    if err != nil {
        os.Exit(1)
    }

	// Mostrar resultados finales según nivel de salida
	if _DEPURACION > 0 && len(resultados) > 0 {
    	core.EmitirDepuracion("main", 0, 6100, nucleo.MENSAJES_DEPURACION[6100])
    	for k, v := range resultados {
        	core.EmitirDepuracion("main", 0, 6101, fmt.Sprintf(nucleo.MENSAJES_DEPURACION[6101], k, v))
    	}
	}

	if _DETALLE > 0 && len(resultados) > 0 {
    	core.EmitirDetalle("main", 0, 7000, nucleo.MENSAJES_DETALLE[7000])
    	for k, v := range resultados {
        	core.EmitirDetalle("main", 0, 7001, fmt.Sprintf(nucleo.MENSAJES_DETALLE[7001], k, v))
    	}
	}
}
