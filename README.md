# X-Sky

It’s a website based on a fictional company, containing a homepage, registration page, and login page.

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

**1** - Install the necessary technologies.

 - Node JS: https://nodejs.org/en/download/
 - Golang: https://go.dev/dl/
 - Docker-compose: https://docs.docker.com/compose/install/

**2** - Clone the repository `git clone "https://github.com/GustavoHenriqueSchmitz/X-Sky.git"`

### Back-end

**1** - Open a terminal, navigate to the back-end folder of the project.

**2** - Run `sudo docker-compose up`, and wait for the container to initialize.

**3** - In another terminal, navigate to the same folder and run `go run main.go`

**4** Code to create the database table.

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

**1** - Open a terminal, navigate to the front-end folder of the project.

**2** - Run `node init_server.js`

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
