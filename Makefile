# Variables
SWAG_BIN=~/go/bin/swag
MAIN_PATH=cmd/api/main.go

.PHONY: doc run build

# Generar documentación
doc:
	$(SWAG_BIN) init -g $(MAIN_PATH) --parseDependency --dir ./

# Correr la aplicación generando la doc antes
run: doc
	go run $(MAIN_PATH)

# Compilar el proyecto generando la doc antes
build: doc
	go build -o bin/api $(MAIN_PATH)