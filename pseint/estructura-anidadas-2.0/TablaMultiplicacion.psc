// Elabora un algoritmo que genere e imprima
// las tablas de multiplicar del 1 al 10.
// Se espera que en la salida por pantalla
// se presente cada tabla de multiplicar de la siguiente manera:

// 1 x 1 = 1
// 1 x 2 = 2
// 1 x 3 = 3
// ...
// 1 x 10 = 10
// -----------
// 2 x 1 = 2
// 2 x 2 = 4
// 2 x 3 = 6
// ...
// 2 x 10 = 20
// -----------
// ...

Algoritmo TablaMultiplicacion
	Definir FACTOR_INICIAL, FACTOR_FINAL Como Entero
	Definir factor1, factor2 Como Entero
	
	FACTOR_INICIAL = 1
	FACTOR_FINAL = 10
	
	Para factor1 <- FACTOR_INICIAL Hasta FACTOR_FINAL Hacer
		Para factor2 <- FACTOR_INICIAL Hasta FACTOR_FINAL Hacer
			Escribir factor1, " x ", factor2, " = ", factor1 * factor2
		FinPara
		
		Escribir "----------------------------"
	FinPara
FinAlgoritmo
