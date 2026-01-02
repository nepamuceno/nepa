package evaluador

import (
    "fmt"
    "reflect"

    "nepa/desarrollo/interno/bloque"
    "nepa/desarrollo/interno/parser"
)

// Contexto de ejecución: variables locales, globales, constantes y funciones
type Contexto struct {
    Variables  map[string]interface{}                      // locales
    Globales   map[string]interface{}                      // compartidas entre todos los programas
    Constantes map[string]interface{}                      // inmutables
    Funciones  map[string]func(...interface{}) interface{} // primitivas y módulos
}

// SolicitudEjecutar representa una instrucción de ejecutar(...) que debe manejar el kernel
type SolicitudEjecutar struct {
    Archivo    string
    Argumentos []interface{}
}

func (s SolicitudEjecutar) Error() string {
    return fmt.Sprintf("Solicitud de ejecutar(%s, %v)", s.Archivo, s.Argumentos)
}

// EjecutarConContexto ejecuta un AST con un contexto inicial
func EjecutarConContexto(ast []parser.Nodo, args map[string]interface{},
    globales map[string]interface{}, constantes map[string]interface{}) (map[string]interface{}, error) {

    ctx := Contexto{
        Variables:  map[string]interface{}{},
        Globales:   globales,
        Constantes: constantes,
        Funciones:  globales["__funciones"].(map[string]func(...interface{}) interface{}),
    }

    // Cargar argumentos iniciales como variables locales
    for k, v := range args {
        ctx.Variables[k] = v
    }

    // Recorrer AST
    for _, nodo := range ast {
        if err := ejecutarNodo(&ctx, nodo); err != nil {
            return nil, err
        }
    }

    return ctx.Variables, nil
}

// ejecutarNodo procesa un nodo del AST
func ejecutarNodo(ctx *Contexto, nodo parser.Nodo) error {
    switch nodo.Tipo {
    case "variable":
        nombre := nodo.Nombre
        valor := nodo.Valor

        // Validar que no sea palabra reservada
        for _, p := range bloque.PalabrasReservadas {
            if nombre == p {
                return fmt.Errorf("Uso inválido de palabra reservada '%s' como variable", nombre)
            }
        }

        ctx.Variables[nombre] = valor

    case "global":
        nombre := nodo.Nombre
        valor := nodo.Valor
        ctx.Globales[nombre] = valor

    case "constante":
        nombre := nodo.Nombre
        valor := nodo.Valor
        if _, existe := ctx.Constantes[nombre]; existe {
            return fmt.Errorf("Constante '%s' ya definida y no puede modificarse", nombre)
        }
        ctx.Constantes[nombre] = valor

    case "llamada":
        fn := nodo.Nombre
        args := nodo.Args

        if fn == "imprimir" {
            for _, a := range args {
                fmt.Print(FormatearValor(a), " ")
            }
            fmt.Println()
        } else if fn == "ejecutar" {
            if len(args) < 1 {
                return fmt.Errorf("La llamada a ejecutar requiere al menos un archivo .nepa")
            }
            archivo := fmt.Sprintf("%v", args[0])
            return SolicitudEjecutar{Archivo: archivo, Argumentos: args[1:]}
        } else if f, ok := ctx.Funciones[fn]; ok {
            resultado := f(args...)
            fmt.Println(FormatearValor(resultado))
        } else {
            return fmt.Errorf("Función '%s' no definida", fn)
        }

    case "bloque":
        return nil

    case "expresion":
        fmt.Println(FormatearValor(nodo.Valor))
        return nil
    }

    return nil
}

// FormatearValor convierte cualquier tipo a string legible (exportada)
func FormatearValor(v interface{}) string {
    if v == nil {
        return "nulo"
    }
    switch val := v.(type) {
    case bool:
        if val {
            return "true"
        }
        return "false"
    case string:
        return val
    case []interface{}:
        var out string
        out = "["
        for i, elem := range val {
            out += FormatearValor(elem)
            if i < len(val)-1 {
                out += ", "
            }
        }
        out += "]"
        return out
    default:
        return fmt.Sprintf("%v", val)
    }
}

// debugTipo devuelve el tipo interno (para depuración opcional)
func debugTipo(v interface{}) string {
    if v == nil {
        return "nil"
    }
    return reflect.TypeOf(v).String()
}
