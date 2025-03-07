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
