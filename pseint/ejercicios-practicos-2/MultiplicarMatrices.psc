// Desarrolla un programa que calcule la multiplicación de dos matrices de
// enteros de tamaño 3x3.
// Asegúrate de inicializar las matrices para evitar tener que ingresar datos desde el teclado.
// La multiplicación se almacenará en una tercera matriz,
// donde cada elemento será el resultado de multiplicar los elementos
// correspondientes en la misma posición de las matrices A y B.
// Por ejemplo, el elemento en la posición (0,0) de la matriz C será el resultado
// de multiplicar el elemento en la posición (0,0) de la matriz A con el elemento
// en la posición (0,0) de la matriz B.

Algoritmo MultiplicarMatrices
	Definir LIMITE_INFERIOR, LIMITE_SUPERIOR, TAMANO_LADO Como Entero
	Definir matrizA, matrizB, matrizC Como Entero
	
	LIMITE_INFERIOR = -99
	LIMITE_SUPERIOR = 99
	TAMANO_LADO = 3
	
	Dimension matrizA[TAMANO_LADO, TAMANO_LADO], matrizB[TAMANO_LADO, TAMANO_LADO], matrizC[TAMANO_LADO, TAMANO_LADO]
	
	llenarMatriz2DAleatoriamenteEnteros(matrizA, TAMANO_LADO, TAMANO_LADO, LIMITE_INFERIOR, LIMITE_SUPERIOR)
	llenarMatriz2DAleatoriamenteEnteros(matrizB, TAMANO_LADO, TAMANO_LADO, LIMITE_INFERIOR, LIMITE_SUPERIOR)
	
	multiplicarElementoMatrices(matrizA, matrizB, matrizC, TAMANO_LADO)
	
	Escribir "La matriz A es:"
	mostrarMatriz2D(matrizA, TAMANO_LADO, TAMANO_LADO)
	
	Escribir "La matriz B es:"
	mostrarMatriz2D(matrizB, TAMANO_LADO, TAMANO_LADO)
	
	Escribir "La matriz resultante C es:"
	mostrarMatriz2D(matrizC, TAMANO_LADO, TAMANO_LADO)
FinAlgoritmo

SubProceso multiplicarElementoMatrices(matrizA, matrizB, matrizC, tamanoLado)
	Definir columna, fila Como Entero
	
	Para fila = 0 Hasta tamanoLado - 1 Con Paso 1 Hacer		
		Para columna = 0 Hasta tamanoLado - 1 Con Paso 1 Hacer
			matrizC[fila, columna] = matrizA[fila, columna] * matrizB[fila, columna]
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
