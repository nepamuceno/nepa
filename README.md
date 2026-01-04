# **nepa**

**nepa** es un lenguaje de programación **nuevo y experimental**, acompañado de su propio intérprete escrito en **Go**.  
Su diseño se inspira en la claridad estructurada de lenguajes modernos como Python, pero con una identidad única: utiliza **acrónimos y palabras reservadas en español** en lugar de inglés, creando una experiencia más cercana y accesible.

---

## ✨ Características principales
- **Lenguaje original**: no es un clon, sino una propuesta propia y experimental.
- **Sintaxis estructurada**: bloques y reglas claras, sin numeración de líneas.
- **Arquitectura modular**: cada tipo de variable, bloque y función vive en su propio paquete.
- **Soporte matemático amplio**: álgebra, estadísticas, finanzas y física.
- **Palabras reservadas en español**: pensado para ser más intuitivo y culturalmente relevante.
- **Extensible**: se dejan hooks para futuros módulos externos y nuevas funcionalidades.

---

## 📂 Estructura del proyecto
- `variables/` → Tipos básicos y complejos (entero, real, texto, lista, matriz, objeto, etc.)
- `matematicas/` → Funciones de álgebra, estadísticas, finanzas y física.
- `bloque/` → Palabras y bloques reservados del lenguaje.
- `sintaxis/` → Reglas de interpretación y validación.
- `core/` → Utilidades y funciones centrales del sistema.

---

## 📂 Instalacion
Compilar e instalar **nepa** es sencillo. Existen dos formas principales: 

### 🔨 Usando `make`
```
make
make install
```
Usando compila.sh

```
bash compila.sh
```
El binario estara en dist/bin/nepa

## uso: 
```
./dist/bin/nepa <programa.nepa>
```
### Ejemplo: ./dist/bin/nepa test.nepa
---

## 🚀 Objetivo
El objetivo de **nepa** es servir como base para un **lenguaje modular en español**, fácil de extender y mantener, que permita experimentar con nuevas ideas de sintaxis y ejecución.  
Es un proyecto en evolución, pensado para crecer paso a paso y dejar siempre espacio para futuras expansiones.

---

## 🌱 Filosofía
**nepa** busca ser un **laboratorio de ideas**:  
- Simple en el presente, para que cualquiera pueda probarlo.  
- Modular y estructurado, para que sea mantenible.  
- Con identidad propia en español, para demostrar que los lenguajes de programación también pueden hablar nuestro idioma.  

---

💡 * **nepa** es más que un intérprete: es una propuesta experimental para imaginar cómo podrían ser los lenguajes del futuro, diseñados desde nuestra lengua y cultura.*
