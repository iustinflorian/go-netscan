package checker

import (
	"net/http"
	"time"
)

type HTTPChecker struct {
	URL     string
	Timeout time.Duration
}

func (h HTTPChecker) Check() Result {
	startTime := time.Now()

	client := http.Client{
		Timeout: h.Timeout,
	}

	resp, err := client.Get(h.URL)

	if err != nil {
		return Result{
			Target:  h.URL,
			Status:  "FAILED",
			Latency: time.Since(startTime),
			Err:     nil,
		}
	}

	defer resp.Body.Close()

	return Result{
		Target:  h.URL,
		Status:  resp.Status,
		Latency: time.Since(startTime),
		Err:     nil,
	}
}
