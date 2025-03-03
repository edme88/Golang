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
