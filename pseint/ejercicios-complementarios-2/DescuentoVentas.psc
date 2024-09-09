// Una tienda ofrece un descuento del 10% sobre el total de la compra realizada por un cliente en los meses de septiembre, octubre y noviembre.
// El programa solicitará al usuario que ingrese el mes y el importe de la compra.
// Luego, calculará y mostrará por pantalla el monto total que se debe cobrar al cliente, considerando el descuento si corresponde.

Algoritmo DescuentoVentas
	Definir mes Como Entero
	Definir compra, precioPorcentaje Como Real
	
	precioPorcentaje = 1
	
	Escribir "Ingrese el numero del mes de la compra"
	Leer mes
	
	Escribir "Ingrese el valor de la compra"
	Leer compra
	
	Si mes >= 9 Y mes <= 11 Entonces
		precioPorcentaje = 0.9
	FinSi
	
	Escribir "Debe pagar ", compra * precioPorcentaje
	
FinAlgoritmo
