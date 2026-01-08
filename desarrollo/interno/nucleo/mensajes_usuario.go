package nucleo

// CATÁLOGO CENTRALIZADO DE MENSAJES PARA EL USUARIO EN NEPA
// Se usa para comunicación directa con la persona:
// bienvenida, despedida, confirmaciones y avisos personalizados.

var MENSAJES_USUARIO = map[int]string{
    // --- BIENVENIDA (100–199) ---
    100: "Bienvenido a Nepa",
    101: "Hola [%s], tu sesión está lista",
    102: "Inicio exitoso, disfruta tu experiencia con Nepa",

    // --- DESPEDIDA (200–299) ---
    200: "Gracias por usar Nepa",
    201: "Tu sesión ha terminado, hasta pronto",
    202: "Cierre exitoso, vuelve cuando quieras",

    // --- CONFIRMACIONES (300–399) ---
    300: "Acción [%s] realizada correctamente",
    301: "Tu solicitud [%s] fue completada",
    302: "Confirmación: [%s] aplicado con éxito",

    // --- INTERACCIÓN (400–499) ---
    400: "¿Deseas continuar con [%s]?",
    401: "Tu entrada [%s] fue recibida",
    402: "Procesando tu solicitud [%s]",
    403: "Resultado disponible para [%s]",

    // --- USUARIO GENÉRICO (999) ---
    999: "Mensaje genérico para el usuario",
}
