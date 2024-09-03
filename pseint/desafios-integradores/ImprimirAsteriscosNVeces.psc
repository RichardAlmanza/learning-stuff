// Realiza un programa que lea cinco n�meros enteros,
// cada uno comprendido entre 1 y 20,
// e imprima cada n�mero seguido de una
// cantidad de asteriscos equivalente a su valor.
// Por ejemplo:

// Para el n�mero 5, imprimir: 5 *****
// Para el n�mero 3, imprimir: 3 ***
// Para el n�mero 11, imprimir: 11 ***********
// Para el n�mero 2, imprimir: 2 **
// Para el n�mero 9, imprimir: 9 *********

Algoritmo ImprimirAsteriscosNVeces
	Definir LIMITE_SUPERIOR, LIMITE_INFERIOR, CANTIDAD_NUMEROS Como Entero
	Definir CARACTER_IMPRIMIR Como Caracter
	
	Definir numeroUsuario Como Entero
	Definir iteracionNumero, iteracionCaracter Como Entero
	
	LIMITE_SUPERIOR = 20
	LIMITE_INFERIOR = 1
	CANTIDAD_NUMEROS = 5
	
	CARACTER_IMPRIMIR = "*"
	
	Para iteracionNumero <- 1 Hasta CANTIDAD_NUMEROS Hacer
		Repetir
			Escribir "Ingrese un numero dentro del rango [ ", LIMITE_INFERIOR, ", ", LIMITE_SUPERIOR, " ]"
			Leer numeroUsuario
		Hasta Que LIMITE_INFERIOR <= numeroUsuario Y numeroUsuario <= LIMITE_SUPERIOR
		
		Imprimir Sin Saltar numeroUsuario, " "
		
		Para iteracionCaracter <- 1 Hasta numeroUsuario Hacer
			Imprimir Sin Saltar CARACTER_IMPRIMIR
		FinPara
		
		Escribir ""
	FinPara
FinAlgoritmo
