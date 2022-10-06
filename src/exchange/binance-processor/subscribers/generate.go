package subscribers

//go:generate mockgen -source=order-book-subscriber.go -destination=mock/order-book-subscriber_gen.go
//go:generate mockgen -source=synthetic-order-subscriber.go -destination=mock/synthetic-order-subscriber_gen.go
//go:generate mockgen -source=tradingpairs-subscriber.go -destination=mock/tradingpairs-subscriber_gen.go
