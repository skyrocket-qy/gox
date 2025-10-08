package service

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"github.com/skyrocket-qy/gox/lifecyclex"
)

func NewKafkaReader(lc *lifecyclex.ConcurrentLifecycle, host, port string) *kafka.Reader {
	log.Info().Msgf("start kafka reader on %s:%s", host, port)

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{fmt.Sprintf("%s:%s", host, port)},
		Topic:   "pg.public.tuples",
	})

	lc.Add(r, func(c context.Context) error {
		return r.Close()
	})

	return r
}
