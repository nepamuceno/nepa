package nucleo

// CATÁLOGO CENTRALIZADO DE MENSAJES DE DEPURACIÓN EN NEPA
// Niveles progresivos:
//   Nivel 1 → Básico (estado general)
//   Nivel 2 → Ubicación exacta
//   Nivel 3 → Ubicación + valores recibidos
//   Nivel 4 → Todos los detalles internos (estado, pila, contexto)

var MENSAJES_DEPURACION = map[int]string{
    // --- SÍMBOLOS (1000–1099) ---
    1000: "Símbolos: inicio de revisión interna",
    1001: "Símbolos: símbolo [%s] en línea %d",
    1002: "Símbolos: símbolo [%s] en línea %d con valor [%s]",
    1003: "Símbolos: símbolo [%s] en línea %d con valor [%s] y estado interno [%s]",

    // --- SINTAXIS (2000–2099) ---
    2000: "Sintaxis: inicio de bloque [%s]",
    2001: "Sintaxis: instrucción [%s] en línea %d",
    2002: "Sintaxis: instrucción [%s] en línea %d con argumentos [%s]",
    2003: "Sintaxis: instrucción [%s] en línea %d con argumentos [%s] y contexto interno [%s]",

    // --- ESTRUCTURAS (2100–2199) ---
    2100: "Estructuras: declaración [%s]",
    2101: "Estructuras: asignación [%s] en línea %d",
    2102: "Estructuras: asignación [%s] en línea %d con valor [%s]",
    2103: "Estructuras: asignación [%s] en línea %d con valor [%s] y tipo [%s]",

    // --- CONSTANTES (2200–2299) ---
    2200: "Constantes: constante [%s]",
    2201: "Constantes: constante [%s] en línea %d",
    2202: "Constantes: constante [%s] en línea %d con valor [%s]",
    2203: "Constantes: constante [%s] en línea %d con valor [%s] y tipo [%s]",

    // --- PUNTEROS (2300–2399) ---
    2300: "Punteros: puntero [%s]",
    2301: "Punteros: puntero [%s] en línea %d",
    2302: "Punteros: puntero [%s] en línea %d apunta a [%s]",
    2303: "Punteros: puntero [%s] en línea %d apunta a [%s] con estado interno [%s]",

    // --- LISTAS Y DICCIONARIOS (2400–2499) ---
    2400: "Listas/Diccionarios: estructura [%s]",
    2401: "Listas/Diccionarios: estructura [%s] en línea %d",
    2402: "Listas/Diccionarios: estructura [%s] en línea %d con elementos [%s]",
    2403: "Listas/Diccionarios: estructura [%s] en línea %d con elementos [%s] y contexto interno [%s]",

    // --- COMANDOS (2500–2599) ---
    2500: "Comandos: comando [%s]",
    2501: "Comandos: comando [%s] en línea %d",
    2502: "Comandos: comando [%s] en línea %d con argumentos [%s]",
    2503: "Comandos: comando [%s] en línea %d con argumentos [%s] y estado interno [%s]",

    // --- BLOQUES Y CONDICIONALES (2600–2699) ---
    2600: "Bloques: bloque [%s]",
    2601: "Bloques: bloque [%s] en línea %d",
    2602: "Bloques: bloque [%s] en línea %d con condición [%s]",
    2603: "Bloques: bloque [%s] en línea %d con condición [%s] y contexto interno [%s]",

    // --- EVALUADOR (5000–5099) ---
    5000: "Evaluador: inicio de ejecución interna",
    5001: "Evaluador: nodo [%s] en línea %d",
    5002: "Evaluador: nodo [%s] en línea %d con valor [%s]",
    5003: "Evaluador: nodo [%s] en línea %d con valor [%s] y contexto interno [%s]",

    // --- ENTORNO (6000–6099) ---
    6000: "Entorno: inicializando entorno [%s]",
    6001: "Entorno: activando entorno [%s] en línea %d",
    6002: "Entorno: entorno [%s] en línea %d con variables [%s]",
    6003: "Entorno: entorno [%s] en línea %d con variables [%s] y estado interno [%s]",
    6099: "%s[%d]: #%d Depuración desconocida", // genérico para depuración
    
    // --- RESULTADOS FINALES (6100–6199) ---
	6100: "Resultados: inicio de impresión de resultados finales",
	6101: "Resultados: variable [%s] = [%s]",
}
