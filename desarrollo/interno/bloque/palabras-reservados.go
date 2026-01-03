package bloque

var PalabrasReservadas = []string{
    // --- Operadores lógicos ---
    "y_es", "o_es", "no_es",

    // --- Constantes matemáticas y físicas ---
    "pi", "euler", "phi", "gravedad", "luz", "planck", "boltzmann",

    // --- Literales especiales ---
    "verdadero", "falso", "nulo",

    // --- Funciones matemáticas internas ---
    "abs", "raiz", "potencia", "exp", "log", "log10",
    "seno", "coseno", "tangente", "arcseno", "arccoseno", "arctan",
    "senh", "cosh", "tanh",

    // --- Estadística ---
    "media", "mediana", "moda", "varianza", "desviacion",
    "maximo", "minimo",

    // --- Física ---
    "energia", "trabajo", "densidad", "presion", "fuerza",
    "momento", "impulso",

    // --- Finanzas ---
    "interes_simple", "interes_compuesto", "valor_presente", "valor_futuro",

    // --- Operaciones bit a bit ---
    "bit_y", "bit_o", "bit_xor", "bit_no",
    "bit_desplazar_izq", "bit_desplazar_der",

    // --- Tipos de variables ---
    "entero", "decimal", "binario", "cadena",
    "booleano", "lista", "mapa", "objeto",

    // --- Manejo de cadenas (futuro) ---
    "sin_espacio", "espacios_izquierda", "espacios_derecha",
    "mayusculas", "minusculas", "longitud",
    "subcadena", "reemplazar", "dividir", "unir",
}
