package evaluador

import (
    "fmt"
)

// fnBinario convierte un número real a su representación binaria.
func fnBinario(args ...interface{}) interface{} {
    if len(args) < 1 {
        return NuevaErrorConversion("binario", ayudaBinario, nil)
    }
    n, err := ConvertirAReal(args[0])
    if err != nil {
        return NuevaErrorConversion("binario", ayudaBinario, args[0])
    }
    return fmt.Sprintf("%b", int(n))
}

// fnHexadecimal convierte un número real a su representación hexadecimal.
func fnHexadecimal(args ...interface{}) interface{} {
    if len(args) < 1 {
        return NuevaErrorConversion("hexadecimal", ayudaHexadecimal, nil)
    }
    n, err := ConvertirAReal(args[0])
    if err != nil {
        return NuevaErrorConversion("hexadecimal", ayudaHexadecimal, args[0])
    }
    return fmt.Sprintf("%x", int(n))
}

// fnCelsiusAFarenheit convierte grados Celsius a Farenheit.
func fnCelsiusAFarenheit(args ...interface{}) interface{} {
    if len(args) < 1 {
        return NuevaErrorConversion("celsius_a_farenheit", ayudaCelsiusAFarenheit, nil)
    }
    n, err := ConvertirAReal(args[0])
    if err != nil {
        return NuevaErrorConversion("celsius_a_farenheit", ayudaCelsiusAFarenheit, args[0])
    }
    return (n * 9.0 / 5.0) + 32.0
}

// fnFarenheitACelsius convierte grados Farenheit a Celsius.
func fnFarenheitACelsius(args ...interface{}) interface{} {
    if len(args) < 1 {
        return NuevaErrorConversion("farenheit_a_celsius", ayudaFarenheitACelsius, nil)
    }
    n, err := ConvertirAReal(args[0])
    if err != nil {
        return NuevaErrorConversion("farenheit_a_celsius", ayudaFarenheitACelsius, args[0])
    }
    return (n - 32.0) * 5.0 / 9.0
}

// Ayudas integradas para cada función
const ayudaBinario = `
binario(valor) → cadena
Convierte un número real o entero a su representación binaria.
Ejemplo: binario(10) → "1010"
`

const ayudaHexadecimal = `
hexadecimal(valor) → cadena
Convierte un número real o entero a su representación hexadecimal.
Ejemplo: hexadecimal(255) → "ff"
`

const ayudaCelsiusAFarenheit = `
celsius_a_farenheit(valor) → real
Convierte grados Celsius a Farenheit.
Ejemplo: celsius_a_farenheit(0) → 32.0
`

const ayudaFarenheitACelsius = `
farenheit_a_celsius(valor) → real
Convierte grados Farenheit a Celsius.
Ejemplo: farenheit_a_celsius(32) → 0.0
`

// RegistrarFuncionesConversiones agrega las funciones de conversión al contexto.
func RegistrarFuncionesConversiones(ctx *Contexto) {
    ctx.Funciones["binario"] = fnBinario
    ctx.Funciones["hexadecimal"] = fnHexadecimal
    ctx.Funciones["celsius_a_farenheit"] = fnCelsiusAFarenheit
    ctx.Funciones["farenheit_a_celsius"] = fnFarenheitACelsius
}
