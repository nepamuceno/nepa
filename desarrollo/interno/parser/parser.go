package parser

import (
    "fmt"
    "nepa/desarrollo/interno/ast"
    "nepa/desarrollo/interno/lexer"
    "os"
    "strconv"
)

type Parser struct {
    tokens  []lexer.Token
    pos     int
    archivo string
}

func Nuevo(tokens []lexer.Token, archivo string) *Parser {
    return &Parser{tokens, 0, archivo}
}

func (p *Parser) getPos() ast.Posicion {
    if p.pos >= len(p.tokens) {
        return ast.Posicion{Linea: 0, Archivo: p.archivo}
    }
    return ast.Posicion{Linea: p.actual().Linea, Archivo: p.archivo}
}

// miraSiguiente actúa como un peek del token que viene
func (p *Parser) mirarSiguiente() lexer.Token {
    if p.pos+1 >= len(p.tokens) {
        return lexer.Token{Tipo: lexer.T_EOF}
    }
    return p.tokens[p.pos+1]
}

func (p *Parser) Parsear() []ast.Nodo {
    var programa []ast.Nodo
    for !p.esFin() {
        // CORRECCIÓN: Saltamos líneas vacías Y DEDENTs/INDENTs huérfanos entre instrucciones principales
        // Esto evita que el parser se trabe al salir de una función y encontrar una variable global.
        // Se añade soporte para limpiar tokens de indentación en el scope global.
        tipoActual := p.actual().Tipo
        if tipoActual == lexer.T_LINEA || tipoActual == lexer.T_DEDENT || tipoActual == lexer.T_INDENT {
            p.avanzar()
            continue
        }
        
        nodo := p.parsearInstruccion()
        if nodo != nil {
            programa = append(programa, nodo)
        } else {
            p.avanzar()
        }
    }
    return programa
}

func (p *Parser) parsearInstruccion() ast.Nodo {
    t := p.actual()
    
    // CORRECCIÓN: Usamos el Literal para identificar palabras clave
    if t.Tipo == lexer.T_KEYWORD || t.Tipo == lexer.T_IDENT {
        switch t.Literal {
        case "func":
            return p.parsearDefFuncion()
        case "si":
            return p.parsearSi()
        case "para":
            return p.parsearPara()
        case "retorna":
            return p.parsearRetornar()
        case "variable", "var":
            return p.parsearAsignacion()
        }
    }

    // Manejo de identificadores y expresiones
    switch t.Tipo {
    case lexer.T_IDENT: 
        if p.mirarSiguiente().Tipo == lexer.T_IGUAL {
            return p.parsearAsignacionDirecta()
        }
        return p.parsearExpresion()
    default:
        return p.parsearExpresion()
    }
}

func (p *Parser) parsearAsignacionDirecta() ast.Nodo {
    pos := p.getPos()
    nombre := p.consumir(lexer.T_IDENT).Literal
    p.consumir(lexer.T_IGUAL)
    return ast.Asignacion{Base: ast.Base{Pos: pos}, Nombre: nombre, Valor: p.parsearExpresion()}
}

func (p *Parser) parsearRetornar() ast.Nodo {
    pos := p.getPos()
    p.avanzar() // saltar 'retorna'
    return ast.Retornar{Base: ast.Base{Pos: pos}, Valor: p.parsearExpresion()}
}

func (p *Parser) parsearBloque() []ast.Nodo {
    var nodos []ast.Nodo
    
    // Tras un ':', DEBE haber un salto de línea seguido de un INDENT
    // Limpiamos líneas acumuladas antes de validar el INDENT obligatorio
    for p.actual().Tipo == lexer.T_LINEA {
        p.avanzar()
    }
    
    if p.actual().Tipo != lexer.T_INDENT {
        t := p.actual()
        // MODIFICACIÓN: Se utiliza %v para representar el tipo de token correctamente en el error
        fmt.Printf("❌ Error de Sintaxis Nepa (%s:%d): Se esperaba bloque indentado (4 espacios) tras ':'. Encontrado: %v\n", p.archivo, t.Linea, t.Tipo)
        os.Exit(1)
    }

    p.avanzar() // Consumir T_INDENT

    for !p.esFin() && p.actual().Tipo != lexer.T_DEDENT {
        if p.actual().Tipo == lexer.T_LINEA {
            p.avanzar()
            continue
        }

        nodo := p.parsearInstruccion()
        if nodo != nil {
            nodos = append(nodos, nodo)
        } else {
            if !p.esFin() && p.actual().Tipo != lexer.T_DEDENT {
                p.avanzar()
            }
        }
    }

    if len(nodos) == 0 {
        fmt.Printf("❌ Error de Sintaxis Nepa (%s): Bloque indentado vacío ilegal.\n", p.archivo)
        os.Exit(1)
    }

    if p.actual().Tipo == lexer.T_DEDENT {
        p.avanzar()
    } else if !p.esFin() {
        fmt.Printf("❌ Error de Sintaxis Nepa (%s): Fallo de indentación al cerrar bloque.\n", p.archivo)
        os.Exit(1)
    }
    return nodos
}

// --- JERARQUÍA DE EXPRESIONES ---

func (p *Parser) parsearExpresion() ast.Nodo {
    izq := p.parsearComparacion()
    for p.actual().Tipo == lexer.T_KEYWORD && (p.actual().Literal == "y" || p.actual().Literal == "o") {
        pos := p.getPos(); op := p.actual().Literal; p.avanzar()
        der := p.parsearComparacion()
        izq = ast.OperacionBinaria{Base: ast.Base{Pos: pos}, Izquierda: izq, Operador: op, Derecha: der}
    }
    return izq
}

func (p *Parser) parsearComparacion() ast.Nodo {
    izq := p.parsearSumaResta()
    for p.actual().Tipo == lexer.T_OP {
        op := p.actual().Literal
        if op == "==" || op == "!=" || op == "<" || op == ">" || op == "<=" || op == ">=" {
            pos := p.getPos(); p.avanzar()
            der := p.parsearSumaResta()
            izq = ast.OperacionBinaria{Base: ast.Base{Pos: pos}, Izquierda: izq, Operador: op, Derecha: der}
        } else {
            break 
        }
    }
    return izq
}

func (p *Parser) parsearSumaResta() ast.Nodo {
    izq := p.parsearMultDiv()
    for p.actual().Tipo == lexer.T_OP && (p.actual().Literal == "+" || p.actual().Literal == "-") {
        pos := p.getPos(); op := p.actual().Literal; p.avanzar()
        der := p.parsearMultDiv()
        izq = ast.OperacionBinaria{Base: ast.Base{Pos: pos}, Izquierda: izq, Operador: op, Derecha: der}
    }
    return izq
}

func (p *Parser) parsearMultDiv() ast.Nodo {
    izq := p.parsearPrimario()
    for p.actual().Tipo == lexer.T_OP && (p.actual().Literal == "*" || p.actual().Literal == "/") {
        pos := p.getPos(); op := p.actual().Literal; p.avanzar()
        der := p.parsearPrimario()
        izq = ast.OperacionBinaria{Base: ast.Base{Pos: pos}, Izquierda: izq, Operador: op, Derecha: der}
    }
    return izq
}

func (p *Parser) parsearPrimario() ast.Nodo {
    pos := p.getPos()
    t := p.actual()

    switch t.Tipo {
    case lexer.T_NUMERO:
        v, _ := strconv.ParseFloat(t.Literal, 64)
        p.avanzar()
        return ast.Literal{Base: ast.Base{Pos: pos}, Valor: v}
    case lexer.T_CADENA:
        v := t.Literal; p.avanzar()
        return ast.Literal{Base: ast.Base{Pos: pos}, Valor: v}
    case lexer.T_IDENT:
        nombre := t.Literal; p.avanzar()
        if p.actual().Tipo == lexer.T_PUNTO {
            p.avanzar()
            prop := p.consumir(lexer.T_IDENT).Literal
            if p.actual().Tipo == lexer.T_PARENT_A {
                return p.parsearLlamadaModulo(nombre, prop)
            }
            return ast.LlamadaModulo{Base: ast.Base{Pos: pos}, Modulo: nombre, Funcion: prop}
        }
        if p.actual().Tipo == lexer.T_PARENT_A {
            return p.parsearLlamada(nombre)
        }
        return ast.Identificador{Base: ast.Base{Pos: pos}, Nombre: nombre}
    case lexer.T_PARENT_A:
        p.avanzar()
        exp := p.parsearExpresion()
        p.consumir(lexer.T_PARENT_C)
        return exp
    case lexer.T_OP:
        if t.Literal == "{" || t.Literal == "}" {
            fmt.Printf("❌ Error de Sintaxis Nepa (%s:%d): Carácter prohibido '%s'. Use indentación para bloques.\n", p.archivo, t.Linea, t.Literal)
            os.Exit(1)
        }
    }
    return nil
}

func (p *Parser) parsearLlamada(nombre string) ast.Nodo {
    pos := p.getPos(); p.consumir(lexer.T_PARENT_A)
    var args []ast.Nodo
    for p.actual().Tipo != lexer.T_PARENT_C && !p.esFin() {
        args = append(args, p.parsearExpresion())
        if p.actual().Tipo == lexer.T_COMA { p.avanzar() }
    }
    p.consumir(lexer.T_PARENT_C)
    return ast.LlamadaFuncion{Base: ast.Base{Pos: pos}, Nombre: nombre, Args: args}
}

func (p *Parser) parsearLlamadaModulo(mod, fn string) ast.Nodo {
    pos := p.getPos(); p.consumir(lexer.T_PARENT_A)
    var args []ast.Nodo
    for p.actual().Tipo != lexer.T_PARENT_C && !p.esFin() {
        args = append(args, p.parsearExpresion())
        if p.actual().Tipo == lexer.T_COMA { p.avanzar() }
    }
    p.consumir(lexer.T_PARENT_C)
    return ast.LlamadaModulo{Base: ast.Base{Pos: pos}, Modulo: mod, Funcion: fn, Args: args}
}

func (p *Parser) parsearSi() ast.Nodo {
    pos := p.getPos(); p.avanzar()
    cond := p.parsearExpresion()
    p.consumir(lexer.T_DOSPUNTOS)
    cuerpo := p.parsearBloque()
    var sino []ast.Nodo
    
    // Limpiar líneas antes de buscar un posible 'sino'
    for p.actual().Tipo == lexer.T_LINEA { p.avanzar() }

    if p.actual().Literal == "sino" {
        p.avanzar()
        if p.actual().Literal == "si" {
            sino = append(sino, p.parsearSi())
        } else {
            // El sino también requiere ':' y un bloque indentado
            if p.actual().Tipo == lexer.T_DOSPUNTOS { p.avanzar() }
            sino = p.parsearBloque()
        }
    }
    return ast.Si{Base: ast.Base{Pos: pos}, Condicion: cond, Cuerpo: cuerpo, Sino: sino}
}

func (p *Parser) parsearDefFuncion() ast.Nodo {
    pos := p.getPos(); p.avanzar()
    nombre := p.consumir(lexer.T_IDENT).Literal
    p.consumir(lexer.T_PARENT_A)
    params := []string{}
    for p.actual().Tipo != lexer.T_PARENT_C && !p.esFin() {
        params = append(params, p.consumir(lexer.T_IDENT).Literal)
        if p.actual().Tipo == lexer.T_COMA { p.avanzar() }
    }
    p.consumir(lexer.T_PARENT_C)
    p.consumir(lexer.T_DOSPUNTOS)
    return ast.FuncionDef{Base: ast.Base{Pos: pos}, Nombre: nombre, Parametros: params, Cuerpo: p.parsearBloque()}
}

func (p *Parser) parsearAsignacion() ast.Nodo {
    pos := p.getPos(); p.avanzar()
    nombre := p.consumir(lexer.T_IDENT).Literal
    p.consumir(lexer.T_IGUAL)
    return ast.Asignacion{Base: ast.Base{Pos: pos}, Nombre: nombre, Valor: p.parsearExpresion()}
}

func (p *Parser) actual() lexer.Token { 
    if p.pos >= len(p.tokens) { return lexer.Token{Tipo: lexer.T_EOF} }
    return p.tokens[p.pos] 
}

func (p *Parser) avanzar() { p.pos++ }

func (p *Parser) esFin() bool { return p.actual().Tipo == lexer.T_EOF }

func (p *Parser) consumir(t lexer.TipoToken) lexer.Token { 
    // CORRECCIÓN: Manejo selectivo de saltos de línea para evitar saltar el token objetivo.
    // Solo saltamos LINEA o DEDENT si NO es lo que estamos buscando.
    for !p.esFin() && t != lexer.T_LINEA && t != lexer.T_DEDENT && 
        (p.actual().Tipo == lexer.T_LINEA || p.actual().Tipo == lexer.T_DEDENT) {
        p.avanzar()
    }

    tok := p.actual()
    if tok.Tipo != t {
        fmt.Printf("❌ Error de Sintaxis Nepa (%s:%d): Se esperaba %v, pero se encontró %v\n", p.archivo, tok.Linea, t, tok.Tipo)
        os.Exit(1)
    }
    p.avanzar() 
    return tok 
}

func (p *Parser) parsearPara() ast.Nodo {
    pos := p.getPos()
    p.avanzar() 

    tokenVar := p.consumir(lexer.T_IDENT)
    variable := tokenVar.Literal

    var origen, fin, incremento ast.Nodo
    actual := p.actual().Literal
    switch actual {
    case "desde":
        p.avanzar()
        origen = p.parsearExpresion()
        // Limpiamos posibles saltos de línea antes de buscar 'hasta'
        for p.actual().Tipo == lexer.T_LINEA { p.avanzar() }
        if p.actual().Literal != "hasta" {
            fmt.Printf("❌ Error de Sintaxis (%d): Bucle 'para' incompleto. Se esperaba 'hasta'.\n", p.actual().Linea)
            os.Exit(1)
        }
        p.avanzar()
        fin = p.parsearExpresion()
        if p.actual().Literal == "incremento" {
            p.avanzar()
            incremento = p.parsearExpresion()
        }
    case "en":
        p.avanzar()
        origen = p.parsearExpresion()
    default:
        fmt.Printf("❌ Error de Sintaxis (%d): Uso incorrecto de 'para'. Se esperaba 'desde' o 'en'.\n", p.actual().Linea)
        os.Exit(1)
    }

    p.consumir(lexer.T_DOSPUNTOS)
    cuerpo := p.parsearBloque()

    return ast.Para{
        Base:        ast.Base{Pos: pos},
        Variable:    variable,
        Origen:      origen,
        Fin:         fin,
        Incremento:  incremento,
        Cuerpo:      cuerpo,
    }
}
