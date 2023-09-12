# JUS BRASIL GATEWAY

## RESUMO

* Gateway para requisicoes aos sites de primeira e segunda instância.
  * TJAL
    * 1º grau - <https://www2.tjal.jus.br/cpopg/open.do>
    * 2º grau - <https://www2.tjal.jus.br/cposg5/open.do>
  * TJCE
    * 1º grau - <https://esaj.tjce.jus.br/cpopg/open.do>
    * 2º grau - <https://esaj.tjce.jus.br/cposg5/open.do>

## Execucão

### Requirementos

* Golang >= 1.21.1
* Docker Engine ou Docker Desktop

### Instalar dependências

```bash
go mod tidy
```

### Gerar Binários

```bash
make build
```

### Run

```bash
go run cmd/main.go
```

### Docker Compose

```bash
# iniciar
$ docker-compose up
```

```bash
# encerrar
$ docker-compose down
```

### Rotas

Por default, a aplicacao roda na porta `8080`

* POST

```bash
curl --location 'http://localhost:8080/' \
--header 'Content-Type: application/json' \
--data '{
    "url": "https://www2.tjal.jus.br/cpopg/search.do?cbPesquisa=NUMPROC&numeroDigitoAnoUnificado=0710802-55.2018&foroNumeroUnificado=0001&dadosConsulta.valorConsultaNuUnificado=0710802-55.2018.8.02.0001&dadosConsulta.valorConsultaNuUnificado=UNIFICADO&dadosConsulta.tipoNuProcesso=UNIFICADO"
}'
```
