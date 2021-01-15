package util

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"task5/pkg/setting"
)

type rocketMqMessage struct {
	messageBytes []byte
	messageId    int64
	tag          string
}

var RocketMqProducer rocketmq.Producer

var rocketMqInChan chan rocketMqMessage

func init() {
	fmt.Println("初始化了rocketMqInChan")
	rocketMqInChan = make(chan rocketMqMessage, 10000)
	go func() {
		for {
			select {
			case one := <-rocketMqInChan:
				SendToRocketMq(one)
			}
		}
	}()
}

func SendMessageToRocketChan(messageBytes []byte, messageId int64, tag string) {
	rocketMqInChan <- rocketMqMessage{
		messageBytes: messageBytes,
		messageId:    messageId,
		tag:          tag,
	}
}

func SendToRocketMq(one rocketMqMessage) {
	msg := &primitive.Message{
		Topic: "Wuxi",
		Body:  one.messageBytes,
	}
	msg.WithTag(one.tag)
	result, err := RocketMqProducer.SendSync(context.Background(), msg)
	if err != nil {
		StartRocketMQ()
		return
	}
	if one.messageId == 0 {
		//发送成功
		fmt.Println(result.MsgID)
	}
}

func StartRocketMQ() {
	//记录一些日志信息
	p, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{setting.RocketAddress}),
		producer.WithRetry(3),
	)
	if err != nil {
		//记录出错信息
		return
	}
	err = p.Start()
	if err != nil {
		//同样记录出错信息
		return
	}
	RocketMqProducer = p
}
