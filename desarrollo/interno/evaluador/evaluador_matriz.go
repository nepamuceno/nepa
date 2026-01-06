package evaluador

import "fmt"

// fnDimension devuelve las dimensiones de una matriz como [filas, columnas].
func fnDimension(args ...interface{}) interface{} {
    if len(args) < 1 {
        return nil
    }
    matriz, ok := args[0].([][]interface{})
    if !ok {
        return nil
    }
    filas := len(matriz)
    columnas := 0
    if filas > 0 {
        columnas = len(matriz[0])
    }
    return []int{filas, columnas}
}

// fnTranspuesta devuelve la transpuesta de una matriz.
func fnTranspuesta(args ...interface{}) interface{} {
    if len(args) < 1 {
        return nil
    }
    matriz, ok := args[0].([][]interface{})
    if !ok {
        return nil
    }
    if len(matriz) == 0 {
        return [][]interface{}{}
    }
    filas := len(matriz)
    columnas := len(matriz[0])
    transpuesta := make([][]interface{}, columnas)
    for i := 0; i < columnas; i++ {
        transpuesta[i] = make([]interface{}, filas)
        for j := 0; j < filas; j++ {
            transpuesta[i][j] = matriz[j][i]
        }
    }
    return transpuesta
}

// fnElemento devuelve el elemento en la posición [fila, columna] de la matriz.
func fnElemento(args ...interface{}) interface{} {
    if len(args) < 3 {
        return nil
    }
    matriz, ok := args[0].([][]interface{})
    if !ok {
        return nil
    }
    fila, okFila := args[1].(int)
    columna, okCol := args[2].(int)
    if !okFila || !okCol {
        return nil
    }
    if fila < 0 || fila >= len(matriz) {
        return fmt.Errorf("índice de fila fuera de rango")
    }
    if columna < 0 || columna >= len(matriz[fila]) {
        return fmt.Errorf("índice de columna fuera de rango")
    }
    return matriz[fila][columna]
}

// RegistrarFuncionesMatriz agrega las funciones relacionadas con matrices al contexto.
func RegistrarFuncionesMatriz(ctx *Contexto) {
    ctx.Funciones["dimension"] = fnDimension
    ctx.Funciones["transpuesta"] = fnTranspuesta
    ctx.Funciones["elemento"] = fnElemento
}
