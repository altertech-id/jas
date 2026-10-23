package middleware

import (
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

type clientRecord struct {
	count     int
	resetTime time.Time
}

type RateLimiter struct {
	mu      sync.Mutex
	clients map[string]*clientRecord
	limit   int
	window  time.Duration
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	rl := &RateLimiter{
		clients: make(map[string]*clientRecord),
		limit:   limit,
		window:  window,
	}

	// Periodically cleanup expired clients
	go func() {
		ticker := time.NewTicker(rl.window * 2)
		for range ticker.C {
			rl.mu.Lock()
			now := time.Now()
			for ip, record := range rl.clients {
				if now.After(record.resetTime) {
					delete(rl.clients, ip)
				}
			}
			rl.mu.Unlock()
		}
	}()

	return rl
}

func (rl *RateLimiter) Middleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ip := c.IP()
		now := time.Now()

		rl.mu.Lock()
		record, exists := rl.clients[ip]
		if !exists || now.After(record.resetTime) {
			rl.clients[ip] = &clientRecord{
				count:     1,
				resetTime: now.Add(rl.window),
			}
			rl.mu.Unlock()
			return c.Next()
		}

		if record.count >= rl.limit {
			rl.mu.Unlock()
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Terlalu banyak permintaan dari perangkat Anda. Mohon tunggu beberapa saat sebelum mencoba kembali.",
			})
		}

		record.count++
		rl.mu.Unlock()

		return c.Next()
	}
}
