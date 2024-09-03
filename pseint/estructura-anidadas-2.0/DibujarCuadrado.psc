// Desarrollar un programa que lea un n�mero entero (que representa el tama�o del lado)
// y genere un cuadrado de asteriscos con ese tama�o.
// Los asteriscos s�lo se ubicar�n en el borde del cuadrado, no en su interior.
// Por ejemplo, si se ingresa el n�mero 4, se mostrar�:

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
		Escribir "Ingrese el tama�o del lado del cuadrado"
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
