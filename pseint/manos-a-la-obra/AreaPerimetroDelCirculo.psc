// Solicita al usuario ingresar el valor del radio de una circunferencia y luego calcula y muestra por pantalla el área y perímetro.
// Para calcular estos valores, puedes usar las siguientes fórmulas:
// --  Area = PI * radio2
// --  Perimetro = 2 * PI * radio
// Recuerda que en matemáticas, ? (PI) representa el número aproximado de 3.14
// Recuerda que al realizar cálculos, tienes la opción de almacenar el resultado en una variable para su uso posterior,
// o bien ejecutar la operación directamente en una instrucción de salida, como por ejemplo, al escribir el resultado.

Algoritmo AreaPerimetroDelCirculo
	Definir radio, perimetro, area Como Real
	
	Escribir "Ingrese el Radio del Circulo"
	Leer radio
	
	perimetro = 2 * PI * radio
	area = PI * radio ^ 2
	
	Escribir "El Perimetro del Circulo es: ", perimetro
	Escribir "El Area del Circulo es: ", area
FinAlgoritmo
