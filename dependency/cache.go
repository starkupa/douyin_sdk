package dependency

import (
	"context"
	"time"
)

type Cache interface {
	Get(ctx context.Context, key string) string
	Set(ctx context.Context, key string, val interface{}, timeout time.Duration) error
	IsExist(ctx context.Context, key string) bool
	Delete(ctx context.Context, key string) error
}
