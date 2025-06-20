package goalex

import "github.com/Sunhill666/goalex/pkg/core"

type Client = core.Client

var PolitePool = core.PolitePool

func NewClient(opts ...core.Option) *Client {
	return core.New(opts...)
}
