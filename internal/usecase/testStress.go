package usecase

import (
	"net/http"
	"sync"
	"time"

	"github.com/YuriBertoldi/Go-TestStress/internal/model"
)

func ExecuteTestStress(c *model.Config) model.ReportTestStress {
	var successCount int
	var errorStatusCodes = make(map[int]int)
	var wg sync.WaitGroup

	startTime := time.Now()
	done := make(chan bool)

	ExecRequest := func() {
		defer wg.Done()

		client := http.Client{Timeout: time.Second * 10}

		for i := 0; i < c.GetQtdeResquesLoop(); i++ {
			resp, _ := client.Get(c.GetURL())

			if resp.StatusCode == http.StatusOK {
				successCount++
			} else {
				errorStatusCodes[resp.StatusCode]++
			}

			resp.Body.Close()
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
		TotalRequests:      c.GetQtdeRequests(),
		SuccessfulRequests: successCount,
		ErrorStatusCodes:   errorStatusCodes,
	}

	return report
}
