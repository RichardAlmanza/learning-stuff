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

# VENTAS_POR_MES: Diccionario anidado con estructura {mes: {nombre_vendedor: monto_ventas}}
# - Clave principal: número de mes (1-12)
# - Clave secundaria: nombre del vendedor
# - Valor: monto de ventas del vendedor en ese mes
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

def agregar_ventas(datos_actuales, mes, monto, nombre):
    """Agrega un nuevo mes de ventas al diccionario."""
    # Initialize the month if it doesn't exist
    if mes not in datos_actuales:
        datos_actuales[mes] = {}

    # Store salesman's sales for this month
    datos_actuales[mes][nombre] = monto
    return datos_actuales

def revisar_bono(ventas_por_mes, limite, nombre_vendedor):
    """Verifica si el vendedor califica para un bono."""
    total_ventas_vendedor = calcular_total_ventas(ventas_por_mes, nombre_vendedor)

    if total_ventas_vendedor > limite:
        print(f"¡Felicidades {nombre_vendedor}! Ganas el bono por superar el límite.")
        return True
    else:
        print(f"Siga esforzándose {nombre_vendedor} para el bono.")
        return False


def calcular_total_ventas(ventas_por_mes, nombre_vendedor):
    """Calcula el total de ventas de un vendedor específico."""
    total_ventas = 0
    for mes_vendedor in ventas_por_mes.values():
        if nombre_vendedor in mes_vendedor:
            total_ventas += mes_vendedor[nombre_vendedor]
    return total_ventas


def menu_principal():
    """Menú principal que orquesta la ejecución del programa."""
    global VENTAS_POR_MES
    print("=" * 50)
    print("Sistema de Gestión de Ventas y Bono")
    print("=" * 50)

    while True:
        print("\n--- Menú Principal ---")
        print("1. Registrar nuevas ventas")
        print("2. Ver reportes de ventas")
        print("3. Salir del programa")
        opcion = input("Seleccione una opción (1-3): ")

        if opcion == "1":
            print("\n--- Registrar Nuevas Ventas ---")
            vendedor, mes, nuevas_ventas = solicitar_datos()
            VENTAS_POR_MES = agregar_ventas(VENTAS_POR_MES, mes, nuevas_ventas, vendedor)
            total = calcular_total_ventas(VENTAS_POR_MES, vendedor)
            revisar_bono(VENTAS_POR_MES, LIMITE_BONO, vendedor)
            print(f"Ventas totales de {vendedor}: ${total:.2f}")
            print(f"Ventas de {mes}: {VENTAS_POR_MES.get(mes, {})}")

        elif opcion == "2":
            print("\n--- Reportes de Ventas ---")
            print(VENTAS_POR_MES)
            if VENTAS_POR_MES:
                print(f"Total de ventas anuales: ${sum(sum(v.values()) for v in VENTAS_POR_MES.values()):.2f}")
                vendedores = {nombre for mes_vendedores in VENTAS_POR_MES.values() for nombre in mes_vendedores.keys()}
                print(f"Número de vendedores registrados: {len(vendedores)}")

        elif opcion == "3":
            print("\n¡Gracias por usar el sistema de gestión de ventas!")
            break
        else:
            print("Opción inválida. Por favor seleccione 1, 2 o 3.")

menu_principal()
