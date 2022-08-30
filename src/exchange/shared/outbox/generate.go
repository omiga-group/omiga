package outbox

//go:generate mockgen -source=outbox-background-service.go -destination=mock/outbox-background-service_gen.go
//go:generate mockgen -source=outbox-publisher.go -destination=mock/outbox-publisher_gen.go
