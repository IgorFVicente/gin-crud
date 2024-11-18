package routes

import (
	"github.com/IgorFVicente/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorId)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.Run()
}
