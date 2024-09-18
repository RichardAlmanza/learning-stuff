// Dise�a un algoritmo que permita al usuario ingresar una frase.
// Deber�s asegurarte de que la longitud de la frase no exceda los 100 caracteres;
// en caso de superar este l�mite, se solicitar� al usuario que vuelva a ingresar
// la frase hasta que cumpla con esta condici�n. 

// Una vez ingresada la frase, el programa separar� los caracteres en dos vectores:
// uno para almacenar exclusivamente las vocales (A-E-I-O-U) y otro para las
// consonantes, as� como para los caracteres especiales y los espacios. 

// Los elementos en cada vector se acomodar�n inicialmente por orden de aparici�n en la frase.
// Adem�s, se requiere la implementaci�n de un subprograma que permita mostrar los vectores creados,
// as� como el recuento total de elementos en cada uno de ellos.

Algoritmo FiltrarTextoEnDos
	Definir LIMITE_TAMANO Como Entero
	Definir FILTRO Como Caracter
	Definir textoUsuario Como Caracter
	
	LIMITE_TAMANO = 100
	FILTRO = "aeiou�����"
	FILTRO = Concatenar(FILTRO, Mayusculas(FILTRO))
	
	Repetir
		Escribir "Ingrese una frase que no exceda el limite de 100 caracteres"
		Leer textoUsuario
	Hasta Que Longitud(textoUsuario) <= LIMITE_TAMANO
	
	Escribir "Frase Original"
	Escribir textoUsuario
	
	Escribir "Filtro a utilizar"
	Escribir FILTRO
	
	Escribir "Caracteres Filtrados"
	Escribir filtrarTexto(textoUsuario, FILTRO, Verdadero)
	
	Escribir "Caracteres sin filtrar"
	Escribir filtrarTexto(textoUsuario, FILTRO, Falso)
FinAlgoritmo

Funcion textoFiltrado <- filtrarTexto(textoAFiltrar, cadenaFiltro, esIncluyente)
	Definir indice Como Entero
	Definir caracterTexto, textoFiltrado Como Caracter
	Definir esValido Como Logico
	
	textoFiltrado = ""
	
	Para indice = 0 Hasta Longitud(textoAFiltrar) - 1 Con Paso 1 Hacer
		caracterTexto = Subcadena(textoAFiltrar, indice, indice)
		esValido = caracterEnCadena(caracterTexto, cadenaFiltro)
		
		Si NO esIncluyente Entonces
			esValido = !esValido
		FinSi
		
		Si esValido Entonces
			textoFiltrado = Concatenar(textoFiltrado, caracterTexto)
		FinSi
	FinPara
FinSubProceso


Funcion encontrado <- caracterEnCadena(caracterABuscar, opciones)
	Definir indice Como Entero
	Definir encontrado Como Logico
	
	encontrado = Falso
	indice = 0
	
	Mientras indice < Longitud(opciones) Y NO encontrado Hacer
		encontrado = Subcadena(opciones, indice, indice) == caracterABuscar
		
		indice = indice + 1
	FinMientras
FinFuncion
