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
