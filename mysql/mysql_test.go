package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConn(t *testing.T) {
	conf := my{
		User:     "root",
		Password: "",
	}
	DB, err := conf.Conn()

	err = DB.Ping()
	if err != nil {
		t.Error("PING - Error ao conectar == > ", err)
	}
}

func TestGetDatabase(t *testing.T) {
	conf := my{
		User:     "root",
		Password: "",
	}
	DB, err := conf.Conn()

	err = DB.Ping()
	if err != nil {
		t.Error("PING - Error ao conectar == > ", err)
	}

	database, err := conf.GetDatabase()
	if err != nil {
		t.Error("Falha >>>>", err)
	}

	assert.Contains(t, database, "mysql")

}
