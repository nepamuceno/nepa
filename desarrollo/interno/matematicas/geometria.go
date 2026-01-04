package matematicas

import (
	"math"
	"nepa/desarrollo/interno/evaluador"
)

func inyectarGeometriaGlobal() {

	// --- 1. GEOMETRÍA 2D (ÁREAS Y PERÍMETROS) ---

	// area_circulo(radio)
	evaluador.Funciones["area_circulo"] = func(args ...interface{}) (interface{}, error) {
		r, err := validar1("area_circulo", args); if err != nil { return nil, err }
		return finalizar("area_circulo", math.Pi*math.Pow(r, 2))
	}

	// perimetro_circulo(radio)
	evaluador.Funciones["perimetro_circulo"] = func(args ...interface{}) (interface{}, error) {
		r, err := validar1("perimetro_circulo", args); if err != nil { return nil, err }
		return finalizar("perimetro_circulo", 2*math.Pi*r)
	}

	// area_triangulo(base, altura)
	evaluador.Funciones["area_triangulo"] = func(args ...interface{}) (interface{}, error) {
		b, a, err := validar2("area_triangulo", args); if err != nil { return nil, err }
		return finalizar("area_triangulo", (b*a)/2)
	}

	// area_heron(lado_a, lado_b, lado_c) -> Área de triángulo sin conocer la altura
	evaluador.Funciones["area_heron"] = func(args ...interface{}) (interface{}, error) {
		a, b, c, err := validar3("area_heron", args); if err != nil { return nil, err }
		s := (a + b + c) / 2 // Semiperímetro
		return finalizar("area_heron", math.Sqrt(s*(s-a)*(s-b)*(s-c)))
	}

	// area_poligono_regular(num_lados, longitud_lado)
	evaluador.Funciones["area_poligono_regular"] = func(args ...interface{}) (interface{}, error) {
		n, l, err := validar2("area_poligono_regular", args); if err != nil { return nil, err }
		// Fórmula: (n * l^2) / (4 * tan(pi/n))
		denominador := 4 * math.Tan(math.Pi/n)
		return finalizar("area_poligono_regular", (n*math.Pow(l, 2))/denominador)
	}

	// --- 2. TEOREMAS Y DISTANCIAS ---

	// pitagoras_hipotenusa(cateto_a, cateto_b)
	evaluador.Funciones["pitagoras_hipotenusa"] = func(args ...interface{}) (interface{}, error) {
		a, b, err := validar2("pitagoras_hipotenusa", args); if err != nil { return nil, err }
		return finalizar("pitagoras_hipotenusa", math.Hypot(a, b))
	}

	// distancia_2d(x1, y1, x2, y2)
	evaluador.Funciones["distancia_2d"] = func(args ...interface{}) (interface{}, error) {
		x1, y1, x2, y2, err := validar4("distancia_2d", args); if err != nil { return nil, err }
		return finalizar("distancia_2d", math.Sqrt(math.Pow(x2-x1, 2)+math.Pow(y2-y1, 2)))
	}

	// --- 3. GEOMETRÍA 3D (VOLÚMENES Y SUPERFICIES) ---

	// volumen_esfera(radio)
	evaluador.Funciones["volumen_esfera"] = func(args ...interface{}) (interface{}, error) {
		r, err := validar1("volumen_esfera", args); if err != nil { return nil, err }
		return finalizar("volumen_esfera", (4.0/3.0)*math.Pi*math.Pow(r, 3))
	}

	// area_superficie_esfera(radio)
	evaluador.Funciones["area_superficie_esfera"] = func(args ...interface{}) (interface{}, error) {
		r, err := validar1("area_superficie_esfera", args); if err != nil { return nil, err }
		return finalizar("area_superficie_esfera", 4*math.Pi*math.Pow(r, 2))
	}

	// volumen_cilindro(radio, altura)
	evaluador.Funciones["volumen_cilindro"] = func(args ...interface{}) (interface{}, error) {
		r, h, err := validar2("volumen_cilindro", args); if err != nil { return nil, err }
		return finalizar("volumen_cilindro", math.Pi*math.Pow(r, 2)*h)
	}

	// volumen_cono(radio, altura)
	evaluador.Funciones["volumen_cono"] = func(args ...interface{}) (interface{}, error) {
		r, h, err := validar2("volumen_cono", args); if err != nil { return nil, err }
		return finalizar("volumen_cono", (1.0/3.0)*math.Pi*math.Pow(r, 2)*h)
	}

	// volumen_piramide(area_base, altura)
	evaluador.Funciones["volumen_piramide"] = func(args ...interface{}) (interface{}, error) {
		ab, h, err := validar2("volumen_piramide", args); if err != nil { return nil, err }
		return finalizar("volumen_piramide", (ab*h)/3.0)
	}

	// --- 4. CONVERSIONES ---

	// grados_a_rad(grados)
	evaluador.Funciones["grados_a_rad"] = func(args ...interface{}) (interface{}, error) {
		g, err := validar1("grados_a_rad", args); if err != nil { return nil, err }
		return finalizar("grados_a_rad", g*(math.Pi/180))
	}

	// rad_a_grados(radianes)
	evaluador.Funciones["rad_a_grados"] = func(args ...interface{}) (interface{}, error) {
		r, err := validar1("rad_a_grados", args); if err != nil { return nil, err }
		return finalizar("rad_a_grados", r*(180/math.Pi))
	}
}
