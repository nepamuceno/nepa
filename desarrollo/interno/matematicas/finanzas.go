package matematicas

import (
    "errors"
    "math"

    "nepa/desarrollo/interno/evaluador"
)

// InyectarFinanzas agrega funciones financieras al contexto
func InyectarFinanzas(ctx *evaluador.Contexto) {
    if ctx.Funciones == nil {
        ctx.Funciones = map[string]func(...interface{}) interface{}{}
    }

    reg := func(n string, f func(...interface{}) interface{}) {
        ctx.Funciones[n] = f
    }

    // --- Interés compuesto ---
    reg("interes_compuesto", func(args ...interface{}) interface{} {
        if len(args) != 3 {
            return errors.New("❌ ERROR FATAL: interes_compuesto requiere 3 argumentos (capital, tasa, tiempo)")
        }
        c, _ := evaluador.ConvertirAReal(args[0])
        i, _ := evaluador.ConvertirAReal(args[1])
        t, _ := evaluador.ConvertirAReal(args[2])
        return c * math.Pow(1+i, t)
    })

    // --- Interés simple ---
    reg("interes_simple", func(args ...interface{}) interface{} {
        if len(args) != 3 {
            return errors.New("❌ ERROR FATAL: interes_simple requiere 3 argumentos (capital, tasa, tiempo)")
        }
        c, _ := evaluador.ConvertirAReal(args[0])
        i, _ := evaluador.ConvertirAReal(args[1])
        t, _ := evaluador.ConvertirAReal(args[2])
        return c * i * t
    })

    // --- Valor presente ---
    reg("valor_presente", func(args ...interface{}) interface{} {
        if len(args) != 3 {
            return errors.New("❌ ERROR FATAL: valor_presente requiere 3 argumentos (valor futuro, tasa, tiempo)")
        }
        vf, _ := evaluador.ConvertirAReal(args[0])
        i, _ := evaluador.ConvertirAReal(args[1])
        t, _ := evaluador.ConvertirAReal(args[2])
        return vf / math.Pow(1+i, t)
    })

    // --- Valor futuro de una anualidad ---
    reg("valor_futuro_anualidad", func(args ...interface{}) interface{} {
        if len(args) != 3 {
            return errors.New("❌ ERROR FATAL: valor_futuro_anualidad requiere 3 argumentos (pago, tasa, periodos)")
        }
        p, _ := evaluador.ConvertirAReal(args[0])
        i, _ := evaluador.ConvertirAReal(args[1])
        n, _ := evaluador.ConvertirAReal(args[2])
        if i == 0 {
            return p * n
        }
        return p * (math.Pow(1+i, n) - 1) / i
    })

    // --- Valor presente de una anualidad ---
    reg("valor_presente_anualidad", func(args ...interface{}) interface{} {
        if len(args) != 3 {
            return errors.New("❌ ERROR FATAL: valor_presente_anualidad requiere 3 argumentos (pago, tasa, periodos)")
        }
        p, _ := evaluador.ConvertirAReal(args[0])
        i, _ := evaluador.ConvertirAReal(args[1])
        n, _ := evaluador.ConvertirAReal(args[2])
        if i == 0 {
            return p * n
        }
        return p * (1 - math.Pow(1+i, -n)) / i
    })

    // --- Pago de anualidad (PMT) ---
    reg("pago_anualidad", func(args ...interface{}) interface{} {
        if len(args) != 3 {
            return errors.New("❌ ERROR FATAL: pago_anualidad requiere 3 argumentos (principal, tasa, periodos)")
        }
        p, _ := evaluador.ConvertirAReal(args[0])
        i, _ := evaluador.ConvertirAReal(args[1])
        n, _ := evaluador.ConvertirAReal(args[2])
        if i == 0 {
            return p / n
        }
        return (p * i) / (1 - math.Pow(1+i, -n))
    })

    // --- Valor Neto Actual (VAN) ---
    reg("van", func(args ...interface{}) interface{} {
        if len(args) < 2 {
            return errors.New("❌ ERROR FATAL: van requiere al menos 2 argumentos (tasa, flujos...)")
        }
        i, _ := evaluador.ConvertirAReal(args[0])
        flujos := args[1:]
        van := 0.0
        for t, f := range flujos {
            val, err := evaluador.ConvertirAReal(f)
            if err != nil {
                return err
            }
            van += val / math.Pow(1+i, float64(t+1))
        }
        return van
    })

    // --- Tasa Interna de Retorno (TIR) ---
    reg("tir", func(args ...interface{}) interface{} {
        if len(args) < 2 {
            return errors.New("❌ ERROR FATAL: tir requiere al menos 2 argumentos (costo inicial, flujos...)")
        }
        c0, _ := evaluador.ConvertirAReal(args[0])
        flujos := args[1:]
        // Método de bisección simple
        low, high := -0.99, 1.0
        for iter := 0; iter < 100; iter++ {
            mid := (low + high) / 2
            npv := -c0
            for t, f := range flujos {
                val, _ := evaluador.ConvertirAReal(f)
                npv += val / math.Pow(1+mid, float64(t+1))
            }
            if npv > 0 {
                low = mid
            } else {
                high = mid
            }
        }
        return (low + high) / 2
    })

    // --- Tasa efectiva anual (TEA) ---
    reg("tea", func(args ...interface{}) interface{} {
        if len(args) != 2 {
            return errors.New("❌ ERROR FATAL: tea requiere 2 argumentos (tasa nominal, número de periodos por año)")
        }
        i, _ := evaluador.ConvertirAReal(args[0])
        m, _ := evaluador.ConvertirAReal(args[1])
        return math.Pow(1+i/m, m) - 1
    })

    // --- Tasa nominal anual (TNA) ---
    reg("tna", func(args ...interface{}) interface{} {
        if len(args) != 2 {
            return errors.New("❌ ERROR FATAL: tna requiere 2 argumentos (tasa periódica, número de periodos por año)")
        }
        i, _ := evaluador.ConvertirAReal(args[0])
        m, _ := evaluador.ConvertirAReal(args[1])
        return i * m
    })
}
