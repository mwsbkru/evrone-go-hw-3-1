# Упражнение 3

Простой веб-сервер с конфигурацией через env-переменные.

Кодген для openapi
```shell
oapi-codegen -config configs/oapi-gen-cfg.yaml api/openapi.json
```

запуск сервера
```shell
PORT=8081 HOST=0.0.0.0 GREETING=привет go run .
```
