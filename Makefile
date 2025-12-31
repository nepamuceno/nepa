# Nombres de los binarios
INTERPRETE=nepa
GENERADOR=nepa_lib

# Rutas de origen
SRC_MAIN=./desarrollo/cmd/nepa/main.go
SRC_LIB=./desarrollo/cmd/nepa_lib/main.go

# Ruta de destino
DEST=./dist/bin

all: clean build install

build:
	@echo "🔨 Compilando Nepa (Intérprete)..."
	@go build -o $(INTERPRETE) $(SRC_MAIN)
	@echo "🔨 Compilando Generador de Librerías..."
	@go build -o $(GENERADOR) $(SRC_LIB)

install:
	@mkdir -p $(DEST)
	@mkdir -p ./dist/lib
	@mv $(INTERPRETE) $(DEST)/
	@mv $(GENERADOR) $(DEST)/
	@echo "✅ Binarios instalados en $(DEST)/"

clean:
	@rm -f $(INTERPRETE) $(GENERADOR)
	@rm -rf ./dist/bin/*
	@echo "🧹 Limpieza completada."




