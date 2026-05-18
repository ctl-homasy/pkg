// pkg/middleware/rate_limiter.go

package middleware

import (
	"context"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RateLimiter implements a token bucket rate limiter
type RateLimiter struct {
	tokens         float64
	maxTokens      float64
	refillRate     float64
	lastRefillTime time.Time
	mu             sync.Mutex
	clientLimiters map[string]*RateLimiter // For per-client rate limiting
}

// NewRateLimiter creates a new rate limiter with the specified rate and burst
func NewRateLimiter(rate float64, burst int) *RateLimiter {
	return &RateLimiter{
		tokens:         float64(burst),
		maxTokens:      float64(burst),
		refillRate:     rate,
		lastRefillTime: time.Now(),
		clientLimiters: make(map[string]*RateLimiter),
	}
}

// Allow checks if a request should be allowed based on the rate limit
func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(rl.lastRefillTime).Seconds()
	rl.tokens = min(rl.maxTokens, rl.tokens+(elapsed*rl.refillRate))
	rl.lastRefillTime = now

	if rl.tokens >= 1 {
		rl.tokens--
		return true
	}
	return false
}

// GetClientLimiter returns a rate limiter for a specific client
func (rl *RateLimiter) GetClientLimiter(clientID string) *RateLimiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if limiter, exists := rl.clientLimiters[clientID]; exists {
		return limiter
	}

	// Create a new limiter for this client
	clientLimiter := NewRateLimiter(rl.refillRate, int(rl.maxTokens))
	rl.clientLimiters[clientID] = clientLimiter
	return clientLimiter
}

// RateLimitInterceptor creates a gRPC interceptor for rate limiting
func RateLimitInterceptor(limiter *RateLimiter) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Extract client ID from context or metadata if needed
		// For simplicity, we're using a global rate limit here
		if !limiter.Allow() {
			return nil, status.Errorf(codes.ResourceExhausted, "rate limit exceeded")
		}
		return handler(ctx, req)
	}
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}