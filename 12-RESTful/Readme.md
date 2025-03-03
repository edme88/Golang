## API - RESTfull con Gin

- Crear proyecto
- Instalar Gin
- Obtener albums
- Crear album
- Obtener álbum específico

#### Diseño de los puntos finales del API

GET /albums - Devolver todos los albums en formato JSON
POST /albums - Agregar un nuevo album en JSON
GET /albums/:id - Obtener un album específico
PUT /albums/:id - Modificar un album específico

1. `go mod init vinyl-api` dentro de /vinyl-api
2. Ingresar a https://gin-gonic.com/docs/quickstart
3. Ejecutar `go get -u github.com/gin-gonic/gin`
4. Importar en el documento
5. Levantar el servidor
6. Crear struc para los ambums
7. Revisar doc https://developer.mozilla.org/es/docs/Web/HTTP/Status
8. Leer doc https://pkg.go.dev/net/http
9. Entrar en el navegador a `http://localhost:8080/albums`
10. Usar postman o `Thunder Client` (extension del VSC)

Lo guardaremos en la memoria en lugar de la BD para simplificar temporalmente

Resumen
En esta sección, aprendimos a desarrollar una API RESTful utilizando el lenguaje de programación Go y el framework Gin. Comenzamos creando un nuevo proyecto Go y configuramos el entorno de desarrollo necesario. Instalamos el framework Gin y configuramos la estructura básica del proyecto.

A continuación, creamos datos de ejemplo y exploramos cómo almacenarlos en memoria. Implementamos controladores para manejar las solicitudes GET y POST. El controlador de GET nos permitió devolver todos los elementos de la API en formato JSON, mientras que el controlador de POST nos permitió agregar nuevos elementos a la colección de datos.

Además, implementamos un controlador para manejar solicitudes GET específicas, donde utilizamos parámetros en la URL para identificar y devolver un elemento específico de la API.

Durante este proceso, aprendimos cómo utilizar las funciones y características proporcionadas por el framework Gin, como la serialización de JSON y la gestión de rutas.

Con estos conocimientos, hemos sido capaces de construir una API RESTful básica y comprender los fundamentos de Go y Gin en el contexto de la creación de una API.

[Materia de esta sección en GitHub](https://github.com/alexroel/curso-golang/blob/main/sections/11-api-rest-gin.md)

[Documentación de Go](https://go.dev/doc/tutorial/web-service-gin)
