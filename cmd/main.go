package main

import (
	"github.com/YuriBertoldi/Go-TestStress/internal/model"
	"github.com/YuriBertoldi/Go-TestStress/internal/usecase"
)

func main() {
	config := model.CarregarParams()
	report := usecase.ExecuteTestStress(&config)
	report.ExibirReport()
}
