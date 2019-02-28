// Package mysql resposanvel por pegar nome das bases, tabelas e functions/procedures no servidor mysql
// Vai pegar todas as bases do servidor local
package mysql

// package responsável por obter informações do banco, como nome das bases e tabelas nas bases

// Interface
//_ "github.com/go-sql-driver/mysql"

// GetData function get name databse
// retorna um array com essas informações
/*func GetData(cf *DataBase) ([]DataBase, error) {

	base, err := GetDatabase(cf)
	if err != nil {
		return []DataBase{}, err
	}

	//Alimenta cada posição do slice indexado com name dos schemas, com nome das tabelas e procedures/functions
	db := make([]DataBase, len(base))

	for i, v := range base {
		db[i].Name = v
	}

	return db, nil

}
*/
