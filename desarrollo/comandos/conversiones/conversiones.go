package conversiones

import "nepa/desarrollo/interno/evaluador"

// RegistrarConversionesBasicas registra todas las funciones de conversión y punteros en el contexto.
func RegistrarConversionesBasicas(ctx *evaluador.Contexto) {
    // Conversión básica de tipos primarios
    RegistrarConvertirEntero(ctx)
    RegistrarConvertirReal(ctx)
    RegistrarConvertirCadena(ctx)
    RegistrarConvertirBooleano(ctx)

    // Conversión utilitaria y formatos
    RegistrarConvertirBinario(ctx)
    RegistrarConvertirHexadecimal(ctx)
    RegistrarConvertirFecha(ctx)
    RegistrarConvertirHora(ctx)
    RegistrarConvertirTiempo(ctx)
    RegistrarConvertirMatriz(ctx)

    // Conversión de punteros como comando explícito
    RegistrarConvertirPuntero(ctx)

    // Funciones explícitas de puntero (además de la sintaxis &x y *p)
    evaluador.RegistrarFuncionesPuntero(ctx)
}

// 🔧 Este init conecta el módulo al ciclo global del evaluador
func init() {
    evaluador.RegistrarModulo(RegistrarConversionesBasicas)
}
