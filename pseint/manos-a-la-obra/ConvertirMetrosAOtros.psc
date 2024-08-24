// A partir de una conocida cantidad de metros que el usuario ingresa a través del teclado
// se debe obtener su equivalente en centímetros, en milímetros y en pulgadas.

// Equivalencias:
//   1 metro equivale a 100 centímetros.
//   1 metro equivale a 1000 milímetros.
//   1 pulgada equivale a 2.54 centímetros.

// Recuerda nombrar y guardar tu algoritmo.

Algoritmo ConvertirMetrosAOtros
	Definir metros, centimetros, milimetros, pulgadas Como Real
	
	Escribir "Ingrese la cantidad de metros"
	Leer metros
	
	centimetros = 100 * metros
	milimetros = 1000 * metros
	pulgadas = centimetros / 2.54
	
	Escribir "Son ", centimetros, " centimetros"
	Escribir "Son ", milimetros, " milimetros"
	Escribir "Son ", pulgadas, " pulgadas"
	
FinAlgoritmo
