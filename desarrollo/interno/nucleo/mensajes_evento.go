package nucleo

// CATÁLOGO CENTRALIZADO DE MENSAJES DE EVENTOS EN NEPA
// Se usa para reflejar sucesos importantes dentro del sistema:
// inicio de sesión, cambios de configuración, ejecución de tareas, etc.

var MENSAJES_EVENTO = map[int]string{
    // --- SESIONES (100–199) ---
    100: "Evento: inicio de sesión de usuario [%s]",
    101: "Evento: cierre de sesión de usuario [%s]",
    102: "Evento: sesión expirada para usuario [%s]",

    // --- CONFIGURACIÓN (200–299) ---
    200: "Evento: configuración modificada [%s]",
    201: "Evento: configuración restaurada a valores predeterminados",
    202: "Evento: nueva configuración aplicada desde [%s]",

    // --- ARCHIVOS (300–399) ---
    300: "Evento: archivo [%s] creado",
    301: "Evento: archivo [%s] eliminado",
    302: "Evento: archivo [%s] renombrado a [%s]",

    // --- EJECUCIÓN DE TAREAS (400–499) ---
    400: "Evento: tarea [%s] programada",
    401: "Evento: tarea [%s] iniciada",
    402: "Evento: tarea [%s] completada",
    403: "Evento: tarea [%s] cancelada",

    // --- ENTORNO (500–599) ---
    500: "Evento: entorno [%s] creado",
    501: "Evento: entorno [%s] eliminado",
    502: "Evento: entorno [%s] reiniciado",

    // --- USUARIO (600–699) ---
    600: "Evento: usuario [%s] registrado",
    601: "Evento: usuario [%s] eliminado",
    602: "Evento: usuario [%s] actualizado",

    // --- EVENTO GENÉRICO (999) ---
    999: "Evento desconocido",
}
