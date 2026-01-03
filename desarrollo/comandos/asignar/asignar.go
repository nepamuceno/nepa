package asignar

import (
    "errors"
    "fmt"
    "strings"

    "nepa/desarrollo/interno/administrador"
    "nepa/desarrollo/interno/evaluador"
    "nepa/desarrollo/interno/parser"
)

var (
    ErrSintaxisInvalida = errors.New("sintaxis inválida: use 'asignar <nombre> = <expresión>' o '<nombre> = <expresión>' o '<nombre>++/--'")
    ErrVariableNoExiste = errors.New("la variable no existe")
    ErrNombreVacio      = errors.New("nombre vacío en asignación")
    ErrValorVacio       = errors.New("valor vacío en asignación")
)

// Ejecutar soporta:
//   - asignar a = expr
//   - a = expr
//   - a++ / a--
//   - asignar a++ / asignar a--
func Ejecutar(linea string) error {
    linea = strings.TrimSpace(linea)
    if linea == "" {
        return ErrSintaxisInvalida
    }

    // Quitar palabra clave "asignar" si existe
    if strings.HasPrefix(strings.ToLower(linea), "asignar") {
        linea = strings.TrimSpace(linea[len("asignar"):])
    }

    // Incremento/decremento
    if strings.HasSuffix(linea, "++") || strings.HasSuffix(linea, "--") {
        nombre := strings.TrimSpace(strings.TrimSuffix(strings.TrimSuffix(linea, "++"), "--"))
        if nombre == "" {
            return ErrNombreVacio
        }
        v, err := administrador.ObtenerVariable(nombre)
        if err != nil {
            return fmt.Errorf("❌ %w: '%s'", ErrVariableNoExiste, nombre)
        }
        delta := 1.0
        if strings.HasSuffix(linea, "--") {
            delta = -1.0
        }
        actual, err := toFloat64(v.ValorComoInterface())
        if err != nil {
            return fmt.Errorf("⚠️ incremento/decremento inválido para tipo=%s: %v", v.Tipo(), err)
        }
        nuevo := actual + delta
        if err := administrador.ModificarVariable(nombre, nuevo); err != nil {
            return fmt.Errorf("⚠️ error asignando a '%s' (tipo=%s): %w", nombre, v.Tipo(), err)
        }
        fmt.Printf("✔ %s '%s' %s → %v\n", strings.ToUpper(v.Tipo()), nombre, tern(delta > 0, "++", "--"), nuevo)
        return nil
    }

    // Asignación con '='
    partes := strings.SplitN(linea, "=", 2)
    if len(partes) != 2 {
        return ErrSintaxisInvalida
    }
    nombre := strings.TrimSpace(partes[0])
    expr := strings.TrimSpace(partes[1])
    if nombre == "" {
        return ErrNombreVacio
    }
    if expr == "" {
        return ErrValorVacio
    }

    // Verificar existencia
    v, err := administrador.ObtenerVariable(nombre)
    if err != nil {
        return fmt.Errorf("❌ %w: '%s'", ErrVariableNoExiste, nombre)
    }

    // Evaluar expresión completa (variables, literales, funciones, etc.)
    resultado, err := evaluador.Eval(expr)
    if err != nil {
        // Fallback: pasar crudo, el tipo decidirá si es válido
        resultado = expr
    }

    // Asignar resultado
    if err := administrador.ModificarVariable(nombre, resultado); err != nil {
        return fmt.Errorf("⚠️ error asignando a '%s' (tipo=%s) con '%v': %w", nombre, v.Tipo(), resultado, err)
    }

    fmt.Printf("✔ %s '%s' ← %s → %v\n", strings.ToUpper(v.Tipo()), nombre, expr, resultado)
    return nil
}

func tern(cond bool, a, b string) string {
    if cond {
        return a
    }
    return b
}

// Conversión genérica a float64 (ahora soporta bit y bool)
func toFloat64(v interface{}) (float64, error) {
    switch n := v.(type) {
    case int:
        return float64(n), nil
    case int64:
        return float64(n), nil
    case float64:
        return n, nil
    case uint8:
        return float64(n), nil
    case bool:
        if n {
            return 1.0, nil
        }
        return 0.0, nil
    default:
        return 0, fmt.Errorf("no se puede convertir %T a float64", v)
    }
}

// init registra este comando en el ejecutor universal
func init() {
    evaluador.Registrar("asignar", func(n parser.Nodo, ctx *evaluador.Contexto) {
        // Verificar existencia
        v, err := administrador.ObtenerVariable(n.Nombre)
        if err != nil {
            fmt.Printf("❌ Variable '%s' no existe\n", n.Nombre)
            return
        }

        // Asignar valor directamente desde el nodo
        if err := administrador.ModificarVariable(n.Nombre, n.Valor); err != nil {
            fmt.Printf("⚠️ Error asignando a '%s': %v\n", n.Nombre, err)
            return
        }

        fmt.Printf("✔ %s '%s' ← %v\n", strings.ToUpper(v.Tipo()), n.Nombre, n.Valor)
    })
}
