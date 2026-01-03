package matematicas

import (
    "errors"
    "math"

    "nepa/desarrollo/interno/evaluador"
)

// InyectarBasicas agrega funciones matemáticas básicas al contexto
func InyectarBasicas(ctx *evaluador.Contexto) {
    if ctx.Funciones == nil {
        ctx.Funciones = map[string]func(...interface{}) interface{}{}
    }

    reg := func(n string, f func(...interface{}) interface{}) {
        ctx.Funciones[n] = f
    }

    // --- Raíz cuadrada ---
    reg("raiz", func(args ...interface{}) interface{} {
        if len(args) != 1 {
            return errors.New("❌ ERROR FATAL: raiz requiere 1 argumento")
        }
        f, err := evaluador.ConvertirAReal(args[0])
        if err != nil {
            return err
        }
        return math.Sqrt(f)
    })

    // --- Potencia ---
    reg("potencia", func(args ...interface{}) interface{} {
        if len(args) != 2 {
            return errors.New("❌ ERROR FATAL: potencia requiere 2 argumentos")
        }
        base, err := evaluador.ConvertirAReal(args[0])
        if err != nil {
            return err
        }
        exp, err := evaluador.ConvertirAReal(args[1])
        if err != nil {
            return err
        }
        return math.Pow(base, exp)
    })

    // --- Seno ---
    reg("seno", func(args ...interface{}) interface{} {
        if len(args) != 1 {
            return errors.New("❌ ERROR FATAL: seno requiere 1 argumento")
        }
        f, err := evaluador.ConvertirAReal(args[0])
        if err != nil {
            return err
        }
        return math.Sin(f)
    })

    // --- Coseno ---
    reg("coseno", func(args ...interface{}) interface{} {
        if len(args) != 1 {
            return errors.New("❌ ERROR FATAL: coseno requiere 1 argumento")
        }
        f, err := evaluador.ConvertirAReal(args[0])
        if err != nil {
            return err
        }
        return math.Cos(f)
    })

    // --- Tangente ---
    reg("tangente", func(args ...interface{}) interface{} {
        if len(args) != 1 {
            return errors.New("❌ ERROR FATAL: tangente requiere 1 argumento")
        }
        f, err := evaluador.ConvertirAReal(args[0])
        if err != nil {
            return err
        }
        return math.Tan(f)
    })

    // --- Logaritmo natural ---
    reg("log", func(args ...interface{}) interface{} {
        if len(args) != 1 {
            return errors.New("❌ ERROR FATAL: log requiere 1 argumento")
        }
        f, err := evaluador.ConvertirAReal(args[0])
        if err != nil {
            return err
        }
        return math.Log(f)
    })

    // --- Valor absoluto ---
    reg("abs", func(args ...interface{}) interface{} {
        if len(args) != 1 {
            return errors.New("❌ ERROR FATAL: abs requiere 1 argumento")
        }
        f, err := evaluador.ConvertirAReal(args[0])
        if err != nil {
            return err
        }
        return math.Abs(f)
    })
}
