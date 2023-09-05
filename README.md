# chronos
Chronos is a web application for SISAE

## Quick start
Download it:
```
git clone https://github.com/JotaEspig/chronos
cd chronos
```
Without .env file:
```
CHRONOS_PORT=8080 CHRONOS_ROOT_DIR=. go run main.go
```
With .env file:
```
echo -e "CHRONOS_PORT=8080\nCHRONOS_ROOT_DIR=." > .env
go run main.go
```
