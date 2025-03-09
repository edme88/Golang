## Go Web - Introducción (119)

1. Iniciamos el manejador
   `go mod init rps`
   rock-paper-scisor

2. Crear las routas y handlers

3. Ingresar a https://github.com/air-verse/air e instalar con
   `go install github.com/air-verse/air@latest`
   (para evitar tener que bajar y re-levantar el servidor con cada cambio)

Comprobar
`air -v`

4. Ejecutar la config inicial
   `air init`

5. Ejecutar el servidor con
   `air`

(Si no anda, ir a .profile o en .bashrc)

```bash
cd..
cd..
nano .profile
```

```bash
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

O sino en Windows revisar las variables de entorno

### Resumen (123)

En esta sección del curso de Go web, exploramos los fundamentos esenciales para crear un proyecto web utilizando el lenguaje de programación Go. Nos centramos en tres temas principales: la creación de un proyecto web, el enrutamiento y los manejadores, y el modo de depuración con Air.

Comenzamos estableciendo una estructura básica de carpetas y archivos para nuestro proyecto web y aprendimos a manejar las dependencias utilizando Go Modules. También nos familiarizamos con las mejores prácticas para la organización del código y la estructura del proyecto.

Luego, nos sumergimos en el enrutamiento y los manejadores en Go. Utilizamos el enrutador "mux" de Gorilla, un enrutador popular y potente, para dirigir las solicitudes entrantes a las funciones de manejo adecuadas. Aprendimos cómo configurar rutas y manipular las solicitudes y respuestas HTTP.

Finalmente, nos adentramos en el modo de depuración utilizando Air. Configuramos esta herramienta en nuestro proyecto y aprovechamos sus capacidades para facilitar el desarrollo y la corrección de errores. El modo de depuración nos permitió detectar y solucionar problemas de manera eficiente durante el desarrollo de nuestra aplicación web.

Al final de la sección, adquirimos los conocimientos necesarios para crear un proyecto web en Go, implementar enrutadores y manejar solicitudes entrantes. Además, aprendimos a utilizar el modo de depuración con Air para acelerar el proceso de desarrollo y solucionar rápidamente cualquier problema que surgiera.

En resumen, esta sección del curso nos proporcionó una base sólida para construir aplicaciones web utilizando Go. Aprendimos a crear proyectos, implementar enrutadores y manejar solicitudes, así como a utilizar el modo de depuración para un desarrollo eficiente. Con estos conocimientos, podemos aprovechar todo el potencial del desarrollo web con Go.

## TEMAS (124)

- Renderizar plantillas
- Enviar datos
- Declaración de variables
- Composicion de plantillas
- Uso de Must
- Enlaces a páginas

Usaremos https://pkg.go.dev/html/template

## Resumen (131)

En la sección de manejo de plantillas del curso de Go web, exploramos las capacidades de rendirizar plantillas en aplicaciones web. Aprendimos a cargar y rendirizar plantillas HTML utilizando el paquete html/template de Go. Aprendimos a integrar las plantillas en nuestro código y enviarlas al cliente para su visualización.

También aprendimos cómo enviar datos desde el código Go a las plantillas. Exploramos diferentes métodos para pasar variables, estructuras o mapas a las plantillas, permitiendo generar contenido dinámico.

En cuanto a la declaración de variables en las plantillas, aprendimos cómo declarar y utilizar variables, asignarles valores, realizar operaciones básicas y mostrar su contenido en la página generada.

Además, exploramos la composición de plantillas, donde aprendimos a dividir nuestras plantillas en partes reutilizables y combinarlas para construir páginas completas.

En cuanto al uso de template.Must, aprendimos a utilizar esta función para cargar y analizar nuestras plantillas de forma segura. Esto nos permitió manejar errores de manera más eficiente y garantizar que nuestras plantillas se cargaran correctamente.

Finalmente, aprendimos a crear enlaces dinámicos a diferentes páginas en nuestra aplicación web. Aprendimos cómo generar los enlaces en nuestras plantillas y cómo manejar las solicitudes correspondientes en nuestro código Go.

Con estos conocimientos, hemos adquirido habilidades para aprovechar al máximo el manejo de plantillas en nuestros proyectos de desarrollo web con Go. Ahora estamos preparados para crear páginas dinámicas y atractivas en nuestras aplicaciones web.

## TEMAS (132)

Archivos estáticos: CSS, Imágenes

- Configurar y Servir archivos estáticos
- Agregar CSS y Bootstrap y Js
