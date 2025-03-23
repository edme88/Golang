package models

type User struct{
	Id 		 int
	Username string
	Password string
	Email	 string
}

const UserSchema string = `CREATE TABLE users (
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(60) NOT NULL,
	email VARCHAR(100),
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)`