// Escribe un programa que calcule cu�ntos litros de combustible consumi� un autom�vil.
// El usuario ingresar� una cantidad de litros de combustible cargados en la estaci�n y una cantidad de kil�metros recorridos,
// despu�s, el programa calcular� el consumo (km/lt) y se lo mostrar� al usuario.

Algoritmo ConsumoKilometrosPorLitro
	Definir kilometros, litros, rendimiento Como Real
	
	Escribir "Ingrese la cantidad de litros cargados"
	Leer litros
	
	Escribir "Ingrese la cantidad de kilometros recorridos"
	Leer  kilometros
	
	rendimiento = kilometros / litros
	
	Escribir "El rendimiento es de ", rendimiento, " Kilometros/Litro"
FinAlgoritmo
