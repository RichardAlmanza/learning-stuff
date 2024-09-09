// Dise�a un algoritmo que lea un n�mero de tres cifras correspondiente a la posici�n t�ctica del equipo (433, 343, 253, etc.)
// y determine si es o no capic�a. Un n�mero es capic�a cuando se lee igual de izquierda a derecha que de derecha a izquierda.
// Recuerda nombrar y guardar tu algoritmo.

Algoritmo Capicua
	Definir numero_, digitoInicial, digitoFinal Como Entero
	
	Escribir "Ingrese el numero de 3 digitos"
	Leer numero_
	
	digitoInicial = trunc(numero_ / 100)
	digitoFinal = numero_ mod 10
	
	SI digitoInicial == digitoFinal Entonces
		Escribir "El numero es capicua"
	SiNo
		Escribir "El numero NO es capicua"
	FinSi
	
FinAlgoritmo
