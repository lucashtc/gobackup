# Gobackup
[![Go Report Card](https://goreportcard.com/badge/github.com/lucashtc/gobackup)](https://goreportcard.com/report/github.com/lucashtc/gobackup)

Ferramenta para criação de backups de Banco MySQL localmente.

## Em Desenvolvimento

## Installation

```bash
go get github.com/lucashtc/gobackup 
```

## Usage

```bash
cd github.com/lucashtc/gobackup 

go run gobackup.go -l /backup/ -u userDatabase -p password
#OR 
go build 
gobackup -l /backup/ -u userDatabase -p password
```

## License
[MIT](https://github.com/lucashtc/gobackup/blob/master/LICENSE)
