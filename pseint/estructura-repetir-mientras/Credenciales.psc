// Desarrolla un programa que solicite al usuario su código
// de usuario (un número entero mayor que cero)
// y su contraseña numérica (otro número entero positivo).
// El programa no permitirá continuar hasta que el usuario
// introduzca el código 1024 y la contraseña 4567.
// El programa finaliza cuando se ingresan los datos correctos.
// Recuerda nombrar y guardar tu algoritmo.

Algoritmo Credenciales
	Definir CODIGO, CONTRASENIA Como Entero
	Definir  codigoUsuario, contraseniaUsuario Como Entero
	
	CODIGO = 1024
	CONTRASENIA = 4567
	
	Repetir
		Repetir
			Escribir "Ingresar su código de usuario (un número entero mayor que cero)."
			Leer codigoUsuario
		Mientras Que codigoUsuario <= 0
		
		Repetir
			Escribir "Ingresar su contraseña numérica (otro número entero positivo)."
			Leer contraseniaUsuario
		Mientras Que contraseniaUsuario <= 0
	Mientras Que NO (codigoUsuario == CODIGO Y contraseniaUsuario == contraseniaUsuario)
	
	Escribir "Usted ha ingresado... tal vez a fuerza bruta"
FinAlgoritmo
