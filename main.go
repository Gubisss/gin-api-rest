package main

import (
	"github.com/gubisss/api-go-gin/models"
	"github.com/gubisss/api-go-gin/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Gustavo", CPF: "1234565781", RG: "12300033300"},
		{Nome: "Ana", CPF: "44411122233", RG: "1777033300"},
	}
	routes.HandleRequests()
}
