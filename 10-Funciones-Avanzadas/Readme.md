## Funciones Avanzadas

- Funciones variádicas (reciben n parámetros)
- Funciones recursivas
- Funciones anónimas
- Funciones de orden superior
- Cosures en Go

Pregunta 1:
¿Cuál de las siguientes opciones describe mejor las funciones variádicas en Go?

- Funciones que pueden recibir n número de variables de argumentos del mismo tipo o diferente. (correcto)
- Funciones que solo pueden recibir un argumento y devolver un valor.
- Funciones que pueden recibir argumentos de diferentes tipos.
- Funciones que no pueden recibir argumentos.

Pregunta 2:
¿Cuál de las siguientes afirmaciones es cierta sobre las funciones recursivas en Go?

- Las funciones recursivas no son admitidas en Go.
- Las funciones recursivas permiten llamarse a sí mismas dentro de su cuerpo. (correcto)
- Las funciones recursivas solo pueden contener un solo bloque de código.
- Las funciones recursivas siempre producen un desbordamiento de pila (stack overflow)

Pregunta 3:
En Go, ¿cuál es la forma correcta de declarar una función anónima y asignarla a una variable llamada "miFuncion"?

- func miFunction() {}
- miFuncion := func() {} (correcto)
- anonymous miFuncion() {}
- func() miFuncion {}

Pregunta 4:
¿Cuál de las siguientes opciones describe mejor las funciones de orden superior en Go?

- Funciones que se pueden pasar como argumentos y/o ser devueltas como resultados. (correcto)
- Funciones que solo pueden recibir una cantidad fija de argumentos.
- Funciones que solo pueden realizar operaciones matemáticas simples.
- Funciones que no pueden llamar a otras funciones.

Pregunta 5:
¿Cuál de las siguientes afirmaciones es cierta sobre los closures en Go?

- Los closures son funciones que no tienen acceso a variables fuera de su alcance.
- Los closures solo se pueden utilizar en funcoines recursivas.
- Los closures en Go se definen con la palabra clave "closure".
- Los closures permiten que una función acceda y mantenga referencias variables fuera de su propio alcance léxico. (correcto)

Resumen
En la sección de Funciones Avanzadas del curso de Go, exploramos una variedad de conceptos poderosos que nos permiten llevar nuestras habilidades de programación a un nivel superior.

Comenzamos aprendiendo sobre las funciones variádicas, que nos permiten definir funciones capaces de aceptar un número variable de argumentos. Esta característica resulta especialmente útil cuando no conocemos de antemano cuántos argumentos necesitaremos pasar a una función.

Continuamos con las funciones recursivas, que son funciones que se llaman a sí mismas. Aprendimos a utilizar la recursión de manera efectiva para resolver problemas complejos, evitando trampas comunes como la recursión infinita.

Posteriormente, exploramos las funciones anónimas, también conocidas como closures. Estas son funciones sin nombre que pueden asignarse a variables o pasarse como argumentos a otras funciones. Descubrimos cómo las funciones anónimas nos permiten escribir código más conciso y expresivo.

Luego nos sumergimos en las funciones de orden superior, que son funciones capaces de tomar otras funciones como argumentos o devolver funciones como resultados. Aprendimos a utilizar estas funciones para implementar patrones de programación comunes, como el patrón de estrategia o el patrón de iterador.

Por último, abordamos el tema de los closures en Go. Un closure es una función que captura y mantiene referencias a variables locales dentro de su ámbito. Aprendimos cómo los closures nos permiten crear funciones que retienen el estado incluso después de que la función original haya terminado de ejecutarse.

A lo largo de la sección, pusimos en práctica los conceptos aprendidos mediante ejemplos y ejercicios que nos desafiaron a aplicarlos de manera creativa. Al finalizar, adquirimos un conjunto de herramientas poderosas que ampliaron nuestras capacidades como programadores en Go.

[Materia de esta sección en GitHub](https://github.com/alexroel/curso-golang/blob/main/sections/09-functions-go.md) y Tambien puedes descargar en PDF.

Sigue aprendiedo con los siguientes tutoriales sobre esta sección:

[Ejemplo de funciones avanzadas](https://gobyexample.com/variadic-functions)
