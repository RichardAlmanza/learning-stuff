// Crea un programa que solicite al usuario ingresar un l�mite positivo.
// Luego, pide n�meros al usuario hasta que la suma de los n�meros
// introducidos supere este l�mite inicial.
// Recuerda nombrar y guardar tu algoritmo

Algoritmo SumaAlLimite
	Definir limitePositivo, sumatoria, nuevoNumero Como Real
	
	limitePositivo = -1
	sumatoria = 0
	
	Mientras limitePositivo <= 0 Hacer
		Escribir "Ingrese un numero positivo para ser el numero limite"
		Leer limitePositivo
	FinMientras
	
	Mientras sumatoria < limitePositivo Hacer
		Escribir "Sub total: ", sumatoria
		Escribir "Limite Positivo: ", limitePositivo
		Escribir "Ingrese un numero a sumar hasta llegar o superar el limite"
		Leer nuevoNumero
		
		sumatoria = sumatoria + nuevoNumero
	FinMientras
	
	Escribir "Total es ", sumatoria
FinAlgoritmo
