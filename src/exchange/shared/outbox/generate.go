package outbox

//go:generate mockgen -source=outbox-background-service.go -destination=mock/mock-outbox-background-service.go
//go:generate mockgen -source=outbox-publisher.go -destination=mock/mock-outbox-publisher.go
