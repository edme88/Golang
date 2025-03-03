Pregunta 1:
¿Cuál es la forma idiomática de controlar los errores en Go?

- Use un bloque try/catch
- Use una condición if para una función que devuelva varios valores. EL último valor es el error. (correcto)
- Use una condición if y compruebe si la respuesta es nil.
- Deje que el programa finalice.

Pregunta 2:
¿Qué función puede usar para crear una variable de error?

- errors.New('Employee nor found!') (correcto)
- fmt.Errorf('Employee nor found!')
- error.New('Employee nor found!')
- log.Error('Employee nor found!')

Pregunta 3:
¿Cuál de las siguientes afirmaciones sobre el uso de defer en Go es verdadera?

- defer se utiliza para finalizar la ejecución de un programa inmediatamente.
- defer se utiliza para posponer la ejecución de una función hasta que la función actual termine. (correcto)
- defer se utiliza para capturar y manejar errores en Go.
- defer se utiliza para controlar el flujo de ejecución en Go.

Pregunta 4:
¿Cuál de las siguientes opciones describe correctamente el propósito de panic y recover en Go?

- panic se utiliza para manejar errores en tiempo de compilación, mientras que recover se utiliza para manejar errores en tiempo de ejecución.
- panic se utiliza para finalizar la ejecución de un programa, mientras que recover se utiliza para capturar y recuperarse de un panic. (correcto)
- panic se utiliza para reportar errores, mientras que recover se utiiza para registrar errores.
- panic y recover son sinónimos y se utilizan indistintamente para manejar errores en Go.

Pregunta 5:
¿Cuál de las siguientes opciones describe correctamente el registro de errores en Go?

- El registro de errores en Go se realiza utilizando el paquete log y la función Error. (correcto)
- El registro de errores en Go se realiza utilizando la función panic y capturando los errores con recover.
- El registro de errores en Go se realiza utilizando la función defer y el paquete log.
- El registro de errores en Go se realiza utilizando el paquete errores y la función New.

Resumen
En la sección del curso de Go dedicada al manejo de errores, exploramos diferentes aspectos del manejo de errores en Go y cómo aplicarlos de manera efectiva en nuestras aplicaciones.

Comenzamos aprendiendo sobre el manejo de errores en Go, donde vimos cómo utilizar la sintaxis if err != nil para verificar y manejar errores de forma concisa y eficiente. También exploramos las mejores prácticas para devolver y manejar errores adecuadamente en nuestras funciones.

Luego profundizamos en un ejemplo práctico de manejo de errores en Go, donde analizamos cómo implementar el manejo de errores en una aplicación real. Exploramos técnicas como el uso de variables de error, el registro de errores y la propagación adecuada de errores en diferentes capas de nuestra aplicación.

Uno de los aspectos más interesantes del manejo de errores en Go es el uso de la declaración defer. Aprendimos cómo usar defer para asegurarnos de que ciertas operaciones se realicen antes de salir de una función, incluso en caso de error. Vimos ejemplos prácticos que demostraron cómo defer puede facilitar el manejo de tareas como la liberación de recursos o el cierre adecuado de archivos.

También exploramos el uso de panic y recover. Aprendimos cómo utilizar estas dos funciones para manejar situaciones excepcionales y errores graves en nuestras aplicaciones. Vimos cómo panic puede interrumpir el flujo normal del programa y cómo recover nos permite recuperarnos de dichos errores y continuar la ejecución de manera controlada.

En el proyecto de esta sección, creamos un gestor de contactos utilizando Go. Aplicamos todos los conceptos que aprendimos sobre el manejo de errores en Go para implementar una funcionalidad robusta para agregar, mostrar y almacenar contactos, asegurándonos de manejar los posibles errores que pudieran ocurrir durante el proceso.

Concluimos esta sección con una sólida comprensión del manejo de errores en Go. Ahora estamos preparados para crear aplicaciones confiables y resistentes que manejen adecuadamente situaciones inesperadas.

[Materia de esta sección en GitHub](https://github.com/alexroel/curso-golang/blob/main/sections/05-error-handling.md) y Tambien puedes descargar en PDF.

Sigue aprendiedo con los siguientes tutoriales sobre esta sección:

[Implementación del control de errores y el registro en Go](https://learn.microsoft.com/es-es/training/modules/go-errors-logs/)
