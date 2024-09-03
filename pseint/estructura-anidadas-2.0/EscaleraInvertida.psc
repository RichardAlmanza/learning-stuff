// Crear un programa que lea un número entero (que represente la altura)
// y que genere una escalera invertida de asteriscos con esa altura.
// Por ejemplo, si ingresamos una altura de 5, se deberá mostrar:

// *****
// ****
// ***
// **
// *

Algoritmo EscaleraInvertida
	Definir CARACTER_DIBUJAR Como Caracter
	Definir altura, fila, columna, maxColumnas Como Entero
	
	CARACTER_DIBUJAR = "*"
	
	Repetir
		Escribir "Ingrese la altura de la escalera"
		Leer altura
	Hasta Que altura > 0
	
	Para fila <- 1 Hasta altura Hacer
		maxColumnas = altura - (fila - 1)
		
		Para columna <- 1 Hasta maxColumnas Hacer
			Escribir Sin Saltar CARACTER_DIBUJAR
		FinPara
		
		Escribir ""
	FinPara
FinAlgoritmo
