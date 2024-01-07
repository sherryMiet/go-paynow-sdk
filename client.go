package go_paynow_sdk

type Client struct {
	Password string
	Account  string
}

func NewClient() (c *Client) {
	return &Client{}
}