## Как запустить
Если у вас есть docker, то вначале надо сбилдить ```docker build -t opa-app . ```, после запустить с примонтированной папкой ```docker run -p 8080:8080 \-v $(pwd)/policy:/app/policy \opa-app```
Если у вас на компьютере есть go, то вы можете сбилдить ```go build -o cmd/main.go``` и запустить ```./app```. Или сразу запустить ```go run cmd/main.go```.
Работает hot reload, вы можете изменить конфигурацию в auth.rego, программа вычитает её и применит, но спустя время (обновление раз в 5 секунд).

## Как протестить c помощью curl
``` 
curl -X GET http://localhost:8080/resource   -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ1c2VyLTEyMyIsIm5hbWUiOiJUZXN0IFVzZXIiLCJlbWFpbCI6InRlc3RAZXhhbXBsZS5jb20iLCJyb2xlcyI6WyJyZWFkZXIiXX0.signature"
```
``` 
curl -X POST http://localhost:8080/resource \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ1c2VyLTQ1NiIsIm5hbWUiOiJBZG1pbiIsImVtYWlsIjoiYWRtaW5AZXhhbXBsZS5jb20iLCJyb2xlcyI6WyJhZG1pbiJdfQ.signature"
```

## Как запустить тесты
```go test ./...```