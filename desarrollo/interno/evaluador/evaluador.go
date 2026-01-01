package evaluador

import (
    "fmt"
    "nepa/desarrollo/interno/ast"
    "nepa/desarrollo/interno/modulo"
    "strconv"
    "strings"
)

type Interpretador struct {
    Entorno    map[string]interface{}
    Modulos    map[string]*modulo.LibreriaNepa
    Protegidos map[string]bool 
}

func NuevoEntorno() *Interpretador {
    i := &Interpretador{
        Entorno:    make(map[string]interface{}),
        Modulos:    make(map[string]*modulo.LibreriaNepa),
        Protegidos: make(map[string]bool),
    }

    // Registro de funciones nativas
    i.registrarNativa("imprime", func(args ...interface{}) interface{} {
        for idx, arg := range args {
            if s, ok := arg.(string); ok {
                args[idx] = strings.ReplaceAll(s, "\\n", "\n")
            }
        }
        fmt.Println(args...)
        return nil
    })

    i.registrarNativa("ayuda", func(args ...interface{}) interface{} {
        fmt.Println("--- SISTEMA DE AYUDA NEPA ---")
        return nil
    })

    i.registrarNativa("formatear", func(args ...interface{}) interface{} {
        if len(args) < 2 { return args[0] }
        val, _ := toFloatSafe(args[0])
        prec, _ := toFloatSafe(args[1])
        return fmt.Sprintf("%.*f", int(prec), val)
    })

    // Carga automática del SDK
    sdk, err := modulo.Cargar("matematicas")
    if err == nil {
        for nombre, fn := range sdk.Funciones { i.registrarNativa(nombre, fn) }
        for nombre, val := range sdk.Variables { i.registrarNativa(nombre, val) }
    }

    return i
}

func (i *Interpretador) registrarNativa(nombre string, valor interface{}) {
    i.Entorno[nombre] = valor
    i.Protegidos[nombre] = true
}

func toFloatSafe(v interface{}) (float64, bool) {
    switch t := v.(type) {
    case float64: return t, true
    case int: return float64(t), true
    case int64: return float64(t), true
    case string:
        if f, err := strconv.ParseFloat(t, 64); err == nil { return f, true }
    }
    return 0, false
}

func (i *Interpretador) Ejecutar(nodos []ast.Nodo) {
    for _, nodo := range nodos {
        i.Evaluar(nodo)
    }
}

func (i *Interpretador) Evaluar(nodo ast.Nodo) interface{} {
    if nodo == nil { return nil }

    switch n := nodo.(type) {

    case *ast.Importar, ast.Importar:
        var nombre string
        if v, ok := n.(*ast.Importar); ok { nombre = v.Nombre } else { nombre = n.(ast.Importar).Nombre }
        mod, err := modulo.Cargar(nombre)
        if err != nil { panic(fmt.Sprintf("Error al cargar módulo '%s': %v", nombre, err)) }
        i.Modulos[nombre] = mod
        return nil

    case *ast.Literal, ast.Literal:
        if v, ok := n.(*ast.Literal); ok { return v.Valor }
        return n.(ast.Literal).Valor

    case *ast.Identificador, ast.Identificador:
        var nombre string
        if v, ok := n.(*ast.Identificador); ok { nombre = v.Nombre } else { nombre = n.(ast.Identificador).Nombre }
        if val, ok := i.Entorno[nombre]; ok { return val }
        return nombre

    case *ast.Asignacion, ast.Asignacion:
        var nombre string
        var valorNodo ast.Nodo
        if v, ok := n.(*ast.Asignacion); ok { nombre, valorNodo = v.Nombre, v.Valor } else { 
            nombre, valorNodo = n.(ast.Asignacion).Nombre, n.(ast.Asignacion).Valor 
        }

        if i.Protegidos[nombre] {
            panic(fmt.Sprintf("Error: No puedes reasignar la función o constante nativa '%s'", nombre))
        }

        valor := i.Evaluar(valorNodo)
        if f, ok := toFloatSafe(valor); ok { i.Entorno[nombre] = f } else { i.Entorno[nombre] = valor }
        return valor

    case *ast.OperacionBinaria, ast.OperacionBinaria:
        var op string
        var izqN, derN ast.Nodo
        if v, ok := n.(*ast.OperacionBinaria); ok { op, izqN, derN = v.Operador, v.Izquierda, v.Derecha } else {
            op, izqN, derN = n.(ast.OperacionBinaria).Operador, n.(ast.OperacionBinaria).Izquierda, n.(ast.OperacionBinaria).Derecha
        }

        if op == "." {
            modNombre := ""
            if id, ok := izqN.(*ast.Identificador); ok { modNombre = id.Nombre } else if id, ok := izqN.(ast.Identificador); ok { modNombre = id.Nombre }
            if mod, existe := i.Modulos[modNombre]; existe {
                miembro := ""
                switch m := derN.(type) {
                case *ast.Identificador: miembro = m.Nombre
                case ast.Identificador: miembro = m.Nombre
                }
                if fn, ok := mod.Funciones[miembro]; ok { return fn }
                if val, ok := mod.Variables[miembro]; ok { return val }
            }
            return nil
        }

        izq, der := i.Evaluar(izqN), i.Evaluar(derN)
        
        if op == "+" {
            _, ok1 := toFloatSafe(izq)
            _, ok2 := toFloatSafe(der)
            if !ok1 || !ok2 {
                return fmt.Sprintf("%v%v", izq, der)
            }
        }

        v1, ok1 := toFloatSafe(izq)
        v2, ok2 := toFloatSafe(der)

        if ok1 && ok2 {
            switch op {
            case "+": return v1 + v2
            case "-": return v1 - v2
            case "*": return v1 * v2
            case "/": 
                if v2 == 0 { return 0.0 }
                return v1 / v2
            case "==": return v1 == v2
            case "!=": return v1 != v2
            case ">":  return v1 > v2
            case "<":  return v1 < v2
            case ">=": return v1 >= v2
            case "<=": return v1 <= v2
            }
        }
        return nil

    case *ast.Si, ast.Si:
        var c, cuerpo, sino ast.Nodo
        if v, ok := n.(*ast.Si); ok { c, cuerpo, sino = v.Condicion, v.Cuerpo, v.Sino } else {
            c, cuerpo, sino = n.(ast.Si).Condicion, n.(ast.Si).Cuerpo, n.(ast.Si).Sino
        }
        if val := i.Evaluar(c); val == true || val == 1.0 { return i.Evaluar(cuerpo) } else if sino != nil { return i.Evaluar(sino) }
        return nil

    case *ast.Mientras, ast.Mientras:
        var c, cuerpo ast.Nodo
        if v, ok := n.(*ast.Mientras); ok { c, cuerpo = v.Condicion, v.Cuerpo } else { c, cuerpo = n.(ast.Mientras).Condicion, n.(ast.Mientras).Cuerpo }
        for {
            val := i.Evaluar(c)
            if val != true && val != 1.0 { break }
            i.Evaluar(cuerpo)
        }
        return nil

    // --- BLOQUE AÑADIDO PARA EL PARA UNIVERSAL ---
    case *ast.Para, ast.Para:
        var p ast.Para
        if v, ok := n.(*ast.Para); ok { p = *v } else { p = n.(ast.Para) }
        return i.EjecutarPara(p)

    case *ast.LlamadaFuncion, ast.LlamadaFuncion:
        var nombre string
        var argsN []ast.Nodo
        if v, ok := n.(*ast.LlamadaFuncion); ok { 
            nombre, argsN = v.Nombre, v.Args 
        } else { 
            nombre, argsN = n.(ast.LlamadaFuncion).Nombre, n.(ast.LlamadaFuncion).Args 
        }

        obj := i.Entorno[nombre]
        if obj == nil { return nil }

        if fn, ok := obj.(func(...interface{}) interface{}); ok {
            args := make([]interface{}, len(argsN))
            for idx, arg := range argsN { args[idx] = i.Evaluar(arg) }
            return fn(args...)
        }
        if fDef, ok := obj.(*ast.FuncionDef); ok { return i.Evaluar(fDef.Cuerpo) }
        
        return nil

    case *ast.FuncionDef, ast.FuncionDef:
        var nombre string
        if v, ok := n.(*ast.FuncionDef); ok { nombre = v.Nombre } else { nombre = n.(ast.FuncionDef).Nombre }
        if i.Protegidos[nombre] { panic(fmt.Sprintf("Error: No puedes usar el nombre nativo '%s'", nombre)) }
        i.Entorno[nombre] = n
        return nil

    case []ast.Nodo:
        var last interface{}
        for _, sn := range n { last = i.Evaluar(sn) }
        return last

    default:
        return nil
    }
}
