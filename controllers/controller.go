package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gubisss/api-go-gin/models"
)

func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos)

}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "Eae " + nome + ", tudo certo?",
	})
}
