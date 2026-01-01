// =========================================================
// TEST DE ESTRES INTEGRAL: SDK MATEMATICAS NEPA
// =========================================================
# Calculo de area y condicional en Nepa
radio = 5
area = pi * radio * radio

si (area > 50) {
    imprime("El area es grande:", area)
} sino {
    imprime("El area es pequeña:", area)
}

x = 0
mientras (x < 3) {
    imprime("Contador:", x)
    x = x + 1
}

imprime("--- INICIANDO DIAGNOSTICO DEL SISTEMA ---")

// TEST 1 SISTEMA DE AYUDA
imprime("\n[1] Verificando Sistema de Ayuda...")
imprime(ayuda("vol_cono"))
imprime(ayuda("proyectil_pos"))
imprime(ayuda("es_primo"))

// TEST 2 CONSTANTES
imprime("\n[2] Validando Constantes...")
imprime("PI: " + pi)
imprime("E: " + e)
imprime("Gravedad: " + gravedad)
imprime("PHI: " + phi)
imprime("Velocidad de la Luz: " + luz)

// TEST 3 GEOMETRIA 3D
imprime("\n[3] Calculando Volumenes Complejos...")
radio = 5.5
altura = 12.0
v_esfera = vol_esfera(radio)
v_cono = vol_cono(radio, altura)
v_cil = vol_cilindro(radio, altura)

imprime("Esfera (r=5.5): " + formatear(v_esfera, 4))
imprime("Cono (r=5.5, h=12): " + formatear(v_cono, 4))
imprime("Cilindro (r=5.5, h=12): " + formatear(v_cil, 4))

// TEST 4 TEORIA DE NUMEROS
imprime("\n[4] Stress de Teoria de Numeros (Buscando Primos hasta 50)...")
n = 1
mientras n <= 50 {
    si es_primo(n) {
        imprime("Encontrado primo: " + n)
    }
    n = n + 1
}

// TEST 5 ESTADISTICA
imprime("\n[5] Stress de Estadistica...")
avg = media(10, 20, 30, 40, 50, 60, 70, 80, 90, 100)
var_val = varianza(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
desv = desviacion_est(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

imprime("Media de 10-100: " + avg)
imprime("Varianza de 1-10: " + formatear(var_val, 2))
imprime("Desviacion Estandar: " + formatear(desv, 4))

// TEST 6 FISICA CINEMATICA
imprime("\n[6] Simulando Trayectoria de Proyectil (V0=50m/s, Ang=45)...")
t = 0.0
mientras t <= 5.0 {
    pos = proyectil_pos(50, 45, t)
    imprime("Tiempo " + t + "s -> Posicion [x,y]: " + pos)
    t = t + 1.0
}

// TEST 7 BITWISE Y BASES
imprime("\n[7] Operaciones de Bajo Nivel...")
num = 255
imprime("Decimal: " + num)
imprime("Binario: " + binario(num))
imprime("Hexadecimal: " + hex(num))
imprime("Bit XOR (255, 170): " + bit_xor(255, 170))

// TEST 8 FINANZAS
imprime("\n[8] Test de Interes y Potencia...")
capital = 1000
tasa = 0.05
tiempo = 10
total = interes_compuesto(capital, tasa, tiempo)
imprime("Capital Final (1000 al 5% x 10 años): " + formatear(total, 2))

// TEST 9 TRIGONOMETRIA
imprime("\n[9] Funciones Hiperbolicas...")
val = 1.0
imprime("senoh(1.0): " + formatear(senoh(val), 6))
imprime("cosenoh(1.0): " + formatear(cosenoh(val), 6))

imprime("\n--- TEST DE ESTRES FINALIZADO CON EXITO ---")
