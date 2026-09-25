package checker

import (
	"net"
	"time"
)

type TCPChecker struct {
	Address string
	Timeout time.Duration
}

func (t TCPChecker) Check() Result {
	startTime := time.Now()

	conn, err := net.DialTimeout("tcp", t.Address, t.Timeout)

	if err != nil {
		return Result{
			Target:  t.Address,
			Status:  "CLOSED",
			Latency: time.Since(startTime),
			Err:     err,
		}
	}

	defer conn.Close()

	return Result{
		Target:  t.Address,
		Status:  "OPEN",
		Latency: time.Since(startTime),
		Err:     nil,
	}
}
