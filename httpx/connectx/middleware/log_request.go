package middleware

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	"github.com/rs/zerolog/log"
)

func NewLogRequest() connect.UnaryInterceptorFunc {
	return connect.UnaryInterceptorFunc(
		func(next connect.UnaryFunc) connect.UnaryFunc {
			return connect.UnaryFunc(func(
				ctx context.Context,
				req connect.AnyRequest,
			) (connect.AnyResponse, error) {
				// Call the real handler
				res, err := next(ctx, req)
				if err != nil {
					// Try to see if it's a Connect error (to get code)
					cerr := &connect.Error{}
					if errors.As(err, &cerr) {
						log.Printf("Request %s failed: code=%s, err=%v\n",
							req.Spec().Procedure,
							cerr.Code(),
							cerr.Message(),
						)
					} else {
						log.Printf("Request %s failed: err=%v\n",
							req.Spec().Procedure, err)
					}
				}

				return res, err
			})
		},
	)
}
