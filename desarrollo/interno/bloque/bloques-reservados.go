package bloque

var BloquesReservados = []string{
    // --- Control de flujo ---
    "si_es", "si_no", "opcion_en", "entonces",

    // --- Bucles ---
    "mientras", "para", "por_cada", "hasta",
    "desde", "incremento",

    // --- Definiciones ---
    "funcion", "modulo", "clase", "estructura", "interfaz",

    // --- Manejo de errores ---
    "intentar", "capturar", "finalmente",

    // --- Concurrencia / asincronía ---
    "async", "esperar", "concurrente",

    // --- Contextos especiales ---
    "usar", "transaccion",
}
