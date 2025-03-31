# Упражнение 3

Простой веб-сервер с конфигурацией через env-переменные.

Кодген для openapi
```shell
oapi-codegen -config configs/oapi-gen-cfg.yaml api/openapi.json
```

запуск сервера
```shell
port=8080 host=0.0.0.0 greeting=привет go run .
```
