// Package cache provides an optional read-through cache.
//
// Optional is the important word. Skriptra is meant to be self-hosted, often on
// one machine, and requiring a second server to answer a question would be a
// poor trade. With REDIS_URL unset the application runs identically, just
// without the savings, so the no-op implementation is a supported deployment
// rather than a fallback for failure.
package cache

import (
	"context"
	"time"
)

// Cache is the whole surface the application needs.
//
// Deliberately not a Redis client. Handlers should not be able to reach for
// pipelines or pub/sub on a whim, and a three-method interface is trivial to
// fake in a test.
type Cache interface {
	Get(ctx context.Context, key string) ([]byte, bool)
	Set(ctx context.Context, key string, value []byte, ttl time.Duration)
	// DeletePrefix drops every key under a prefix, used to invalidate a
	// course's derived data when a document finishes indexing.
	DeletePrefix(ctx context.Context, prefix string)
	Enabled() bool
}

// CourseScope is the key prefix for everything derived from one course's
// questions. Ingesting a document invalidates the whole scope in a single call,
// so a new key never has to be remembered to be added to an invalidation list.
func CourseScope(courseID string) string { return "course:" + courseID + ":" }

// AnalyticsKey names the cached chapter-frequency aggregate.
func AnalyticsKey(courseID, variant string) string {
	return CourseScope(courseID) + "analytics:" + variant
}

// NoOp is used when no cache is configured.
//
// It never stores and never returns a hit, so every code path behaves exactly
// as it would with a cold cache. That equivalence is what makes running without
// Redis safe rather than merely tolerated.
type NoOp struct{}

func (NoOp) Get(context.Context, string) ([]byte, bool)          { return nil, false }
func (NoOp) Set(context.Context, string, []byte, time.Duration)  {}
func (NoOp) DeletePrefix(context.Context, string)                {}
func (NoOp) Enabled() bool                                       { return false }
