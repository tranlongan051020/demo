package googlechat

type Chat interface {
	SendText(text string) error
}

type ChatConfig struct {
	WebhookURL string
	AppName    string
	Env        string
	Version    string
}

type instance struct {
	config ChatConfig
}

func New(config ChatConfig) Chat {
	return &instance{
		config: config,
	}
}
