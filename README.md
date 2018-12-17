# Gobackup
## Em Desenvolvimento

[![Go Report Card](https://goreportcard.com/badge/github.com/lucashtc/gobackup)](https://goreportcard.com/report/github.com/lucashtc/gobackup)

** Ferramenta para criação de backups de Banco MySQL localmente

## Instalação
go get github.com/lucashtc/gobackup
cd github.com/lucashtc/gobackup
go run gobackup.go -l /backup/ -u userDatabase -p password ALL
OR 
go build gobackup -l /backup/ -u userDatabase -p password ALL

## Next Functions
Dump only one Database
