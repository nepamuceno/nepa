package evaluador

import (
    "fmt"
    "nepa/desarrollo/interno/ast"
    "nepa/desarrollo/interno/modulo"
    "strconv"
    "strings"
)

// NepaRetorno es necesario para capturar el valor de 'retorna' en funciones
type NepaRetorno struct{ Valor interface{} }

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

    // --- REGISTRO DE FUNCIONES NATIVAS ---
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

    // Carga de SDK Matemáticas por defecto
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
    if r, ok := v.(NepaRetorno); ok { v = r.Valor }
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
        res := i.Evaluar(nodo)
        // Si hay un error o un retorno fuera de función, manejarlo aquí
        if _, ok := res.(NepaRetorno); ok {
            return 
        }
    }
}

func (i *Interpretador) Evaluar(nodo interface{}) interface{} {
    if nodo == nil { return nil }

    switch n := nodo.(type) {

    case *ast.FuncionDef, ast.FuncionDef:
        var nombre string
        var fDef *ast.FuncionDef
        if v, ok := n.(*ast.FuncionDef); ok { 
            nombre = v.Nombre
            fDef = v 
        } else { 
            d := n.(ast.FuncionDef)
            nombre = d.Nombre
            fDef = &d
        }
        i.Entorno[nombre] = fDef
        return nil

    case *ast.Identificador, ast.Identificador:
        var nombre string
        if v, ok := n.(*ast.Identificador); ok { nombre = v.Nombre } else { nombre = n.(ast.Identificador).Nombre }
        if val, ok := i.Entorno[nombre]; ok {
            return val
        }
        return nil

    case *ast.Literal, ast.Literal:
        if v, ok := n.(*ast.Literal); ok { return v.Valor }
        return n.(ast.Literal).Valor

    case *ast.Asignacion, ast.Asignacion:
        var nombre string
        var valorNodo ast.Nodo
        if v, ok := n.(*ast.Asignacion); ok { nombre, valorNodo = v.Nombre, v.Valor } else { 
            nombre, valorNodo = n.(ast.Asignacion).Nombre, n.(ast.Asignacion).Valor 
        }
        valor := i.Evaluar(valorNodo)
        if r, ok := valor.(NepaRetorno); ok { valor = r.Valor }
        i.Entorno[nombre] = valor
        return valor

    case *ast.OperacionBinaria, ast.OperacionBinaria:
        var op string
        var izqN, derN ast.Nodo
        if v, ok := n.(*ast.OperacionBinaria); ok { op, izqN, derN = v.Operador, v.Izquierda, v.Derecha } else {
            op, izqN, derN = n.(ast.OperacionBinaria).Operador, n.(ast.OperacionBinaria).Izquierda, n.(ast.OperacionBinaria).Derecha
        }

        resIzq := i.Evaluar(izqN)
        resDer := i.Evaluar(derN)

        // IMPORTANTE: Extraer valores reales de retornos para operaciones matemáticas
        if r, ok := resIzq.(NepaRetorno); ok { resIzq = r.Valor }
        if r, ok := resDer.(NepaRetorno); ok { resDer = r.Valor }

        v1, ok1 := toFloatSafe(resIzq)
        v2, ok2 := toFloatSafe(resDer)

        if ok1 && ok2 {
            switch op {
            case "+": return v1 + v2
            case "-": return v1 - v2
            case "*": return v1 * v2
            case "/": if v2 != 0 { return v1 / v2 }; return 0.0
            case "==": return v1 == v2
            case "!=": return v1 != v2
            case "<=": return v1 <= v2
            case ">=": return v1 >= v2
            case "<": return v1 < v2
            case ">": return v1 > v2
            }
        }
        return nil

    case *ast.LlamadaFuncion, ast.LlamadaFuncion:
        var nombre string
        var argsN []ast.Nodo
        if v, ok := n.(*ast.LlamadaFuncion); ok { nombre, argsN = v.Nombre, v.Args } else { nombre, argsN = n.(ast.LlamadaFuncion).Nombre, n.(ast.LlamadaFuncion).Args }
        
        obj, existe := i.Entorno[nombre]
        if !existe {
            return nil
        }

        argsValores := make([]interface{}, len(argsN))
        for idx, arg := range argsN {
            val := i.Evaluar(arg)
            if r, ok := val.(NepaRetorno); ok { val = r.Valor }
            argsValores[idx] = val
        }

        if fDef, ok := obj.(*ast.FuncionDef); ok {
            // Clonar entorno para alcance local (Scope)
            respaldo := make(map[string]interface{})
            for k, v := range i.Entorno { respaldo[k] = v }
            
            for idx, param := range fDef.Parametros {
                if idx < len(argsValores) { i.Entorno[param] = argsValores[idx] }
            }
            
            resultado := i.Evaluar(fDef.Cuerpo)
            i.Entorno = respaldo
            
            if r, ok := resultado.(NepaRetorno); ok { return r.Valor }
            return resultado
        }

        if fn, ok := obj.(func(...interface{}) interface{}); ok {
            return fn(argsValores...)
        }
        return nil

    case *ast.Si, ast.Si:
        var cond ast.Nodo
        var cuerpo, sino []ast.Nodo
        if v, ok := n.(*ast.Si); ok { 
            cond, cuerpo, sino = v.Condicion, v.Cuerpo, v.Sino 
        } else {
            s := n.(ast.Si)
            cond, cuerpo, sino = s.Condicion, s.Cuerpo, s.Sino
        }
        valCond := i.Evaluar(cond)
        if r, ok := valCond.(NepaRetorno); ok { valCond = r.Valor }

        esVerdad := false
        if b, ok := valCond.(bool); ok { esVerdad = b } else if f, ok := toFloatSafe(valCond); ok { esVerdad = f != 0 }

        if esVerdad { 
            return i.Evaluar(cuerpo) 
        } else if sino != nil { 
            return i.Evaluar(sino) 
        }
        return nil

    case *ast.Retornar, ast.Retornar:
        var valNodo ast.Nodo
        if v, ok := n.(*ast.Retornar); ok { valNodo = v.Valor } else { valNodo = n.(ast.Retornar).Valor }
        return NepaRetorno{ Valor: i.Evaluar(valNodo) }

    case []ast.Nodo:
        var ultimo interface{}
        for _, sn := range n {
            ultimo = i.Evaluar(sn)
            if _, ok := ultimo.(NepaRetorno); ok { return ultimo }
        }
        return ultimo

    case *ast.Para, ast.Para:
        var pDef *ast.Para
        if v, ok := n.(*ast.Para); ok { pDef = v } else { d := n.(ast.Para); pDef = &d }
        // Se llama al método EjecutarPara que reside en el archivo para.go del mismo paquete
        return i.EjecutarPara(*pDef)

    default:
        return nil
    }
}
