# X-Sky

It’s a website based on a fictional company, containing a homepage, registration page, and login page.

## Project Video
[![Xsky System Demo](https://img.youtube.com/vi/Z2GD6-xEVr0/0.jpg)](https://www.youtube.com/watch?v=Z2GD6-xEVr0)

## Structure

### Website

- Main page
- Login page
- Registration page

### Technologies

- HTML5
- CSS3
- Java Script
- Golang
- Node.js 
- PostgresSQL
- Docker Compose

## Project Startup Guide
**1** - Install the necessary technologies:

 - [Node.js V16.17.1](https://nodejs.org/en/download)
 - [Golang V1.18.1](https://go.dev/dl)
 - [Docker-Compose](https://docs.docker.com/compose/install)

**2** - Clone the repository:
```
git clone "https://github.com/GustavoHenriqueSchmitz/X-Sky.git"
```

### Back-end
**1** - Open a terminal, navigate to the `back-end` directory of the project and run:
```
docker-compose up
```

**2** - Open another terminal, navigate to the same directory and run:
```
go run main.go
```

**3** To create the database table, use the following SQL command:
```
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
```

### Front-end
**1** - Open a terminal, navigate to the `front-end` directory of the project and run:
```
npm i
```

**2** - After that, run:
```
node init_server.js
```

## Usage Information

- Front-end initialized on port 3000.
- Back-end initialized on port 5000.
- PostgreSQL database initialized on port 5432.
- Database user: xsky
- Database password: xsky
- Database name: X-Sky

## Author
**Gustavo Henrique Schmitz**

**Linkedin:** https://www.linkedin.com/in/gustavo-henrique-schmitz  
**Portfolio:** https://gustavohenriqueschmitz.com  
**Email:** gustavohenriqueschmitz568@gmail.com  
