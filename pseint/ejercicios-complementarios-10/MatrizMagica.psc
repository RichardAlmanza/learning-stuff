// Un ejemplo de una matriz mágica es una matriz cuadrada,
// donde el número de filas es igual al número de columnas,
// y que tiene la propiedad especial de que la suma de sus filas,
// columnas y diagonales es igual. Por ejemplo:

// 2 7 6
// 9 5 1
// 4 3 8

// En el ejemplo dado, las sumas siempre dan como resultado 15.
// Se plantea el problema de construir un algoritmo que verifique
// si una matriz de datos enteros es mágica o no.
// En caso de serlo, el programa debe escribir la suma total.
// Además, el programa debe validar que los números introducidos
// estén en el rango de 1 a 9. El usuario define el tamaño de la matriz,
// que no debe exceder un orden de 10.


Algoritmo MatrizMagica
	Definir matriz2D, tamanoLado Como Entero
	
	Repetir
		Escribir "Ingrese el tamano del lado de la mariz cuadrada"
		Leer tamanoLado
	Hasta Que tamanoLado > 0
	
	Dimension matriz2D[tamanoLado, tamanoLado]
	
	llenarMatriz2DManualmente(matriz2D, tamanoLado, tamanoLado)
	mostrarMatriz2D(matriz2D, tamanoLado, tamanoLado)
	
	Si esMatrizMagica(matriz2D, tamanoLado) Entonces
		Escribir "La matriz es magica y la suma de la diagonal es ", sumaDiagonalMatriz(matriz2D, tamanoLado, Verdadero)
	SiNo
		Escribir "El numero no es magico"
	FinSi
FinAlgoritmo

SubProceso llenarMatriz2DManualmente(matriz, tamanoColumna, tamanoFila)
	Definir LIMITE_INFERIOR, LIMITE_SUPERIOR Como Entero
	Definir columna, fila Como Entero
	
	LIMITE_INFERIOR = 1
	LIMITE_SUPERIOR = 9
	
	Para fila = 0 Hasta tamanoFila - 1 Con Paso 1 Hacer		
		Para columna = 0 Hasta tamanoColumna - 1 Con Paso 1 Hacer
			Repetir
				Escribir "Ingrese un numero para un elemento de la matriz (el valor debe estar en el rango [", LIMITE_INFERIOR, ", ", LIMITE_SUPERIOR, "])"
				Escribir "Posicion [", fila + 1, ", ", columna + 1, "](Fila, columna) de ", tamanoFila, "x", tamanoColumna
				Leer matriz[fila, columna]
			Hasta Que LIMITE_INFERIOR <= matriz[fila, columna] Y matriz[fila, columna] <= LIMITE_SUPERIOR			
		FinPara
	FinPara
FinSubProceso

Funcion esMagica <- esMatrizMagica(matriz, tamanoLado)
	Definir columna, fila Como Entero
	Definir sumaFila, sumaColumna, sumaDiagonal Como Real
	Definir esMagica Como Logico
	
	sumaDiagonal = sumaDiagonalMatriz(matriz, tamanoLado, Verdadero)
	esMagica = sumaDiagonal == sumaDiagonalMatriz(matriz, tamanoLado, Falso)
	
	Para fila = 0 Hasta tamanoLado - 1 Con Paso 1 Hacer
		columna = fila
		
		Si esMagica Entonces
			sumaFila = sumaRegionEnMatriz2DNumeros(matriz, fila, 0, fila, tamanoLado - 1)
			sumaColumna = sumaRegionEnMatriz2DNumeros(matriz, 0, columna, tamanoLado - 1, columna)
			
			esMagica = sumaFila == sumaColumna Y sumaFila == sumaDiagonal
		FinSi
	FinPara
FinFuncion

Funcion suma <- sumaDiagonalMatriz(matriz, tamanoLado, diagonalPrincipal)
	Definir fila, columna Como Entero
	Definir suma Como Real
	
	suma = 0
	
	Para fila = 0 Hasta tamanoLado - 1 Con Paso 1 Hacer
		Si diagonalPrincipal Entonces
			columna = fila
		SiNo
			columna = tamanoLado - fila - 1
		FinSi
		
		suma = suma + matriz[fila, columna]
	FinPara
FinFuncion

Funcion suma <- sumaRegionEnMatriz2DNumeros(matriz, filaInicial, columnaInicial, filaFinal, columnaFinal)
	Definir columna, fila, pasoFila, pasoColumna Como Entero
	Definir suma Como Real
	
	suma = 0
	pasoFila = 1
	pasoColumna = 1
	
	Si filaInicial > filaFinal Entonces
		pasoFila = -1
	FinSi
	
	Si columnaInicial > columnaFinal Entonces
		pasoColumna = -1
	FinSi
	
	Para fila = filaInicial Hasta filaFinal Con Paso pasoFila Hacer		
		Para columna = columnaInicial Hasta columnaFinal Con Paso pasoColumna Hacer
			suma = suma + matriz[fila, columna]
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
