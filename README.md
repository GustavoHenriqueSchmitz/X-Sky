# X-Sky

Este projeto é um desafio fornecido pela **ELITI**

É um site baseado em uma empresa fictícia.
Contendo uma página inicial, cadastro e login.

## Estrutura

### Site

* Página principal
* Página de login
* Página de cadastro

### Tecnologias

* HTML5
* CSS3
* Java Script
* Golang
* Node JS 
* PostgresSQL
* Docker-compose

## Guia de inicialização do projeto 

**1** - Instale as Tecnologias necessárias.

 * Node JS: https://nodejs.org/en/download/
 * Golang: https://go.dev/dl/
 * Docker-compose: https://docs.docker.com/compose/install/

**2** - Clone o repositório `git clone "https://github.com/GustavoHenriqueSchmitz/X-Sky.git"`

### Back-end

**1** - Abra um terminal, navegue até a pasta **back-end** do projeto.

**2** - Execute `sudo docker-compose up`, e espere o container ser inicializado.

**3** - Em outro terminal, navegue até a mesma pasta e execute `go run main.go`

### Front-end

**1** - Abra um terminal, navegue até a pasta **front-end** do projeto.

**2** - E execute `node init_server.js`

## Informações para uso

 * Front-end inicializado na porta 3000.
 * Back-end inicializado na porta 5000.
 * Banco de dados Postgres inicializado na porta 5432.
 * Usuário do banco de dados: xsky
 * Senha do banco de dados: xsky
 * Nome do banco de dados: X-Sky
