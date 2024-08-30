// Escribe un programa que:

// - Pida por teclado un número entero positivo.
// - Pregunte al usuario si desea introducir otro número.
// - Repita los pasos anteriores mientras el usuario no responda "n" o "N" (no).

// Muestre por pantalla la suma de los números introducidos por el usuario cuando determine NO ingresar más números. 

Algoritmo SumatoriaInfinita
	Definir sumatoria, nuevoNumero Como Entero
	Definir eleccionUsuario Como Caracter
	Definir continuar Como Logico
	
	sumatoria = 0
	continuar = Verdadero
	
	Repetir
		Escribir "Sub total: ", sumatoria
		
		Repetir
			Escribir "Ingresar un numero entero positivo para sumar"
			Leer nuevoNumero
		Hasta Que nuevoNumero > 0
		
		sumatoria = sumatoria + nuevoNumero
		
		Repetir
			Escribir "Desea ingresar mas numeros? SI/NO / S/N"
			Leer eleccionUsuario
			
			eleccionUsuario = Mayusculas(eleccionUsuario)
		Hasta Que eleccionUsuario == "N" O eleccionUsuario == "NO" O eleccionUsuario == "S" O eleccionUsuario == "SI"
		
		Si eleccionUsuario == "NO" O eleccionUsuario == "N" Entonces
			continuar = Falso
		FinSi
	Mientras Que continuar
	
	Escribir "La suma total es ", sumatoria
FinAlgoritmo
