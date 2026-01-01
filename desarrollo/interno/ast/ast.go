package ast

// Posicion representa la ubicación de un nodo en el código fuente
type Posicion struct {
	Linea   int
	Columna int
	Archivo string
}

// Base contiene la información común a todos los nodos
type Base struct {
	Pos Posicion
}

// Nodo es la interfaz base para todos los elementos del AST
type Nodo interface{}

// Importar representa la sentencia 'incluir' (Añadido para el motor de módulos)
type Importar struct {
	Base
	Nombre string
}

// Identificador representa nombres de variables o funciones
type Identificador struct {
	Base
	Nombre string
}

// Literal representa valores constantes (números, strings, booleanos)
type Literal struct {
	Base
	Valor interface{}
}

// Asignacion representa la creación o actualización de una variable
type Asignacion struct {
	Base
	Nombre string
	Valor  Nodo
}

// OperacionBinaria representa cálculos como suma, resta, comparaciones
type OperacionBinaria struct {
	Base
	Izquierda Nodo
	Operador  string
	Derecha   Nodo
}

// LlamadaFuncion representa la ejecución de una función nativa o de usuario
type LlamadaFuncion struct {
	Base
	Nombre string
	Args   []Nodo
}

// LlamadaModulo representa el acceso a funciones/variables de un SDK (ej. matematicas.raiz)
type LlamadaModulo struct {
	Base
	Modulo  string
	Funcion string
	Args    []Nodo
}

// Si representa la estructura condicional
type Si struct {
	Base
	Condicion Nodo
	Cuerpo    Nodo
	Sino      Nodo
}

// Mientras representa el bucle iterativo
type Mientras struct {
	Base
	Condicion Nodo
	Cuerpo    Nodo
}

// FuncionDef representa la declaración de una nueva función
type FuncionDef struct {
	Base
	Nombre     string
	Parametros []string
	Cuerpo     Nodo
}

// Retornar representa la sentencia de salida de una función
type Retornar struct {
	Base
	Valor Nodo
}

// Para representa el bucle universal (rangos, cadenas, matrices, archivos)
type Para struct {
    Base
    Variable   string // Nombre de la variable iteradora
    Origen     Nodo   // Inicio del rango o la colección a recorrer
    Fin        Nodo   // Límite final (opcional, usado en rangos)
    Incremento Nodo   // Salto dinámico (opcional)
    Cuerpo     Nodo   // Bloque de código a ejecutar
}
