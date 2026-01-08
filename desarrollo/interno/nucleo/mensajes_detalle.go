package nucleo

// CATÁLOGO CENTRALIZADO DE MENSAJES DE DETALLE EN NEPA
// Niveles progresivos:
//   Nivel 1 → Básico (entrada general)
//   Nivel 2 → Ubicación exacta
//   Nivel 3 → Ubicación + argumentos
//   Nivel 4 → Todos los detalles (contexto, valores, estado interno)

var MENSAJES_DETALLE = map[int]string{
    // --- SÍMBOLOS (1000–1099) ---
    1000: "Símbolos: iniciando revisión de archivo [%s]",
    1001: "Símbolos: revisando símbolo [%s] en línea %d",
    1002: "Símbolos: símbolo [%s] en línea %d con valor [%s]",
    1003: "Símbolos: símbolo [%s] en línea %d con valor [%s] y estado [%s]",

    // --- SINTAXIS (2000–2099) ---
    2000: "Sintaxis: iniciando bloque [%s]",
    2001: "Sintaxis: procesando instrucción [%s] en línea %d",
    2002: "Sintaxis: instrucción [%s] en línea %d con argumentos [%s]",
    2003: "Sintaxis: instrucción [%s] en línea %d con argumentos [%s] y contexto [%s]",

    // --- ESTRUCTURAS (2100–2199) ---
    2100: "Estructuras: detectada declaración [%s]",
    2101: "Estructuras: asignando [%s] en línea %d",
    2102: "Estructuras: asignando [%s] en línea %d con valor [%s]",
    2103: "Estructuras: asignando [%s] en línea %d con valor [%s] y tipo [%s]",

    // --- CONSTANTES (2200–2299) ---
    2200: "Constantes: detectada constante [%s]",
    2201: "Constantes: constante [%s] en línea %d",
    2202: "Constantes: constante [%s] en línea %d con valor [%s]",
    2203: "Constantes: constante [%s] en línea %d con valor [%s] y tipo [%s]",

    // --- PUNTEROS (2300–2399) ---
    2300: "Punteros: detectado puntero [%s]",
    2301: "Punteros: puntero [%s] en línea %d",
    2302: "Punteros: puntero [%s] en línea %d apunta a [%s]",
    2303: "Punteros: puntero [%s] en línea %d apunta a [%s] con estado [%s]",

    // --- LISTAS Y DICCIONARIOS (2400–2499) ---
    2400: "Listas/Diccionarios: detectada estructura [%s]",
    2401: "Listas/Diccionarios: estructura [%s] en línea %d",
    2402: "Listas/Diccionarios: estructura [%s] en línea %d con elementos [%s]",
    2403: "Listas/Diccionarios: estructura [%s] en línea %d con elementos [%s] y contexto [%s]",

    // --- COMANDOS (2500–2599) ---
    2500: "Comandos: detectado comando [%s]",
    2501: "Comandos: comando [%s] en línea %d",
    2502: "Comandos: comando [%s] en línea %d con argumentos [%s]",
    2503: "Comandos: comando [%s] en línea %d con argumentos [%s] y estado [%s]",

    // --- BLOQUES Y CONDICIONALES (2600–2699) ---
    2600: "Bloques: detectado bloque [%s]",
    2601: "Bloques: bloque [%s] en línea %d",
    2602: "Bloques: bloque [%s] en línea %d con condición [%s]",
    2603: "Bloques: bloque [%s] en línea %d con condición [%s] y contexto [%s]",

    // --- EVALUADOR (5000–5099) ---
    5000: "Evaluador: iniciando ejecución de árbol de instrucciones",
    5001: "Evaluador: ejecutando nodo [%s] en línea %d",
    5002: "Evaluador: nodo [%s] en línea %d con valor [%s]",
    5003: "Evaluador: nodo [%s] en línea %d con valor [%s] y contexto [%s]",

    // --- ENTORNO (6000–6099) ---
    6000: "Entorno: inicializando entorno [%s]",
    6001: "Entorno: activando entorno [%s] en línea %d",
    6002: "Entorno: entorno [%s] en línea %d con variables [%s]",
    6003: "Entorno: entorno [%s] en línea %d con variables [%s] y estado [%s]",

    // --- DETALLE GENÉRICO (6999) ---
    6999: "%s[%d]: #%d Detalle desconocido",
    
    // --- RESULTADOS FINALES (7000–7099) --- 
    7000: "Resultados finales del programa:", 7001: "Variable %s = %v",
}
