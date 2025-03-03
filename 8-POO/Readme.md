Pregunta 1:
En Go, ¿cuál es la forma correcta de definir un método llamado "Imprimir" para la estructura "Persona"?

- func (p Persona) Imprimir() {} (correcto)
- func Imprimir(Perosna) {}
- func Perosna.Imprimir() {}
- func Imprimir() Persona {}

Pregunta 2:
En Go, ¿cuál de las siguientes opciones muestra la forma correcta de definir funciones y variables privadas?

- Prefijar el nombre de la función o variable con un quión bajo, por ejemplo "\_miFuncion".
- Utilizar el modificador "private" antes del nombre de la función o variable, por ejemplo, "private miFuncion"
- Utilizar la palabra pace "private" antes del nombre de la función o variable, por ejemplo, "private func miFuncion() {}"
- En Go no existe un mecanismo directo para definir funcione so variables privadas como en otros lenguajes de programación. Sin embargo, Go proporciona una convención par alograr el mismo efecto. (correcto)

Pregunta 3:
En Go, ¿cuál es el mecanismo utilizado para lograr el polimorfismo?

- Clases abstractas
- Herencia múltiple
- Interfaces (correcto)
- Sobrecarga de operadodes

Pregunta 4:
En Go, ¿cuál es la forma correcta de implementar una interfaz llamada "Ordenable" en una estructura "Producto"?

- type Producto implements Ordenable {}
- type Ordenable implements Producto {}
- type Producto struct { Ordenable }
- type Producto struct { } //no se requiere una implementación explícita (correcta)

Resumen
En esta sección del curso, nos adentramos en la programación orientada a objetos (POO) y las interfaces en el lenguaje de programación Go. Exploramos conceptos clave como métodos y constructores, encapsulamiento, composición, polimorfismo e interfaces.

Aprendimos cómo los métodos en Go son funciones asociadas a un tipo de datos en particular, lo que nos permite definir comportamientos y operaciones específicas para los objetos. También vimos cómo crear constructores para inicializar los objetos de manera conveniente.

Además, exploramos el encapsulamiento en Go, que consiste en ocultar los detalles internos de una entidad y exponer una interfaz pública para interactuar con ella. Utilizamos la convención de nomenclatura de mayúsculas/minúsculas para lograr el encapsulamiento de manera efectiva.

En cuanto a la composición, vimos cómo utilizarla en Go para crear relaciones entre objetos y reutilizar código de manera eficiente. También exploramos el concepto de polimorfismo, que nos permite que un objeto tome diferentes formas y responda a las mismas operaciones de manera diferente. Utilizamos interfaces en Go para lograr el polimorfismo, definiendo conjuntos de métodos que deben ser implementados por otros tipos de datos.

A lo largo de esta sección, adquirimos un sólido conocimiento de la POO en Go y cómo utilizar las interfaces para lograr una programación más modular y flexible. Estos conceptos nos permiten escribir código más estructurado y escalable en Go. ¡Hemos explorado con éxito la programación orientada a objetos y las interfaces en Go!

[Materia de esta sección en GitHub](https://github.com/alexroel/curso-golang/blob/main/sections/07-poo-interfaces.md) y Tambien puedes descargar en PDF.

Sigue aprendiedo con los siguientes tutoriales sobre esta sección:

[Uso de métodos e interfaces en Go](https://learn.microsoft.com/es-es/training/modules/go-methods-interfaces/)
