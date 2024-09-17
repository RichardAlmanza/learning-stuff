// Desarrolla un programa que genere una matriz de tamaño 5x15.
// Tu tarea consiste en llenar la matriz con unos y ceros,
// donde los 1 ocupan el borde externo de la matriz y los 0 llenarán el área interior.

// Por ejemplo, el aspecto final de tu matriz deberá ser el siguiente:

//	111111111111111
//	100000000000001
//	100000000000001
//	100000000000001
//	111111111111111

Algoritmo BordesMatriz
	Definir BORDE, RELLENO, TAMANO_FILA, TAMANO_COLUMNA Como Entero
	Definir matriz2D Como Entero
	
	BORDE = 1
	RELLENO = 0
	
	TAMANO_FILA = 5
	TAMANO_COLUMNA = 15
	
	Dimension matriz2D[TAMANO_FILA, TAMANO_COLUMNA]
	
	llenarMatriz2DBordes(matriz2D, TAMANO_FILA, TAMANO_COLUMNA, BORDE, RELLENO)
	mostrarMatriz2D(matriz2D, TAMANO_FILA, TAMANO_COLUMNA)
FinAlgoritmo

SubProceso llenarMatriz2DBordes(matriz, tamanoFila, tamanoColumna, borde, relleno)
	Definir columna, fila Como Entero
	Definir esBorde Como Logico
	
	Para fila = 0 Hasta tamanoFila - 1 Con Paso 1 Hacer		
		Para columna = 0 Hasta tamanoColumna - 1 Con Paso 1 Hacer
			esBorde = columna == 0 O columna == tamanoColumna - 1
			esBorde = esBorde O fila == 0 O fila == tamanoFila - 1
			
			Si esBorde Entonces
				matriz[fila, columna] = borde
			SiNo
				matriz[fila, columna] = relleno
			FinSi
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
