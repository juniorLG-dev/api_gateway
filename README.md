# API-Gateway
Uma API Gateway escrita em Golang fácil de se usar.

**Atuais funcionalidades:**
- Criação de serviço
- Registro dos endpoints/rotas
- Roteamento dinâmico

## Criação de um serviço:
Um serviço é um contexto que sua API Gateway irá conter. O objetivo é fazer menção ao contexto da sua API Concreta. 
Exemplo: se minha API Concreta trabalha em cima de um contexto de login, o nome do serviço poderá ser "login". 

---
> OBS: o nome do serviço servirá de prefixo durante a etapa de roteamento dinâmico.
---

### Rotas:
* Endpoint: **/service**
* Método: **POST**
---

### Request:

**DATA:**
| Campo | Tipo | Obrigatório | Descrição
|:------|:-----|:------------|:---------------------|
|name   |string|    SIM      | nome do seu serviço

**Exemplo:**
```json
{"name": "example"}
```
**Exemplo de requisição:**
```sh
curl -i -X POST -d '{"name": "example"}' localhost:8080/service
```

### Response:

**Status:**
* `201` - Created
* `400` - Bad Request
* `500` - Internal Error

**Body:**
```json
{"message":"service created"}
```
**Header:**
```
HTTP/1.1 201 Created
Authorization: <seu_token_JWT>
Content-Type: application/json; charset=utf-8
```

## Registro dos endpoints:
Os endpoints são criados de maneira diferente do serviço. Para a criação é necessário que o usuário mande um arquivo JSON contendo as informações necessárias. 

### Rotas:
- Criar endpoint:
  * Endpoint: **/routes**
  * Método: **POST**
- Buscar endpoints:
  * Endpoint: **/routes**
  * Método: GET

---
 
### Request [POST] /routes:
**Data:**
|  Campo      |  Tipo  |  Obrigatório  |  Descrição  |
|:------------|:-------|:--------------|:------------|
| path        | string |      SIM      | Seu endpoint
| service_url | string |      SIM      | URL da sua api (com http ou https)
| method      | string |      SIM      | Método do seu endpoint

> Seu arquivo json deverá seguir essa estrutura

**Exemplo:**
```json
{
  "path": "hello-post",
  "service_url": "http://localhost:3333/",
  "method": "POST"
}

{
  "path": "hello-get",
  "service_url": "http://localhost:4444/",
  "method": "GET"
}
```

**Header:**
```
Authorization: Bearer <seu_token_JWT>
Content-Type: multipart/form-data
```

**Exemplo de requisição:**
```sh
curl -X POST -H "Authorization: Bearer <seu_token_JWT>" -H "Content-Type: multipart/form-data" -F file=@/caminho/do/arquivo.json localhost:8080/routes
```

### Response [POST] /routes:

**Status:**
* `201` - Created
* `400` - Bad Request
* `500` - Internal Error

**Body:**
```json
{"message":"route created"}
```
**Header:**
```
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
```

### Request [GET] /routes:

**Header**
```
Authorization: Bearer <seu_token_JWT>
```

**Exemplo de requisição:**
```sh
curl -H "Authorization: Bearer <seu_token_JWT> localhost:8080/routes" 
```

### Response [GET] /routes:

**Status:**
* `200`- OK
* `400` - Bad Request
* `404` - Not Found
* `500` - Internal Error

**Body:**
```json
[
    {
        "id": "8d71e5ee-0a17-45e3-afb7-3b29f43b210e",
        "path": "hello",
        "service_url": "http://localhost:3333/",
        "method": "POST"
    },
    {
        "id": "12c09690-cc1e-469f-8e20-9b7fdcc09fd2",
        "path": "greeting",
        "service_url": "http://localhost:3333/",
        "method": "GET"
    }
]
```
  
**Header:**
```
Content-Type: application/json; charset=utf-8
Content-Length: 235
```

## Roteamento dinâmico
O roteamento dinâmico é a etapa na qual o usuário vai utilizar o seu serviço criado pra enviar as requisições pra diferentes rotas e servidores de um mesmo contexto de API. 

### Rotas:
Como as rotas são dinâmicas, todas as rotas vão ser construídas a partir do seu serviço (contexto) + rota que será acessada. 

**Exemplo:**
`localhost:8080/service/route`
* service = meu serviço que criei nas etapas anteriores
* route = rota que desejo acessar relacionada ao meu serviço

> OBS: o usuário deverá realizar a requisição com o método, corpo e cabeçalho que aquela determinada rota solicita.

**Exemplo de requisição:**
```sh
curl -s -X POST -d '{"name": "John Doe"}' localhost:8080/test/register
```

### Response:
A API não tem umaresposta para suas rotas, ela encaminha a mensagem retornada pelas suas APIs clientes. 

> OBS: a API pode retornar uma mensagem de erro caso a requisição pras APIs clientes falhem. 
