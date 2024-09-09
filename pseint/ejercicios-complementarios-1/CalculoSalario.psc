// Desarrolla un programa que permita calcular la remuneración total que recibirá un vendedor al finalizar el mes.
// El vendedor recibe un salario base más un 10% adicional por comisión sobre sus ventas.
// El programa debe solicitar al usuario el salario base del vendedor y luego pedir el valor de cada una de las ventas realizadas en el mes.
// Después, calculará y mostrará el monto total que recibirá el vendedor, considerando tanto su salario base como las comisiones por las ventas realizadas en el período.

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
