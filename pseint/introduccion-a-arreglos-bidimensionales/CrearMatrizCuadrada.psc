// Crea un programa que solicite al usuario ingresar 9 valores.
// Los valores ingresados deben ser almacenados en un único
// arreglo bidimensional (matriz).
// y debe mostrarlos posteriormente por pantalla.
// Sigue estos pasos:

// - Declara el tipo de dato que almacenará la matriz.

// - Define la dimensión del arreglo, en este caso,
//   3X3 (ya que precisamos almacenar 9 datos).

// - Utiliza un bucle para recorrer el arreglo recién creado,
//   posición por posición, y solicita al usuario que
//   introduzca un dato.
//   Puedes emplear una estructura de bucle "Para" para esta tarea.
//   Recuerda que necesitarás bucles anidados para recorrer cada
//   fila y cada columna, siendo el bucle externo para las filas
//   y el interno para las columnas.

// Puntos importantes:

// - Definición de la matriz y su tamaño:
//   Se establece una matriz de caracteres con un tamaño
//   predefinido de 3 filas por 3 columnas.

// - Recorrido de la matriz para ingresar datos:
//   Se utilizan bucles anidados para iterar sobre cada posición
//   de la matriz y se solicita al usuario que ingrese un valor
//   de tipo caracter para esa posición.
	
// - Método para mostrar la matriz:
//   Se define un subproceso llamado "MUESTRA" que recibe
//   la matriz como parámetro y muestra su contenido en la consola.


Algoritmo CrearMatrizCuadrada
	Definir TAMANO_LADO Como Entero
	Definir matriz2D Como Entero
	
	TAMANO_LADO = 3
	
	Dimension matriz2D[TAMANO_LADO, TAMANO_LADO]
	
	llenarMatriz2DManualmente(matriz2D, TAMANO_LADO, TAMANO_LADO, "Ingrese un numero entero para el elemento de la matriz")
	
	mostrarMatriz2D(matriz2D, TAMANO_LADO, TAMANO_LADO)
FinAlgoritmo

SubProceso llenarMatriz2DManualmente(matriz, tamanoColumna, tamanoFila, mensaje)
	Definir columna, fila Como Entero
	
	Para fila = 0 Hasta tamanoFila - 1 Con Paso 1 Hacer		
		Para columna = 0 Hasta tamanoColumna - 1 Con Paso 1 Hacer			
			Escribir mensaje
			Escribir "Posicion [", fila + 1, ", ", columna + 1, "](Fila, columna) de ", tamanoFila, "x", tamanoColumna
			Leer matriz[fila, columna]
		FinPara
	FinPara
FinSubProceso

SubProceso mostrarMatriz2D(matriz, tamanoColumna, tamanoFila)
	Definir columna, fila Como Entero
	Definir ultimaColumna, ultimaFila Como Entero
	
	ultimaColumna = tamanoColumna - 1
	ultimaFila = tamanoFila - 1
	
	Para fila = 0 Hasta ultimaFila Con Paso 1 Hacer
		Escribir Sin Saltar "["
		
		Para columna = 0 Hasta ultimaColumna Con Paso 1 Hacer
			Escribir Sin Saltar matriz[fila, columna]
			
			Si columna < ultimaColumna Entonces
				Escribir Sin Saltar ", "
			FinSi
		FinPara
		
		Escribir "]"
	FinPara
	
	Escribir ""
FinSubProceso