package rmq

type BatchConsumer interface {
	Consume(batch Deliveries)
	OnPanic(err error)
}
