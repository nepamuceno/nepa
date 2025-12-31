#!/bin/bash

GIST_INPUT="gist.txt"

if [ ! -f "$GIST_INPUT" ]; then
    echo "Error: No se encuentra $GIST_INPUT"
    exit 1
fi

# Usamos AWK para segmentar el archivo de forma precisa
awk '
/^(\/.*):$/ {
    # Cuando encontramos una nueva ruta de archivo
    if (file != "") {
        close(file);
    }
    
    # Limpiamos los dos puntos de la ruta
    file = $0;
    sub(/:$/, "", file);
    
    # Creamos el directorio si no existe (llamada al sistema)
    system("mkdir -p \"$(dirname \""file"\")\"");
    
    print "Restaurando: " file;
    next;
}

{
    # Si tenemos un archivo abierto, escribimos la línea actual
    if (file != "") {
        print $0 > file;
    }
}
' "$GIST_INPUT"

echo "--- Restauración finalizada ---"
