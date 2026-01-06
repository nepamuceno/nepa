package comandos

// Este paquete centraliza la carga de todos los comandos.
// Al importar "nepa/desarrollo/comandos" en main.go,
// se ejecutarán los init() de cada subcomando.

import (
    _ "nepa/desarrollo/comandos/variable"
    _ "nepa/desarrollo/comandos/imprimir"
    _ "nepa/desarrollo/comandos/asignar"
    _ "nepa/desarrollo/comandos/bloque"
    _ "nepa/desarrollo/comandos/expresion"
    _ "nepa/desarrollo/comandos/llamada"
    _ "nepa/desarrollo/comandos/inyectar_todas_variables" // inyecta todos los tipos de variables
    _ "nepa/desarrollo/comandos/conversiones"             // 👈 nuevo: dispara init() de conversiones
    _ "nepa/desarrollo/interno/matematicas"               // 🔥 dispara init() matematicas
)
