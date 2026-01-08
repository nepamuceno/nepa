package nucleo

// CATÁLOGO CENTRALIZADO DE MENSAJES DE ESTADO EN NEPA
// Se usa para reflejar el estado actual del sistema o de la ejecución.
// No son errores ni avisos, solo estados transitorios o permanentes.

var MENSAJES_ESTADO = map[int]string{
    // --- INICIO DEL SISTEMA (100–199) ---
    100: "Estado: cargando intérprete Nepa",
    101: "Estado: configuración en proceso de carga",
    102: "Estado: versión [%s] lista",

    // --- ARCHIVOS (200–299) ---
    200: "Estado: archivo [%s] abierto",
    201: "Estado: archivo [%s] en proceso de guardado",
    202: "Estado: archivo [%s] cerrado",

    // --- EJECUCIÓN (300–399) ---
    300: "Estado: programa [%s] en ejecución",
    301: "Estado: programa [%s] detenido",
    302: "Estado: programa [%s] en pausa",
    303: "Estado: programa [%s] esperando entrada",

    // --- VARIABLES Y CONSTANTES (400–499) ---
    400: "Estado: variable [%s] activa",
    401: "Estado: constante [%s] registrada",
    402: "Estado: variable [%s] en espera de asignación",

    // --- COMANDOS (500–599) ---
    500: "Estado: comando [%s] en ejecución",
    501: "Estado: comando [%s] completado",
    502: "Estado: comando [%s] en espera",

    // --- ENTORNO (600–699) ---
    600: "Estado: entorno [%s] inicializado",
    601: "Estado: entorno [%s] activo",
    602: "Estado: entorno [%s] cerrado",
    603: "Estado: entorno [%s] en espera",

    // --- ESTADO GENÉRICO (999) ---
    999: "Estado desconocido",
}
