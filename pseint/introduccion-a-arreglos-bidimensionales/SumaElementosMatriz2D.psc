// Dada una matriz de tamaño n x m, donde n y m
// son valores ingresados por el usuario,
// se requiere implementar dos subprogramas:

// - El primer subprograma se encargará de llenar
//   la matriz con números aleatorios.

// - El segundo subprograma calculará y mostrará
//   la suma de todos los elementos de la matriz.

// Después de ejecutar ambos subprogramas,
// se mostrará la matriz generada junto con
// los resultados de la suma por pantalla.

Algoritmo SumaElementosMatriz2D
	Definir LIMITE_INFERIOR, LIMITE_SUPERIOR Como Entero
	Definir matriz2DNumeros, tamanoFila, tamanoColumna Como Entero
	
	LIMITE_INFERIOR = 0
	LIMITE_SUPERIOR = 999
	
	pedirTamanoMatriz2D(tamanoFila, tamanoColumna)
	
	Dimension matriz2DNumeros[tamanoFila, tamanoColumna]
	
	llenarMatriz2DAleatoriamenteEnteros(matriz2DNumeros, tamanoFila, tamanoColumna, LIMITE_INFERIOR, LIMITE_SUPERIOR)
	
	Escribir "La matriz generada es"
	mostrarMatriz2D(matriz2DNumeros, tamanoFila, tamanoColumna)
	
	Escribir "La suma de todos los elementos es: ", sumaElmentosMatriz2DNumeros(matriz2DNumeros, tamanoFila, tamanoColumna)	
FinAlgoritmo

SubProceso pedirTamanoMatriz2D(tamanoFila Por Referencia, tamanoColumna Por Referencia)
	Repetir
		Escribir "Ingrese cuantas Filas quiere para la matriz (solo numero naturales)"
		Leer tamanoFila
	Hasta Que tamanoFila > 0
	
	Repetir
		Escribir "Ingrese cuantas Columnas quiere para la matriz (solo numero naturales)"
		Leer tamanoColumna
	Hasta Que tamanoColumna > 0
FinSubProceso

Funcion suma <- sumaElmentosMatriz2DNumeros(matriz, tamanoFila, tamanoColumna)
	Definir columna, fila Como Entero
	Definir suma Como Real
	
	suma = 0
	
	Para fila = 0 Hasta tamanoFila - 1 Con Paso 1 Hacer		
		Para columna = 0 Hasta tamanoColumna - 1 Con Paso 1 Hacer
			suma = suma + matriz[fila, columna]
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
