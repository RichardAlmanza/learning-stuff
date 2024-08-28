// Desarrolla un programa que solicite al usuario ingresar un nombre para su competencia, el cual puede ser una frase o palabra.
// Si la frase o palabra tiene una longitud de 4 caracteres, el programa concatenará un signo de exclamación al final de la cadena; 
// de lo contrario, añadirá un signo de interrogación. Posteriormente, el programa mostrará la frase final resultante.
// Se recomienda investigar las funciones Longitud() y Concatenar() de PSeInt para realizar esta tarea de manera eficiente

Algoritmo NombreDe4Letras
	Definir alias, resultado Como Caracter
	
	Escribir "Ingrese un nombre para la competencia"
	Leer alias
	
	Si Longitud(alias) == 4 Entonces
		resultado = Concatenar(alias, "!")
	SiNo
		resultado = Concatenar(alias, "?")
	FinSi
	
	Escribir "Su nombre sera ", resultado
	
FinAlgoritmo
