package equipe

import modelApresentacao "gerenciadorDeProjetos/domain/equipe/model"

type IEquipe interface {
	NovaEquipe(req *modelApresentacao.ReqEquipe)
	ListarEquipes()
}
