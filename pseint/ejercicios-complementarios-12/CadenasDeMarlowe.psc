Algoritmo CadenasDeMarlowe
	Definir FILTRO Como Caracter
	Definir cadenaADN, matrizADN Como Caracter
	Definir tamanoLado Como Entero
	
	FILTRO = "ABCD"
	
	Repetir
		Escribir "Ingrese la cadena de ADN valida"
		Leer cadenaADN
		
		cadenaADN = Mayusculas(cadenaADN)
	Hasta Que esCadenaADNValida(cadenaADN)
	
	tamanoLado = raiz(Longitud(cadenaADN))
	
	Dimension matrizADN[tamanoLado, tamanoLado]
	
	cadenaADNAMatriz(matrizADN, tamanoLado, cadenaADN)
	imprimirMatriz(matrizADN, tamanoLado, tamanoLado)
	
	Si diagonalMatrizValida(matrizADN, tamanoLado, Verdadero) y diagonalMatrizValida(matrizADN, tamanoLado, Falso) Entonces
		Escribir "¡Felicidades! La muestra contiene patrones específicos en sus diagonales"
	SiNo
		Escribir "La muestra NO contiene los patrones especificados en sus diagonales"
	FinSi
FinAlgoritmo


Funcion esValida <- diagonalMatrizValida(matriz, tamanoLado, diagonalPrincipal)
	Definir fila, columna Como Entero
	Definir anterior Como Caracter
	Definir esValida Como Logico
	
	esValida = Verdadero
	
	Si diagonalPrincipal Entonces
		anterior = matriz[0, 0]
	SiNo
		anterior = matriz[0, tamanoLado - 1]
	FinSi
	
	Para fila = 0 Hasta tamanoLado - 1 Con Paso 1 Hacer
		Si diagonalPrincipal Entonces
			columna = fila
		SiNo
			columna = tamanoLado - fila - 1
		FinSi
		
		esValida = esValida Y anterior == matriz[fila, columna]
	FinPara
FinFuncion


SubProceso cadenaADNAMatriz(matriz, tamanoLado, cadenaADN)
	Definir columna, fila, indice Como Entero
	
	Para fila = 0 Hasta tamanoLado - 1 Con Paso 1 Hacer		
		Para columna = 0 Hasta tamanoLado - 1 Con Paso 1 Hacer
			indice = fila * tamanoLado + columna
			
			matriz[fila, columna] = Subcadena(cadenaADN, indice, indice)
		FinPara
	FinPara
FinSubProceso


Funcion valida <- esCadenaADNValida(cadenaADN)
	Definir FILTRO Como Caracter
	Definir valida Como Logico
	Definir raizLargo Como Real
	
	FILTRO = "ABCD"
	
	raizLargo = raiz(longitud(cadenaADN))
	
	valida = raizLargo == trunc(raizLargo)
	valida = valida Y raizLargo > 2
	valida = valida Y filtrarTexto(cadenaADN, FILTRO, Verdadero) == cadenaADN
FinFuncion


Funcion textoFiltrado <- filtrarTexto(textoAFiltrar, cadenaFiltro, esIncluyente)
	Definir indice Como Entero
	Definir caracterTexto, textoFiltrado Como Caracter
	Definir esValido Como Logico
	
	textoFiltrado = ""
	
	Para indice = 0 Hasta Longitud(textoAFiltrar) - 1 Con Paso 1 Hacer
		caracterTexto = Subcadena(textoAFiltrar, indice, indice)
		esValido = caracterEnCadena(caracterTexto, cadenaFiltro)
		
		Si NO esIncluyente Entonces
			esValido = !esValido
		FinSi
		
		Si esValido Entonces
			textoFiltrado = Concatenar(textoFiltrado, caracterTexto)
		FinSi
	FinPara
FinSubProceso


Funcion encontrado <- caracterEnCadena(caracterABuscar, opciones)
	Definir indice Como Entero
	Definir encontrado Como Logico
	
	encontrado = Falso
	indice = 0
	
	Mientras indice < Longitud(opciones) Y NO encontrado Hacer
		encontrado = Subcadena(opciones, indice, indice) == caracterABuscar
		
		indice = indice + 1
	FinMientras
FinFuncion


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
