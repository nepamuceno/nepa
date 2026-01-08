package nucleo

// CATÁLOGO CENTRALIZADO DE MENSAJES DEL SISTEMA EN NEPA
// Se usa para reflejar estados generales del sistema:
// inicialización, apagado, reinicio, carga de módulos, etc.

var MENSAJES_SISTEMA = map[int]string{
    // --- INICIALIZACIÓN (100–199) ---
    100: "Sistema: inicializando intérprete Nepa",
    101: "Sistema: cargando configuración desde [%s]",
    102: "Sistema: versión [%s] lista",
    103: "Sistema: módulos inicializados",

    // --- APAGADO (200–299) ---
    200: "Sistema: apagando intérprete Nepa",
    201: "Sistema: cerrando entornos activos",
    202: "Sistema: liberando memoria",
    203: "Sistema: apagado completo",

    // --- REINICIO (300–399) ---
    300: "Sistema: reiniciando intérprete Nepa",
    301: "Sistema: reinicio de configuración",
    302: "Sistema: reinicio de módulos",
    303: "Sistema: reinicio completado",

    // --- MÓDULOS (400–499) ---
    400: "Sistema: cargando módulo [%s]",
    401: "Sistema: módulo [%s] listo",
    402: "Sistema: módulo [%s] desactivado",
    403: "Sistema: módulo [%s] reiniciado",

    // --- ESTADO GENERAL (500–599) ---
    500: "Sistema: en espera",
    501: "Sistema: en ejecución",
    502: "Sistema: en pausa",
    503: "Sistema: detenido",

    // --- SISTEMA GENÉRICO (999) ---
    999: "Sistema desconocido",
}
