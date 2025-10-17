package service

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type KafkaConnection interface {
	Close() error
}

type KafkaDialer interface {
	DialLeader(
		ctx context.Context,
		network, address, topic string,
		partition int,
	) (KafkaConnection, error)
}

type kafkaDialer struct{}

func NewKafkaDialer() KafkaDialer {
	return &kafkaDialer{}
}

func (d *kafkaDialer) DialLeader(
	ctx context.Context,
	network, address, topic string,
	partition int,
) (KafkaConnection, error) {
	return kafka.DialLeader(ctx, network, address, topic, partition)
}
