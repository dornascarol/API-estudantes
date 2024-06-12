# API Estudantes 👩‍🎓 👨‍🎓
API para gerenciar os estudantes do curso Golang do Zero

## Rotas

| Método | URL             | Descrição                                        |
| ------ | --------------  | -------------------------------------------------|
| GET    | /estudantes     | Listar todos os alunos                           |
| POST   | /estudantes     | Cadastrar um estudante                           |
| GET    | /estudantes/:id | Pegar a informação de um estudante específico    |
| PUT    | /estudantes/:id | Atualizar informações de um estudante específico |
| DELETE | /estudantes/:id | Deletar um estudante específico                  |


## Estrutura do estudante:
- Nome
- CPF
- Email
- Idade
- Ativo
