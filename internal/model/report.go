package model

import (
	"fmt"
	"time"
)

type ReportTestStress struct {
	TotalTime          time.Duration
	TotalRequests      int
	SuccessfulRequests int
	ErrorStatusCodes   map[int]int
}

func (r *ReportTestStress) ExibirReport() {
	fmt.Println("Relatório do teste de stress:")
	fmt.Printf("* Tempo total em milisegundos: %d\n", int(r.TotalTime/time.Millisecond))
	fmt.Printf("* Quantidade total de requisições: %d\n", r.TotalRequests)
	fmt.Printf("* Quantidade total de requisições com status code 200: %d\n", r.SuccessfulRequests)

	for code, count := range r.ErrorStatusCodes {
		fmt.Printf("* Status Code %d: %d\n", code, count)
	}
}
