# API Estudantes 👩‍🎓 👨‍🎓
<br>

<h2> Este foi o projeto que aprendi no Curso Golang do Zero </h2>
API para gerenciar os estudantes cadastrados no Curso Golang do Zero.

## Tecnologias utilizadas
* VS Code
* Golang (Go)
* Echo
* GORM
* SQLite


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
