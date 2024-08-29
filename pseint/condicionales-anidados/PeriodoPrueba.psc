// Escribe un programa para obtener el grado de eficiencia de un operario de una fábrica de tornillos,
// de acuerdo a las siguientes dos condiciones que se le imponen para un período de prueba:

// - Producir menos de 200 tornillos defectuosos.

// - Producir más de 10000 tornillos sin defectos.

// El grado de eficiencia se determina de la siguiente manera:
	
// - GRADO 5: Si no cumple ninguna de las condiciones.
// - GRADO 6: Si sólo cumple la primera condición
// - GRADO 7: Si sólo cumple la segunda condición.
// - GRADO 8: Si cumple las dos condiciones.

// Muestra un mensaje acorde al grado de eficiencia de un operario, luego de evaluar las condiciones. 
// Recuerda nombrar y guardar tu algoritmo.

// Nota: Para abordar este ejercicio de manera ordenada y eficiente, se recomienda probar cada inciso
// por separado en lugar de intentar resolver todos los puntos simultáneamente.
// Esto permite identificar y corregir posibles errores de forma más sencilla, ya que abordar todas las 
// tareas a la vez puede dificultar la identificación de problemas específicos

Algoritmo PeriodoPrueba
	Definir tornillosDefectuosos, tornillosPerfectos, banderas, grado Como Entero
	
	banderas = 0
	grado = 5
	
	// Lo siguiente no es recomendado por complicar la lectura, pero estaba aburrido
	
	// Recuerden Facil de entender/leer es mejor que `tener menos lineas`
	
	Escribir "Ingrese la cantidad de tornillos defectuosos"
	Leer tornillosDefectuosos
	
	Escribir "Ingrese la cantidad de tornillos perfectos"
	Leer tornillosPerfectos
	
	Si tornillosDefectuosos < 200 Entonces
		banderas = banderas + 1
	FinSi
	
	Si tornillosPerfectos > 1000 Entonces
		banderas = banderas + 2
	FinSi
	
	grado = grado + banderas
	
	Escribir "El grado es ", grado
FinAlgoritmo
