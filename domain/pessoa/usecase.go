package pessoa

import (
	"gerenciadorDeProjetos/config/database"
	"gerenciadorDeProjetos/infra/equipe"

	"github.com/gin-gonic/gin"
)

func NovaPessoa(c *gin.Context) {
	db := database.Conectar()
	defer db.Close()
	equipeRepo := equipe.NovoRepo(db)
	equipeRepo.ListarEquipes()
}
