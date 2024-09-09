// Escribe un programa que calcule cuántos litros de combustible consumió un automóvil.
// El usuario ingresará una cantidad de litros de combustible cargados en la estación y una cantidad de kilómetros recorridos,
// después, el programa calculará el consumo (km/lt) y se lo mostrará al usuario.

Algoritmo ConsumoKilometrosPorLitro
	Definir kilometros, litros, rendimiento Como Real
	
	Escribir "Ingrese la cantidad de litros cargados"
	Leer litros
	
	Escribir "Ingrese la cantidad de kilometros recorridos"
	Leer  kilometros
	
	rendimiento = kilometros / litros
	
	Escribir "El rendimiento es de ", rendimiento, " Kilometros/Litro"
FinAlgoritmo
