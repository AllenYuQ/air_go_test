package util

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"os"
)

var p, _ = rocketmq.NewProducer(
	producer.WithNameServer([]string{"127.0.0.1:9876"}),
	producer.WithRetry(2),
)

// Package main implements a simple producer to send message.

func init() {
	err := p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		os.Exit(1)
	}
}

func SimpleSendMsgToRocket() {

	topic := "test"

	msg := &primitive.Message{
		Topic: topic,
		Body:  []byte("Hello RocketMQ Go Client!"),
	}
	res, err := p.SendSync(context.Background(), msg)

	if err != nil {
		fmt.Printf("send message error: %s\n", err)
	} else {
		fmt.Printf("send message success: result=%s\n", res.String())
	}
}
