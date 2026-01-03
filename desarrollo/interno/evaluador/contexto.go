package evaluador

// Contexto representa el entorno de ejecución del intérprete.
type Contexto struct {
    Variables  map[string]interface{}                       // Variables locales
    Globales   map[string]interface{}                       // Variables globales
    Constantes map[string]interface{}                       // Constantes definidas
    Funciones  map[string]func(...interface{}) interface{}  // Funciones registradas
}
