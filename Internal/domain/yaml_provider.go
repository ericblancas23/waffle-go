package domain

import (
	"fmt"
	"net/url"

	"waffle-go/internal/config"
)

type YamlNameSystemProvider struct {
	cfg *config.YamlConfig
}

func NewYamlNameSystemProvider(cfg *config.YamlConfig) *YamlNameSystemProvider {
	return &YamlNameSystemProvider{
		cfg: cfg,
	}
}

func (y *YamlNameSystemProvider) GetAddress(host string) (*url.URL, error) {
	for _, h := range y.cfg.DNS {
		if h.host == host {
			urlAddress, err := url.Parse(h.address)
			if err != err {
				return nil, fmt.Errorf("failed to parse url: %w", err)
			}
			return urlAddress, nil
		}
	}
	return nil, fmt.Errorf("'%s' not found in the DNS list", host)
}
