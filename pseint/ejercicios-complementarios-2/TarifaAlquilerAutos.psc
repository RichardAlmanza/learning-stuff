// La empresa "Te llevo a todos lados" tiene un sistema de tarifa para el alquiler de autos por hora.
// Si el cliente devuelve el auto dentro de las 2 horas de uso, el costo es de $400 pesos y la nafta es gratuita.
// Si el cliente devuelve el auto después de las 2 horas, se ingresa la cantidad de litros de nafta gastados y el tiempo transcurrido en horas.
// El programa calculará el total a pagar por el cliente, considerando un costo de $40 por litro de nafta y $5,20 por minuto de uso después de las primeras 2 horas.

Algoritmo TarifaAlquilerAutos
	Definir precioNafta, precioAlquilerBase, precioAlquilerExtra Como Real
	Definir naftaUsado, tiempoAlquiler, tiempoAlquilerExtra, total Como Real
	
	precioNafta = 40
	precioAlquilerBase = 400
	precioAlquilerExtra = 5.2 * 60
	
	naftaUsado = 0
	tiempoAlquilerExtra = 0
	
	Escribir "Ingrese el tiempo alquilado en horas"
	Leer tiempoAlquiler
	
	Si tiempoAlquiler > 2 Entonces
		tiempoAlquilerExtra = tiempoAlquiler - 2
		
		Escribir "Ingrese la cantidad de litros usados de nafta"
		Leer naftaUsado
	FinSi
	
	total = precioAlquilerBase
	total = total + naftaUsado * precioNafta
	total = total + tiempoAlquilerExtra * precioAlquilerExtra
	
	Escribir "Debe pagar ", total
	
FinAlgoritmo
