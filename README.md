# 🚀 Nepa: Lenguaje de Programación en Español (Científico)

**Nepa** es un lenguaje de programación de sintaxis nativa en español, desarrollado íntegramente en **Go**. Su objetivo es eliminar la barrera del idioma en el desarrollo lógico y científico, permitiendo la ejecución de cálculos complejos, fórmulas físicas y algoritmos mediante comandos intuitivos en nuestro propio idioma.

---

## 🔬 ¿Qué es Nepa?
Nepa es un intérprete diseñado para el ámbito científico y académico. Proporciona un entorno donde las expresiones matemáticas y las constantes universales se manejan de forma natural, ofreciendo una estructura clara para el análisis de datos y simulaciones físicas.

---

## 📖 Sintaxis del Lenguaje
El lenguaje utiliza palabras clave en español que facilitan la comprensión de la lógica del programa:

* **Estructuras de Control:** `si`, `sino`, `mientras`, `para`.
* **Definiciones:** `funcion`, `retornar`.
* **Salida de Datos:** `imprimir`.
* **Constantes Integradas:** `PI`, `E`, `GRAVEDAD`, `LUZ`, `PHI`.
* **Funciones Científicas:** `seno`, `coseno`, `raiz`, `es_primo`, `vol_cono`, `proyectil_pos`.


## 📖 Compilación y Uso

El flujo de trabajo en Nepa está optimizado mediante scripts de automatización.
Ejecución con probar.sh

Para compilar y probar el sistema rápidamente, ejecuta:
Bash

###
bash probar.sh

## 📖 ¿Qué hace este script?

    Limpia: Ejecuta make clean para eliminar binarios obsoletos.

    Compila: Construye el motor mediante el Makefile.

    Genera SDK: Invoca a nepa_lib para regenerar las librerías matemáticas dinámicas.

    Ejecuta: Lanza el intérprete con el archivo de prueba inicio.nepa.

## 📖  Configuración de Seguridad (.env)

Este proyecto utiliza un archivo de configuración local para manejar credenciales de forma segura:

    Archivo .env: Debe existir en la raíz (ignorado por Git).

    Contenido: Debe incluir la variable NEPA_TOKEN="tu_token_de_github".

    Despliegue: El script subir_nepa.sh lee este token para realizar subidas seguras y etiquetado (tags) automático de versiones sin exponer claves en el historial público.

## 📖 Estructura del Proyecto

    desarrollo/: Núcleo del lenguaje (Lexer, Parser, Evaluador y AST).

    dist/bin/: Binarios finales (nepa, nepa_lib).

    history/: Registro de comandos utilizados en el desarrollo.

    gist.txt: Respaldo maestro del código fuente.

    probar.sh: Automatización de pruebas y compilación.

    subir_nepa.sh: Sincronización segura con el repositorio.

### Ejemplo de Código
```nepa
# Cálculo de área y condicional en Nepa
radio = 5
area = PI * radio * radio

si (area > 50) {
    imprimir("El área es grande:", area)
} sino {
    imprimir("El área es pequeña:", area)
}

x = 0
mientras (x < 3) {
    imprimir("Contador:", x)
    x = x + 1
}

## 📖 Hecho con ❤️ para la comunidad de programación científica.

