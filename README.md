# customer-api

Estructura actual de directorios y archivos del proyecto.

```text
customer-api/
├── .air.toml
├── .gitignore
├── Makefile
├── go.mod
├── go.sum
├── README.md
├── cmd/
│   └── api/
│       └── main.go
├── docs/
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── internal/
│   ├── customer/
│   │   ├── dto.go
│   │   ├── entity.go
│   │   ├── handler.go
│   │   ├── repository.go
│   │   └── service.go
│   ├── phone/
│   └── shared/
│       └── errors.go
├── platform/
│   └── storage/
│       ├── database.go
│       └── generic_repository.go
└── tmp/
    ├── build-errors.log
    └── main
```
