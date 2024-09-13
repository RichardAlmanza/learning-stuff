// Crea un programa que cumpla con las siguientes condiciones:
// - Define y dimensiona una matriz de 5x5 para almacenar datos numéricos enteros.

// - Rellena la matriz de manera aleatoria con números comprendidos entre 10 y 40.

// - Permite al usuario ingresar un número para buscarlo dentro de la matriz.

// - Si el número se encuentra, muestra en pantalla un mensaje adecuado junto
//   con las coordenadas en la matriz (fila y columna).
//   En caso de que el número esté repetido,
//   solo se mostrará la posición de la primera ocurrencia.

// - Si el número no se encuentra, informa por pantalla.

Algoritmo BuscarNumeroEnMatriz2D
	Definir TAMANO_LADO, LIMITE_INFERIOR, LIMITE_SUPERIOR Como Entero
	Definir matriz2DNumeros, numeroABuscar, matrizPosiciones, tamanoPosiciones Como Entero
	
	TAMANO_LADO = 5
	LIMITE_INFERIOR = 10
	LIMITE_SUPERIOR = 40
	
	Dimension matriz2DNumeros[TAMANO_LADO, TAMANO_LADO], matrizPosiciones[TAMANO_LADO * TAMANO_LADO, 2]
	
	llenarMatriz2DAleatoriamenteEnteros(matriz2DNumeros, TAMANO_LADO, TAMANO_LADO, LIMITE_INFERIOR, LIMITE_SUPERIOR)
	
	Escribir "La matriz generada es"
	mostrarMatriz2D(matriz2DNumeros, TAMANO_LADO, TAMANO_LADO)
	
	Escribir "Ingrese el numero entero que desea buscar en la matriz"
	Leer numeroABuscar
	
	Si buscarEnMatriz2D(numeroABuscar, matriz2DNumeros, matrizPosiciones, TAMANO_LADO, TAMANO_LADO, tamanoPosiciones) Entonces
		Escribir "Se encontro el numero en las siguientes posiciones"
		mostrarMatriz2D(matrizPosiciones, tamanoPosiciones, 2)
	SiNo
		Escribir "El numero no se encontro"
	FinSi
FinAlgoritmo

Funcion encontrado <- buscarEnMatriz2D(elementoBuscado, matrizABuscar, matrizPosiciones, tamanoFila, tamanoColumna, tamanoPosiciones Por Referencia)
	Definir fila, columna Como Entero
	Definir indice Como Entero
	Definir encontrado Como Logico
	
	encontrado = Falso
	tamanoPosiciones = 0
	
	Para fila = 0 Hasta tamanoFila - 1 Con Paso 1 Hacer		
		Para columna = 0 Hasta tamanoColumna - 1 Con Paso 1 Hacer
			Si matrizABuscar[fila, columna] == elementoBuscado Entonces
				matrizPosiciones[tamanoPosiciones, 0] = fila + 1
				matrizPosiciones[tamanoPosiciones, 1] = columna + 1
				tamanoPosiciones = tamanoPosiciones + 1
			FinSi
		FinPara
	FinPara
	
	Si tamanoPosiciones > 0 Entonces
		encontrado = Verdadero
	FinSi
FinFuncion

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

