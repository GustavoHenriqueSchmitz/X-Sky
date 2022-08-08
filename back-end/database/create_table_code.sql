/* Code for create the users table in database*/

create table users (
	id serial primary key,
	name varchar(255) not null,
	last_name varchar(255) not null,
	email varchar(255) unique not null,
	password varchar(255) not null,
	area_code varchar(2),
	phone varchar(25),
	perfil_photo text
);
