# Actividad - Desafío integrador Grupal

> 🚀 **RECUERDEN**: Antes de comenzar, asegúrense de leer todo el enunciado para comprender el alcance completo de la actividad.
---
>💡 **Para empezar**, les recomendamos ordenar el diagrama de flujo de las acciones necesarias que deben cubrir. Pueden utilizar herramientas como compartir pantalla o pizarras compartidas, como Jamboard o Excalidraw, para facilitar esta tarea.
---
> 🔍 Es importante **reconocer qué variables serán necesarias** y de qué tipo deben ser definidas para llevar a cabo la actividad de manera eficiente.

## Sistema de Reciclaje de Botellas Automático

Se requiere desarrollar un sistema para una máquina de reciclaje de botellas que recompensará a los usuarios por la cantidad de plástico reciclado. El sistema solicitará a los usuarios que ingresen su nombre de usuario y contraseña para cargar el saldo correspondiente a su cuenta.

1. **Condición Simple Anidada**:  En esta estructura, se verificará inicialmente el nombre de usuario. Si el nombre de usuario coincide con "Albus_D", se procederá a validar la contraseña como "caramelosDeLimon". Si la contraseña es correcta, se establecerá la variable lógica "Login" como verdadera. En resumen, la variable "Login" solo se activará cuando tanto el nombre de usuario como la contraseña coincidan con los esperados.

2. **Bucle Mientras**: La validación de la contraseña se llevará a cabo dentro de un bucle Mientras, lo que permitirá al usuario tener un máximo de tres intentos para ingresar la contraseña correctamente.

3. **Bucle Hacer Mientras (Repetir)**: Una vez que el inicio de sesión sea exitoso (es decir, cuando "Login" sea verdadero), el usuario podrá acceder al menú de opciones, que incluye:

   1. Ingresar Botellas: Inicialmente,el usuario ingresará la cantidad de botellas a reciclar, y se generará un peso aleatorio para cada botella (entre 100 y 3000 gramos). Según el peso unitario, se asignará un valor monetario utilizando un condicional múltiple.

      - Si el peso es menos de 500 gr, corresponden $50
      - Si el peso es entre 501 gr y 1500 gr, corresponden $125
      - Si el peso es más de 1501 gr, corresponden $200

       Hecho esto, el programa debe informar al usuario por pantalla el valor que se le ofrece por el total de las botellas ingresadas. Si el usuario acepta, lo acreditamos a su saldo, si no se debe devolver el material (sólo mostrar en pantalla “Devolviendo material”). Para esto se debe utilizar un condicional doble.

   2. Consultar Saldo: Se mostrará al usuario el valor monetario actual en su saldo, en una variable definida previamente que “acumula” dicho importe. .

   3. Salir: Permitirá al usuario salir del sistema.

4. **Finalización de Opciones**: Después de completar las acciones de "Ingresar Botellas" o "Consultar Saldo", el programa volverá al menú principal automáticamente.

5. **Bucle Para (Incorporación opcional)**:El bucle Para se utilizará específicamente dentro de la opción "Ingresar Botellas". Permitirá al usuario ingresar múltiples botellas en una misma sesión, si así lo desea.

Si lo desean, pueden incorporar alguna otra funcionalidad o validación que consideren que optimizará el funcionamiento de su sistema.

> 🤝**Aunque no hayan logrado completar todos los puntos de las actividades en equipo, cada integrante puede luego terminarlas de manera individual**. Esta experiencia colaborativa les permitirá afianzar los conocimientos adquiridos, desarrollar habilidades de pensamiento lógico y creativo, así como también fomentar la cooperación entre los miembros del equipo. Además, al trabajar juntos, tendrán la oportunidad de enriquecerse académicamente al compartir y discutir diversas opiniones y perspectivas. Si bien el resultado final es importante, el proceso de trabajo en equipo es igualmente valioso para el aprendizaje y el crecimiento personal.
