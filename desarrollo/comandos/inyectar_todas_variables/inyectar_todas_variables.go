package comandos

import (
    "nepa/desarrollo/interno/administrador"

    // Importa aquí todos los paquetes de tipos
    "nepa/desarrollo/interno/variables/bit"
  
    // "nepa/desarrollo/comandos/entero"
    // "nepa/desarrollo/comandos/cadena"
    // etc...
)

func init() {
    administrador.RegistrarConstructores(map[string]func(string, interface{}) (administrador.Variable, error){
        "bit":    bit.CrearBit,
        // "entero": entero.CrearEntero,
        // "cadena": cadena.CrearCadena,
        // agrega más tipos aquí
    })
}
