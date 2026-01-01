package parser

import (
    "nepa/desarrollo/interno/ast"
    "nepa/desarrollo/interno/lexer"
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

func (p *Parser) Parsear() []ast.Nodo {
    var programa []ast.Nodo
    for !p.esFin() {
        if p.actual().Tipo == lexer.T_LINEA {
            p.avanzar()
            continue
        }
        nodo := p.parsearInstruccion()
        if nodo != nil {
            programa = append(programa, nodo)
        } else {
            // EVITAR LOOP INFINITO: Si nada parsea, avanzamos sí o sí
            p.avanzar()
        }
    }
    return programa
}

func (p *Parser) parsearInstruccion() ast.Nodo {
    t := p.actual()
    
    // Caso: Asignación directa (identificador = ...)
    if t.Tipo == lexer.T_IDENT && p.pos+1 < len(p.tokens) && p.tokens[p.pos+1].Tipo == lexer.T_IGUAL {
        return p.parsearAsignacionDirecta()
    }

    if t.Tipo == lexer.T_KEYWORD {
        switch t.Literal {
        case "variable": return p.parsearAsignacion()
        case "incluir":  return p.parsearIncluir()
        case "imprime":  
            p.avanzar()
            return p.parsearLlamada("imprime")
        case "si":       return p.parsearSi()
        case "mientras": return p.parsearMientras()
        case "para":     return p.parsearPara()
        case "funcion":  return p.parsearDefFuncion()
        case "retorna":
            pos := p.getPos(); p.avanzar()
            return ast.Retornar{Base: ast.Base{Pos: pos}, Valor: p.parsearExpresion()}
        }
    }
    return p.parsearExpresion()
}

func (p *Parser) parsearAsignacionDirecta() ast.Nodo {
    pos := p.getPos()
    nombre := p.consumir(lexer.T_IDENT).Literal
    p.consumir(lexer.T_IGUAL)
    return ast.Asignacion{Base: ast.Base{Pos: pos}, Nombre: nombre, Valor: p.parsearExpresion()}
}

func (p *Parser) parsearMientras() ast.Nodo {
    pos := p.getPos()
    p.avanzar() // saltar 'mientras'
    cond := p.parsearExpresion()
    p.consumir(lexer.T_DOSPUNTOS)
    return ast.Mientras{Base: ast.Base{Pos: pos}, Condicion: cond, Cuerpo: p.parsearBloque()}
}

func (p *Parser) parsearBloque() []ast.Nodo {
    if p.actual().Tipo == lexer.T_LINEA { p.avanzar() }
    p.consumir(lexer.T_INDENT)
    var nodos []ast.Nodo
    for !p.esFin() && p.actual().Tipo != lexer.T_DEDENT {
        if p.actual().Tipo == lexer.T_LINEA {
            p.avanzar()
            continue
        }
        nodo := p.parsearInstruccion()
        if nodo != nil {
            nodos = append(nodos, nodo)
        } else {
            p.avanzar()
        }
    }
    p.consumir(lexer.T_DEDENT)
    return nodos
}

// --- JERARQUÍA DE EXPRESIONES (Para evitar 1 + 2 * 3 = 0) ---

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
        // Caso: modulo.funcion(...)
        if p.actual().Tipo == lexer.T_PUNTO {
            p.avanzar()
            prop := p.consumir(lexer.T_IDENT).Literal
            if p.actual().Tipo == lexer.T_PARENT_A {
                return p.parsearLlamadaModulo(nombre, prop)
            }
            return ast.LlamadaModulo{Base: ast.Base{Pos: pos}, Modulo: nombre, Funcion: prop}
        }
        // Caso: funcion(...)
        if p.actual().Tipo == lexer.T_PARENT_A {
            return p.parsearLlamada(nombre)
        }
        return ast.Identificador{Base: ast.Base{Pos: pos}, Nombre: nombre}
    case lexer.T_PARENT_A:
        p.avanzar()
        exp := p.parsearExpresion()
        p.consumir(lexer.T_PARENT_C)
        return exp
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
    if p.actual().Literal == "sino" {
        p.avanzar()
        if p.actual().Literal == "si" {
            sino = append(sino, p.parsearSi())
        } else {
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
    p.consumir(lexer.T_PARENT_C); p.consumir(lexer.T_DOSPUNTOS)
    return ast.FuncionDef{Base: ast.Base{Pos: pos}, Nombre: nombre, Parametros: params, Cuerpo: p.parsearBloque()}
}

func (p *Parser) parsearAsignacion() ast.Nodo {
    pos := p.getPos(); p.avanzar()
    nombre := p.consumir(lexer.T_IDENT).Literal
    p.consumir(lexer.T_IGUAL)
    return ast.Asignacion{Base: ast.Base{Pos: pos}, Nombre: nombre, Valor: p.parsearExpresion()}
}

func (p *Parser) parsearIncluir() ast.Nodo {
    pos := p.getPos(); p.avanzar()
    nombre := p.consumir(lexer.T_CADENA).Literal
    return ast.LlamadaFuncion{Base: ast.Base{Pos: pos}, Nombre: "incluir", Args: []ast.Nodo{ast.Literal{Base: ast.Base{Pos: pos}, Valor: nombre}}}
}

func (p *Parser) actual() lexer.Token { 
    if p.pos >= len(p.tokens) { return lexer.Token{Tipo: lexer.T_EOF} }
    return p.tokens[p.pos] 
}
func (p *Parser) avanzar() { p.pos++ }
func (p *Parser) esFin() bool { return p.actual().Tipo == lexer.T_EOF }
func (p *Parser) consumir(t lexer.TipoToken) lexer.Token { 
    tok := p.actual()
    p.avanzar() 
    return tok 
}

func (p *Parser) parsearPara() ast.Nodo {
    pos := p.getPos()
    p.avanzar() // saltar 'para'
    variable := p.consumir(lexer.T_IDENT).Literal
    
    var origen, fin, incremento ast.Nodo
    
    if p.actual().Literal == "desde" {
        p.avanzar()
        origen = p.parsearExpresion()
        if p.actual().Literal == "hasta" {
            p.avanzar()
            fin = p.parsearExpresion()
        }
        if p.actual().Literal == "incremento" {
            p.avanzar()
            incremento = p.parsearExpresion()
        }
    } else if p.actual().Literal == "en" {
        p.avanzar()
        origen = p.parsearExpresion()
    }

    p.consumir(lexer.T_DOSPUNTOS)
    cuerpo := p.parsearBloque()
    
    return ast.Para{
        Base: ast.Base{Pos: pos},
        Variable: variable,
        Origen: origen,
        Fin: fin,
        Incremento: incremento,
        Cuerpo: cuerpo,
    }
}
