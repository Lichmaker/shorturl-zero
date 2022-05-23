package kafkahelper

import (
	"time"

	"github.com/Shopify/sarama"
)

type Config struct {
	Addr []string
	ClientID      string
	ProducerCount int64
	ConsumerGroup string
	Topic string
}

func GetSaramaConfig(c Config) *sarama.Config {
	clientConfig := sarama.NewConfig()
	clientConfig.ClientID = c.ClientID     // client 名称，用于给broker中日志记录
	clientConfig.Version = sarama.V3_1_0_0 // kafka server的版本号
	clientConfig.Producer.Return.Successes = true
	clientConfig.Producer.RequiredAcks = sarama.WaitForAll // 也就是等待foolower同步，才会返回
	clientConfig.Producer.Return.Errors = true
	clientConfig.Consumer.Return.Errors = true
	clientConfig.Metadata.Full = false                                                // 不用拉取全部的信息
	clientConfig.Consumer.Offsets.AutoCommit.Enable = true                            // 自动提交偏移量，默认开启
	clientConfig.Consumer.Offsets.AutoCommit.Interval = time.Second                   // 这个看业务需求，commit提交频率，不然容易down机后造成重复消费。
	clientConfig.Consumer.Offsets.Initial = sarama.OffsetOldest                       // 从最开始的地方消费，业务中看有没有需求，新业务重跑topic。
	clientConfig.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin // rb策略，默认就是range
	return clientConfig
}
