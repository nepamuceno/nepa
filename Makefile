# Nombres de los binarios
INTERPRETE=nepa

# Rutas de origen
SRC_MAIN=./desarrollo/cmd/nepa

# Ruta de destino
DEST=./dist/bin

.PHONY: all build install clean

all: clean build install

build:
	@echo "Compilando Nepa (Intérprete)..."
	@mkdir -p $(DEST)
	@go build -o $(DEST)/$(INTERPRETE) $(SRC_MAIN)

install:
	@echo "Binarios instalados en $(DEST)/"

clean:
	@rm -rf $(DEST)/*
	@echo "Limpieza completada."
