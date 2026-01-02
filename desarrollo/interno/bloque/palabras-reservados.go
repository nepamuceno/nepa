package bloque

// PalabrasReservadas contiene todas las palabras reservadas globales del lenguaje Nepa.
// Se usan en sintaxis.go para validar que no se asignen ni se redefinan.
var PalabrasReservadas = []string{
    // --- Control de flujo ---
    "si",
    "sino",
    "mientras",
    "para",

    // --- Definiciones y módulos ---
    "funcion",
    "modulo",
    "retornar",
    "importar",

    // --- Sintaxis de bucles ---
    "desde",
    "hasta",
    "incremento",

    // --- Operadores lógicos ---
    "and",
    "or",
    "not",

    // --- Constantes matemáticas y físicas nativas ---
    "pi",       // π
    "e",        // número de Euler
    "phi",      // número áureo
    "gravedad", // aceleración de la gravedad
    "c",        // velocidad de la luz
    "h",        // constante de Planck
    "k",        // constante de Boltzmann

    // --- Literales especiales ---
    "true",
    "false",
    "nulo",

    // --- Funciones matemáticas básicas (core_math.go) ---
    "abs",
    "sqrt",
    "pow",
    "exp",
    "log",
    "log10",
    "sin",
    "cos",
    "tan",
    "asin",
    "acos",
    "atan",
    "sinh",
    "cosh",
    "tanh",

    // --- Estadística ---
    "media",
    "mediana",
    "moda",
    "varianza",
    "desviacion",
    "max",
    "min",

    // --- Física ---
    "energia",
    "trabajo",
    "potencia",
    "densidad",
    "presion",
    "fuerza",
    "momento",
    "impulso",

    // --- Finanzas ---
    "interes_simple",
    "interes_compuesto",
    "valor_presente",
    "valor_futuro",

    // --- Bitwise ---
    "bit_and",
    "bit_or",
    "bit_xor",
    "bit_not",
    "bit_shift_left",
    "bit_shift_right",

    // --- Primitivas del lenguaje ---
    "imprimir", // salida estándar
    "ejecutar", // invocar otro programa .nepa
    "ayuda",    // sistema de ayuda/documentación

    // --- Definiciones de espacio ---
    "global",   // variables compartidas
    "const",    // constantes inmutables
}
