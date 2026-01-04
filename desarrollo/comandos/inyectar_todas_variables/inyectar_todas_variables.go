package inyectar_todas_variables

import (
	"nepa/desarrollo/interno/administrador"
	"nepa/desarrollo/interno/variables/bit"
	"nepa/desarrollo/interno/variables/booleano"
	"nepa/desarrollo/interno/variables/cadena"
	"nepa/desarrollo/interno/variables/caracter"
	"nepa/desarrollo/interno/variables/complejo"
	"nepa/desarrollo/interno/variables/decimal"
	"nepa/desarrollo/interno/variables/diccionario"
	"nepa/desarrollo/interno/variables/entero"
	"nepa/desarrollo/interno/variables/fecha"
	"nepa/desarrollo/interno/variables/hora"
	"nepa/desarrollo/interno/variables/lista"
	"nepa/desarrollo/interno/variables/matriz"
	"nepa/desarrollo/interno/variables/objeto"
	"nepa/desarrollo/interno/variables/puntero"
	"nepa/desarrollo/interno/variables/real"
	"nepa/desarrollo/interno/variables/texto"
	"nepa/desarrollo/interno/variables/tiempo"
)

func init() {
	administrador.RegistrarConstructores(map[string]func(string, interface{}) (administrador.Variable, error){
		"bit":         bit.CrearBit,
		"booleano":    booleano.CrearBooleano,
		"cadena":      cadena.CrearCadena,
		"caracter":    caracter.CrearCaracter,
		"complejo":    complejo.CrearComplejo,
		"decimal":     decimal.CrearDecimal,
		"diccionario": diccionario.CrearDiccionario,
		"entero":      entero.CrearEntero,
		"fecha":       fecha.CrearFecha,
		"hora":        hora.CrearHora,
		"lista":       lista.CrearLista,
		"matriz":      matriz.CrearMatriz,
		"objeto":      objeto.CrearObjeto,
		"puntero":     puntero.CrearPuntero,
		"real":        real.CrearReal,
		"texto":       texto.CrearTexto,
		"tiempo":      tiempo.CrearTiempo,
	})
}
