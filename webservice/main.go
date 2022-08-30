package main

import (
	"gerenciadorDeProjetos/webservice/equipe"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	equipes := r.Group("equipes")
	//pessoas := r.Group("pessoas")

	equipe.Router(equipes)
	// cadastro.Router(pessoas)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
