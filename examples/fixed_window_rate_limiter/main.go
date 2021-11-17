package main

import (
	"time"

	"github.com/mattuskoo/go_rate_limiter"
)

func main() {
	r, err := go_rate_limiter.NewFixedWindowRateLimiter(&go_rate_limiter.Config{
		Limit:         5,
		FixedInterval: 15 * time.Second,
	})

	if err != nil {
		panic(err)
	}

	go_rate_limiter.DoWork(r, 10)
}
