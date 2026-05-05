# ==================================================
# CLASE PRÁCTICA: INTRODUCCIÓN A TKINTER
# ==================================================
# Tkinter es la librería estándar de Python para crear
# interfaces gráficas de usuario (ventanas, botones, etc.)
# ==================================================

# ------------------------------------------
# PASO 1: Importar la librería tkinter
# ------------------------------------------
# tkinter viene incluido con Python, no necesita instalación adicional
import tkinter as tk
# Importamos messagebox para ventanas de diálogo (alertas, confirmaciones)
from tkinter import messagebox

# ------------------------------------------
# PASO 2: Crear la ventana principal
# ------------------------------------------
# Tk() crea la ventana raíz de la aplicación
ventana = tk.Tk()

# Configuramos el título que aparece en la barra de la ventana
ventana.title("Mi primera aplicación con Tkinter")

# Configuramos el tamaño de la ventana (ancho x alto)
ventana.geometry("500x400")

# Configuramos el color de fondo de la ventana
ventana.configure(bg="#f0f0f0")  # Código hexadecimal para gris claro

# ------------------------------------------
# PASO 3: Crear etiquetas (Labels)
# ------------------------------------------
# Las etiquetas se usan para mostrar texto en la ventana
# Nota: aún no se colocan en la ventana, solo se crean

# Etiqueta de título (grande y en negrita)
label_titulo = tk.Label(
    ventana,                    # ventana padre (donde va)
    text="Bienvenido a Tkinter", # texto a mostrar
    font=("Arial", 20, "bold"), # (fuente, tamaño, estilo)
    fg="#2c3e50",               # foreground = color del texto
    bg="#f0f0f0",               # background = color de fondo
    pady=20                     # padding vertical interno
)

# Etiqueta de instrucción
label_instruccion = tk.Label(
    ventana,
    text="Escribe tu nombre y presiona el botón:",
    font=("Arial", 12),
    fg="#34495e",
    bg="#f0f0f0"
)

# ------------------------------------------
# PASO 4: Crear campos de entrada (Entry)
# ------------------------------------------
# Los campos de entrada permiten al usuario escribir texto

# Variable que almacenará el texto ingresado
# StringVar() es un tipo especial de variable que Tkinter puede vigilar
nombre_var = tk.StringVar()

# Campo de entrada
entry_nombre = tk.Entry(
    ventana,
    textvariable=nombre_var,   # vincula la variable al campo
    font=("Arial", 12),
    width=30,                   # ancho en caracteres
    bd=3,                       # border width (grosor del borde)
    relief="solid",             # tipo de borde (solid, groove, ridge, etc.)
    fg="#2c3e50"
)

# Texto de ayuda dentro del campo (placeholder)
entry_nombre.insert(0, "Ej: Ana Pérez")

# ------------------------------------------
# PASO 5: FUNCIÓN que se ejecutará al presionar el botón
# ------------------------------------------
def saludar():
    """
    Esta función se ejecuta cuando el usuario hace clic en el botón.
    Toma el nombre ingresado y muestra un mensaje de saludo.
    """
    # Obtenemos el texto del campo de entrada
    nombre = nombre_var.get()
    
    # Verificamos si el usuario escribió algo
    if nombre and nombre != "Ej: Ana Pérez":
        # Mostramos un mensaje emergente (popup)
        messagebox.showinfo("Saludo", f"¡Hola {nombre}! Bienvenido a Tkinter")
        
        # También actualizamos la etiqueta de resultado
        label_resultado.config(text=f"Te saludo, {nombre}")
        
        # Limpiamos el campo de entrada para la siguiente persona
        entry_nombre.delete(0, tk.END)
        entry_nombre.insert(0, "Ej: Ana Pérez")
    else:
        # Si no escribió nada, mostramos un mensaje de advertencia
        messagebox.showwarning("Campo vacío", "Por favor, escribe tu nombre.")

# ------------------------------------------
# PASO 6: Crear botones (Button)
# ------------------------------------------
boton_saludar = tk.Button(
    ventana,
    text="¡Saludar!",           # texto del botón
    command=saludar,            # función que se ejecuta al hacer clic
    font=("Arial", 12, "bold"),
    bg="#3498db",               # color de fondo del botón
    fg="white",                 # color del texto
    activebackground="#2980b9", # color al hacer clic
    activeforeground="white",   # color del texto al hacer clic
    padx=20,                    # padding horizontal interno
    pady=8,                     # padding vertical interno
    cursor="hand2"              # cambia el cursor a mano al pasar sobre el botón
)

# ------------------------------------------
# PASO 7: Crear etiqueta para mostrar resultados
# ------------------------------------------
label_resultado = tk.Label(
    ventana,
    text="",                    # empieza vacío
    font=("Arial", 12, "italic"),
    fg="#27ae60",               # color verde
    bg="#f0f0f0",
    pady=15
)

# ------------------------------------------
# PASO 8: ORGANIZAR los elementos en la ventana
# ------------------------------------------
# MÉTODO pack(): organiza elementos en orden vertical u horizontal
# Es el método más simple para empezar

label_titulo.pack()           # se coloca arriba
label_instruccion.pack(pady=10)  # pady = espacio vertical externo
entry_nombre.pack(pady=5)
boton_saludar.pack(pady=15)
label_resultado.pack()

# ------------------------------------------
# PASO 9: EJECUTAR la aplicación
# ------------------------------------------
# mainloop() inicia el bucle de eventos de Tkinter
# Esto mantiene la ventana abierta y responde a las interacciones del usuario
# Es OBLIGATORIO para que la ventana se muestre
ventana.mainloop()

# El programa se queda aquí hasta que el usuario cierre la ventana
# Después de cerrar, continúa con el código siguiente (si hay)
print("La aplicación se ha cerrado")
