# API Estudantes 👩‍🎓 👨‍🎓
<br>

<h2> Este foi o projeto que aprendi no Curso Golang do Zero </h2>
API para gerenciar os estudantes no Curso Golang do Zero.

<h2> Introdução </h2>
O objetivo da aplicação é criar um sistema para controlar os estudantes cadastrados e ativos na plataforma do curso. 

## Tecnologias utilizadas
* VS Code
* Golang (Go)
* Insomnia
* Echo v4
* GORM
* Zerolog

## Ferramentas
Foi instalado a extensão Go no VS Code com a versão 0.41.2

Foi instalado a extensão SQLite no VS Code com a versão 0.14.1

Foi instalado a extensão SQLite Viewer no VS Code com a versão 0.5.8

Foi instalado a extensão C/C++ no VS Code com a versão 1.20.5

Foi usado o programa <a href="https://insomnia.rest/download" target="_blank" > Insomnia </a> para testar as requisições das rotas simulando o Front-end.

Foi usado o framework web <a href="https://github.com/labstack/echo" target="_blank"> ECHO </a> na versão 4.

Foi usado a ferramenta de mapeamento relacional de objetos (ORM) <a href="https://gorm.io/docs/connecting_to_the_database.html" target="_blank"> GORM </a> com conexão ao banco de dados SQLite.

Foi usado o pacote <a href="https://github.com/rs/zerolog" target="_blank"> Zerolog </a> .

## Rotas

| Método | URL             | Descrição                                        |
| ------ | --------------  | -------------------------------------------------|
| GET    | /estudantes     | Listar todos os alunos                           |
| POST   | /estudantes     | Cadastrar um estudante                           |
| GET    | /estudantes/:id | Pegar a informação de um estudante específico    |
| PUT    | /estudantes/:id | Atualizar informações de um estudante específico |
| DELETE | /estudantes/:id | Deletar um estudante específico                  |


## Estrutura do estudante
- Nome
- CPF
- Email
- Idade
- Ativo
