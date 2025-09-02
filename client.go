package go_fivepay

import (
	"github.com/asaka1234/go-fivepay/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	Params *FivePayInitParams

	ryClient  *resty.Client
	debugMode bool
	logger    utils.Logger
}

func NewClient(logger utils.Logger, params *FivePayInitParams) *Client {
	return &Client{
		Params: params,

		ryClient:  resty.New(), //client实例
		debugMode: false,
		logger:    logger,
	}
}

func (cli *Client) SetDebugModel(debugMode bool) {
	cli.debugMode = debugMode
}
