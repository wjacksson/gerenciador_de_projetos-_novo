package equipe

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.POST("/nova/equipe", novaEquipe)
	r.GET("")
}
