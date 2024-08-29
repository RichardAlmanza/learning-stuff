// Crea un programa que solicite al usuario ingresar un límite positivo.
// Luego, pide números al usuario hasta que la suma de los números
// introducidos supere este límite inicial.
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
