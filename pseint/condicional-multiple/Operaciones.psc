// Desarrolla un programa que simule un menú de opciones para realizar
// las cuatro operaciones aritméticas básicas (suma, resta, multiplicación y división) con dos valores numéricos enteros.
// El usuario deberá especificar la operación deseada utilizando el primer carácter de la operación: 'S' para suma,
// 'R' para resta, 'M' para multiplicación, y 'D'  para división. Recuerda nombrar y guardar tu algoritmo

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
	
	Escribir "Ingrese la inicial de la operacion que desea realizar (suma, resta, multiplicación y división)"
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
