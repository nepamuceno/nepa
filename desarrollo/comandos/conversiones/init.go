package conversiones

import "nepa/desarrollo/interno/evaluador"

// init registra automáticamente las conversiones cuando el paquete se importa.
func init() {
    evaluador.RegistrarModulo(func(ctx *evaluador.Contexto) {
        RegistrarConversionesBasicas(ctx)
    })
}
