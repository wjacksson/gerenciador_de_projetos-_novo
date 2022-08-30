package postgres

import (
	"database/sql"
	"fmt"
	modelData "gerenciadorDeProjetos/infra/equipe/model"
)

type DBEquipe struct {
	DB *sql.DB
}

func (postgres *DBEquipe) NovaEquipe(req *modelData.Equipe) {
	sqlStatement := `INSERT INTO equipes
	(nome)
	VALUES($1::TEXT)`
	_, err := postgres.DB.Exec(sqlStatement, req.Nome)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("deu tudo certo")
}
