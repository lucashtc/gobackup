package mysql

import (
	"testing"
)

func TestConn(t *testing.T) {
	conf := my{
		User:     "root",
		Password: "",
	}
	DB, err := conf.Conn()

	err = DB.Ping()
	if err != nil {
		t.Log("PING - Error ao conectar == > ", err)
	}
}
