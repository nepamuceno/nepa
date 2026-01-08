package nucleo

// CATÁLOGO CENTRALIZADO DE ERRORES EN NEPA
// Cada subsistema tiene su rango de códigos reservado.
// Todos los mensajes están en Español Mexicano y con placeholders.

var MENSAJES_ERROR = map[int]string{
    // --- SÍMBOLOS (1000–1099) ---
    1000: "%s[%d]: #%d Archivo no encontrado [%s]",
    1001: "%s[%d]: #%d Error de lectura: %s",
    1002: "%s[%d]: #%d Símbolo inesperado [%s]",
    1003: "%s[%d]: #%d Fin de archivo inesperado",
    1004: "%s[%d]: #%d Carácter inválido [%s]",

    // --- SINTAXIS (2000–2099) ---
    2000: "%s[%d]: #%d Sintaxis inválida en [%s]",
    2001: "%s[%d]: #%d Bloque sin cierre correcto",
    2002: "%s[%d]: #%d Indentación incorrecta",
    2003: "%s[%d]: #%d Expresión incompleta [%s]",
    2004: "%s[%d]: #%d Función no reconocida [%s]",

    // --- ESTRUCTURAS (2100–2199) ---
    2100: "%s[%d]: #%d Variable no definida [%s]",
    2101: "%s[%d]: #%d Variable ya existe [%s]",
    2102: "%s[%d]: #%d Tipo de variable inválido [%s]",
    2103: "%s[%d]: #%d Asignación inválida a variable [%s]",

    // --- CONSTANTES (2200–2299) ---
    2200: "%s[%d]: #%d Constante no modificable [%s]",
    2201: "%s[%d]: #%d Constante ya existe [%s]",

    // --- PUNTEROS (2300–2399) ---
    2300: "%s[%d]: #%d Puntero inválido [%s]",
    2301: "%s[%d]: #%d Puntero fuera de rango [%s]",
    2302: "%s[%d]: #%d Puntero nulo [%s]",

    // --- LISTAS Y DICCIONARIOS (2400–2499) ---
    2400: "%s[%d]: #%d Clave duplicada [%s]",
    2401: "%s[%d]: #%d Índice fuera de rango [%d]",
    2402: "%s[%d]: #%d Tipo de elemento inválido [%s]",

    // --- COMANDOS (2500–2599) ---
    2500: "%s[%d]: #%d Comando no reconocido [%s]",
    2501: "%s[%d]: #%d Argumentos insuficientes para comando [%s]",
    2502: "%s[%d]: #%d Argumentos inválidos para comando [%s]",

    // --- BLOQUES Y CONDICIONALES (2600–2699) ---
    2600: "%s[%d]: #%d Condicional inválido [%s]",
    2601: "%s[%d]: #%d Bloque vacío no permitido",
    2602: "%s[%d]: #%d Expresión lógica inválida [%s]",

    // --- EVALUADOR (5000–5099) ---
    5000: "%s[%d]: #%d Error en evaluación de expresión [%s]",
    5001: "%s[%d]: #%d División entre cero",
    5002: "%s[%d]: #%d Tipo incompatible en operación [%s]",
    5003: "%s[%d]: #%d Función no retornó valor [%s]",

    // --- ENTORNO (6000–6099) ---
    6000: "%s[%d]: #%d Variable no accesible en este entorno [%s]",
    6001: "%s[%d]: #%d Conflicto de nombres en entorno [%s]",
    6002: "%s[%d]: #%d Entorno no inicializado",
    6099: "%s[%d]: #%d Depuración desconocida", // genérico para depuración

    // --- DETALLE (7000–7099) ---
    6999: "%s[%d]: #%d Detalle desconocido", // genérico para detalle

    // --- CORE GENÉRICO (9000–9999) ---
    9000: "%s[%d]: #%d Uso incorrecto del intérprete: %s",
    9999: "%s[%d]: #%d Error desconocido", // genérico para errores
}
