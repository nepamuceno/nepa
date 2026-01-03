package bloque

// BloquesReservados contiene todas las sentencias que abren bloques en Nepa.
// Se usan en sintaxis.go para validar que terminen con ":" y tengan indentación correcta.
var BloquesReservados = []string{
    // --- Control de flujo ---
    "si",        // bloque condicional principal
    "sino",      // bloque alternativo, dependiente de "si"

    // --- Bucles ---
    "mientras",  // bucle mientras condición sea verdadera
    "para",      // bucle universal (rangos, listas, strings, mapas, iteradores)

    // --- Definiciones ---
    "funcion",   // definición de funciones
    "modulo",    // definición de módulos / librerías

    // --- Manejo de errores / excepciones ---
    "intentar",   // bloque de manejo de excepciones (equivalente a try)
    "capturar",   // bloque para capturar errores (equivalente a catch)
    "finalmente", // bloque de limpieza tras intentar/capturar (equivalente a finally)

    // --- Otros bloques previstos ---
    "clase",      // definición de clases (futuro)
    "estructura", // definición de estructuras (futuro)
    "ayuda",      // bloque/llamada especial para sistema de ayuda
}
