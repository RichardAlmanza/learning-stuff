// La entrada al cine cuesta 5 d�lares por persona, sin embargo. Si la edad de la persona es menor a 12 a�os se le aplica un descuento del 30%.
// Escribir el algoritmo que calcule y muestre lo que pagar� la entrada al cine seg�n su edad.

Algoritmo BoletaCine
	Definir precioBoleta, edadUmbral, edad Como Entero
	Definir porcentaje, precioTotal Como Real
	
	precioBoleta = 5
	edadUmbral = 12
	porcentaje = 0.7
	
	Escribir "Ingrese la edad de la persona"
	Leer edad
	
	Si edad < edadUmbral Entonces
		precioTotal = precioBoleta * porcentaje
	SiNo
		precioTotal = precioBoleta
	FinSi
	
	Escribir "El precio a pagar es ", precioTotal
FinAlgoritmo
