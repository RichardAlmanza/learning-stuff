// Solicita al usuario ingresar el valor del radio de una circunferencia y luego calcula y muestra por pantalla el �rea y per�metro.
// Para calcular estos valores, puedes usar las siguientes f�rmulas:
// --  Area = PI * radio2
// --  Perimetro = 2 * PI * radio
// Recuerda que en matem�ticas, ? (PI) representa el n�mero aproximado de 3.14
// Recuerda que al realizar c�lculos, tienes la opci�n de almacenar el resultado en una variable para su uso posterior,
// o bien ejecutar la operaci�n directamente en una instrucci�n de salida, como por ejemplo, al escribir el resultado.

Algoritmo AreaPerimetroDelCirculo
	Definir radio, perimetro, area Como Real
	
	Escribir "Ingrese el Radio del Circulo"
	Leer radio
	
	perimetro = 2 * PI * radio
	area = PI * radio ^ 2
	
	Escribir "El Perimetro del Circulo es: ", perimetro
	Escribir "El Area del Circulo es: ", area
FinAlgoritmo
