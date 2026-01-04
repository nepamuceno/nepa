package evaluador

import (
	"strings"
	"unicode"
)

type TipoToken string

const (
	TokenNumero        TipoToken = "NUMERO"
	TokenIdentificador TipoToken = "IDENTIFICADOR"
	TokenOperador      TipoToken = "OPERADOR"
	TokenParenIzq      TipoToken = "PAREN_IZQ"
	TokenParenDer      TipoToken = "PAREN_DER"
	TokenComa          TipoToken = "COMA"
	TokenDesconocido   TipoToken = "DESCONOCIDO"
)

type Token struct {
	Tipo  TipoToken
	Valor string
}

// Lexer convierte el string de la expresión en un slice de tokens.
func Lexer(input string) []Token {
	var tokens []Token
	runas := []rune(input)
	n := len(runas)

	for i := 0; i < n; i++ {
		r := runas[i]

		if unicode.IsSpace(r) {
			continue
		}

		// 1. Números
		if unicode.IsDigit(r) || r == '.' {
			inicioNum := i
			for i+1 < n && (unicode.IsDigit(runas[i+1]) || runas[i+1] == '.') {
				i++
			}
			tokens = append(tokens, Token{TokenNumero, string(runas[inicioNum : i+1])})
			continue
		}

		// 2. Identificadores
		if unicode.IsLetter(r) || r == '_' {
			inicioId := i
			for i+1 < n && (unicode.IsLetter(runas[i+1]) || unicode.IsDigit(runas[i+1]) || runas[i+1] == '_') {
				i++
			}
			tokens = append(tokens, Token{TokenIdentificador, string(runas[inicioId : i+1])})
			continue
		}

		// 3. Paréntesis y Comas
		switch r {
		case '(':
			tokens = append(tokens, Token{TokenParenIzq, "("})
			continue
		case ')':
			tokens = append(tokens, Token{TokenParenDer, ")"})
			continue
		case ',':
			tokens = append(tokens, Token{TokenComa, ","})
			continue
		}

		// 4. Operadores (Simples y Compuestos)
		if strings.ContainsRune("+-*/^%!&|<>= ", r) {
			if i+1 < n {
				siguiente := runas[i+1]
				combinado := string(r) + string(siguiente)
				operadoresDobles := []string{"==", "!=", "<=", ">=", "&&", "||"}
				
				esDoble := false
				for _, op := range operadoresDobles {
					if combinado == op {
						esDoble = true
						break
					}
				}

				if esDoble {
					tokens = append(tokens, Token{TokenOperador, combinado})
					i++ 
					continue
				}
			}
			tokens = append(tokens, Token{TokenOperador, string(r)})
			continue
		}

		// 5. Desconocido
		tokens = append(tokens, Token{TokenDesconocido, string(r)})
	}
	return tokens
}
