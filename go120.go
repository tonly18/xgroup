//go:build go1.20
// +build go1.20

package xgroup

import "context"

func withCancelCause(parent context.Context) (context.Context, func(error)) {
	return context.WithCancelCause(parent)
}
