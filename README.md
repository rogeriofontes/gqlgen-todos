mutation {
  createTodo(input: { text: "Ler um livro", userId: "123" }) {
    id
    text
    done
    userID
  }
}
---
{
  "data": {
    "createTodo": {
      "id": "T94",
      "text": "Ler um livro 2",
      "done": false,
      "userID": "123"
    }
  }
}
===
query {
  todos {
    id
    text
    done
    userID
  }
}
---
{
  "data": {
    "todos": [
      {
        "id": "T7",
        "text": "Ler um livro",
        "done": false,
        "userID": "123"
      }
    ]
  }
}
===
Resumo dos Comandos
Comando	                                          O que faz?
go run github.com/99designs/gqlgen generate 	    Gera modelos e resolvers automaticamente
go mod tidy	                                      Atualiza e limpa dependências
rm -rf graph/generated graph/model/models_gen.go	Remove arquivos gerados para recriação
go run server.go	                                Inicia o servidor GraphQL
curl -X POST http://localhost:8080/query ...	    Testa a API GraphQL via terminal
lsof -i :8080 ou `netstat -ano	                  findstr :8080`
kill -9 <PID> ou taskkill /F /PID <PID>	          Libera a porta se estiver bloqueada
http://localhost:8080/	                          Abre o Playground GraphQL no navegador

====
https://gqlgen.com/getting-started/
====