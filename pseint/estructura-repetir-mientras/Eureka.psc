// Considerando que la clave de acceso a un sistema es "EUREKA",
// desarrolla un programa que solicite al usuario ingresar una clave.
// Se cuenta con tres intentos para adivinarla; si se fallan los tres intentos,
// se mostrará un mensaje indicando que se han agotado los intentos.
// En caso de acertar la clave, se mostrará un mensaje indicando que se
// ha ingresado al sistema correctamente. Recuerda nombrar y guardar tu algoritmo.

Algoritmo Eureka
	Definir CLAVE Como Caracter
	Definir MAXIMOS_INTENTOS Como Entero
	
	Definir claveUsuario Como Caracter
	Definir intentos Como Entero
	Definir autenticado Como Logico
	
	CLAVE = "EUREKA"
	MAXIMOS_INTENTOS = 3
	
	intentos = 0
	autenticado = Falso
	
	Repetir		
		Escribir "Ingrese la clave para ingresar"
		Leer claveUsuario
		
		Si claveUsuario == CLAVE Entonces
			autenticado = Verdadero	
		FinSi
		
		intentos = intentos + 1
	Mientras Que intentos < MAXIMOS_INTENTOS Y NO autenticado
	
	Si autenticado Entonces
		Escribir "Ha ingresado correctamente!"
	SiNo
		Escribir "Se han agotado los intentos"
	FinSi
	
FinAlgoritmo
