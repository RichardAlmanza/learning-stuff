// Crea una funci�n esImpar que determine si un n�mero es impar.
// Si es impar, la funci�n debe devolver True; en caso contrario, debe devolver False.
// Nota: la funci�n no debe incluir mensajes que indiquen si el n�mero es par o impar;
//        esto debe manejarse en el programa principal
// Recuerda nombrar y guardar tu algoritmo.

Algoritmo Impares
	Definir iter, numeroUsuario Como Entero
	
	Para iter <- 1 Hasta 5 Hacer
		Escribir "Ingrese un numero entero que desea saber si es impar"
		Leer numeroUsuario
		
		Escribir "El resultado es ", esImpar(numeroUsuario)
	FinPara
FinAlgoritmo

Funcion resultado <- esImpar (numero_)
	Definir resultado Como Logico
	
	resultado = numero_ mod 2 <> 0
FinFuncion
	