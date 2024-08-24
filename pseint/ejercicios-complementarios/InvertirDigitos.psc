// Diseña un algoritmo que permita obtener el número invertido de un número de dos cifras.
// Ejemplo, si se introduce 23 que muestre 32.

Algoritmo InvertirDigitos
	Definir num, numInvertido Como Entero
	numInvertido = 0
	
	Escribir "Ingrese un numero entero para ser invertido"
	Leer num
	
	num = abs(num)
	
	Mientras num > 0
		numInvertido = 10 * numInvertido
		numInvertido = numInvertido + num mod 10
		num = trunc(num / 10)
	FinMientras
	
	Escribir "El numero invertido es ", numInvertido
FinAlgoritmo
