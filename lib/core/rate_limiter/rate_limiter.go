package rate_limiter

import (
	"sync"
	"time"
)

// Limiter is a thread-safe, TTL-based rate limiter.
// It allows up to Max events per Window duration for any given key,
// using a sliding-window (timestamp bucket) strategy.
type Limiter struct {
	mu      sync.Mutex
	buckets map[string][]time.Time
	max     int
	window  time.Duration
	ttl     time.Duration
	now     func() time.Time

	stopCh   chan struct{}
	stopOnce sync.Once
}

// Option customizes the Limiter.
type Option func(*Limiter)

// WithTTL sets how long an idle key's bucket is kept before cleanup.
// Defaults to the window duration if not set.
func WithTTL(ttl time.Duration) Option {
	return func(l *Limiter) {
		l.ttl = ttl
	}
}

// WithClock overrides the time source (useful for testing).
func WithClock(now func() time.Time) Option {
	return func(l *Limiter) {
		l.now = now
	}
}

// New creates a Limiter allowing `max` events per `window` per key.
func New(max int, window time.Duration, opts ...Option) *Limiter {
	if max <= 0 {
		panic("rate_limiter: max must be > 0")
	}
	if window <= 0 {
		panic("rate_limiter: window must be > 0")
	}

	l := &Limiter{
		buckets: make(map[string][]time.Time),
		max:     max,
		window:  window,
		now:     time.Now,
		stopCh:  make(chan struct{}),
	}

	for _, opt := range opts {
		opt(l)
	}

	if l.ttl <= 0 {
		l.ttl = window
	}

	return l
}

// Allow reports whether an event for key is permitted right now,
// and records it if so.
func (l *Limiter) Allow(key string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := l.now()
	cutoff := now.Add(-l.window)

	times := l.buckets[key]
	times = pruneBefore(times, cutoff)

	if len(times) >= l.max {
		l.buckets[key] = times
		return false
	}

	times = append(times, now)
	l.buckets[key] = times
	return true
}

// Remaining returns how many more events key may make right now.
func (l *Limiter) Remaining(key string) int {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := l.now()
	cutoff := now.Add(-l.window)

	times := pruneBefore(l.buckets[key], cutoff)
	l.buckets[key] = times

	remaining := l.max - len(times)
	if remaining < 0 {
		remaining = 0
	}
	return remaining
}

// Reset clears the state for a specific key.
func (l *Limiter) Reset(key string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	delete(l.buckets, key)
}

// ResetAll clears state for all keys.
func (l *Limiter) ResetAll() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.buckets = make(map[string][]time.Time)
}

// pruneBefore removes timestamps older than cutoff, keeping the slice sorted.
// Assumes times is sorted ascending (true here since we always append now-ordered).
func pruneBefore(times []time.Time, cutoff time.Time) []time.Time {
	idx := 0
	for idx < len(times) && times[idx].Before(cutoff) {
		idx++
	}
	if idx == 0 {
		return times
	}
	// shift remaining elements down to reuse backing array
	n := copy(times, times[idx:])
	return times[:n]
}

// StartCleanup launches a background goroutine that periodically purges
// keys whose most recent event is older than the TTL. Call Stop to end it.
func (l *Limiter) StartCleanup(interval time.Duration) {
	if interval <= 0 {
		interval = l.ttl
	}
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				l.cleanup()
			case <-l.stopCh:
				return
			}
		}
	}()
}

func (l *Limiter) cleanup() {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := l.now()
	ttlCutoff := now.Add(-l.ttl)

	for key, times := range l.buckets {
		if len(times) == 0 {
			delete(l.buckets, key)
			continue
		}
		last := times[len(times)-1]
		if last.Before(ttlCutoff) {
			delete(l.buckets, key)
		}
	}
}

// Stop terminates the background cleanup goroutine, if running.
func (l *Limiter) Stop() {
	l.stopOnce.Do(func() {
		close(l.stopCh)
	})
}