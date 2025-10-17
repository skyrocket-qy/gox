package middleware

import (
	"context"

	"connectrpc.com/connect"
	"github.com/skyrocket-qy/erx"
	"github.com/skyrocket-qy/gox/errcode"
)

type ReqPool struct {
	ch             chan struct{}
	isReturnOnFull bool
}

func NewReqPool(size int, isReturnOnFull bool) *ReqPool {
	return &ReqPool{
		ch:             make(chan struct{}, size),
		isReturnOnFull: isReturnOnFull,
	}
}

func (m *ReqPool) UnaryInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(
			func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {

				if m.isReturnOnFull {
					select {
					case m.ch <- struct{}{}:
					default:
						return nil, connect.NewError(
							connect.CodeResourceExhausted,
							erx.New(errcode.ErrTooManyRequest),
						)
					}
				} else {
					select {
					case m.ch <- struct{}{}:
					case <-ctx.Done():
						return nil, connect.NewError(
							connect.CodeCanceled,
							ctx.Err(),
						)
					}
				}

				defer func() {
					<-m.ch
				}()

				return next(ctx, req)
			},
		)
	}
}
