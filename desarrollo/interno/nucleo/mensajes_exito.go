package nucleo

// CATÁLOGO CENTRALIZADO DE MENSAJES DE ÉXITO EN NEPA
// Se usa para confirmar operaciones realizadas correctamente.
// No son errores ni avisos, solo confirmaciones positivas.

var MENSAJES_EXITO = map[int]string{
    // --- INICIO DEL SISTEMA (100–199) ---
    100: "Éxito: intérprete Nepa cargado correctamente",
    101: "Éxito: configuración aplicada desde [%s]",
    102: "Éxito: versión [%s] inicializada",

    // --- ARCHIVOS (200–299) ---
    200: "Éxito: archivo [%s] abierto",
    201: "Éxito: archivo [%s] guardado",
    202: "Éxito: archivo [%s] cerrado",

    // --- EJECUCIÓN (300–399) ---
    300: "Éxito: programa [%s] iniciado",
    301: "Éxito: programa [%s] finalizado",
    302: "Éxito: resultados generados",

    // --- VARIABLES Y CONSTANTES (400–499) ---
    400: "Éxito: variable [%s] definida",
    401: "Éxito: constante [%s] registrada",

    // --- COMANDOS (500–599) ---
    500: "Éxito: comando [%s] ejecutado",
    501: "Éxito: comando [%s] completado",

    // --- ENTORNO (600–699) ---
    600: "Éxito: entorno [%s] inicializado",
    601: "Éxito: entorno [%s] activado",
    602: "Éxito: entorno [%s] cerrado",

    // --- ÉXITO GENÉRICO (999) ---
    999: "Éxito: operación completada",
}
