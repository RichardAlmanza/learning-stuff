// Escribe un programa que lea un número entero y
// devuelva la cantidad de dígitos que lo componen.
// Por ejemplo, si ingresamos el número 12345, el programa deberá devolver 5.
// Este cálculo se realizará utilizando operaciones matemáticas,
// teniendo en cuenta que las variables de tipo entero truncan los números o resultados.

Algoritmo ContarDigitos
	Definir numeroUsuario, numeroDigitos, numero_ Como Entero
	
	numeroDigitos = 0
	
	Escribir "Ingrese un numero que desea contar los digitos"
	Leer numeroUsuario
	
	numero_ = abs(numeroUsuario)
	
	Mientras numero_ > 0 Hacer
		numero_ = trunc(numero_ / 10)
		numeroDigitos = numeroDigitos + 1
	FinMientras
	
	Si numeroUsuario == 0 Entonces
		numeroDigitos = 1
	FinSi
	
	Escribir "Para el numero ", numeroUsuario, " se representa con ", numeroDigitos, " digito/s"
FinAlgoritmo
