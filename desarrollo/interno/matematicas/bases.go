package matematicas

import (
    "errors"
    "fmt"
    "strconv"

    "nepa/desarrollo/interno/evaluador"
)

// InyectarBases agrega funciones de sistemas de numeración y operaciones bit a bit al contexto
func InyectarBases(ctx *evaluador.Contexto) {
    if ctx.Funciones == nil {
        ctx.Funciones = map[string]func(...interface{}) interface{}{}
    }

    reg := func(n string, f func(...interface{}) interface{}) {
        ctx.Funciones[n] = f
    }

    // --- Conversiones entre bases ---
    reg("decimal_a_binario", func(args ...interface{}) interface{} {
        if len(args) != 1 {
            return errors.New("❌ ERROR FATAL: decimal_a_binario requiere 1 argumento")
        }
        n, err := evaluador.ConvertirAReal(args[0])
        if err != nil {
            return err
        }
        return fmt.Sprintf("%b", int64(n))
    })

    reg("binario_a_decimal", func(args ...interface{}) interface{} {
        if len(args) != 1 {
            return errors.New("❌ ERROR FATAL: binario_a_decimal requiere 1 argumento")
        }
        s, ok := args[0].(string)
        if !ok {
            return errors.New("❌ ERROR FATAL: argumento debe ser cadena binaria")
        }
        val, err := strconv.ParseInt(s, 2, 64)
        if err != nil {
            return err
        }
        return val
    })

    reg("decimal_a_hexadecimal", func(args ...interface{}) interface{} {
        if len(args) != 1 {
            return errors.New("❌ ERROR FATAL: decimal_a_hexadecimal requiere 1 argumento")
        }
        n, err := evaluador.ConvertirAReal(args[0])
        if err != nil {
            return err
        }
        return fmt.Sprintf("%X", int64(n))
    })

    reg("hexadecimal_a_decimal", func(args ...interface{}) interface{} {
        if len(args) != 1 {
            return errors.New("❌ ERROR FATAL: hexadecimal_a_decimal requiere 1 argumento")
        }
        s, ok := args[0].(string)
        if !ok {
            return errors.New("❌ ERROR FATAL: argumento debe ser cadena hexadecimal")
        }
        val, err := strconv.ParseInt(s, 16, 64)
        if err != nil {
            return err
        }
        return val
    })

    reg("decimal_a_octal", func(args ...interface{}) interface{} {
        if len(args) != 1 {
            return errors.New("❌ ERROR FATAL: decimal_a_octal requiere 1 argumento")
        }
        n, err := evaluador.ConvertirAReal(args[0])
        if err != nil {
            return err
        }
        return fmt.Sprintf("%o", int64(n))
    })

    reg("octal_a_decimal", func(args ...interface{}) interface{} {
        if len(args) != 1 {
            return errors.New("❌ ERROR FATAL: octal_a_decimal requiere 1 argumento")
        }
        s, ok := args[0].(string)
        if !ok {
            return errors.New("❌ ERROR FATAL: argumento debe ser cadena octal")
        }
        val, err := strconv.ParseInt(s, 8, 64)
        if err != nil {
            return err
        }
        return val
    })

    // --- Operaciones bit a bit ---
    reg("bit_and", func(args ...interface{}) interface{} {
        if len(args) != 2 {
            return errors.New("❌ ERROR FATAL: bit_and requiere 2 argumentos")
        }
        a, _ := evaluador.ConvertirAReal(args[0])
        b, _ := evaluador.ConvertirAReal(args[1])
        return int64(a) & int64(b)
    })

    reg("bit_or", func(args ...interface{}) interface{} {
        if len(args) != 2 {
            return errors.New("❌ ERROR FATAL: bit_or requiere 2 argumentos")
        }
        a, _ := evaluador.ConvertirAReal(args[0])
        b, _ := evaluador.ConvertirAReal(args[1])
        return int64(a) | int64(b)
    })

    reg("bit_xor", func(args ...interface{}) interface{} {
        if len(args) != 2 {
            return errors.New("❌ ERROR FATAL: bit_xor requiere 2 argumentos")
        }
        a, _ := evaluador.ConvertirAReal(args[0])
        b, _ := evaluador.ConvertirAReal(args[1])
        return int64(a) ^ int64(b)
    })

    reg("bit_not", func(args ...interface{}) interface{} {
        if len(args) != 1 {
            return errors.New("❌ ERROR FATAL: bit_not requiere 1 argumento")
        }
        a, _ := evaluador.ConvertirAReal(args[0])
        return ^int64(a)
    })

    reg("bit_shift_left", func(args ...interface{}) interface{} {
        if len(args) != 2 {
            return errors.New("❌ ERROR FATAL: bit_shift_left requiere 2 argumentos")
        }
        a, _ := evaluador.ConvertirAReal(args[0])
        n, _ := evaluador.ConvertirAReal(args[1])
        return int64(a) << uint(n)
    })

    reg("bit_shift_right", func(args ...interface{}) interface{} {
        if len(args) != 2 {
            return errors.New("❌ ERROR FATAL: bit_shift_right requiere 2 argumentos")
        }
        a, _ := evaluador.ConvertirAReal(args[0])
        n, _ := evaluador.ConvertirAReal(args[1])
        return int64(a) >> uint(n)
    })
}
