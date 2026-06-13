# Variables
SWAG_BIN=~/go/bin/swag
AIR_BIN=~/go/bin/air
MAIN_PATH=cmd/api/main.go

.PHONY: doc run build

# Generar documentación manualmente si es necesario
doc:
	$(SWAG_BIN) init -g $(MAIN_PATH) --parseDependency --dir ./

# Corre Air (Live-Reloading), el cual se encargará de compilar,
# generar la documentación y reiniciar el servidor en cada guardado.
run:
	$(AIR_BIN)

# Compilar el proyecto final para producción generando la doc antes
build: doc
	go build -o bin/api $(MAIN_PATH)