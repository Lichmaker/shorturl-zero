package logic

import (
	"context"

	"github.com/Shopify/sarama"
	"github.com/lichmaker/short-url-micro/model/shorts"
	"github.com/lichmaker/short-url-micro/pkg/statistics"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type ConsumerGroupHandler struct {
	GormDB *gorm.DB
	Ctx    context.Context
}

var _ sarama.ConsumerGroupHandler = ConsumerGroupHandler{}

func (h ConsumerGroupHandler) Setup(s sarama.ConsumerGroupSession) error {
	logx.Info("kafka连接完成")
	return nil
}

func (h ConsumerGroupHandler) Cleanup(s sarama.ConsumerGroupSession) error {
	logx.Info("kafka消费 cleanup")
	return nil
}

func (h ConsumerGroupHandler) ConsumeClaim(sSession sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
FORLOOP:
	for msg := range claim.Messages() { // 接受topic消息
		select {
		case <-h.Ctx.Done():
			logx.WithContext(h.Ctx).Info("ConsumeClaim接收到信号，结束kafka消费")
			break FORLOOP
		default:
			logx.Infof("[Consumer] Message topic:%q partition:%d offset:%d add:%d body:%s \n", msg.Topic, msg.Partition, msg.Offset, claim.HighWaterMarkOffset()-msg.Offset, msg.Value)
			m, err := shorts.GetByShort(h.GormDB, string(msg.Value))
			if err != nil {
				logx.Errorf("消息处理异常，查询数据失败. msg %s , value %s", err.Error(), msg.Value)
			} else {
				statistics.Do(h.GormDB, m)
			}
			sSession.MarkMessage(msg, "") // 必须设置这个，不然你的偏移量无法提交。
		}
	}
	return nil
}
