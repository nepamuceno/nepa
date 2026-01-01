package ast

type Posicion struct {
    Linea   int
    Columna int
    Archivo string
}

type Base struct {
    Pos Posicion
}

// Nodo es una interfaz general para todos los elementos del árbol.
// IMPORTANTE: Se mantiene como interfaz vacía para asegurar la compatibilidad 
// de []ast.Nodo con la estructura de bloques indentados del evaluador.
type Nodo interface{}

type Importar struct {
    Base
    Nombre string
}

type Identificador struct {
    Base
    Nombre string
}

type Literal struct {
    Base
    Valor interface{}
}

type Asignacion struct {
    Base
    Nombre string
    Valor  Nodo
}

type OperacionBinaria struct {
    Base
    Izquierda Nodo
    Operador  string
    Derecha   Nodo
}

type LlamadaFuncion struct {
    Base
    Nombre string
    Args   []Nodo
}

type LlamadaModulo struct {
    Base
    Modulo  string
    Funcion string
    Args    []Nodo
}

type Si struct {
    Base
    Condicion Nodo
    Cuerpo    []Nodo // MODIFICADO: Ahora es una lista para soportar bloques indentados
    Sino      []Nodo // MODIFICADO: Ahora es una lista para soportar bloques indentados
}

type Mientras struct {
    Base
    Condicion Nodo
    Cuerpo    []Nodo // MODIFICADO: Ahora es una lista para soportar bloques indentados
}

type FuncionDef struct {
    Base
    Nombre     string
    Parametros []string
    Cuerpo     []Nodo // MODIFICADO: Ahora es una lista para soportar bloques indentados
}

type Retornar struct {
    Base
    Valor Nodo
}

type Para struct {
    Base
    Variable   string
    Origen     Nodo
    Fin        Nodo
    Incremento Nodo
    Cuerpo     []Nodo // MODIFICADO: Ahora es una lista para soportar bloques indentados
}
