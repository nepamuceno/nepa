package matematicas

import (
    "errors"
    "math"
    "sort"

    "nepa/desarrollo/interno/evaluador"
)

// InyectarEstadistica agrega funciones estadísticas al contexto
func InyectarEstadistica(ctx *evaluador.Contexto) {
    if ctx.Funciones == nil {
        ctx.Funciones = map[string]func(...interface{}) interface{}{}
    }

    reg := func(n string, f func(...interface{}) interface{}) {
        ctx.Funciones[n] = f
    }

    // --- Media ---
    reg("media", func(args ...interface{}) interface{} {
        if len(args) == 0 {
            return errors.New("❌ ERROR FATAL: media requiere al menos 1 argumento")
        }
        sum := 0.0
        for _, v := range args {
            f, err := evaluador.ConvertirAReal(v)
            if err != nil {
                return err
            }
            sum += f
        }
        return sum / float64(len(args))
    })

    // --- Mediana ---
    reg("mediana", func(args ...interface{}) interface{} {
        if len(args) == 0 {
            return errors.New("❌ ERROR FATAL: mediana requiere al menos 1 argumento")
        }
        vals := []float64{}
        for _, v := range args {
            f, err := evaluador.ConvertirAReal(v)
            if err != nil {
                return err
            }
            vals = append(vals, f)
        }
        sort.Float64s(vals)
        l := len(vals)
        if l%2 == 0 {
            return (vals[l/2-1] + vals[l/2]) / 2
        }
        return vals[l/2]
    })

    // --- Varianza ---
    reg("varianza", func(args ...interface{}) interface{} {
        if len(args) == 0 {
            return errors.New("❌ ERROR FATAL: varianza requiere al menos 1 argumento")
        }
        var sum, sumSq float64
        for _, v := range args {
            f, err := evaluador.ConvertirAReal(v)
            if err != nil {
                return err
            }
            sum += f
            sumSq += f * f
        }
        n := float64(len(args))
        return (sumSq / n) - math.Pow(sum/n, 2)
    })

    // --- Desviación estándar ---
    reg("desviacion", func(args ...interface{}) interface{} {
        if len(args) == 0 {
            return errors.New("❌ ERROR FATAL: desviacion requiere al menos 1 argumento")
        }
        var sum, sumSq float64
        for _, v := range args {
            f, err := evaluador.ConvertirAReal(v)
            if err != nil {
                return err
            }
            sum += f
            sumSq += f * f
        }
        n := float64(len(args))
        return math.Sqrt((sumSq / n) - math.Pow(sum/n, 2))
    })

    // --- Rango ---
    reg("rango", func(args ...interface{}) interface{} {
        if len(args) == 0 {
            return errors.New("❌ ERROR FATAL: rango requiere al menos 1 argumento")
        }
        vals := []float64{}
        for _, v := range args {
            f, err := evaluador.ConvertirAReal(v)
            if err != nil {
                return err
            }
            vals = append(vals, f)
        }
        sort.Float64s(vals)
        return vals[len(vals)-1] - vals[0]
    })
}
