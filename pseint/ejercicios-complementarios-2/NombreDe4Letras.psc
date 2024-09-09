// Desarrolla un programa que solicite al usuario ingresar un nombre para su competencia, el cual puede ser una frase o palabra.
// Si la frase o palabra tiene una longitud de 4 caracteres, el programa concatenar� un signo de exclamaci�n al final de la cadena; 
// de lo contrario, a�adir� un signo de interrogaci�n. Posteriormente, el programa mostrar� la frase final resultante.
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
