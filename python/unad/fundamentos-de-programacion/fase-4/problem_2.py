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
# Contenedor:      Podman con Dockerfile para despliegue

VENTAS_POR_MES = {"enero": 1500, "febrero": 2200, "marzo": 1800}
LIMITE_BONO = 5000
def solicitar_datos():
    """Solicita un nombre y una cantidad al usuario."""
    nombre_vendedor = input("Ingrese su nombre: ")
    try:
        cantidad_nueva = int(input("Ingrese las ventas de abril: "))
    except:
        print("Entrada inválida, usando 0.")
        cantidad_nueva = 0
    return nombre_vendedor, cantidad_nueva

def agregar_ventas(datos_actuales, mes, monto):
    """Agrega un nuevo mes de ventas al diccionario."""
    datos_actuales[mes] = monto
    return datos_actuales  # ← FIX: Return the dictionary, not a list

def revisar_bono(ventas_totales, limite):
    """Verifica si el vendedor califica para un bono."""
    if ventas_totales > limite:
        # ← FIX: Calculate bonus as a percentage or fixed amount
        monto_bono = ventas_totales * 0.10  # 10% of total sales as bonus
        print(f"¡Felicidades! Gana un bono de: ${monto_bono:.2f}")
    else:
        print("Siga esforzándose para el bono.")

contador = 1
while contador < 3:
    print(f"\n--- Iteración {contador} ---")
    vendedor, nuevas_ventas = solicitar_datos()
    VENTAS_POR_MES = agregar_ventas(VENTAS_POR_MES, "abril", nuevas_ventas)  # ← FIX: Use returned dict
    total_anual = sum(VENTAS_POR_MES.values())
    try:
        revisar_bono(total_anual, LIMITE_BONO)
        print(f"Ventas de {vendedor}: {total_anual}. Ventas de abril: {VENTAS_POR_MES['abril']}")
    except Exception as e:
        print("Ocurrió un problema en el cálculo final.")