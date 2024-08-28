// Pide al usuario que ingrese dos números enteros y determina si ambos son pares o impares.
// Si ambos números son pares, el programa imprimirá "Ambos números son pares"; de lo contrario, imprimirá "Los números no son pares, o uno de ellos no es par".

Algoritmo DoblePar
	Definir numero1, numero2 Como Entero
	
	Escribir "Ingrese el primer numero"
	Leer numero1
	
	Escribir "Ingrese el segundo numero"
	Leer numero2
	
	Si numero1 mod 2 == 0 Y numero2 mod 2 == 0 Entonces
		Escribir "Ambos números son pares"
	SiNo
		Escribir "Los números no son pares, o uno de ellos no es par"
	FinSi
FinAlgoritmo
