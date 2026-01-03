package variable

import (
    "fmt"
    "regexp"
    "strings"

    "nepa/desarrollo/interno/administrador"
    "nepa/desarrollo/interno/bloque"
    "nepa/desarrollo/interno/evaluador"
    "nepa/desarrollo/interno/parser"
)

// Validación de nombres
func esNombreValido(nombre string) bool {
    re := regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*$`)
    if !re.MatchString(nombre) {
        return false
    }
    for _, r := range bloque.PalabrasReservadas {
        if nombre == r {
            return false
        }
    }
    return true
}

func init() {
    evaluador.Registrar("variable", func(n parser.Nodo, ctx *evaluador.Contexto) {
        // Tipo desde Args[0]; si no viene, lo marcamos como desconocido
        tipo := ""
        if len(n.Args) > 0 {
            if t, ok := n.Args[0].(string); ok {
                tipo = strings.ToLower(strings.TrimSpace(t))
            }
        }
        if tipo == "" {
            fmt.Printf("⚠️ Tipo de variable no especificado para '%s'\n", n.Nombre)
            return
        }

        // Soportar múltiples nombres: "x,y,z"
        nombres := strings.Split(n.Nombre, ",")
        for i := range nombres {
            nombres[i] = strings.TrimSpace(nombres[i])
            if !esNombreValido(nombres[i]) {
                fmt.Printf("⚠️ Nombre de variable inválido: %s\n", nombres[i])
                return
            }
        }

        // Buscar constructor del tipo (universal: hoy puede existir solo 'bit')
        constructor, ok := administrador.Constructores[tipo]
        if !ok {
            fmt.Printf("⚠️ Tipo de variable no implementado: %s (omitido)\n", tipo)
            return
        }

        // Crear y registrar cada variable
        for _, nombre := range nombres {
            v, err := constructor(nombre, n.Valor)
            if err != nil {
                fmt.Printf("⚠️ Error creando variable '%s' (%s): %v\n", nombre, tipo, err)
                continue
            }
            administrador.RegistrarVariable(nombre, v)
            fmt.Printf("✔ Variable creada: %s\n", v.Mostrar())
        }
    })
}
