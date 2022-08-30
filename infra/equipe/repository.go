package equipe

import (
	"database/sql"
	modelApresentacao "gerenciadorDeProjetos/domain/equipe/model"
	modelData "gerenciadorDeProjetos/infra/equipe/model"
	"gerenciadorDeProjetos/infra/equipe/postgres"
)

type respositorio struct {
	Data *postgres.DBEquipe
}

func novoRepo(novoDB *sql.DB) *respositorio {
	return &respositorio{
		Data: &postgres.DBEquipe{DB: novoDB},
	}
}

func (r *respositorio) NovaEquipe(req *modelApresentacao.ReqEquipe) {
	r.Data.NovaEquipe(&modelData.Equipe{Nome: req.Nome})
}
func (r *respositorio) ListarEquipes() {

}
