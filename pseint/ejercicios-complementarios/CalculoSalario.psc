// Desarrolla un programa que permita calcular la remuneraci�n total que recibir� un vendedor al finalizar el mes.
// El vendedor recibe un salario base m�s un 10% adicional por comisi�n sobre sus ventas.
// El programa debe solicitar al usuario el salario base del vendedor y luego pedir el valor de cada una de las ventas realizadas en el mes.
// Despu�s, calcular� y mostrar� el monto total que recibir� el vendedor, considerando tanto su salario base como las comisiones por las ventas realizadas en el per�odo.

Algoritmo CalculoSalario
	Definir salarioBase, ventas, comision, salarioFinal Como Real
	
	Escribir "Ingrese el salario base del vendedor"
	Leer salarioBase
	
	Escribir "Ingrese las el valor de las ventas del vendedor"
	Leer ventas
	
	comision = 0.1 * ventas
	salarioFinal = salarioBase + comision
	
	Escribir "El salario final del vendedor es ", salarioFinal
FinAlgoritmo
