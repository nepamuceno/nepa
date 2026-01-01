package lexer

import (
    "fmt"
    "os"
    "unicode"
)

type TipoToken int

const (
    T_EOF TipoToken = iota
    T_IDENT
    T_NUMERO
    T_CADENA
    T_OP
    T_IGUAL
    T_PARENT_A
    T_PARENT_C
    T_COMA
    T_PUNTO
    T_DOSPUNTOS
    T_KEYWORD
    T_LINEA
    T_INDENT
    T_DEDENT
)

type Token struct {
    Tipo    TipoToken
    Literal string
    Linea   int
}

type Lexer struct {
    entrada    string
    pos        int
    linea      int
    indents    []int
    pendientes []Token
}

func Nuevo(entrada string) *Lexer {
    l := &Lexer{
        entrada: entrada,
        linea:   1,
        indents: []int{0},
    }
    // IMPORTANTE: Procesar la indentación de la primerísima línea del archivo
    l.procesarIndentacion()
    return l
}

func (l *Lexer) Tokenizar() []Token {
    var tokens []Token
    for {
        t := l.SiguienteToken()
        tokens = append(tokens, t)
        if t.Tipo == T_EOF {
            break
        }
    }
    return tokens
}

func (l *Lexer) SiguienteToken() Token {
    if len(l.pendientes) > 0 {
        t := l.pendientes[0]
        l.pendientes = l.pendientes[1:]
        return t
    }

    l.saltarEspacios()

    if l.pos >= len(l.entrada) {
        // Si todavía hay niveles de indentación (más allá del nivel 0 inicial)
        if len(l.indents) > 1 {
            l.indents = l.indents[:len(l.indents)-1]
            return Token{T_DEDENT, "", l.linea}
        }
        return Token{T_EOF, "", l.linea}
    }

    char := l.entrada[l.pos]

    // Soporte para comentarios (ignora hasta el final de la línea)
    if char == '#' {
        for l.pos < len(l.entrada) && l.entrada[l.pos] != '\n' {
            l.pos++
        }
        return l.SiguienteToken()
    }

    // Manejo de nuevas líneas e indentación
    if char == '\n' {
        l.pos++
        l.linea++
        
        lineaActual := l.linea - 1
        
        // Antes de devolver la línea, procesamos la indentación de la SIGUIENTE línea
        l.procesarIndentacion()
        
        // CORRECCIÓN: Si procesarIndentacion generó tokens (INDENT/DEDENT), 
        // debemos entregarlos ANTES del T_LINEA o asegurar que el orden no rompa el Parser.
        // En Nepa, el T_LINEA ayuda a separar instrucciones, pero si hay pendientes, 
        // priorizamos vaciar la cola de indentación.
        if len(l.pendientes) > 0 {
            // Guardamos el T_LINEA al final de los pendientes para que se procese tras los DEDENTS/INDENTS
            l.pendientes = append(l.pendientes, Token{T_LINEA, "\n", lineaActual})
            t := l.pendientes[0]
            l.pendientes = l.pendientes[1:]
            return t
        }
        
        return Token{T_LINEA, "\n", lineaActual}
    }

    // Identificadores y Palabras Clave
    if unicode.IsLetter(rune(char)) || char == '_' {
        inicio := l.pos
        for l.pos < len(l.entrada) && (unicode.IsLetter(rune(l.entrada[l.pos])) || unicode.IsDigit(rune(l.entrada[l.pos])) || l.entrada[l.pos] == '_') {
            l.pos++
        }
        lit := l.entrada[inicio:l.pos]
        if l.esKeyword(lit) {
            return Token{T_KEYWORD, lit, l.linea}
        }
        return Token{T_IDENT, lit, l.linea}
    }

    // Números
    if unicode.IsDigit(rune(char)) {
        inicio := l.pos
        for l.pos < len(l.entrada) && (unicode.IsDigit(rune(l.entrada[l.pos])) || l.entrada[l.pos] == '.') {
            l.pos++
        }
        return Token{T_NUMERO, l.entrada[inicio:l.pos], l.linea}
    }

    // Cadenas
    if char == '"' {
        l.pos++
        inicio := l.pos
        for l.pos < len(l.entrada) && l.entrada[l.pos] != '"' {
            l.pos++
        }
        lit := l.entrada[inicio:l.pos]
        l.pos++
        return Token{T_CADENA, lit, l.linea}
    }

    // Operadores y Símbolos
    l.pos++
    switch char {
    // --- BLOQUE DE SEGURIDAD: PROHIBICIÓN DE LLAVES Y CORCHETES ---
    case '{', '}', '[', ']':
        fmt.Printf("❌ Error de Sintaxis Nepa (Línea %d): Carácter prohibido '%c'. Nepa utiliza indentación profunda para definir bloques.\n", l.linea, char)
        os.Exit(1)

    case '=':
        if l.pos < len(l.entrada) && l.entrada[l.pos] == '=' {
            l.pos++
            return Token{T_OP, "==", l.linea}
        }
        return Token{T_IGUAL, "=", l.linea}
    case '+': 
        return Token{T_OP, "+", l.linea}
    case '-': 
        return Token{T_OP, "-", l.linea}
    case '*': 
        return Token{T_OP, "*", l.linea}
    case '/': 
        return Token{T_OP, "/", l.linea}
    case '>':
        if l.pos < len(l.entrada) && l.entrada[l.pos] == '=' {
            l.pos++
            return Token{T_OP, ">=", l.linea}
        }
        return Token{T_OP, ">", l.linea}
    case '<':
        if l.pos < len(l.entrada) && l.entrada[l.pos] == '=' {
            l.pos++
            return Token{T_OP, "<=", l.linea}
        }
        return Token{T_OP, "<", l.linea}
    case '!':
        if l.pos < len(l.entrada) && l.entrada[l.pos] == '=' {
            l.pos++
            return Token{T_OP, "!=", l.linea}
        }
        return Token{T_OP, "!", l.linea}
    case '(': 
        return Token{T_PARENT_A, "(", l.linea}
    case ')': 
        return Token{T_PARENT_C, ")", l.linea}
    case ',': 
        return Token{T_COMA, ",", l.linea}
    case '.': 
        return Token{T_PUNTO, ".", l.linea}
    case ':': 
        return Token{T_DOSPUNTOS, ":", l.linea}
    }

    return l.SiguienteToken()
}

func (l *Lexer) procesarIndentacion() {
    espacios := 0
    // Contamos espacios/tabs de la nueva línea
    for l.pos < len(l.entrada) && (l.entrada[l.pos] == ' ' || l.entrada[l.pos] == '\t') {
        if l.entrada[l.pos] == '\t' {
            espacios += 4
        } else {
            espacios++
        }
        l.pos++
    }

    // Si la línea está vacía o es un comentario, ignorar procesamiento de indentación
    if l.pos < len(l.entrada) && (l.entrada[l.pos] == '\n' || l.entrada[l.pos] == '#') {
        return
    }

    ultimo := l.indents[len(l.indents)-1]
    if espacios > ultimo {
        // Validación de Nepa: Solo permitimos incrementos de exactamente 4 espacios
        if espacios-ultimo != 4 {
            fmt.Printf("❌ Error de Sintaxis Nepa (Línea %d): Indentación ilegal. Se esperan exactamente 4 espacios adicionales.\n", l.linea)
            os.Exit(1)
        }
        l.indents = append(l.indents, espacios)
        l.pendientes = append(l.pendientes, Token{T_INDENT, "", l.linea})
    } else if espacios < ultimo {
        // Generar múltiples DEDENTS si bajamos varios niveles
        for len(l.indents) > 1 && espacios < l.indents[len(l.indents)-1] {
            l.indents = l.indents[:len(l.indents)-1]
            l.pendientes = append(l.pendientes, Token{T_DEDENT, "", l.linea})
        }
        // Validación estricta de Nepa: Si los espacios no coinciden con un nivel previo
        if espacios != l.indents[len(l.indents)-1] {
            fmt.Printf("❌ Error de Sintaxis Nepa (Línea %d): Indentación inconsistente. Se esperaban %d espacios.\n", l.linea, l.indents[len(l.indents)-1])
            os.Exit(1)
        }
    }
}

func (l *Lexer) saltarEspacios() {
    // Solo saltamos espacios si no estamos al inicio de una línea 
    for l.pos < len(l.entrada) && l.entrada[l.pos] == ' ' {
        l.pos++
    }
}

func (l *Lexer) esKeyword(lit string) bool {
    ks := []string{"func", "var", "variable", "si", "entonces", "sino", "mientras", "funcion", "retorna", "imprime", "leer", "incluir", "y", "o", "no", "para", "desde", "hasta", "en", "incremento"}
    for _, k := range ks {
        if k == lit { 
            return true 
        }
    }
    return false
}
