package lexer

import (
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
    entrada     string
    pos         int
    linea       int
    indents     []int
    pendientes  []Token
}

func Nuevo(entrada string) *Lexer {
    return &Lexer{
        entrada: entrada,
        linea:   1,
        indents: []int{0},
    }
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
        // Al final, cerramos todos los niveles de indentación abiertos
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
        l.procesarIndentacion()
        return Token{T_LINEA, "\n", l.linea - 1}
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
    case '=':
        if l.pos < len(l.entrada) && l.entrada[l.pos] == '=' {
            l.pos++
            return Token{T_OP, "==", l.linea}
        }
        return Token{T_IGUAL, "=", l.linea}
    case '+': return Token{T_OP, "+", l.linea}
    case '-': return Token{T_OP, "-", l.linea}
    case '*': return Token{T_OP, "*", l.linea}
    case '/': return Token{T_OP, "/", l.linea}
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
        // Si es solo !, lo tratamos como operador aunque falte el =
        return Token{T_OP, "!", l.linea}
    case '(': return Token{T_PARENT_A, "(", l.linea}
    case ')': return Token{T_PARENT_C, ")", l.linea}
    case ',': return Token{T_COMA, ",", l.linea}
    case '.': return Token{T_PUNTO, ".", l.linea}
    case ':': return Token{T_DOSPUNTOS, ":", l.linea}
    }

    return l.SiguienteToken()
}

func (l *Lexer) procesarIndentacion() {
    espacios := 0
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
        l.indents = append(l.indents, espacios)
        l.pendientes = append(l.pendientes, Token{T_INDENT, "", l.linea})
    } else {
        for espacios < ultimo {
            l.indents = l.indents[:len(l.indents)-1]
            ultimo = l.indents[len(l.indents)-1]
            l.pendientes = append(l.pendientes, Token{T_DEDENT, "", l.linea})
        }
    }
}

func (l *Lexer) saltarEspacios() {
    // Solo saltamos espacios si no estamos al inicio de una línea procesando indentación
    for l.pos < len(l.entrada) && l.entrada[l.pos] == ' ' {
        l.pos++
    }
}

func (l *Lexer) esKeyword(lit string) bool {
    ks := []string{"variable", "si", "entonces", "sino", "mientras", "funcion", "retorna", "imprime", "leer", "incluir", "y", "o", "no", "para", "desde", "hasta", "en", "incremento"}
    for _, k := range ks {
        if k == lit {
            return true
        }
    }
    return false
}
