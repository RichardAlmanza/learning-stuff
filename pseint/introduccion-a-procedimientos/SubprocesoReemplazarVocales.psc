// Escribe un programa que procese una secuencia de caracteres ingresada
// por teclado y terminada en punto, y luego codifique la palabra o frase
// ingresada de la siguiente manera: cada vocal se reemplaza por un carácter
// según la tabla:

// a -> @
// e -> #
// i -> $
// o -> %
// u -> *

// Realiza un subprograma que reciba una secuencia de caracteres
// y retorne la codificación correspondiente, utilizando la
// estructura "según" para la transformación.
// Por ejemplo, si el usuario ingresa:
// "Ayer, lunes, salimos a las once y 10.",
// la salida del programa debería ser:
// "@y#r, l*n#s, s@l$m%s @ l@s %nc# y 10."
// Considera repasar el uso de la función concatenar de 
// PSeInt para armar la palabra/frase.
// Recuerda nombrar y guardar tu algoritmo.

Algoritmo SubprocesoReemplazarVocales
	Definir fraseUsuario, nuevaFrase Como Caracter
	
	Escribir "Ingrese una frase que desee reemplazar las vocales"
	Leer fraseUsuario
	
	reemplazarVocales(nuevaFrase, fraseUsuario)
	
	Escribir "La nueva frase es:"
	Escribir nuevaFrase
FinAlgoritmo

SubProceso reemplazarVocales (nuevaFrase Por Referencia, fraseOriginal)
	Definir ultimoIndiceCaracter, indiceCaracter Como Entero
	Definir nuevoCaracter Como Caracter
	
	ultimoIndiceCaracter = Longitud(fraseOriginal) - 1
	
	nuevaFrase = ""
	
	Para indiceCaracter <- 0 Hasta ultimoIndiceCaracter Con Paso 1 Hacer 
		nuevoCaracter = Subcadena(fraseOriginal, indiceCaracter, indiceCaracter)
		
		Segun Minusculas(nuevoCaracter)
			"a":
				nuevoCaracter = "@"
			"e":
				nuevoCaracter = "#"
			"i":
				nuevoCaracter = "$"
			"o":
				nuevoCaracter = "%"
			"u":
				nuevoCaracter = "*"
		FinSegun
		
		nuevaFrase = Concatenar(nuevaFrase, nuevoCaracter)
	FinPara
FinSubProceso
	