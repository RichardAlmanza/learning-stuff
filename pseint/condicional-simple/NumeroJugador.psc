// Crea un programa que permita ingresar un número de jugador, si el número es mayor de 10, se debe calcular y mostrar en pantalla el 18% de este.
// Recuerda nombrar y guardar tu algoritmo.

Algoritmo NumeroJugador
	Definir numero_, porcentaje Como Real
	porcentaje = 0.18
	
	Escribir "Ingrese el numero de jugador"
	Leer numero_
	
	Si numero_ > 10 Entonces
		Escribir "El 18% es ", numero_ * porcentaje
	FinSi
FinAlgoritmo
