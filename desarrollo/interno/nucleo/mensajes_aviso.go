package nucleo

// CATÁLOGO CENTRALIZADO DE MENSAJES DE AVISO EN NEPA
// Se usa para advertencias que no son errores fatales,
// pero requieren atención del usuario.

var MENSAJES_AVISO = map[int]string{
    // --- ARCHIVOS (100–199) ---
    100: "Aviso: el archivo [%s] ya existe, se sobrescribirá",
    101: "Aviso: el archivo [%s] está vacío",
    102: "Aviso: el archivo [%s] no tiene extensión reconocida",

    // --- SÍMBOLOS (200–299) ---
    200: "Aviso: símbolo [%s] redefinido",
    201: "Aviso: símbolo [%s] sin uso posterior",
    202: "Aviso: símbolo [%s] podría ser ambiguo",

    // --- SINTAXIS (300–399) ---
    300: "Aviso: instrucción [%s] sin efecto",
    301: "Aviso: bloque [%s] sin contenido útil",
    302: "Aviso: condición [%s] siempre verdadera o falsa",

    // --- ESTRUCTURAS (400–499) ---
    400: "Aviso: variable [%s] declarada pero no utilizada",
    401: "Aviso: variable [%s] podría sobrescribir otra existente",
    402: "Aviso: tipo de variable [%s] no recomendado",

    // --- CONSTANTES (500–599) ---
    500: "Aviso: constante [%s] declarada pero no utilizada",
    501: "Aviso: constante [%s] podría duplicar otra existente",

    // --- PUNTEROS (600–699) ---
    600: "Aviso: puntero [%s] sin inicializar",
    601: "Aviso: puntero [%s] apunta a dirección incierta",
    602: "Aviso: puntero [%s] podría causar fuga de memoria",

    // --- LISTAS Y DICCIONARIOS (700–799) ---
    700: "Aviso: lista [%s] vacía",
    701: "Aviso: diccionario [%s] sin claves definidas",
    702: "Aviso: estructura [%s] con elementos duplicados",

    // --- COMANDOS (800–899) ---
    800: "Aviso: comando [%s] sin argumentos",
    801: "Aviso: comando [%s] con argumentos redundantes",

    // --- BLOQUES Y CONDICIONALES (900–999) ---
    900: "Aviso: bloque [%s] nunca ejecutado",
    901: "Aviso: condicional [%s] sin efecto práctico",

    // --- EVALUADOR (1000–1099) ---
    1000: "Aviso: operación [%s] con resultado no utilizado",
    1001: "Aviso: función [%s] retorna valor ignorado",

    // --- ENTORNO (1100–1199) ---
    1100: "Aviso: entorno [%s] sin variables activas",
    1101: "Aviso: entorno [%s] duplicado",

    // --- AVISO GENÉRICO (1999) ---
    1999: "Aviso desconocido",
}
