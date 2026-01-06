package conversiones

import (
    "fmt"
    "strconv"
    "nepa/desarrollo/interno/evaluador"
    "nepa/desarrollo/interno/administrador"
)

const ayudaConvertirEntero = `
convertir_entero(valor) → entero
Alias: a_entero(valor), convertir.entero(valor)

Convierte un valor a entero.
Ejemplo:
    convertir_entero("123") → 123
    convertir_entero(3.14) → 3
`

func fnConvertirEntero(args ...interface{}) interface{} {

	for i, a := range args { 
		fmt.Printf("DEBUG arg[%d]: tipo=%T valor=%#v\n", i, a, a) 
	}


    if len(args) < 1 {
        return evaluador.NuevaErrorConversion("convertir_entero", ayudaConvertirEntero, nil)
    }
	// 👇 Debug: imprime todos los argumentos con su tipo 
	for i, a := range args { 
		fmt.Printf("\n\nDEBUG desarrollo/comandos/convertir_entero.go:\narg[%d]: tipo=%T valor=%#v\n\n\n", i, a, a) 
	}
	
	 
    switch v := args[0].(type) {
    case int:
        return v
    case int64:
        return int(v)
    case uint:
        return int(v)
    case uint64:
        return int(v)
    case float64:
        return int(v)
    case string:
        n, err := strconv.Atoi(v)
        if err != nil {
            return evaluador.NuevaErrorConversion("convertir_entero", ayudaConvertirEntero, v)
        }
        return n
    case administrador.Variable:
        val := v.ValorComoInterface()
        switch vv := val.(type) {
        case string:
            n, err := strconv.Atoi(vv)
            if err != nil {
                return evaluador.NuevaErrorConversion("convertir_entero", ayudaConvertirEntero, vv)
            }
            return n
        case int:
            return vv
        case float64:
            return int(vv)
        default:
            s := fmt.Sprintf("%v", vv)
            n, err := strconv.Atoi(s)
            if err != nil {
                return evaluador.NuevaErrorConversion("convertir_entero", ayudaConvertirEntero, vv)
            }
            return n
        }
    default:
        s := fmt.Sprintf("%v", v)
        n, err := strconv.Atoi(s)
        if err != nil {
            return evaluador.NuevaErrorConversion("convertir_entero", ayudaConvertirEntero, v)
        }
        return n
    }
}

func RegistrarConvertirEntero(ctx *evaluador.Contexto) {
    evaluador.Funciones["convertir_entero"] = func(args ...interface{}) (interface{}, error) {
        r := fnConvertirEntero(args...)
        if err, ok := r.(error); ok {
            return nil, err
        }
        return r, nil
    }
    evaluador.Funciones["a_entero"] = evaluador.Funciones["convertir_entero"]
    evaluador.Funciones["convertir.entero"] = evaluador.Funciones["convertir_entero"]
}
