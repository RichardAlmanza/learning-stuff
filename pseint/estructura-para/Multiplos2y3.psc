// Dise�a un programa que cuente la cantidad de
// n�meros entre 1 y 100 que son m�ltiplos de 2 o de 3.
// Recuerda nombrar y guardar tu algoritmo.
// Nota: Si un n�mero cumple ambas condiciones,
// debe ser contabilizado en ambos casos.
// Por ejemplo, el n�mero 18 es m�ltiplo de 2 y,
// a su vez, es m�ltiplo de 3.

Algoritmo Multiplos2y3
	Definir NUMERO_INICIAL, NUMERO_FINAL Como Entero
	Definir numeroIteracion, contadorMultiplos Como Entero
	
	NUMERO_INICIAL = 1
	NUMERO_FINAL = 100
	
	contadorMultiplos = 0
	
	Para numeroIteracion = NUMERO_INICIAL Hasta NUMERO_FINAL Hacer
		Si numeroIteracion mod 2 == 0 Entonces
			contadorMultiplos =  contadorMultiplos + 1
		FinSi
		
		Si numeroIteracion mod 3 == 0 Entonces
			contadorMultiplos =  contadorMultiplos + 1
		FinSi
	FinPara
	
	Escribir contadorMultiplos
FinAlgoritmo
