# Problema 2: Una pequeña empresa necesita automatizar la gestión y el
# seguimiento del desempeño de sus ventas para motivar a su equipo.
# Actualmente, los datos de ventas se manejan de forma manual, lo que
# dificulta la actualización rápida de los registros y la evaluación oportuna
# de los incentivos. Se requiere un sistema que pueda llevar un registro
# de las ventas por mes y determinar si un vendedor califica para un bono
# basado en un límite de ventas predefinido.

# Nombre:           Richard A. Vega Almanza
# Grupo:            213022_154
# Programa:         Ingenieria Electronica
# Codigo Fuente:    Autoria Propia
# Modelo:           qwen3.5:9b
# Ollama Stack:     Ollama con soporte para modelos locales
# Podman Stack:     Podman para containerization sin root
# Contenedor:       Podman con Dockerfile para despliegue

VENTAS_POR_MES = {}  # Diccionario vacío para llenar durante el proceso
LIMITE_BONO = 5000
def solicitar_datos():
    """Solicita un nombre y una cantidad al usuario."""
    nombre_vendedor = input("Ingrese su nombre: ")
    
    # Validar mes
    while True:
        mes = input("Ingrese el mes (1-12): ")
        try:
            mes_num = int(mes)
            if 1 <= mes_num <= 12:
                mes_valido = mes_num
                break
            else:
                print("Mes inválido. Por favor ingrese un número entre 1 y 12.")
        except ValueError:
            print("Entrada inválida. Por favor ingrese un número entero.")
    
    # Validar cantidad de ventas
    while True:
        try:
            cantidad_nueva = int(input(f"Ingrese las ventas de {mes_valido}: "))
            if cantidad_nueva >= 0:
                break
            else:
                print("La cantidad debe ser 0 o un número positivo.")
        except ValueError:
            print("Entrada inválida. Por favor ingrese un número entero.")
    
    return nombre_vendedor, mes_valido, cantidad_nueva

def agregar_ventas(datos_actuales, mes, monto):
    """Agrega un nuevo mes de ventas al diccionario."""
    datos_actuales[mes] = monto
    return datos_actuales

def revisar_bono(ventas_totales, limite):
    """Verifica si el vendedor califica para un bono."""
    if ventas_totales > limite:
        monto_bono = ventas_totales * 0.10  # 10% of total sales as bonus
        print(f"¡Felicidades! Gana un bono de: ${monto_bono:.2f}")
    else:
        print("Siga esforzándose para el bono.")

contador = 1
while contador < 3:
    print(f"\n--- Iteración {contador} ---")
    vendedor, mes, nuevas_ventas = solicitar_datos()
    VENTAS_POR_MES = agregar_ventas(VENTAS_POR_MES, mes, nuevas_ventas)
    total_anual = sum(VENTAS_POR_MES.values())
    try:
        revisar_bono(total_anual, LIMITE_BONO)
        print(f"Ventas de {vendedor}: {total_anual}. Ventas de {mes}: {VENTAS_POR_MES[mes]}")
    except Exception as e:
        print("Ocurrió un problema en el cálculo final.")