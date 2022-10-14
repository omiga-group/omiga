package subscribers

//go:generate mockgen -source=order_book_subscriber.go -destination=mock/order_book_subscriber_gen.go
//go:generate mockgen -source=synthetic_order_subscriber.go -destination=mock/synthetic_order_subscriber_gen.go
//go:generate mockgen -source=trading_pair_subscriber.go -destination=mock/trading_pair_subscriber_gen.go
