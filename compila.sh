#!/bin/bash

# Colores para la terminal
VERDE='\033[0;32m'
AZUL='\033[0;34m'
ROJO='\033[0;31m'
NC='\033[0m' # Sin color

echo -e "${AZUL}🛠️  Limpiando y Compilando...${NC}"
make clean && make

if [ $? -eq 0 ]; then
    echo -e "${VERDE}✅ Compilación exitosa.${NC}"
    
    echo -e "${AZUL}🚀 Ejecutando inicio.nepa...${NC}"
    echo "--------------------------"
    ./dist/bin/nepa test.nepa
else
    echo -e "${ROJO}❌ Error en la compilación.${NC}"
    exit 1
fi
