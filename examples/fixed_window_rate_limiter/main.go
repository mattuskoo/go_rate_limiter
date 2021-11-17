package main

import (
	"time"

	"github.com/mattuskoo/go_rate_limiter"
)

func main() {
	r, err := ratelimiter.NewFixedWindowRateLimiter(&ratelimiter.Config{
		Limit:         5,
		FixedInterval: 15 * time.Second,
	})

	if err != nil {
		panic(err)
	}

	ratelimiter.DoWork(r, 10)
}
