# **API de gerenciamento de Books em GoLang**

# **Tecnologias Utilizadas**

> **GoLang** - utilizado por conta da segurança e performance </br>
> **MySQL** - usado por conta da consistência, alta performance, confiabilidade e é fácil de usar</br>
> **JWT** - usado na autenticação dos usuários e validações em outras transações</br>


# Como rodar a API

- Para rodar a API, primeiramente, é necessário ter o Docker instalado em sua máquina, e executar o comando:</br>
`docker-compose up --build -V`
- Este comando criara 2 containers, um com a API em GoLang rodando na porta 5000 e outro do banco de dados rodando na porta 3306

# Criar as tabelas no Banco de Dados
- Como a nossa API não possue migrations nem um ORM é necessário criar as tabelas por comando, você pode utilizar o próprio terminal do mysql acessado o seu container:
`docker exec -it mysql mysql -uroot -p`
- Por padrão foi definido a senha: **userroot**

- Apos isso é necessário criar as tabelas executando:
```jsx
USE devbookdb;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    email varchar(50) not null unique,
    password varchar(180) not null,
    createdAt timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE books(
    id int auto_increment,
    title varchar(50) not null unique,
    category varchar(60) not null,
    synopsis varchar(250) not null,
    author_id int not null,
    createdAt timestamp default current_timestamp(),
  
  	PRIMARY KEY (id),
    CONSTRAINT FK_AuthorBook FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=INNODB;
```

* Você também pode realizar a criação das tabelas por meio de ferramentas como o Beekeeper, dbeaver,...


# Como utilizar a API

- Você pode utilizar uma ferramenta como o Postman ou Insomnia para fazer as requisições HTTP

## Rotas da API
## 1) [POST] `http://localhost:5000/users`
-criação do usuario
![image](https://user-images.githubusercontent.com/66275588/188914000-ccbbbe1b-fd15-426f-9296-0785b2745ad4.png)

## 2) [POST] `http://localhost:5000/signIn`
- Login do usuário: recebe o token JWT que será utilizado em todas as rotas autenticadas
![image](https://user-images.githubusercontent.com/66275588/188914746-baf172a4-0b87-49b9-848f-0016c3f1c453.png)

## 3) [GET] `http://localhost:5000/users`
- Visualizar todos os usuários cadastrados: não requer token JWT
![image](https://user-images.githubusercontent.com/66275588/188915335-b717776c-f97b-4481-943a-f614623ebd2e.png)

## 4) [GET] `http://localhost:5000/users/id`
- Visualizar o proprio usuário: requer token JWT
![image](https://user-images.githubusercontent.com/66275588/188917630-a2709b12-7fba-4c1e-8f45-00949f01647b.png)

## 5) [PUT] `http://localhost:5000/users/id`
- Atualizar o proprio usuário: requer token JWT
![image](https://user-images.githubusercontent.com/66275588/188918704-1f6f6913-412f-461b-b3f0-50dc00f04b58.png)

## 6) [DELETE] `http://localhost:5000/users/id`
- Deletar o proprio usuário: requer token JWT
![image](https://user-images.githubusercontent.com/66275588/188918183-a7491f22-e360-4c42-b8d3-4e0faf3b3c93.png)


## 7) [POST] `http://localhost:5000/books`
-criação do BOOK: requer token JWT
![image](https://user-images.githubusercontent.com/66275588/188920063-7342b80f-cce3-48db-81ac-ed69433cade7.png)

## 8) [GET] `http://localhost:5000/books`
- Visualizar todos os BOOK cadastrados: não requer token JWT
![image](https://user-images.githubusercontent.com/66275588/188920294-d4ba4543-8146-4433-b6f6-97e54038f2ba.png)

## 9) [GET] `http://localhost:5000/books/id`
- Visualizar um BOOK: requer token JWT
![image](https://user-images.githubusercontent.com/66275588/188920478-d52242a2-27c6-422b-80dc-a9e3d0973562.png)

## 10) [PUT] `http://localhost:5000/books/id`
- Atualizar um BOOK: requer token JWT
![image](https://user-images.githubusercontent.com/66275588/188920723-98840914-c76a-43df-a867-a9c11d0c5550.png)

## 11) [DELETE] `http://localhost:5000/books/id`
- Deletar um BOOK: requer token JWT
![image](https://user-images.githubusercontent.com/66275588/188920890-74389afe-91cf-4eee-b07d-0a74c401a748.png)

**Feito por Willian Brandão Mendonça**
