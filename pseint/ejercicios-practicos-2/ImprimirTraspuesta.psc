// Crear una matriz de tamaño n x m, donde n y m son valores ingresados por el usuario.
// Llenar la matriz con números aleatorios comprendidos entre 1 y 100,
// luego mostrar su traspuesta. En caso de no estar familiarizado con el concepto de traspuesta,
// puedes consultar la siguiente referencia:Matriz Traspuesta.

Algoritmo ImprimirTraspuesta
	Definir LIMITE_INFERIOR, LIMITE_SUPERIOR Como Entero
	Definir matriz2D, matriz2DTraspuesta, tamanoFila, tamanoColumna Como Entero
	
	LIMITE_INFERIOR = 1
	LIMITE_SUPERIOR = 100
	
	Repetir
		Escribir "Ingrese el # de filas para la matriz"
		Leer tamanoFila
	Hasta Que tamanoFila > 0 
	
	Repetir
		Escribir "Ingrese el # de columnas para la matriz"
		Leer tamanoColumna
	Hasta Que tamanoColumna > 0
	
	Dimension matriz2D[tamanoFila, tamanoColumna], matriz2DTraspuesta[tamanoColumna, tamanoFila]
	
	llenarMatriz2DAleatoriamenteEnteros(matriz2D, tamanoFila, tamanoColumna, LIMITE_INFERIOR, LIMITE_SUPERIOR)
	Escribir "La matriz orginal es:"
	mostrarMatriz2D(matriz2D, tamanoFila, tamanoColumna)
	
	crearMatriz2DTraspuesta(matriz2D, matriz2DTraspuesta, tamanoFila, tamanoColumna)
	Escribir "La matriz traspuesta es:"
	mostrarMatriz2D(matriz2DTraspuesta, tamanoColumna, tamanoFila)
	
FinAlgoritmo

SubProceso crearMatriz2DTraspuesta(matriz, traspuesta, tamanoFila, tamanoColumna)
	Definir fila, columna Como Entero
	
	Para fila = 0 Hasta tamanoFila - 1 Con Paso 1 Hacer		
		Para columna = 0 Hasta tamanoColumna - 1 Con Paso 1 Hacer
			traspuesta[columna, fila] = matriz[fila, columna]
		FinPara
	FinPara
FinSubProceso

SubProceso llenarMatriz2DAleatoriamenteEnteros(matriz, tamanoFila, tamanoColumna, limiteInferior, limiteSuperior)
	Definir columna, fila Como Entero
	
	Para fila = 0 Hasta tamanoFila - 1 Con Paso 1 Hacer		
		Para columna = 0 Hasta tamanoColumna - 1 Con Paso 1 Hacer
			matriz[fila, columna] = Aleatorio(limiteInferior, limiteSuperior)
		FinPara
	FinPara
FinSubProceso

SubProceso mostrarMatriz2D(matriz, tamanoFila, tamanoColumna)
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
