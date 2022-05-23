package server

import (
	"context"

	"github.com/Shopify/sarama"
	"github.com/lichmaker/short-url-micro/kafkaqueue/internal/logic"
	"github.com/lichmaker/short-url-micro/pkg/kafkahelper"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type MyService struct {
	ConsumerG sarama.ConsumerGroup
	Ctx       context.Context
	CtxCancel context.CancelFunc
	GormDB    *gorm.DB
	Config    kafkahelper.Config
}

func (s *MyService) Start() {
	logx.WithContext(s.Ctx).Info("启动咯")
WORKERLOOP:
	for { // for循环的目的是因为存在重平衡，他会重新启动

		select {
		case <-s.Ctx.Done():
			logx.WithContext(s.Ctx).Info("接收到信号，结束kafka消费")
			break WORKERLOOP
		default:
			logx.Info("consume 开始...")
			handler := new(logic.ConsumerGroupHandler)
			handler.GormDB = s.GormDB
			handler.Ctx = s.Ctx                                                  // 必须传递一个handler
			err := s.ConsumerG.Consume(s.Ctx, []string{s.Config.Topic}, handler) // consume 操作，死循环。exampleConsumerGroupHandler的ConsumeClaim不允许退出，也就是操作到完毕。
			if err != nil {
				logx.Error(err)
				panic(err)
			}
		}
	}
}

func (s *MyService) Stop() {
	logx.WithContext(s.Ctx).Info("关闭咯")
	s.CtxCancel()
	logx.Close()
}
