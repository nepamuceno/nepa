package matematicas

import (
    "errors"
    "fmt"
    "math"

    "nepa/desarrollo/interno/evaluador"
)

// InyectarFisica agrega funciones físicas al contexto
func InyectarFisica(ctx *evaluador.Contexto) {
    if ctx.Funciones == nil {
        ctx.Funciones = map[string]func(...interface{}) interface{}{}
    }

    reg := func(n string, f func(...interface{}) interface{}) {
        ctx.Funciones[n] = f
    }

    // --- Energía y movimiento ---
    reg("energia_relativista", func(args ...interface{}) interface{} {
        if len(args) != 1 {
            return errors.New("❌ ERROR FATAL: energia_relativista requiere 1 argumento (masa)")
        }
        m, err := evaluador.ConvertirAReal(args[0])
        if err != nil {
            return err
        }
        return m * math.Pow(299792458.0, 2)
    })

    reg("energia_cinetica", func(args ...interface{}) interface{} {
        if len(args) != 2 {
            return errors.New("❌ ERROR FATAL: energia_cinetica requiere 2 argumentos (masa, velocidad)")
        }
        m, err := evaluador.ConvertirAReal(args[0])
        if err != nil {
            return err
        }
        v, err := evaluador.ConvertirAReal(args[1])
        if err != nil {
            return err
        }
        return 0.5 * m * math.Pow(v, 2)
    })

    reg("trabajo", func(args ...interface{}) interface{} {
        if len(args) != 3 {
            return errors.New("❌ ERROR FATAL: trabajo requiere 3 argumentos (fuerza, distancia, ángulo)")
        }
        f, err := evaluador.ConvertirAReal(args[0])
        if err != nil {
            return err
        }
        d, err := evaluador.ConvertirAReal(args[1])
        if err != nil {
            return err
        }
        ang, err := evaluador.ConvertirAReal(args[2])
        if err != nil {
            return err
        }
        return f * d * math.Cos(ang*math.Pi/180)
    })

    reg("energia_potencial", func(args ...interface{}) interface{} {
        if len(args) != 3 {
            return errors.New("❌ ERROR FATAL: energia_potencial requiere 3 argumentos (masa, gravedad, altura)")
        }
        m, err := evaluador.ConvertirAReal(args[0])
        if err != nil {
            return err
        }
        g, err := evaluador.ConvertirAReal(args[1])
        if err != nil {
            return err
        }
        h, err := evaluador.ConvertirAReal(args[2])
        if err != nil {
            return err
        }
        return m * g * h
    })

    // --- Proyectiles ---
    reg("caida_libre", func(args ...interface{}) interface{} {
        if len(args) != 1 {
            return errors.New("❌ ERROR FATAL: caida_libre requiere 1 argumento (tiempo)")
        }
        t, err := evaluador.ConvertirAReal(args[0])
        if err != nil {
            return err
        }
        return 0.5 * 9.80665 * math.Pow(t, 2)
    })

    reg("proyectil_pos", func(args ...interface{}) interface{} {
        if len(args) != 3 {
            return errors.New("❌ ERROR FATAL: proyectil_pos requiere 3 argumentos (velocidad inicial, ángulo, tiempo)")
        }
        v0, _ := evaluador.ConvertirAReal(args[0])
        ang, _ := evaluador.ConvertirAReal(args[1])
        t, _ := evaluador.ConvertirAReal(args[2])
        rad := ang * math.Pi / 180
        x := v0 * math.Cos(rad) * t
        y := (v0 * math.Sin(rad) * t) - (0.5 * 9.80665 * t * t)
        return fmt.Sprintf("%f,%f", x, y)
    })

    reg("proyectil_x", func(args ...interface{}) interface{} {
        if len(args) != 3 {
            return errors.New("❌ ERROR FATAL: proyectil_x requiere 3 argumentos (velocidad inicial, ángulo, tiempo)")
        }
        v0, _ := evaluador.ConvertirAReal(args[0])
        ang, _ := evaluador.ConvertirAReal(args[1])
        t, _ := evaluador.ConvertirAReal(args[2])
        return v0 * math.Cos(ang*math.Pi/180) * t
    })

    reg("proyectil_y", func(args ...interface{}) interface{} {
        if len(args) != 3 {
            return errors.New("❌ ERROR FATAL: proyectil_y requiere 3 argumentos (velocidad inicial, ángulo, tiempo)")
        }
        v0, _ := evaluador.ConvertirAReal(args[0])
        ang, _ := evaluador.ConvertirAReal(args[1])
        t, _ := evaluador.ConvertirAReal(args[2])
        rad := ang * math.Pi / 180
        return (v0 * math.Sin(rad) * t) - (0.5 * 9.80665 * t * t)
    })

    // --- Física general ---
    reg("densidad", func(args ...interface{}) interface{} {
        if len(args) != 2 {
            return errors.New("❌ ERROR FATAL: densidad requiere 2 argumentos (masa, volumen)")
        }
        m, _ := evaluador.ConvertirAReal(args[0])
        v, _ := evaluador.ConvertirAReal(args[1])
        if v == 0 {
            return errors.New("❌ ERROR FATAL: volumen no puede ser 0")
        }
        return m / v
    })

    reg("presion_gas", func(args ...interface{}) interface{} {
        if len(args) != 3 {
            return errors.New("❌ ERROR FATAL: presion_gas requiere 3 argumentos (n moles, temperatura, volumen)")
        }
        n, _ := evaluador.ConvertirAReal(args[0])
        T, _ := evaluador.ConvertirAReal(args[1])
        V, _ := evaluador.ConvertirAReal(args[2])
        if V == 0 {
            return errors.New("❌ ERROR FATAL: volumen no puede ser 0")
        }
        R := 8.314462
        return (n * R * T) / V
    })

    reg("fuerza_gravitatoria", func(args ...interface{}) interface{} {
        if len(args) != 3 {
            return errors.New("❌ ERROR FATAL: fuerza_gravitatoria requiere 3 argumentos (masa1, masa2, distancia)")
        }
        m1, _ := evaluador.ConvertirAReal(args[0])
        m2, _ := evaluador.ConvertirAReal(args[1])
        r, _ := evaluador.ConvertirAReal(args[2])
        if r == 0 {
            return errors.New("❌ ERROR FATAL: distancia no puede ser 0")
        }
        G := 6.67430e-11
        return G * (m1 * m2) / math.Pow(r, 2)
    })

    reg("ley_ohm_v", func(args ...interface{}) interface{} {
        if len(args) != 2 {
            return errors.New("❌ ERROR FATAL: ley_ohm_v requiere 2 argumentos (corriente, resistencia)")
        }
        I, _ := evaluador.ConvertirAReal(args[0])
        R, _ := evaluador.ConvertirAReal(args[1])
        return I * R
    })

    reg("magnitud_vector", func(args ...interface{}) interface{} {
        if len(args) != 3 {
            return errors.New("❌ ERROR FATAL: magnitud_vector requiere 3 argumentos (x, y, z)")
        }
        x, _ := evaluador.ConvertirAReal(args[0])
        y, _ := evaluador.ConvertirAReal(args[1])
        z, _ := evaluador.ConvertirAReal(args[2])
        return math.Sqrt(x*x + y*y + z*z)
    })
}
