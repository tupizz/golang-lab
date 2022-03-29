```bash
go test (dentro do package)
``` 

roda em todos pacotes

```bash
go test ./...
```

roda teste em todos pacotes em modo verbose

```bash
go test -v ./...
```

roda teste em todos pacotes extraindo o coverage

```bash
go test --cover ./... 
```

roda teste em todos pacotes extraindo o profile coverage em um arquivo

```bash
go test --coverprofile cobertura.txt  ./.... 
```


lemos o arquivo de cobertura de uma forma mais semantica para nós

```bash
go tool cover --func=cobertura.txt
go tool cover --html=cobertura.txt # esse é muito foda!!!!
```