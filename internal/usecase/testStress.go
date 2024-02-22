package usecase

import (
	"crypto/tls"
	"net/http"
	"sync"
	"time"

	"github.com/YuriBertoldi/Go-TestStress/internal/model"
)

func MakeRequest(url string) (*http.Response, error) {
	// Fazendo uma cópia do transporte padrão
	transport := http.DefaultTransport.(*http.Transport)
	// Configurando o TLSClientConfig na cópia
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	// Criando um cliente HTTP personalizado com o transporte modificado
	client := &http.Client{Transport: transport}
	// Fazendo a solicitação com o cliente personalizado
	return client.Get(url)
}

func ExecuteTestStress(c *model.Config) model.ReportTestStress {
	var successCount int
	var qtdeTotalRequests int
	var errorStatusCodes = make(map[int]int)
	var wg sync.WaitGroup

	startTime := time.Now()
	done := make(chan bool)

	ExecRequest := func() {
		defer wg.Done()

		for i := 0; i < c.GetQtdeResquesLoop(); i++ {

			resp, err := MakeRequest(c.GetURL())
			qtdeTotalRequests++
			if err == nil {
				if resp.StatusCode == http.StatusOK {
					successCount++
				} else {
					errorStatusCodes[resp.StatusCode]++
				}
				resp.Body.Close()
			} else {
				errorStatusCodes[0]++
			}
		}

		done <- true
	}

	for i := 0; i < c.GetQtdeConcurrency(); i++ {
		wg.Add(1)
		go ExecRequest()
	}

	go func() {
		wg.Wait()
		close(done)
	}()

	for range done {
		// used for synchronization
	}

	report := model.ReportTestStress{
		TotalTime:          time.Since(startTime),
		TotalRequests:      qtdeTotalRequests,
		SuccessfulRequests: successCount,
		ErrorStatusCodes:   errorStatusCodes,
	}

	return report
}
