# Índice

155. Introducción
156. Temas de la Sección
157. Crear Base de Datos
158. Conexión a MySQL
159. Ping
160. Crear Tabla
161. Verificar si una tabla existe

# Instalación

Ingresar a https://dev.mysql.com/downloads/mysql/ y descargar el más completo
Ingresar a https://go.dev/wiki/SQLDrivers

## 156 Temas de esta Sección

En esta sección veremos como trabajar con base de datos específicamente con el gestor de base de datos que MySQL. En esta sección estaremos aprendiendo lo siguiente:

- Crear una base de datos en MySQL y luego conectarse desde Go
- Haremos todas las consultas desde de Go hacia base de datos
- Desde crear una tabla y verificar todo
- También realizar todo el CRUD

Recursos de esta clase: [Package en Go](https://pkg.go.dev/database/sql)

## MySQL

Abrir MySQL 9.2 Command Line Cli ó

```bash
mysql -h localhost -u root -p
```

GoLang123#

### Comandos básicos de MySQL

Listar Bases de datos:

```bash
show databases;
```

Crear una base de datos:

```bash
create database goweb_db;
```

```bash
use goweb_db;
```

### En go...

1. Ir a la carpeta correspondiente y hacer:

```
go mod init gomysql
```

2. Usar el comando:

```
go get -u github.com/go-sql-driver/mysql
```

## 158

Creamos carpeta de _db_ con el archivo _database.go_ con las funciones de connect y close.
Probamos que funcione.

## 159. Ping

Verificar que la conexión se mantiene.

### 160. Crear tabla

Creamos la carpeta _models_ con un archivo _users.go_.

Listar tablas:

```
show tables;
```

```
show columns from users;
```

Usar comando:

```
show tables like 'users';
```
