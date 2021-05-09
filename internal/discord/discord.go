package discord

type Provider interface {
	Connect() error
	Close() error
}
