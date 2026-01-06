package evaluador

import "fmt"

// ErrorConversion representa un fallo al convertir tipos.
// Incluye el nombre del comando, la ayuda integrada y el valor recibido.
type ErrorConversion struct {
    Comando string
    Ayuda   string
    Valor   interface{}
}

// Error devuelve el mensaje formateado para mostrar en terminal.
// Incluye el nombre del comando, la ayuda y el valor problemático.
func (e ErrorConversion) Error() string {
    return fmt.Sprintf("❌ ERROR en %s\n%s\nValor recibido: %v", e.Comando, e.Ayuda, e.Valor)
}

// NuevaErrorConversion crea un error uniforme de conversión.
// Se usa en cada convertir_* para devolver un error claro y consistente.
func NuevaErrorConversion(comando, ayuda string, valor interface{}) error {
    return ErrorConversion{
        Comando: comando,
        Ayuda:   ayuda,
        Valor:   valor,
    }
}
