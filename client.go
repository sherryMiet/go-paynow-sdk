package go_paynow_sdk

type Client struct {
	Password      string
	Account       string
	EncryptionKey string
	EncryptionIV  string
}

func NewClient() (c *Client) {
	return &Client{}
}
