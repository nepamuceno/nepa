package matematicas

import (
    "errors"
    "math"

    "nepa/desarrollo/interno/evaluador"
)

// InyectarAlgebra agrega funciones de álgebra lineal y vectorial al contexto
func InyectarAlgebra(ctx *evaluador.Contexto) {
    if ctx.Funciones == nil {
        ctx.Funciones = map[string]func(...interface{}) interface{}{}
    }

    reg := func(n string, f func(...interface{}) interface{}) {
        ctx.Funciones[n] = f
    }

    // --- Determinantes ---
    reg("det2x2", func(args ...interface{}) interface{} {
        if len(args) != 4 {
            return errors.New("❌ ERROR FATAL: det2x2 requiere 4 argumentos")
        }
        a, _ := evaluador.ConvertirAReal(args[0])
        b, _ := evaluador.ConvertirAReal(args[1])
        c, _ := evaluador.ConvertirAReal(args[2])
        d, _ := evaluador.ConvertirAReal(args[3])
        return a*d - b*c
    })

    reg("det3x3", func(args ...interface{}) interface{} {
        if len(args) != 9 {
            return errors.New("❌ ERROR FATAL: det3x3 requiere 9 argumentos")
        }
        a, _ := evaluador.ConvertirAReal(args[0])
        b, _ := evaluador.ConvertirAReal(args[1])
        c, _ := evaluador.ConvertirAReal(args[2])
        d, _ := evaluador.ConvertirAReal(args[3])
        e, _ := evaluador.ConvertirAReal(args[4])
        f, _ := evaluador.ConvertirAReal(args[5])
        g, _ := evaluador.ConvertirAReal(args[6])
        h, _ := evaluador.ConvertirAReal(args[7])
        i, _ := evaluador.ConvertirAReal(args[8])
        return a*(e*i-f*h) - b*(d*i-f*g) + c*(d*h-e*g)
    })

    // --- Traza ---
    reg("traza2x2", func(args ...interface{}) interface{} {
        if len(args) != 4 {
            return errors.New("❌ ERROR FATAL: traza2x2 requiere 4 argumentos")
        }
        a, _ := evaluador.ConvertirAReal(args[0])
        d, _ := evaluador.ConvertirAReal(args[3])
        return a + d
    })

    reg("traza3x3", func(args ...interface{}) interface{} {
        if len(args) != 9 {
            return errors.New("❌ ERROR FATAL: traza3x3 requiere 9 argumentos")
        }
        a, _ := evaluador.ConvertirAReal(args[0])
        e, _ := evaluador.ConvertirAReal(args[4])
        i, _ := evaluador.ConvertirAReal(args[8])
        return a + e + i
    })

    // --- Producto punto ---
    reg("producto_punto", func(args ...interface{}) interface{} {
        if len(args) != 6 {
            return errors.New("❌ ERROR FATAL: producto_punto requiere 6 argumentos")
        }
        x1, _ := evaluador.ConvertirAReal(args[0])
        y1, _ := evaluador.ConvertirAReal(args[1])
        z1, _ := evaluador.ConvertirAReal(args[2])
        x2, _ := evaluador.ConvertirAReal(args[3])
        y2, _ := evaluador.ConvertirAReal(args[4])
        z2, _ := evaluador.ConvertirAReal(args[5])
        return x1*x2 + y1*y2 + z1*z2
    })

    // --- Magnitud de vector ---
    reg("magnitud_vector", func(args ...interface{}) interface{} {
        if len(args) != 3 {
            return errors.New("❌ ERROR FATAL: magnitud_vector requiere 3 argumentos")
        }
        x, _ := evaluador.ConvertirAReal(args[0])
        y, _ := evaluador.ConvertirAReal(args[1])
        z, _ := evaluador.ConvertirAReal(args[2])
        return math.Sqrt(x*x + y*y + z*z)
    })

    // --- Producto cruz ---
    reg("producto_cruz", func(args ...interface{}) interface{} {
        if len(args) != 6 {
            return errors.New("❌ ERROR FATAL: producto_cruz requiere 6 argumentos")
        }
        x1, _ := evaluador.ConvertirAReal(args[0])
        y1, _ := evaluador.ConvertirAReal(args[1])
        z1, _ := evaluador.ConvertirAReal(args[2])
        x2, _ := evaluador.ConvertirAReal(args[3])
        y2, _ := evaluador.ConvertirAReal(args[4])
        z2, _ := evaluador.ConvertirAReal(args[5])
        return []float64{
            y1*z2 - z1*y2,
            z1*x2 - x1*z2,
            x1*y2 - y1*x2,
        }
    })

    // --- Normalización de vector ---
    reg("normalizar_vector", func(args ...interface{}) interface{} {
        if len(args) != 3 {
            return errors.New("❌ ERROR FATAL: normalizar_vector requiere 3 argumentos")
        }
        x, _ := evaluador.ConvertirAReal(args[0])
        y, _ := evaluador.ConvertirAReal(args[1])
        z, _ := evaluador.ConvertirAReal(args[2])
        mag := math.Sqrt(x*x + y*y + z*z)
        if mag == 0 {
            return errors.New("❌ ERROR FATAL: no se puede normalizar un vector nulo")
        }
        return []float64{x / mag, y / mag, z / mag}
    })

    // --- Ángulo entre dos vectores (grados) ---
    reg("angulo_vectores", func(args ...interface{}) interface{} {
        if len(args) != 6 {
            return errors.New("❌ ERROR FATAL: angulo_vectores requiere 6 argumentos")
        }
        x1, _ := evaluador.ConvertirAReal(args[0])
        y1, _ := evaluador.ConvertirAReal(args[1])
        z1, _ := evaluador.ConvertirAReal(args[2])
        x2, _ := evaluador.ConvertirAReal(args[3])
        y2, _ := evaluador.ConvertirAReal(args[4])
        z2, _ := evaluador.ConvertirAReal(args[5])
        dot := x1*x2 + y1*y2 + z1*z2
        mag1 := math.Sqrt(x1*x1 + y1*y1 + z1*z1)
        mag2 := math.Sqrt(x2*x2 + y2*y2 + z2*z2)
        if mag1 == 0 || mag2 == 0 {
            return errors.New("❌ ERROR FATAL: no se puede calcular ángulo con vector nulo")
        }
        cos := dot / (mag1 * mag2)
        // Corrección numérica por posibles redondeos
        if cos > 1 {
            cos = 1
        } else if cos < -1 {
            cos = -1
        }
        return math.Acos(cos) * 180 / math.Pi
    })

    // --- Proyección de un vector sobre otro ---
    reg("proyeccion_vector", func(args ...interface{}) interface{} {
        if len(args) != 6 {
            return errors.New("❌ ERROR FATAL: proyeccion_vector requiere 6 argumentos")
        }
        ax, _ := evaluador.ConvertirAReal(args[0])
        ay, _ := evaluador.ConvertirAReal(args[1])
        az, _ := evaluador.ConvertirAReal(args[2])
        bx, _ := evaluador.ConvertirAReal(args[3])
        by, _ := evaluador.ConvertirAReal(args[4])
        bz, _ := evaluador.ConvertirAReal(args[5])
        dot := ax*bx + ay*by + az*bz
        magB2 := bx*bx + by*by + bz*bz
        if magB2 == 0 {
            return errors.New("❌ ERROR FATAL: no se puede proyectar sobre un vector nulo")
        }
        esc := dot / magB2
        return []float64{esc * bx, esc * by, esc * bz}
    })

    // --- Ecuación cuadrática ---
    reg("ecuacion_cuadratica", func(args ...interface{}) interface{} {
        if len(args) != 3 {
            return errors.New("❌ ERROR FATAL: ecuacion_cuadratica requiere 3 argumentos (a, b, c)")
        }
        a, _ := evaluador.ConvertirAReal(args[0])
        b, _ := evaluador.ConvertirAReal(args[1])
        c, _ := evaluador.ConvertirAReal(args[2])
        if a == 0 {
            return errors.New("❌ ERROR FATAL: coeficiente 'a' no puede ser 0")
        }
        disc := b*b - 4*a*c
        if disc < 0 {
            return errors.New("❌ ERROR FATAL: discriminante negativo, raíces complejas")
        }
        r1 := (-b + math.Sqrt(disc)) / (2 * a)
        r2 := (-b - math.Sqrt(disc)) / (2 * a)
        return []float64{r1, r2}
    })

    // --- Multiplicación de matrices 2x2 ---
    // A = [a b; c d], B = [e f; g h] → A*B = [ae+bg, af+bh; ce+dg, cf+dh]
    reg("multiplicar2x2", func(args ...interface{}) interface{} {
        if len(args) != 8 {
            return errors.New("❌ ERROR FATAL: multiplicar2x2 requiere 8 argumentos (a,b,c,d,e,f,g,h)")
        }
        a, _ := evaluador.ConvertirAReal(args[0])
        b, _ := evaluador.ConvertirAReal(args[1])
        c, _ := evaluador.ConvertirAReal(args[2])
        d, _ := evaluador.ConvertirAReal(args[3])
        e, _ := evaluador.ConvertirAReal(args[4])
        f, _ := evaluador.ConvertirAReal(args[5])
        g, _ := evaluador.ConvertirAReal(args[6])
        h, _ := evaluador.ConvertirAReal(args[7])
        return []float64{
            a*e + b*g, a*f + b*h,
            c*e + d*g, c*f + d*h,
        }
    })

    // --- Inversa de matriz 2x2 ---
    // A^{-1} = (1/det) * [d -b; -c a]
    reg("inversa2x2", func(args ...interface{}) interface{} {
        if len(args) != 4 {
            return errors.New("❌ ERROR FATAL: inversa2x2 requiere 4 argumentos")
        }
        a, _ := evaluador.ConvertirAReal(args[0])
        b, _ := evaluador.ConvertirAReal(args[1])
        c, _ := evaluador.ConvertirAReal(args[2])
        d, _ := evaluador.ConvertirAReal(args[3])
        det := a*d - b*c
        if det == 0 {
            return errors.New("❌ ERROR FATAL: la matriz no es invertible (determinante 0)")
        }
        invDet := 1 / det
        return []float64{
            d * invDet, -b * invDet,
            -c * invDet, a * invDet,
        }
    })

    // --- Resolver sistema 2x2 ---
    // a x + b y = e
    // c x + d y = f
    reg("resolver_sistema2x2", func(args ...interface{}) interface{} {
        if len(args) != 6 {
            return errors.New("❌ ERROR FATAL: resolver_sistema2x2 requiere 6 argumentos (a,b,c,d,e,f)")
        }
        a, _ := evaluador.ConvertirAReal(args[0])
        b, _ := evaluador.ConvertirAReal(args[1])
        c, _ := evaluador.ConvertirAReal(args[2])
        d, _ := evaluador.ConvertirAReal(args[3])
        e, _ := evaluador.ConvertirAReal(args[4])
        f, _ := evaluador.ConvertirAReal(args[5])
        det := a*d - b*c
        if det == 0 {
            return errors.New("❌ ERROR FATAL: sistema sin solución única (determinante 0)")
        }
        x := (e*d - b*f) / det
        y := (a*f - e*c) / det
        return []float64{x, y}
    })
}
