// Crea un programa que solicite al usuario un n�mero entero y determine si es par o impar.
// Luego, mostrar en pantalla un mensaje indicando si el n�mero es par o impar.
// Para determinar si un n�mero es par, se debe dividir entre dos y verificar que el resto sea igual a 0.
// Se recomienda investigar la funci�n mod de PSeInt para lograr esta verificaci�n de manera eficiente. Recuerda nombrar y guardar tu algoritmo.

Algoritmo ParImpart
	Definir numero_ Como Entero
	
	Escribir "Ingrese el numero a consultar"
	Leer numero_
	
	Si numero_ mod 2 == 0 Entonces
		Escribir "El numero es Par"
	SiNo
		Escribir "El numero es Impar"
	FinSi
FinAlgoritmo
