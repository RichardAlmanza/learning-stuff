// Crea una matriz con 3 columnas y una cantidad de filas definida por el usuario.
// En las dos primeras columnas, el usuario ingresará valores enteros
// (puede diseñar este ingreso de manera aleatoria para enviar la carga manual).
// En la tercera columna se almacenará el resultado de sumar los números de la
// primera y segunda columna.
// La matriz se mostrará de la siguiente forma:

// 3 | 5 | 8 ? 8 se obtuvo de sumar 3 + 5 

// 4 | 3 | 7 ? 7 se obtuvo de sumar 4 + 3 

// 1 | 4 | 5 ? 5 se obtuvo de sumar 1 +4 

Algoritmo ColumnaResultado
	Definir LIMITE_INFERIOR, LIMITE_SUPERIOR, TAMANO_COLUMNA Como Entero
	Definir matriz2D, tamanoFila, tamanoColumna Como Entero
	
	LIMITE_INFERIOR = 1
	LIMITE_SUPERIOR = 10
	TAMANO_COLUMNA = 3
	
	Repetir
		Escribir "Ingrese el # de filas para la matriz"
		Leer tamanoFila
	Hasta Que tamanoFila > 0 
	
	Dimension matriz2D[tamanoFila, TAMANO_COLUMNA]
	
	llenarMatriz2DAleatoriamenteEnteros(matriz2D, tamanoFila, TAMANO_COLUMNA - 1, LIMITE_INFERIOR, LIMITE_SUPERIOR)
	mostrarMatriz2D(matriz2D, tamanoFila, TAMANO_COLUMNA - 1)
	
	Escribir "El resultado de la suma de los elementos es:"
	sumarElementosDeFilaMatriz2D(matriz2D, tamanoFila, TAMANO_COLUMNA)
	mostrarMatriz2D(matriz2D, tamanoFila, TAMANO_COLUMNA)
FinAlgoritmo

SubProceso sumarElementosDeFilaMatriz2D(matriz, tamanoFila, tamanoColumna)
	Definir fila, columna Como Entero
	
	Para fila = 0 Hasta tamanoFila - 1 Con Paso 1 Hacer		
		matriz[fila, tamanoColumna - 1] = 0
		
		Para columna = 0 Hasta tamanoColumna - 2 Con Paso 1 Hacer
			matriz[fila, tamanoColumna - 1] = matriz[fila, tamanoColumna - 1] + matriz[fila, columna]
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
