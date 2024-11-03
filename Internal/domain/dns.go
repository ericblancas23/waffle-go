package domain

import "net/url"

type NameSystemprovider interface {
	GetAddress(host string) (*url.URL, error)
}
