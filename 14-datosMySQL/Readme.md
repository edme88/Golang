107

- Introducción a MySQL
- Conexión a MySQL desde Go
- Uso de variables de entorno
- Realizar als consultas CRUD
- Proyecto interactivo con base de datos MySQL

1. Verificar si tenemos mySQL descargado
   `mysql --version`

Sino, Descargar mySQL
https://dev.mysql.com/downloads/mysql/

MySQL Community Server
Es gratuito pero requiere una cuenta

Para linux:
`sudo apt install mysql-server`

Luego ejecutar:
`man mysql`

GoLang123#
MySQL92

Reference Manual https://dev.mysql.com/doc/refman/9.2/en/

`show database:`

ctrol+l

CREATE USER 'agus'@'localhost' IDENTIFIED BY 'laContraseña';

GRANT ALL PRIVILEGES ON _._ TO 'agus'@'localhost' WITH GRANT OPTION;

exit;

`sudo mysql -u agus -p`

show databases;

https://dev.mysql.com/doc/refman/8.4/en/tutorial.html

2. Ingresar a MySQL 9.2 Command Line Cli

`show databases;` //para ver cuales vienen por defecto
+--------------------+
| Database |
+--------------------+
| information_schema |
| mysql |
| performance_schema |
| sys |
+--------------------+

Para crear una database:
`create database if not exist db_contacts;`

Para usar:
`use db_contacts;`

Para listar las tablas:
`show tables;`

Para crear una tabla:

```
create table if not exists contact(
id int auto_increment primary key,
name varchar(255) not null,
email varchar(255),
phone varchar(20)
);
```

Para listar las tablas:
show tables

para ver que registros tiene:
`select * from contact;`

Para agregar datos:

```
insert into contact (name, email, phone) values ('Juan Perez', 'juan@perez.com', '3513892222'),('Maria Gomez', 'maria@example.com','987527645'),('Carlos Lopez', null, '663453');
```

(fin 110)

https://pkg.go.dev/database/sql
https://github.com/go-sql-driver/mysql

https://go.dev/wiki/SQLDrivers

Primero:
go mod init datosMySQL

Para instalar el driver
go get -u github.com/go-sql-driver/mysql

111. Variables de entorno

https://pkg.go.dev/os

https://pkg.go.dev/github.com/joho/godotenv //Para trabajar con variables de entorno

https://github.com/joho/godotenv -> https://github.com/joho/godotenv

instalar usando
`go get github.com/joho/godotenv`

Resumen
En esta sección del curso de Golang, exploramos cómo trabajar con una base de datos MySQL para la persistencia de datos en nuestras aplicaciones. A lo largo de esta sección, aprendimos los fundamentos necesarios para interactuar con una base de datos MySQL desde una aplicación Go.

Comenzamos instalando y configurando MySQL en nuestro sistema para prepararnos para trabajar con bases de datos. Luego, creamos una base de datos utilizando comandos SQL básicos y establecimos una conexión entre nuestra aplicación Go y la base de datos MySQL para poder realizar operaciones de lectura y escritura.

Utilizamos la biblioteca godotenv para cargar variables de entorno desde un archivo .env, lo que nos permitió gestionar de manera segura y conveniente la configuración de la conexión a la base de datos.

Implementamos la funcionalidad para listar todos los contactos almacenados en la base de datos y aprendimos a recuperar un contacto específico utilizando su identificador único. También exploramos cómo registrar un nuevo contacto, actualizar la información de un contacto existente y eliminar un contacto de la base de datos cuando fuera necesario.

Finalmente, concluimos la sección construyendo un menú de navegación interactivo utilizando el paquete bufio, que permitió al usuario realizar todas las operaciones CRUD de manera intuitiva.

Al finalizar esta sección, estamos bien equipados para trabajar con bases de datos MySQL en nuestros proyectos Go y gestionar eficazmente la persistencia de datos.

[Materia de esta sección en GitHub](https://github.com/alexroel/curso-golang/blob/main/sections/13-mysql-go-database.md)
