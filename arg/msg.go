package arg

var (
	// HELP for helpe
	HELP = `Parametros disponiveis
	ALL 	Vai fazer backup de toda a base de dados local
	SOME	Vai Fazer backup das base de dados configuradas no arquivo config.yml
	HELP 	Vai mostrar esta mensagem :-)
	`
	// ERROR Mensagem de error caso um parametro incorreto for informado
	ERROR = "%s Não é um param valido\n%s"

	// EMPTY mensagem pra prametro vazio
	EMPTY = "Você precisa informar 1 parametro %s"
)
