package evaluador

import "math"

// fnPromedio calcula el promedio de una lista de valores reales.
func fnPromedio(args ...interface{}) interface{} {
    if len(args) < 1 {
        return nil
    }
    lista, err := ConvertirAListaReal(args[0])
    if err != nil {
        return nil
    }
    if len(lista) == 0 {
        return nil
    }
    suma := 0.0
    for _, v := range lista {
        suma += v
    }
    return suma / float64(len(lista))
}

// fnDesviacion calcula la desviación estándar de una lista de valores reales.
func fnDesviacion(args ...interface{}) interface{} {
    if len(args) < 1 {
        return nil
    }
    lista, err := ConvertirAListaReal(args[0])
    if err != nil {
        return nil
    }
    if len(lista) == 0 {
        return nil
    }

    // Promedio
    prom := 0.0
    for _, v := range lista {
        prom += v
    }
    prom /= float64(len(lista))

    // Varianza
    suma := 0.0
    for _, v := range lista {
        d := v - prom
        suma += d * d
    }
    varianza := suma / float64(len(lista))

    // Desviación estándar
    return math.Sqrt(varianza)
}

// RegistrarFuncionesEstadistica agrega las funciones estadísticas al contexto.
func RegistrarFuncionesEstadistica(ctx *Contexto) {
    ctx.Funciones["promedio"] = fnPromedio
    ctx.Funciones["desviacion"] = fnDesviacion
}
