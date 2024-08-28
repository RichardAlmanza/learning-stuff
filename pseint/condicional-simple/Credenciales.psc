// Diseña un programa que solicite al usuario ingresar su nombre de usuario y contraseña.
// El programa debe verificar si el nombre de usuario es "admin" y si la contraseña es "1234".
// Si ambos son correctos, el programa debe imprimir un mensaje de bienvenida. Recuerda nombrar y guardar tu algoritmo.

Algoritmo Credenciales
	Definir usuario, contrasena Como Caracter
	
	Escribir "Ingrese el usuario"
	Leer usuario
	
	Escribir "Ingrese la contraseña"
	Leer contrasena
	
	Si usuario == "admin" Y contrasena == "1234" Entonces
		Escribir "Bienvenido"
	FinSi
FinAlgoritmo
