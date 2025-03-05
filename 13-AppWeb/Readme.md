## Escribe una aplicación web en go

- Crea rpalicación web
- El paquete net/http
- El paquete html/template
- Manejo de plantillas
- Manejo de errores
- Usar funciones de clausura

Resumen
En esta sección, exploramos cómo escribir aplicaciones web en Go con una serie de temas interesantes. Comenzamos presentando el paquete net/http, que es la base para construir aplicaciones web en Go.

Utilizamos el paquete net/http para servir páginas wiki y aprendimos a configurar un servidor HTTP básico. Aprendimos cómo manejar solicitudes y respuestas utilizando los manejadores proporcionados por este paquete.

Luego nos adentramos en el paquete html/template, que nos permitió generar y mostrar contenido dinámico en nuestras páginas web. Aprendimos a cargar y utilizar plantillas HTML para renderizar contenido de manera flexible.

Abordamos el problema del manejo de páginas inexistentes y aprendimos a redirigir al usuario a una página de edición cuando intenta acceder a una página que no existe.

Además, exploramos cómo guardar las páginas editadas en el servidor. Implementamos un formulario de edición y aprendimos a guardar los cambios realizados por el usuario en un archivo en el servidor.

También abordamos el tema del manejo de errores de manera adecuada. Aprendimos a validar y manejar los errores que pueden ocurrir durante la ejecución de nuestra aplicación web.

Optimizamos nuestra aplicación mediante la caché de plantillas. Aprendimos a compilar y almacenar en memoria las plantillas HTML para mejorar el rendimiento de nuestra aplicación.

Por último, introdujimos los literales de funciones y clausuras. Aprendimos a utilizar estas características de Go para abstraer la funcionalidad repetitiva y simplificar nuestros controladores de manejo de solicitudes.

[Materia de esta sección en GitHub](https://github.com/alexroel/curso-golang/blob/main/sections/12-writing-web-aplications.md)

[Documentación de Go](https://go.dev/doc/articles/wiki/)
