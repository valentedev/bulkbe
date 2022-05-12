package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/tomasen/realip"
	"golang.org/x/time/rate"
)

func (app *application) rateLimit(next http.Handler) http.Handler {

	type client struct {
		limiter  *rate.Limiter
		lastseen time.Time
	}

	var (
		mu      sync.Mutex
		clients = make(map[string]*client)
	)

	go func() {
		for {
			time.Sleep(time.Minute)

			mu.Lock()

			for ip, client := range clients {
				if time.Since(client.lastseen) > 3*time.Minute {
					delete(clients, ip)
				}
			}

			mu.Unlock()

		}
	}()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if app.config.limiter.enabled {

			ip := realip.FromRequest(r)

			mu.Lock()

			if _, found := clients[ip]; !found {
				clients[ip] = &client{limiter: rate.NewLimiter(rate.Limit(app.config.limiter.rps), app.config.limiter.burst)}
			}

			if !clients[ip].limiter.Allow() {
				mu.Unlock()
				app.rateLimitExceedResponse(w, r)
				return
			}

			mu.Unlock()

		}

		next.ServeHTTP(w, r)
	})
}
