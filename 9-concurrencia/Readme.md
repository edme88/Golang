## Concurrencia

Permite aprovechar al máximo los recursos de hardware y mejorar el rednimiento de las aplicaciones

- Concurrencia
- Go run times (como los thread o hilos en otros lenguajes de programación. Ejecución simultaneo o en orden no determinista. Se pueden hacer multi-hilos si es multi procesos)
- Canales (comunicación y la sincronización entre los go run times)

Paquete http

Pregunta 1:
¿Cuál de las siguientes opciones describe mejor el concepto de concurrencia en Go?

- La capacidad de ejecutar múltiples hilos de ejecución simultáneamente.
- La capacidad de ejecutar múltiples instrucciones en paralelo.
- La capacidad de ejecutar múltiples programas independientes al mismo tiempo.
- La capacidad de ejecutar múltiples funciones en diferentes goroutines. (correcta)

Pregunta 2:
En Go, ¿cuál es la forma correcta de lanzar una nueva goroutine para ejecutar la función "miFuncion()"?

- miFuncion()
- go miFuncion() (correcto)
- tart miFuncion()
- newGoroutine miFuncion()

Pregunta 3:
¿Cuál de las siguientes afirmaciones es cierta sobre los canales en Go?

- Los canales solo s epueden utilizar par ala comunicación entre hilos de ejecución.
- Los canales solo se pueden utilizar para enviar valores, no para recibirlos.
- Los canales proporcionan una forma de sincronizar y comunicar datos entre goroutines. (correcto)
- Los canales son exclusivos de la programación secuencial y no se utilizan en la concurrencia.

En esta sección del curso de Go, exploramos cómo Go maneja la concurrencia de manera efectiva. Nos centramos en dos aspectos clave: el uso de goroutines y el manejo de canales.

Las goroutines son funciones independientes que se ejecutan de forma concurrente y ligera dentro de un programa Go. Aprendimos a crear y utilizar goroutines, aprovechando su capacidad para realizar múltiples tareas simultáneamente y maximizar el rendimiento del hardware.

Los canales, por otro lado, son herramientas fundamentales para la comunicación y sincronización entre goroutines. Aprendimos a enviar y recibir datos de manera segura y coordinada utilizando canales, evitando condiciones de carrera y garantizando una ejecución coherente de las goroutines.

Además, exploramos patrones comunes de concurrencia en Go y técnicas avanzadas para optimizar el rendimiento de nuestras aplicaciones.

En resumen, esta sección nos brindó una comprensión profunda de cómo aprovechar la concurrencia en Go, utilizando goroutines y canales para escribir programas rápidos, eficientes y concurrentes.

Materia de esta sección en GitHub y Tambien puedes [descargar en PDF](https://github.com/alexroel/curso-golang/blob/main/sections/08-concurrencia.md).

Sigue aprendiedo con los siguientes tutoriales sobre esta sección:

[Funcionamiento de la simultaneidad en Go](https://learn.microsoft.com/es-es/training/modules/go-concurrency/)
