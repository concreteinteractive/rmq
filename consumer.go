package rmq

type Consumer interface {
	Consume(delivery Delivery)
	OnPanic(err error)
}
