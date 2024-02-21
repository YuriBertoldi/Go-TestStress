package usecase_test

import (
	"testing"
	"time"

	"github.com/YuriBertoldi/Go-TestStress/internal/model"
	"github.com/YuriBertoldi/Go-TestStress/internal/usecase"
)

func TestExecuteTestStress(t *testing.T) {
	config := model.CarregarParamsParaTest()

	startTime := time.Now()
	report := usecase.ExecuteTestStress(&config)
	elapsedTime := time.Since(startTime)

	QtdeRequetsSucesso := config.GetQtdeConcurrency() * config.GetQtdeResquesLoop()
	if report.SuccessfulRequests != QtdeRequetsSucesso {
		t.Errorf("Quantidade total de requisições com sucesso diferente da esperado. Quatidade do relatório igual a %d, esperada é  %d", report.SuccessfulRequests, QtdeRequetsSucesso)
	}

	if report.TotalRequests != config.GetQtdeRequests() {
		t.Errorf("Quantidade total de requests diferente da esperado. Quatidade do relatório igual a %d, esperada é %d", report.TotalRequests, config.GetQtdeRequests())
	}

	if report.TotalTime > elapsedTime {
		t.Errorf("Tempo total do relatório está maior que o esperado. Tempo do relatório igual %d, esperado menor ou igual a %d", report.TotalTime, elapsedTime)
	}

}
