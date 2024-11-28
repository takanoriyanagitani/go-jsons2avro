package util

import (
	"context"
)

type Io[T any] func(context.Context) (T, error)

type Void struct{}

var Empty Void = struct{}{}

func Bind[T, U any](
	i Io[T],
	f func(T) Io[U],
) Io[U] {
	return func(ctx context.Context) (u U, e error) {
		t, e := i(ctx)
		if nil != e {
			return u, e
		}
		return f(t)(ctx)
	}
}

func Lift[T, U any](
	pure func(T) (U, error),
) func(T) Io[U] {
	return func(t T) Io[U] {
		return func(_ context.Context) (U, error) {
			return pure(t)
		}
	}
}

func Of[T any](t T) Io[T] {
	return func(_ context.Context) (T, error) {
		return t, nil
	}
}
