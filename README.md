# Go-TestStress

Script para execução de testes de sobrecargas em endpoints.

# Execução via Docker

1. Realize o build do arquivo dockerfile na raiz do projeto com o comando a baixo: 

```bash
docker build -t go-teststress .
```

2. Realize a execução do container com o comando abaixo, adicionando nos devidos campos a URL, número de requisições a serem realizadas e a quantidade de concorrência desejada.

```bash
docker run --rm go-teststress --url <URL_DO_SERVIÇO> --requests <NÚMERO_DE_REQUESTS> --concurrency <CHAMADAS_SIMULTÂNEAS>
```
