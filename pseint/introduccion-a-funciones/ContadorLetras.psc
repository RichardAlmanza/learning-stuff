// Diseña una función que reciba una frase y una letra proporcionadas
// por el usuario y devuelva la cantidad de veces que la letra aparece en la frase.
// Recuerda nombrar y guardar tu algoritmo.

Algoritmo ContadorLetras
	Definir iter Como Entero
	Definir fraseUsuario, letraUsuario Como Caracter
	
	Para iter <- 1 Hasta 5 Hacer
		Escribir "Ingrese la frase en la que desea buscar"
		Leer fraseUsuario
		
		Escribir "Ingrese la letra que desea contar en la frase anteriormente ingresada"
		Leer letraUsuario
		
		Escribir "La letra aparece ", contarLetraEn(letraUsuario, fraseUsuario), " veces en ", fraseUsuario
	FinPara
FinAlgoritmo

Funcion resultado <- contarLetraEn(letra, frase)
	Definir resultado, indice Como Entero
	Definir fraseInterna, letraInterna Como Caracter
	
	letraInterna = Minusculas(letra)
	fraseInterna = Minusculas(frase)
	
	resultado = 0
	
	Para indice <- 0 Hasta Longitud(fraseInterna) - 1 Hacer
		Si letraInterna == Subcadena(fraseInterna, indice, indice) Entonces
			resultado = resultado + 1
		FinSi
	FinPara
FinFuncion
	