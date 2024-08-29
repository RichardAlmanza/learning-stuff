// Crea un programa que almacene una vocal secreta en una variable.
// Solicita al usuario que intente adivinar la vocal secreta, permitiéndole
// intentar tantas veces como sea necesario hasta que la adivine.
// Recuerda nombrar y guardar tu algoritmo.

Algoritmo AdivinaVocal
	Definir VOCALES, VOCAL_SECRETA Como Caracter
	Definir INDICE_VOCAL Como Entero
	Definir vocalUsuario Como Caracter
	
	VOCALES = "aeiou"
	
	INDICE_VOCAL = Aleatorio(0, Longitud(VOCALES) - 1)
	VOCAL_SECRETA = Subcadena(VOCALES, INDICE_VOCAL, INDICE_VOCAL)
	
	vocalUsuario = ""
	
	Mientras vocalUsuario <> VOCAL_SECRETA Hacer
		Escribir "Adivine la vocal!"
		Leer vocalUsuario
		
		vocalUsuario = Minusculas(vocalUsuario)
	FinMientras
	
	Escribir "Haz acertado!"
FinAlgoritmo
