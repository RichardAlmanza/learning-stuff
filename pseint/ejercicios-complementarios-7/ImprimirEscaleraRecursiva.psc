// Crea un programa que dibuje una escalera de numeros,
// donde cada linea comience en uno y termine en el numero de la linea.
// Solicita al usuario la altura de la escalera al comenzar.
// Ejemplo: si se ingresa el numero 3:

// 1
// 12
// 123

Algoritmo ImprimirEscaleraRecursiva
	Definir numeroUsuario Como Entero
	Definir _ Como Caracter
	
	Repetir
		Escribir "Ingrese el numero de filas para la escalera"
		Leer numeroUsuario
	Hasta Que numeroUsuario > 0
	
	imprimirFilaDesdeHasta(numeroUsuario, 1, -1)
	
	_ = imprimirEscalera(ConvertirATexto(numeroUsuario))
FinAlgoritmo

/// textoFila es tipo caracter y deber ser un numero entero positivo
Funcion salida <- imprimirEscalera (textoFila)
	Definir salida, nuevaFila Como Caracter
	Definir fila Como Entero
	
	Si textoFila == "1" Entonces
		salida = "1"
	SiNo
		fila = ConvertirANumero(textoFila) - 1
		nuevaFila = ConvertirATexto(fila)
		
		salida = Concatenar(imprimirEscalera(nuevaFila), textoFila)
	FinSi
	
	Escribir salida
FinFuncion

/// filaInicial, filaFinal y conPaso son de tipo entero
/// conPaso debe ser 1 o -1
Funcion imprimirFilaDesdeHasta(filaInicial, filaFinal, conPaso)
	imprimirFilaSinConcatenar(filaInicial)
	Escribir ""
	
	Si filaInicial <> filaFinal Entonces
		imprimirFilaDesdeHasta(filaInicial + conPaso, filaFinal, conPaso)
	FinSi
FinFuncion

/// fila es tipo entero y debe ser positivo
Funcion imprimirFilaSinConcatenar (fila)
	Si fila <> 1 Entonces
		imprimirFilaSinConcatenar(fila - 1)
	FinSi
	
	Escribir Sin Saltar fila
FinFuncion
	