package modulo

type LibreriaNepa struct {
	Nombre    string
	Variables map[string]interface{}
	Funciones map[string]func(...interface{}) interface{}
	Valores   map[string]float64
}

func Cargar(nombre string) (*LibreriaNepa, error) {
	lib := &LibreriaNepa{
		Nombre:    nombre,
		Variables: make(map[string]interface{}),
		Funciones: make(map[string]func(...interface{}) interface{}),
		Valores:   make(map[string]float64),
	}

	if nombre == "matematicas" {
		InyectarMatematicas(lib)
		// Pasamos las constantes al mapa de Valores
		for k, v := range lib.Variables {
			if val, ok := v.(float64); ok {
				lib.Valores[k] = val
			}
		}
		return lib, nil
	}
	return lib, nil
}

func (l *LibreriaNepa) Exportar(ruta string) error { return nil }




