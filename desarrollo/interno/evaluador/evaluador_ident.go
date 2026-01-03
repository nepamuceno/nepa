package evaluador

import (
    "errors"
    "go/ast"
    "strings"

    "nepa/desarrollo/interno/administrador"

)

// Error específico para identificadores inexistentes
var ErrIdentificadorNoExiste = errors.New("❌ ERROR FATAL: el identificador no existe")

// evalIdent maneja identificadores:
// - verdadero / falso → booleanos
// - nombres de variables registrados en el administrador
func evalIdent(n *ast.Ident) (interface{}, error) {
    nombre := strings.ToLower(strings.TrimSpace(n.Name))

    switch nombre {
    case "verdadero":
        return true, nil
    case "falso":
        return false, nil
    default:
        v, err := administrador.ObtenerVariable(nombre)
        if err != nil {
            return nil, errors.New(ErrIdentificadorNoExiste.Error() + " → " + nombre)
        }
        return v.ValorComoInterface(), nil
    }
}
