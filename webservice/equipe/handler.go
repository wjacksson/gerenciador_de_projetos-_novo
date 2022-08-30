package equipe

import (
	"fmt"
	"gerenciadorDeProjetos/domain/equipe"
	modelApresentacao "gerenciadorDeProjetos/domain/equipe/model"

	"github.com/gin-gonic/gin"
)

func novaEquipe(c *gin.Context) {
	fmt.Println("tentando adicionar nova equipe")
	req := modelApresentacao.ReqEquipe{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Could not create. Parameters were not passed correctly " + err.Error(),
		})
		return
	}

	equipe.NovaEquipe(c, &req)
}
