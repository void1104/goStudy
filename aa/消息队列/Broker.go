package mq

/**
Broker在MQ中担任消息的中转站（存储，转发）
*/
type Broker interface {
	publish(topic string, msg interface{}) error
	subscribe(topic string) (<-chan interface{}, error)
	unsubscribe(topic string, sub <-chan interface{}) error
	close()
	broadcast(msg interface{}, subscribers []chan interface{})
	setConditions(capacity int)
}

/**
- publish: 进行消息的推送
	- topic：所订阅的主题
	- msg：要传递的消息
- subscribe：消息的订阅，传入订阅的主题即可完成定于，并返回对应的channel通道用来接收数据
- unsubscribe：取消订阅，传入订阅的主题和对应的通道
- close：关闭消息队列
- broadcast：这个属于内部方法，作用是进行广播，对推送的消息进行广播，保证每一个订阅者都可以收到
- setConditions：用来设置条件，（这里的条件就是消息队列的容量）
*/
