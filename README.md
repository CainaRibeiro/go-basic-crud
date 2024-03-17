# CRUD de Usuários
Este é um exemplo básico de um CRUD (Create, Read, Update, Delete) para usuários em uma aplicação Go usando MongoDB como banco de dados. O CRUD possui as seguintes operações:

- Listar todos os usuários
-  Criar um novo usuário
- Remover um usuário
- Atualizar o nome de um usuário
## Listar todos os usuários
### GET /users
Esta rota retorna todos os usuários cadastrados no banco de dados.

### Criar um novo usuário

#### POST /users
Esta rota permite a criação de um novo usuário. O corpo da requisição deve conter os dados do usuário a ser criado.

Exemplo de corpo da requisição:

```json
{
    "name": "Nome do Usuário",
    "age": 30,
    "address": "rua fulano - bairro/estado",
    "cpf": "12312312312",
    "ActiveTransactions": 2
}
```
### Remover um usuário

#### DELETE /users
Esta rota permite remover um usuário do banco de dados. O corpo da requisição deve conter o CPF do usuário a ser removido.

Exemplo de corpo da requisição:

```json
{
    "cpf": "12345678910"
}
```
### Atualizar o nome de um usuário

#### PUT /users/{id}
Esta rota permite atualizar o nome de um usuário existente no banco de dados. O id na URL deve corresponder ao ID do usuário a ser atualizado. O corpo da requisição deve conter o novo nome do usuário.

Exemplo de corpo da requisição:

```json
{
    "name": "Novo Nome do Usuário"
}
```
Executando a Aplicação
Para executar a aplicação, você precisa configurar um servidor MongoDB e configurar as variáveis de ambiente necessárias, como a string de conexão com o banco de dados.

Após configurar o ambiente, execute o seguinte comando para iniciar a aplicação:

```go
go run main.go
```
Isso iniciará o servidor na porta padrão 4000. Você pode acessar as rotas acima utilizando um cliente HTTP como o cURL ou o Postman.
