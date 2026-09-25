package checker

import "time"

type Result struct {
	Target  string        // url/ip:port checked
	Status  string        // result (e.g. "200 OK")
	Latency time.Duration //how long the check took
	Err     error         // any error (nil if successful)
}

type Checker interface {
	Check() Result
}
