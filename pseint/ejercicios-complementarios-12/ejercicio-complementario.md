# Ejercicio Complementario

> ✨ Este ejercicio **es de tipo complementario**. Esto quiere decir que te ayudará a avanzar en profundidad en el tema visto, pero **no es obligatorio**.

El mundo de la investigación se enfrenta a un nuevo desafío científico. Un brillante investigador, el Dr. Marlowe, ha descubierto una intrigante correlación entre secuencias genéticas codificadas y patrones matriciales.

Cada muestra de ADN se presenta como una secuencia de caracteres compuesta por cuatro bases: A, B, C y D. Por ejemplo, una muestra podría ser "ACDDCADBCDABDBBA".

El Dr. Marlowe, empleando un enfoque innovador, convierte estas secuencias en matrices cuadradas (MxM) y busca patrones específicos en las diagonales principales de estas matrices. **El patrón específico consiste en que ambos conjuntos de caracteres en las diagonales principales de la matriz sean idénticos, aunque los caracteres entre sí pueden ser diferentes**. Por ejemplo, una muestra podría generar la siguiente matriz:

**A C D D**

**C A D B**

**C D A B**

**D B B A**

Sin embargo, lo intrigante es que el tamaño de la matriz no se conoce de antemano y debe inferirse de la muestra ingresada.

<u>Tu Misión</u>: Desarrollar un programa que permita al usuario ingresar una muestra de ADN completa y determinar su validez según los patrones identificados por el Dr. Marlowe. Si la muestra es válida, el programa deberá mostrar la matriz generada y un mensaje indicando si se ha encontrado un patrón específico. En caso contrario, se solicitará al usuario que ingrese una nueva muestra.

<u>Para que una muestra sea considerada válida, debe cumplir con los siguientes criterios:</u>

Debe tener una longitud de 9 caracteres para ser almacenada en una matriz de 3x3 o de 16 caracteres para ser almacenada en una matriz de 4x4, respectivamente.

Los caracteres de la muestra deben ser exclusivamente A, B, C o D para ser considerada válida.

<u>Para encontrar el patrón especificado:</u>

Una vez validada la muestra ingresada, se procederá a crear la matriz correspondiente para recorrer sus diagonales. Si todos los caracteres de la diagonal principal 1 son iguales, indicando que cumple con el patrón, se continuará analizando la diagonal principal 2. Si todos los caracteres de la diagonal principal 2 también son iguales, se concluirá que la muestra ingresada presenta los patrones específicos esperados. En ese caso, se mostrará en pantalla el mensaje “¡Felicidades! La muestra contiene patrones específicos en sus diagonales”. En caso contrario, si la muestra no cumple con los patrones especificados en ambas diagonales, se mostrará en pantalla el mensaje “La muestra NO contiene los patrones especificados en sus diagonales”.
