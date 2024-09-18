Algoritmo AlinearPalabrasCaminoFeliz
	Definir FILAS, COLUMNAS Como Entero
	Definir fila Como Entero
	Definir matriz, palabras Como Caracter
	
	FILAS = 9
	COLUMNAS = 12
	
	Dimension matriz[FILAS, COLUMNAS]
	Dimension palabras[FILAS]
	
	palabras[0] = "VECTOR"
	palabras[1] = "MATRIX"
	palabras[2] = "PROGRAMA"
	palabras[3] = "SUBPROGRAMA"
	palabras[4] = "SUBPROCESO"
	palabras[5] = "VARIABLE"
	palabras[6] = "ENTERO"
	palabras[7] = "PARA"
	palabras[8] = "MIENTRAS"
	
	inicializarMatriz(matriz, FILAS, COLUMNAS)
	imprimirMatriz(matriz, FILAS, COLUMNAS)
	
	Para fila = 0 Hasta FILAS - 1 Con Paso 1 Hacer
		Escribir "Presione cualquier tecla para continuar"
		Esperar Tecla
		Limpiar Pantalla
		
		agregarPalabra(matriz, fila, palabras[fila])
		imprimirMatriz(matriz, FILAS, COLUMNAS)
	FinPara
	
	acomodarPalabra(matriz, FILAS, COLUMNAS)
	
FinAlgoritmo


SubProceso acomodarPalabra(matriz, tamanoFila, tamanoColumna)
	Definir fila, columna, iter, maxPosicion, posicionR Como Entero
	Definir letraAuxiliar, letraInsertar Como Caracter
	
	maxPosicion = buscarR(matriz, 0)
	
	Para fila = 0 Hasta tamanoFila - 1 Con Paso 1 Hacer
		posicionR = buscarR(matriz, fila)
		
		Si posicionR > maxPosicion Entonces
			maxPosicion = posicionR
		FinSi
	FinPara
	
	Para fila = 0 Hasta tamanoFila - 1 Con Paso 1 Hacer
		posicionR = buscarR(matriz, fila)
		
		Si posicionR < maxPosicion Entonces	
			Para iter = 1 Hasta maxPosicion - posicionR Con Paso 1 Hacer
				letraInsertar = matriz[fila, tamanoColumna - 1] // 12345 51234 45123 34512 23451
				
				Para columna = 0 Hasta tamanoColumna - 1 Con Paso 1 Hacer
					letraAuxiliar = letraInsertar
					letraInsertar = matriz[fila, columna]
					matriz[fila, columna] = letraAuxiliar
				FinPara
				
				Limpiar Pantalla
				imprimirMatriz(matriz, tamanoFila, tamanoColumna)
				Esperar 500 Milisegundos
			FinPara
		FinSi
		
		Escribir "Presione cualquier tecla para continuar"
		Esperar Tecla
		Limpiar Pantalla
	FinPara
FinSubProceso


Funcion posicionR <- buscarR(matriz, fila)
	Definir posicionR Como Entero
	Definir buscando Como Logico
	
	buscando = Verdadero
	posicionR = 0
	
	Mientras buscando Hacer
		Si matriz[fila, posicionR] == "R" Entonces
			buscando = Falso
		FinSi
		
		posicionR = posicionR + 1
	FinMientras
FinFuncion


SubProceso agregarPalabra(matriz, fila, palabra)
	Definir columna Como Entero
	
	Para columna = 0 Hasta Longitud(palabra) - 1 Con Paso 1 Hacer
		matriz[fila, columna] = Subcadena(palabra, columna, columna)
	FinPara
FinSubProceso


SubProceso inicializarMatriz(matriz, tamanoFila, tamanoColumna)
	Definir RELLENO Como Caracter
	Definir fila, columna Como Entero
	
	RELLENO = "*"
	
	Para fila = 0 Hasta tamanoFila - 1 Con Paso 1 Hacer		
		Para columna = 0 Hasta tamanoColumna - 1 Con Paso 1 Hacer
			matriz[fila, columna] = RELLENO
		FinPara
	FinPara
FinSubProceso


SubProceso imprimirMatriz(matriz, tamanoFila, tamanoColumna)
	Definir fila, columna Como Entero
	
	Para fila = 0 Hasta tamanoFila - 1 Con Paso 1 Hacer
		Escribir Sin Saltar "| "
		
		Para columna = 0 Hasta tamanoColumna - 1 Con Paso 1 Hacer
			Escribir Sin Saltar matriz[fila, columna]
			
			Si columna < tamanoColumna - 1 Entonces
				Escribir Sin Saltar " "
			FinSi
		FinPara
		
		Escribir " |"
	FinPara
	
	Escribir ""
FinSubProceso
