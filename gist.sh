#!/bin/bash

OUTPUT="gist.txt"
> "$OUTPUT"
BASE_DIR=$(pwd)

echo "Generando $OUTPUT..."

find "$BASE_DIR" -type f -not -path '*/.*' -not -path '*/dist/*' -not -name "$OUTPUT" -not -name "gist.sh" | while read -r file; do
    if file "$file" | grep -qE 'text|JSON|source|empty'; then
        echo "$file:" >> "$OUTPUT"
        cat "$file" >> "$OUTPUT"
        echo -e "\n\n\n" >> "$OUTPUT"
    fi
done

# --- NUEVA PARTE: COPIAR AL PORTAPAPELES ---
if command -v xclip &> /dev/null; then
    cat "$OUTPUT" | xclip -selection clipboard
    echo "✅ Archivo generado y COPIADO al portapapeles. ¡Dale Ctrl+V en GitHub!"
else
    echo "✅ Archivo generado en: $BASE_DIR/$OUTPUT"
    echo "⚠️  xclip no está instalado. Instálalo con 'sudo apt install xclip' para copiar automáticamente."
fi
