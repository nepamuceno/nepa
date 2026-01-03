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

    // --- Operadores lógicos (en español) ---
    "y",     // and
    "o",     // or
    "no",    // not

    // --- Constantes matemáticas y físicas nativas ---
    "pi",        // π
    "e",         // número de Euler
    "phi",       // número áureo
    "gravedad",  // aceleración de la gravedad
    "c",         // velocidad de la luz
    "h",         // constante de Planck
    "k",         // constante de Boltzmann

    // --- Literales especiales ---
    "verdadero",
    "falso",
    "nulo",

    // --- Funciones matemáticas básicas ---
    "abs",
    "raiz",       // sqrt
    "potencia",   // pow
    "exp",
    "log",
    "log10",
    "seno",       // sin
    "coseno",     // cos
    "tangente",   // tan
    "arcseno",    // asin
    "arccoseno",  // acos
    "arctan",     // atan
    "senh",       // sinh
    "cosh",       // cosh
    "tanh",       // tanh

    // --- Estadística ---
    "media",
    "mediana",
    "moda",
    "varianza",
    "desviacion",
    "maximo",
    "minimo",

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

    // --- Operaciones bit a bit ---
    "bit_y",          // and
    "bit_o",          // or
    "bit_xor",
    "bit_no",         // not
    "bit_desplazar_izq",
    "bit_desplazar_der",

    // --- Primitivas del lenguaje ---
    "imprimir", // salida estándar
    "ejecutar", // invocar otro programa .nepa
    "ayuda",    // sistema de ayuda/documentación
    "asignar",  // asignación de variables
    "variable", // declaración de variables

    // --- Definiciones de espacio ---
    "global",   // variables compartidas
    "const",    // constantes inmutables
}
