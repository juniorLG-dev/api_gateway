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
  * Endpoint: **/routes/add**
  * Método: **POST**
- Buscar endpoints:
  * Endpoint: **/routes**
  * Método: GET

---
 
### Request /routes/add:
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
curl -X POST -H "Authorization: Bearer <seu_token_JWT>" -H "Content-Type: multipart/form-data" -F file=@/caminho/do/arquivo.json localhost:8080/routes/add
```

### Response /routes/add

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




  
