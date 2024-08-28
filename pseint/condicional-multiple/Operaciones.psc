// Desarrolla un programa que simule un men� de opciones para realizar
// las cuatro operaciones aritm�ticas b�sicas (suma, resta, multiplicaci�n y divisi�n) con dos valores num�ricos enteros.
// El usuario deber� especificar la operaci�n deseada utilizando el primer car�cter de la operaci�n: 'S' para suma,
// 'R' para resta, 'M' para multiplicaci�n, y 'D'  para divisi�n. Recuerda nombrar y guardar tu algoritmo

Algoritmo Operaciones
	Definir numero1, numero2 Como Entero
	Definir resultado Como Real
	Definir operacion, mensajeError Como Caracter
	Definir error Como Logico
	
	error = Falso
	mensajeError = "No ingreso solamente la inicial de una de las operaciones"
	
	Escribir "Ingrese el primer numero"
	Leer numero1
	
	Escribir "Ingrese el segundo numero"
	Leer numero2
	
	Escribir "Ingrese la inicial de la operacion que desea realizar (suma, resta, multiplicaci�n y divisi�n)"
	Leer operacion
	
	operacion = Mayusculas(operacion)
	
	Segun operacion Hacer
		"S":
			resultado = numero1 + numero2
		"R":
			resultado = numero1 - numero2
		"M":
			resultado = numero1 * numero2
		"D":
			Si numero2 == 0 Entonces
				numero1 = 0
				numero2 = 1
				
				error = Verdadero
				mensajeError = "No puede realizar divisiones entre 0"
			FinSi
			
			resultado = numero1 / numero2
		De Otro Modo:
			error = Verdadero
	FinSegun
	
	Si error Entonces
		Escribir mensajeError
	SiNo
		Escribir "El resultado es ", resultado
	FinSi
FinAlgoritmo
