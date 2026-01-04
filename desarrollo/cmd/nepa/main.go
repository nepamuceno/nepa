package main

import (
    "bufio"
    "fmt"
    "os"

    "nepa/desarrollo/interno/core"
    "nepa/desarrollo/interno/sintaxis"
    "nepa/desarrollo/interno/parser"
    "nepa/desarrollo/interno/evaluador"
    "nepa/desarrollo/interno/modulos"

    // 🔑 Importar todos los comandos de una sola vez
    _ "nepa/desarrollo/comandos"
)

// Mapas globales y constantes compartidos por todos los programas
var Globales = map[string]interface{}{}
var Constantes = map[string]interface{}{}

// EjecutarPrograma abre un archivo .nepa, valida, ejecuta y devuelve resultados o error
func EjecutarPrograma(archivo string, args map[string]interface{}) (map[string]interface{}, error) {
    f, err := os.Open(archivo)
    if err != nil {
        return nil, core.ErrorEjecucion{Archivo: archivo, Linea: 0, Mensaje: fmt.Sprintf("No se pudo abrir archivo: %v", err)}
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    buf := make([]byte, 0, 64*1024)
    scanner.Buffer(buf, 1024*1024)

    var lineas []string
    lineaNum := 0

    for scanner.Scan() {
        lineaNum++
        linea := scanner.Text()

        if lineaNum == 1 {
            linea = core.StripBOM(linea)
        }

        if err := sintaxis.ValidarLinea(linea, lineaNum, archivo); err != nil {
            return nil, core.ErrorEjecucion{Archivo: archivo, Linea: lineaNum, Mensaje: err.Error()}
        }

        lineas = append(lineas, linea)
    }

    if err := scanner.Err(); err != nil {
        return nil, core.ErrorEjecucion{Archivo: archivo, Linea: 0, Mensaje: fmt.Sprintf("Error leyendo archivo: %v", err)}
    }

    ast := parser.Parse(lineas)

    resultados, err := evaluador.EjecutarConContexto(ast, args, Globales, Constantes, archivo)
    if err != nil {
        if solicitud, ok := err.(evaluador.SolicitudEjecutar); ok {
            subArgs := map[string]interface{}{}
            for i, arg := range solicitud.Argumentos {
                clave := fmt.Sprintf("arg%d", i+1)
                subArgs[clave] = arg
            }
            return EjecutarPrograma(solicitud.Archivo, subArgs)
        }
        // ✅ No envolver de nuevo, devolver tal cual
        return nil, err
    }

    for k, v := range resultados {
        Globales[k] = v
    }

    return resultados, nil
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Uso: nepa archivo.nepa [argumentos]")
        fmt.Println("Ejemplo: nepa programa1.nepa 45 matriz.json")
        fmt.Println("Opciones: -v/--version muestra versión, -a/--ayuda muestra ayuda")
        os.Exit(1)
    }

    switch os.Args[1] {
    case "-v", "--v", "-version", "--version":
        fmt.Println("Nepa - Intérprete modular y orquestador de programas .nepa")
        fmt.Printf("Versión: %s\n", core.Version)
        os.Exit(0)
    case "-a", "--a", "-ayuda", "--ayuda":
        fmt.Println("Ayuda: Uso básico → nepa archivo.nepa [argumentos]")
        fmt.Println("Ejemplo: nepa programa1.nepa 45 matriz.json")
        fmt.Println("Opciones disponibles: -v/--version muestra versión, -a/--ayuda muestra esta ayuda")
        os.Exit(0)
    }

    archivo := os.Args[1]

    args := map[string]interface{}{}
    if len(os.Args) > 2 {
        for i, arg := range os.Args[2:] {
            clave := fmt.Sprintf("arg%d", i+1)
            args[clave] = arg
        }
    }

    // --- Cargar todo el core al inicio ---
    ctx := evaluador.Contexto{
        Variables:  map[string]interface{}{},
        Globales:   Globales,
        Constantes: Constantes,
        Funciones:  map[string]func(...interface{}) interface{}{},
    }
    modulos.CargarCore(&ctx)

    // 🔑 pasar funciones al mapa global para que el evaluador las use
    Globales["__funciones"] = ctx.Funciones

    resultados, err := EjecutarPrograma(archivo, args)
    if err != nil {
        fmt.Println(err) // ✅ ya formateado por ErrorEjecucion
        os.Exit(1)
    }

    // Mostrar resultados si existen
    if resultados != nil && len(resultados) > 0 {
        fmt.Println("Resultados:")
        for k, v := range resultados {
            fmt.Printf("  %s = %s\n", k, evaluador.FormatearValor(v))
        }
    }

    // Mostrar globales y constantes solo si el usuario definió algo
    if len(Globales) > 1 { // >1 porque siempre existe "__funciones"
        fmt.Println("Globales:")
        for k, v := range Globales {
            if k == "__funciones" {
                continue // no mostrar el mapa de funciones internas
            }
            fmt.Printf("  %s = %s\n", k, evaluador.FormatearValor(v))
        }
    }

    if len(Constantes) > 0 {
        fmt.Println("Constantes:")
        for k, v := range Constantes {
            fmt.Printf("  %s = %s\n", k, evaluador.FormatearValor(v))
        }
    }
}
