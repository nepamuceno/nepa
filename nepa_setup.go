package main

import (
	"fmt"
	"os"
)

func main() {
	// Definimos la carpeta base de desarrollo
	base := "desarrollo"

	// 1. Carpetas de Desarrollo
	devDirs := []string{
		base + "/cmd/nepa",
		base + "/interno/lexer",
		base + "/interno/parser",
		base + "/interno/ast",
		base + "/interno/evaluador",
		base + "/interno/modulo",
		base + "/librerias_estandar",
	}

	// 2. Carpetas de Distribución
	distDirs := []string{
		"dist/bin",
		"dist/lib",
	}

	fmt.Println("🏗️  Configurando entorno Nepa...")

	// Crear carpetas de desarrollo
	for _, dir := range devDirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Printf("❌ Error creando %s: %v\n", dir, err)
		} else {
			fmt.Printf("📂 Creada: %s\n", dir)
		}
	}

	// Crear carpetas de distribución
	for _, dir := range distDirs {
		os.MkdirAll(dir, 0755)
		fmt.Printf("📦 Creada: %s\n", dir)
	}

	// 3. Crear archivos iniciales .go
	archivos := map[string]string{
		base + "/cmd/nepa/main.go":           "main",
		base + "/interno/lexer/lexer.go":      "lexer",
		base + "/interno/parser/parser.go":    "parser",
		base + "/interno/ast/ast.go":          "ast",
		base + "/interno/modulo/cargador.go":  "modulo",
	}

	for ruta, pkg := range archivos {
		contenido := fmt.Sprintf("package %s\n\n// Módulo %s del lenguaje Nepa\n", pkg, pkg)
		os.WriteFile(ruta, []byte(contenido), 0644)
	}

	fmt.Println("\n✅ Estructura creada con éxito.")
}




