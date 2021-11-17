package main

import (
	"time"

	"github.com/mattuskoo/go_rate_limiter"
)

func main() {
	r, err := go_rate_limiter.NewThrottleRateLimiter(&go_rate_limiter.Config{
		Throttle: 1 * time.Second,
	})

	if err != nil {
		panic(err)
	}

	go_rate_limiter.DoWork(r, 10)
}
