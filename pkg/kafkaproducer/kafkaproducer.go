package kafkaproducer

import (
	"context"
	"log"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
)

type MyProducer struct {
	Ch        chan *sarama.ProducerMessage
	Producer  sarama.AsyncProducer
	Ctx       context.Context
	CtxStopFn context.CancelFunc
}

func MustNew(ctx context.Context, cancelFn context.CancelFunc, addr []string, conf *sarama.Config) *MyProducer {
	client, err := sarama.NewClient(addr, conf)
	if err != nil {
		panic(err)
	}

	producer, err := sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		panic(err)
	}

	ch := make(chan *sarama.ProducerMessage, 10)
	instance := &MyProducer{
		Producer:  producer,
		Ch:        ch,
		Ctx:       ctx,
		CtxStopFn: cancelFn,
	}
	return instance
}

func (p *MyProducer) Start() {

	var (
		wg                                  sync.WaitGroup
		enqueued, successes, producerErrors int
	)

	// Trap SIGINT to trigger a graceful shutdown.
	// signals := make(chan os.Signal, 1)
	// signal.Notify(signals, os.Interrupt)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for msg := range p.Producer.Successes() {
			logx.Infof("kafka发送成功 %+v", msg)
			successes++
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for err := range p.Producer.Errors() {
			log.Println(err)
			producerErrors++
		}
	}()

	threading.GoSafe(func() {
		logx.Info("已启动kafkaProducer...")
	ProducerLoop:
		for {
			msg := <-p.Ch
			logx.Infof("读到消息啦，准备写入 %+v", msg)
			// message := &sarama.ProducerMessage{Topic: "my_topic", Value: sarama.StringEncoder("testing 123")}
			select {
			case p.Producer.Input() <- msg:
				enqueued++

			case <-p.Ctx.Done():
				p.Producer.AsyncClose() // Trigger a shutdown of the producer.
				break ProducerLoop
			}
		}
	})

	wg.Wait()
}

func (p *MyProducer) Stop() {
	logx.Info("正在关闭kafka生产者...")
	p.CtxStopFn()
}
