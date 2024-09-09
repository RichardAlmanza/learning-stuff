// Realiza un procedimiento que permita realizar la división entre
// dos números y muestre el cociente y el residuo utilizando
// el método de restas sucesivas.
// Este método consiste en restar el divisor del dividendo repetidamente
// hasta obtener un resultado menor que el divisor.
// El residuo será el resultado final y el número de restas realizadas
// será el cociente. Por ejemplo, para 50 / 13:
// - 50 ? 13 = 37 (una resta)
// - 37 ? 13 = 24 (dos restas)
// - 24 ? 13 = 11 (tres restas)
// Como 11 es menor que 13, el residuo es 11 y el cociente es 3.
// Recuerda nombrar y guardar tu algoritmo.

Algoritmo SubprocesoDisivionConResiduo
	Definir cociente, residuo, dividendoUsuario, divisorUsuario Como Entero
	
	Escribir "Ingrese el numero que desea dividir"
	Leer dividendoUsuario
	
	Repetir
		Escribir "Ingrese el numero por que el cual desea dividir a ", dividendoUsuario
		Leer divisorUsuario
	Hasta Que divisorUsuario <> 0
	
	divisionConResiduo(cociente, residuo, dividendoUsuario, divisorUsuario)
	
	Escribir "El cociente es ", cociente, " y el residuo es ", residuo, " para de division de ", dividendoUsuario, " entre ", divisorUsuario
FinAlgoritmo

SubProceso divisionConResiduo (cociente Por Referencia, residuo Por Referencia, dividendo, divisor)
	Definir esCocienteNegativo, esResiduoNegativo Como Logico
	
	esResiduoNegativo = dividendo < 0
	
	esCocienteNegativo = dividendo < 0 Y divisor > 0
	esCocienteNegativo = esCocienteNegativo O (dividendo > 0 Y divisor < 0)
	
	dividendo = abs(dividendo)
	divisor = abs(divisor)
	cociente = 0
	
	Mientras dividendo >= divisor Hacer
		dividendo = dividendo - divisor
		cociente = cociente + 1
	FinMientras
	
	residuo = dividendo
	
	Si esCocienteNegativo Entonces
		cociente = -1 * cociente
	FinSi
	
	Si esResiduoNegativo Entonces
		residuo = -1 * residuo
	FinSi
FinSubProceso
