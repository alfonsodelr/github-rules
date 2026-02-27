package clients

type Client struct {
	version string
}

func NewClient(version string) *Client {
	return &Client{
		version: version,
	}
}

func (c *Client) GetVersion() string {
	return c.version
}
