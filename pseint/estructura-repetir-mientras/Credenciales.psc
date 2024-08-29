// Desarrolla un programa que solicite al usuario su c�digo
// de usuario (un n�mero entero mayor que cero)
// y su contrase�a num�rica (otro n�mero entero positivo).
// El programa no permitir� continuar hasta que el usuario
// introduzca el c�digo 1024 y la contrase�a 4567.
// El programa finaliza cuando se ingresan los datos correctos.
// Recuerda nombrar y guardar tu algoritmo.

Algoritmo Credenciales
	Definir CODIGO, CONTRASENIA Como Entero
	Definir  codigoUsuario, contraseniaUsuario Como Entero
	
	CODIGO = 1024
	CONTRASENIA = 4567
	
	Repetir
		Repetir
			Escribir "Ingresar su c�digo de usuario (un n�mero entero mayor que cero)."
			Leer codigoUsuario
		Mientras Que codigoUsuario <= 0
		
		Repetir
			Escribir "Ingresar su contrase�a num�rica (otro n�mero entero positivo)."
			Leer contraseniaUsuario
		Mientras Que contraseniaUsuario <= 0
	Mientras Que NO (codigoUsuario == CODIGO Y contraseniaUsuario == contraseniaUsuario)
	
	Escribir "Usted ha ingresado... tal vez a fuerza bruta"
FinAlgoritmo
