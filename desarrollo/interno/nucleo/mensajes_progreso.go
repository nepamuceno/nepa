package nucleo

// CATÁLOGO CENTRALIZADO DE MENSAJES DE PROGRESO EN NEPA
// Se usa para reflejar avances parciales durante la ejecución.
// No son errores ni avisos, solo indicadores de avance.

var MENSAJES_PROGRESO = map[int]string{
    // --- INICIO DEL SISTEMA (100–199) ---
    100: "Progreso: cargando intérprete Nepa (10%%)",
    101: "Progreso: configuración aplicada (25%%)",
    102: "Progreso: entorno inicial listo (50%%)",
    103: "Progreso: módulos cargados (75%%)",
    104: "Progreso: sistema listo (100%%)",

    // --- ARCHIVOS (200–299) ---
    200: "Progreso: archivo [%s] abierto (25%%)",
    201: "Progreso: archivo [%s] procesado (50%%)",
    202: "Progreso: archivo [%s] guardado (75%%)",
    203: "Progreso: archivo [%s] cerrado (100%%)",

    // --- EJECUCIÓN (300–399) ---
    300: "Progreso: programa [%s] iniciado (10%%)",
    301: "Progreso: programa [%s] ejecutando (50%%)",
    302: "Progreso: programa [%s] finalizando (90%%)",
    303: "Progreso: programa [%s] completado (100%%)",

    // --- VARIABLES Y CONSTANTES (400–499) ---
    400: "Progreso: variable [%s] definida (25%%)",
    401: "Progreso: constante [%s] registrada (25%%)",
    402: "Progreso: asignación completada (50%%)",

    // --- COMANDOS (500–599) ---
    500: "Progreso: comando [%s] iniciado (10%%)",
    501: "Progreso: comando [%s] ejecutando (50%%)",
    502: "Progreso: comando [%s] completado (100%%)",

    // --- ENTORNO (600–699) ---
    600: "Progreso: entorno [%s] inicializando (25%%)",
    601: "Progreso: entorno [%s] activando (50%%)",
    602: "Progreso: entorno [%s] listo (100%%)",

    // --- PROGRESO GENÉRICO (999) ---
    999: "Progreso desconocido",
}
