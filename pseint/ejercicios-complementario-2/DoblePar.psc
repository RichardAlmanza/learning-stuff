// Pide al usuario que ingrese dos n�meros enteros y determina si ambos son pares o impares.
// Si ambos n�meros son pares, el programa imprimir� "Ambos n�meros son pares"; de lo contrario, imprimir� "Los n�meros no son pares, o uno de ellos no es par".

Algoritmo DoblePar
	Definir numero1, numero2 Como Entero
	
	Escribir "Ingrese el primer numero"
	Leer numero1
	
	Escribir "Ingrese el segundo numero"
	Leer numero2
	
	Si numero1 mod 2 == 0 Y numero2 mod 2 == 0 Entonces
		Escribir "Ambos n�meros son pares"
	SiNo
		Escribir "Los n�meros no son pares, o uno de ellos no es par"
	FinSi
FinAlgoritmo
