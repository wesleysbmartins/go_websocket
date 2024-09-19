# WebSocket Golang Project
[![My Skills](https://skillicons.dev/icons?i=golang,docker)](https://skillicons.dev)

Este projeto é uma aplicação Golang que implementa um servidor WebSocket. Abaixo estão as instruções sobre como se conectar ao WebSocket e como acessar a documentação Swagger da API.

<details>
    <summary>O que é WebSocket?</summary>
WebSocket é um protocolo de comunicação que permite a comunicação bidirecional em tempo real entre um cliente e um servidor pela web. Diferente do HTTP, que é unidirecional e requer uma nova solicitação para cada troca de dados, o WebSocket mantém uma conexão aberta e persistente que permite a troca de dados em tempo real.
</details>

## Executando com Docker
Execute no seu terminal os seguintes comandos para rodar a aplicação:
```bash
docker build -t go_websocket . 
```
Após gerar o build, execute o container:
```bash
docker run -p 8080:8080 go_websocket
```

## Conectando ao WebSocket

### Endpoint

O WebSocket está disponível no seguinte endpoint:

### Parâmetros de Conexão

Ao conectar-se ao WebSocket, você deve fornecer dois parâmetros na query string:

- **room**: O nome da sala à qual deseja se conectar (string).
- **name**: O nome do usuário (string).

### Exemplo de Conexão

Você pode usar uma ferramenta como **Postman**, **Insomnia**.

Swagger: http://{host}:{porta}/swagger/index.html


Esse `README.md` cobre como conectar ao WebSocket, acessar a documentação Swagger e executar o projeto. Sinta-se à vontade para ajustar qualquer parte de acordo com as necessidades específicas do seu projeto.

