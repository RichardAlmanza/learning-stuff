// Escribe un programa que te pida ingresar los resultados de los �ltimos 3 partidos de tu equipo y valide si la suma de las anotaciones est� entre 1 y 10.
// Si se encuentran dentro de estos par�metros, se debe establecer como verdadera una variable l�gica y, en caso contrario, como falsa. Al final del programa,
// se deber� indicar si los 3 resultados son correctos utilizando la variable l�gica.

Algoritmo SumaPartidos
	Definir resultadoPartido1, resultadoPartido2, resultadoPartido3, sumatoria Como Entero
	Definir esSumaEntre1Y10 Como Logico
	
	Escribir "Ingrese el resultado del partido 1"
	Leer resultadoPartido1
	
	Escribir "Ingrese el resultado del partido 2"
	Leer resultadoPartido2
	
	Escribir "Ingrese el resultado del partido 3"
	Leer resultadoPartido3
	
	sumatoria = resultadoPartido1 + resultadoPartido2 + resultadoPartido3
	esSumaEntre1Y10 = sumatoria >= 1 Y sumatoria <= 10
	
	Si esSumaEntre1Y10 Entonces
		Escribir "Los resultados son correctos"
	SiNo
		Escribir "Los resultados NO son correctos"
	FinSi
FinAlgoritmo
