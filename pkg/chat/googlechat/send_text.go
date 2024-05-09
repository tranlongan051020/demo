package googlechat

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type message struct {
	Text string `json:"text"`
}

func (p *instance) SendText(text string) error {
	sendMsg := fmt.Sprintf(
		"*%s - Environment %s - Version %s*\n%s",
		p.config.AppName,
		p.config.Env,
		p.config.Version,
		text,
	)

	return p.send(message{
		Text: sendMsg,
	})
}

func (p *instance) send(msg message) error {
	client := resty.New()
	resp, err := client.R().SetBody(msg).Post(p.config.WebhookURL)
	if err != nil || resp.StatusCode() != http.StatusOK {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("send error: %d", resp.StatusCode())
	}

	return nil
}
