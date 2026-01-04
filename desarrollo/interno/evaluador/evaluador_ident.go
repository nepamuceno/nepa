package evaluador

import (
	"errors"
	"fmt"
	"go/ast"
	"strings"
	"nepa/desarrollo/interno/administrador" 
)

// ErrIdentificadorNoExiste es el error base
var ErrIdentificadorNoExiste = errors.New("❌ ERROR FATAL: el identificador no existe")

func evaluarIdentificador(n *ast.Ident, ctx *Contexto) (interface{}, error) {
	nombre := strings.ToLower(strings.TrimSpace(n.Name))

	switch nombre {
	case "verdadero":
		return true, nil
	case "falso":
		return false, nil
	default:
		// 1. Intentar buscar en el Contexto local
		v, err := ctx.ObtenerVariable(nombre)
		
		// 2. Si hay error en el contexto, buscamos en el Administrador Global
		if err != nil {
			// En tu administrador, ObtenerVariable parece devolver (Variable, error)
			res, errGlobal := administrador.ObtenerVariable(nombre)
			
			// Si también hay error en el global, entonces no existe
			if errGlobal != nil {
				return nil, fmt.Errorf("%w → %s", ErrIdentificadorNoExiste, nombre)
			}
			v = res
		}

		// 3. Extraer el valor real (Interface)
		if interfaz, ok := v.(interface{ ValorComoInterface() interface{} }); ok {
			return interfaz.ValorComoInterface(), nil
		}

		return v, nil
	}
}
