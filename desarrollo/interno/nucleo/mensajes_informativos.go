package nucleo

// CATÁLOGO CENTRALIZADO DE MENSAJES INFORMATIVOS EN NEPA
// Se usa para avisos generales, estados del sistema y confirmaciones.
// No son errores ni depuración, solo información para el usuario.

var MENSAJES_INFO = map[int]string{
    // --- INICIO DEL SISTEMA (100–199) ---
    100: "Inicio: cargando intérprete Nepa",
    101: "Inicio: configuración cargada desde [%s]",
    102: "Inicio: versión actual [%s]",

    // --- ARCHIVOS (200–299) ---
    200: "Archivo [%s] abierto correctamente",
    201: "Archivo [%s] guardado correctamente",
    202: "Archivo [%s] cerrado",

    // --- EJECUCIÓN (300–399) ---
    300: "Ejecución: programa [%s] iniciado",
    301: "Ejecución: programa [%s] finalizado",
    302: "Ejecución: resultados disponibles",

    // --- VARIABLES Y CONSTANTES (400–499) ---
    400: "Variable [%s] definida correctamente",
    401: "Constante [%s] registrada correctamente",

    // --- COMANDOS (500–599) ---
    500: "Comando [%s] ejecutado correctamente",
    501: "Comando [%s] completado con éxito",

    // --- ENTORNO (600–699) ---
    600: "Entorno [%s] inicializado",
    601: "Entorno [%s] activado",
    602: "Entorno [%s] cerrado",

    // --- INFORMACIÓN GENÉRICA (900–999) ---
    900: "Operación completada",
    901: "Acción realizada correctamente",
    999: "Información desconocida",
}
