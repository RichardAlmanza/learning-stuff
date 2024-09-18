Algoritmo AlinearPalabrasCaminoFelizExtraFunciones
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
		esperarTeclaLimpiar()
		
		agregarPalabra(matriz, fila, palabras[fila])
		imprimirMatriz(matriz, FILAS, COLUMNAS)
	FinPara
	
	acomodarPalabra(matriz, FILAS, COLUMNAS)
	
FinAlgoritmo


SubProceso acomodarPalabra(matriz, tamanoFila, tamanoColumna)
	Definir fila, iter, maxPosicion, posicionR Como Entero
	
	maxPosicion = maxPosicionRDeLasFilas(matriz, tamanoFila)
	
	Para fila = 0 Hasta tamanoFila - 1 Con Paso 1 Hacer
		posicionR = buscarR(matriz, fila)
		
		Si posicionR < maxPosicion Entonces	
			Para iter = 1 Hasta maxPosicion - posicionR Con Paso 1 Hacer
				rotarColumnas(matriz, tamanoColumna, fila)
				
				Limpiar Pantalla
				imprimirMatriz(matriz, tamanoFila, tamanoColumna)
				Esperar 300 Milisegundos
			FinPara
		FinSi
		
		esperarTeclaLimpiar()
	FinPara
FinSubProceso


SubProceso esperarTeclaLimpiar
	Escribir "Presione cualquier tecla para continuar"
	Esperar Tecla
	Limpiar Pantalla
FinSubProceso


Funcion maxPosicionR <- maxPosicionRDeLasFilas(matriz, tamanoFila)
	Definir fila, posicionR, maxPosicionR Como Entero
	
	maxPosicionR = buscarR(matriz, 0)
	
	Para fila = 0 Hasta tamanoFila - 1 Con Paso 1 Hacer
		posicionR = buscarR(matriz, fila)
		
		Si posicionR > maxPosicionR Entonces
			maxPosicionR = posicionR
		FinSi
	FinPara
FinFuncion


/// Rota una sola columna
SubProceso rotarColumnas(matriz, tamanoColumna, fila)
	// 12345(original) 51234 45123 34512 23451(4 llamados)
	Definir columna, iter Como Entero
	Definir letraInsertar, letraAuxiliar Como Caracter
	
	letraInsertar = matriz[fila, tamanoColumna - 1]
	
	Para columna = 0 Hasta tamanoColumna - 1 Con Paso 1 Hacer
		letraAuxiliar = letraInsertar
		letraInsertar = matriz[fila, columna]
		matriz[fila, columna] = letraAuxiliar
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
	
	RELLENO = " " // Cambie el asterisco por un espacio, se ve mas limpia la matriz
	
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
				Escribir Sin Saltar " | " // Pueden cambiarlo por un solo espacio " ", para que se vea como la version AlinearPalabrasCaminoFeliz.psc
			FinSi
		FinPara
		
		Escribir " |"
	FinPara
	
	Escribir ""
FinSubProceso
