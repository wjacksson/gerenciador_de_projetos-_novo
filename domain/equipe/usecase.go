package equipe

import (
	"gerenciadorDeProjetos/config/database"
	modelApresentacao "gerenciadorDeProjetos/domain/equipe/model"
	"gerenciadorDeProjetos/infra/equipe"

	"github.com/gin-gonic/gin"
)

func NovaEquipe(c *gin.Context, req *modelApresentacao.ReqEquipe) {
	db := database.Conectar()
	defer db.Close()
	equipeRepo := equipe.NovoRepo(db)

	str := *req.Nome

	str = "Equipe: " + str

	req.Nome = &str

	equipeRepo.NovaEquipe(req)
}
