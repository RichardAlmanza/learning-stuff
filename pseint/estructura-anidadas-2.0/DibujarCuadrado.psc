// Desarrollar un programa que lea un número entero (que representa el tamaño del lado)
// y genere un cuadrado de asteriscos con ese tamaño.
// Los asteriscos sólo se ubicarán en el borde del cuadrado, no en su interior.
// Por ejemplo, si se ingresa el número 4, se mostrará:

// * * * *
// *     *
// *     *
// * * * *

// Nota: Recordar el uso del escribir sin saltar en PseInt.

Algoritmo DibujarCuadrado
	Definir CARACTER_DIBUJAR, CARACTER_ESPACIO Como Caracter
	Definir largoLado, fila, columna Como Entero
	Definir esBorde Como Logico
	
	CARACTER_DIBUJAR = "*"
	CARACTER_ESPACIO = " "
	
	Repetir
		Escribir "Ingrese el tamaño del lado del cuadrado"
		Leer largoLado
	Hasta Que largoLado > 0
	
	Para fila <- 1 Hasta largoLado Hacer
		Para columna <- 1 Hasta largoLado Hacer
			esBorde = columna == 1 O columna == largoLado
			esBorde = esBorde O fila == 1 O fila == largoLado
			
			Si esBorde Entonces
				Escribir Sin Saltar CARACTER_DIBUJAR
			SiNo
				Escribir Sin Saltar CARACTER_ESPACIO
			FinSi
			
		FinPara
		
		Escribir ""
	FinPara
	
FinAlgoritmo
